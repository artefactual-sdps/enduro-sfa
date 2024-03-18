package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/jonboulle/clockwork"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
	"go.artefactual.dev/amclient"
	"go.artefactual.dev/tools/log"
	temporalsdk_activity "go.temporal.io/sdk/activity"
	temporalsdk_client "go.temporal.io/sdk/client"
	temporalsdk_contrib_opentelemetry "go.temporal.io/sdk/contrib/opentelemetry"
	temporalsdk_interceptor "go.temporal.io/sdk/interceptor"
	temporalsdk_worker "go.temporal.io/sdk/worker"
	"gocloud.dev/blob"

	"github.com/artefactual-sdps/enduro/internal/am"
	"github.com/artefactual-sdps/enduro/internal/api/auth"
	"github.com/artefactual-sdps/enduro/internal/config"
	"github.com/artefactual-sdps/enduro/internal/db"
	"github.com/artefactual-sdps/enduro/internal/event"
	"github.com/artefactual-sdps/enduro/internal/package_"
	"github.com/artefactual-sdps/enduro/internal/sftp"
	"github.com/artefactual-sdps/enduro/internal/storage"
	"github.com/artefactual-sdps/enduro/internal/telemetry"
	"github.com/artefactual-sdps/enduro/internal/temporal"
	"github.com/artefactual-sdps/enduro/internal/version"
	"github.com/artefactual-sdps/enduro/internal/watcher"
	"github.com/artefactual-sdps/enduro/internal/workflow/activities"
)

const (
	appName = "enduro-am-worker"
)

func main() {
	p := pflag.NewFlagSet(appName, pflag.ExitOnError)

	p.String("config", "", "Configuration file")
	p.Bool("version", false, "Show version information")
	_ = p.Parse(os.Args[1:])

	if v, _ := p.GetBool("version"); v {
		fmt.Println(version.Info(appName))
		os.Exit(0)
	}

	var cfg config.Configuration
	configFile, _ := p.GetString("config")
	configFileFound, configFileUsed, err := config.Read(&cfg, configFile)
	if err != nil {
		fmt.Printf("Failed to read configuration: %v\n", err)
		os.Exit(1)
	}

	logger := log.New(os.Stderr,
		log.WithName(appName),
		log.WithDebug(cfg.Debug),
		log.WithLevel(cfg.Verbosity),
	)
	defer log.Sync(logger)

	logger.Info("Starting...", "version", version.Long, "pid", os.Getpid())

	if configFileFound {
		logger.Info("Configuration file loaded.", "path", configFileUsed)
	} else {
		logger.Info("Configuration file not found.")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up the tracer provider.
	tp, shutdown, err := telemetry.TracerProvider(ctx, logger, cfg.Telemetry, appName, version.Long)
	if err != nil {
		logger.Error(err, "Error creating tracer provider.")
		os.Exit(1)
	}
	defer func() { _ = shutdown(ctx) }()

	enduroDatabase, err := db.Connect(ctx, tp, cfg.Database.Driver, cfg.Database.DSN)
	if err != nil {
		logger.Error(err, "Enduro database configuration failed.")
		os.Exit(1)
	}
	_ = enduroDatabase.Ping()

	// Set up the Temporal client.
	tracingInterceptor, err := temporalsdk_contrib_opentelemetry.NewTracingInterceptor(
		temporalsdk_contrib_opentelemetry.TracerOptions{
			Tracer: tp.Tracer("temporal-sdk-go"),
		},
	)
	if err != nil {
		logger.Error(err, "Unable to create OpenTelemetry interceptor.")
		os.Exit(1)
	}
	temporalClient, err := temporalsdk_client.Dial(temporalsdk_client.Options{
		Namespace:    cfg.Temporal.Namespace,
		HostPort:     cfg.Temporal.Address,
		Logger:       temporal.Logger(logger.WithName("temporal-client")),
		Interceptors: []temporalsdk_interceptor.ClientInterceptor{tracingInterceptor},
	})
	if err != nil {
		logger.Error(err, "Error creating Temporal client.")
		os.Exit(1)
	}

	// Set up the watcher service.
	var wsvc watcher.Service
	{
		wsvc, err = watcher.New(ctx, tp, logger.WithName("watcher"), &cfg.Watcher)
		if err != nil {
			logger.Error(err, "Error setting up watchers.")
			os.Exit(1)
		}
	}

	// Set up the event service.
	evsvc, err := event.NewEventServiceRedis(logger.WithName("events"), tp, &cfg.Event)
	if err != nil {
		logger.Error(err, "Error creating Event service.")
		os.Exit(1)
	}

	// Set up the package service.
	var pkgSvc package_.Service
	{
		pkgSvc = package_.NewService(logger.WithName("package"), enduroDatabase, temporalClient, evsvc, &auth.NoopTokenVerifier{}, nil, cfg.Temporal.TaskQueue)
	}

	// Set-up failed transfers bucket.
	var ft *blob.Bucket
	{
		fl, err := storage.NewInternalLocation(&cfg.FailedTransfers)
		if err != nil {
			logger.Error(err, "Error setting up failed transfers location.")
			os.Exit(1)
		}
		ft, err = fl.OpenBucket(ctx)
		if err != nil {
			logger.Error(err, "Error getting failed transfers bucket.")
			os.Exit(1)
		}
	}

	// Set-up failed SIPs bucket.
	var fs *blob.Bucket
	{
		fl, err := storage.NewInternalLocation(&cfg.FailedSIPs)
		if err != nil {
			logger.Error(err, "Error setting up failed SIPs location.")
			os.Exit(1)
		}
		fs, err = fl.OpenBucket(ctx)
		if err != nil {
			logger.Error(err, "Error getting failed SIPs bucket.")
			os.Exit(1)
		}
	}

	var g run.Group

	// Activity worker.
	{
		logger.V(1).Info("AM worker config", "capacity", cfg.AM.Capacity)

		done := make(chan struct{})
		workerOpts := temporalsdk_worker.Options{
			DisableWorkflowWorker:              true,
			EnableSessionWorker:                true,
			MaxConcurrentSessionExecutionSize:  cfg.AM.Capacity,
			MaxConcurrentActivityExecutionSize: 1,
		}
		w := temporalsdk_worker.New(temporalClient, temporal.AmWorkerTaskQueue, workerOpts)
		if err != nil {
			logger.Error(err, "Error creating Temporal worker.")
			os.Exit(1)
		}

		httpClient := cleanhttp.DefaultPooledClient()
		sftpClient := sftp.NewGoClient(logger, cfg.AM.SFTP)
		amc := amclient.NewClient(httpClient, cfg.AM.Address, cfg.AM.User, cfg.AM.APIKey)

		w.RegisterActivityWithOptions(
			activities.NewDownloadActivity(logger, wsvc).Execute,
			temporalsdk_activity.RegisterOptions{Name: activities.DownloadActivityName},
		)
		w.RegisterActivityWithOptions(
			activities.NewUnarchiveActivity(logger).Execute,
			temporalsdk_activity.RegisterOptions{Name: activities.UnarchiveActivityName},
		)
		w.RegisterActivityWithOptions(
			activities.NewBundleActivity(logger).Execute,
			temporalsdk_activity.RegisterOptions{Name: activities.BundleActivityName},
		)
		w.RegisterActivityWithOptions(
			activities.NewZipActivity(logger).Execute, temporalsdk_activity.RegisterOptions{Name: activities.ZipActivityName},
		)
		w.RegisterActivityWithOptions(
			am.NewUploadTransferActivity(logger, sftpClient, cfg.AM.PollInterval).Execute,
			temporalsdk_activity.RegisterOptions{Name: am.UploadTransferActivityName},
		)
		w.RegisterActivityWithOptions(
			am.NewDeleteTransferActivity(logger, sftpClient).Execute,
			temporalsdk_activity.RegisterOptions{Name: am.DeleteTransferActivityName},
		)
		w.RegisterActivityWithOptions(
			am.NewStartTransferActivity(logger, &cfg.AM, amc.Package).Execute,
			temporalsdk_activity.RegisterOptions{Name: am.StartTransferActivityName},
		)
		w.RegisterActivityWithOptions(
			am.NewPollTransferActivity(
				logger,
				&cfg.AM,
				clockwork.NewRealClock(),
				amc.Transfer,
				amc.Jobs,
				pkgSvc,
			).Execute,
			temporalsdk_activity.RegisterOptions{Name: am.PollTransferActivityName},
		)
		w.RegisterActivityWithOptions(
			am.NewPollIngestActivity(
				logger,
				&cfg.AM,
				clockwork.NewRealClock(),
				amc.Ingest,
				amc.Jobs,
				pkgSvc,
			).Execute,
			temporalsdk_activity.RegisterOptions{Name: am.PollIngestActivityName},
		)
		w.RegisterActivityWithOptions(
			activities.NewCleanUpActivity().Execute,
			temporalsdk_activity.RegisterOptions{Name: activities.CleanUpActivityName},
		)
		w.RegisterActivityWithOptions(
			activities.NewSendToFailedBuckeActivity(ft, fs).Execute,
			temporalsdk_activity.RegisterOptions{Name: activities.SendToFailedBucketName},
		)

		g.Add(
			func() error {
				if err := w.Start(); err != nil {
					return err
				}
				<-done
				return nil
			},
			func(err error) {
				w.Stop()
				close(done)
			},
		)
	}

	// Observability server.
	{
		srv := &http.Server{
			Addr:         cfg.DebugListen,
			ReadTimeout:  time.Second * 1,
			WriteTimeout: time.Second * 1,
			IdleTimeout:  time.Second * 30,
		}

		g.Add(func() error {
			mux := http.NewServeMux()

			// Health check.
			mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, "OK")
			})

			// Prometheus metrics.
			mux.Handle("/metrics", promhttp.Handler())

			// Profiling data.
			mux.HandleFunc("/debug/pprof/", pprof.Index)
			mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
			mux.Handle("/debug/pprof/block", pprof.Handler("block"))
			mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
			mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
			mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

			srv.Handler = mux

			return srv.ListenAndServe()
		}, func(error) {
			ctx, cancel := context.WithTimeout(ctx, time.Second*5)
			defer cancel()
			_ = srv.Shutdown(ctx)
		})
	}

	// Signal handler.
	{
		var (
			cancelInterrupt = make(chan struct{})
			ch              = make(chan os.Signal, 2)
		)
		defer close(ch)

		g.Add(
			func() error {
				signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

				select {
				case <-ch:
				case <-cancelInterrupt:
				}

				return nil
			}, func(err error) {
				logger.Info("Quitting...")
				close(cancelInterrupt)
				cancel()
				signal.Stop(ch)
			},
		)
	}

	err = g.Run()
	if err != nil {
		logger.Error(err, "Application failure.")
		os.Exit(1)
	}
	logger.Info("Bye!")
}
