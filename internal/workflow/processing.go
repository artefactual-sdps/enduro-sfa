// Package workflow contains an experimental workflow for Archivemica transfers.
//
// It's not generalized since it contains client-specific activities. However,
// the long-term goal is to build a system where workflows and activities are
// dynamically set up based on user input.
package workflow

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	temporalsdk_temporal "go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	temporalsdk_workflow "go.temporal.io/sdk/workflow"

	"github.com/artefactual-sdps/enduro/internal/a3m"
	"github.com/artefactual-sdps/enduro/internal/am"
	"github.com/artefactual-sdps/enduro/internal/fsutil"
	"github.com/artefactual-sdps/enduro/internal/package_"
	"github.com/artefactual-sdps/enduro/internal/ref"
	sfa_activities "github.com/artefactual-sdps/enduro/internal/sfa/activities"
	"github.com/artefactual-sdps/enduro/internal/temporal"
	"github.com/artefactual-sdps/enduro/internal/watcher"
	"github.com/artefactual-sdps/enduro/internal/workflow/activities"
)

type ProcessingWorkflow struct {
	logger           logr.Logger
	pkgsvc           package_.Service
	wsvc             watcher.Service
	useArchivematica bool
	cleanUpPaths     []string
}

func NewProcessingWorkflow(logger logr.Logger, pkgsvc package_.Service, wsvc watcher.Service, useAm bool) *ProcessingWorkflow {
	return &ProcessingWorkflow{
		logger:           logger,
		pkgsvc:           pkgsvc,
		wsvc:             wsvc,
		useArchivematica: useAm,
	}
}

// TransferInfo is shared state that is passed down to activities. It can be
// useful for hooks that may require quick access to processing state.
type TransferInfo struct {
	// It is populated by the workflow request.
	req package_.ProcessingWorkflowRequest

	// TempFile is the temporary location where the blob is downloaded.
	//
	// It is populated by the workflow with the result of DownloadActivity.
	TempFile string

	// SIPID given by a3m.
	//
	// It is populated by CreateAIPActivity.
	SIPID string

	// Path to the compressed AIP generated by the preservation system.
	//
	// It is populated once the preservation system creates the AIP.
	AIPPath string

	// StoredAt is the time when the AIP is stored.
	//
	// It is populated by PollIngestActivity as long as Ingest completes.
	StoredAt time.Time

	// Information about the bundle (transfer) that we submit to Archivematica.
	// Full path, relative path...
	//
	// It is populated by BundleActivity.
	Bundle activities.BundleActivityResult

	// Identifier of the preservation action that creates the AIP
	//
	// It is populated by createPreservationActionLocalActivity .
	PreservationActionID uint
}

func (t *TransferInfo) Name() string {
	return fsutil.BaseNoExt(t.req.Key)
}

// ProcessingWorkflow orchestrates all the activities related to the processing
// of a SIP in Archivematica, including is retrieval, creation of transfer,
// etc...
//
// Retrying this workflow would result in a new Archivematica transfer. We  do
// not have a retry policy in place. The user could trigger a new instance via
// the API.
func (w *ProcessingWorkflow) Execute(ctx temporalsdk_workflow.Context, req *package_.ProcessingWorkflowRequest) error {
	var (
		logger = temporalsdk_workflow.GetLogger(ctx)

		tinfo = &TransferInfo{
			req: *req,
		}

		// Package status. All packages start in queued status.
		status = package_.StatusQueued

		// Create AIP preservation action status.
		paStatus = package_.ActionStatusUnspecified
	)

	// Persist package as early as possible.
	{
		activityOpts := withLocalActivityOpts(ctx)
		var err error

		if req.PackageID == 0 {
			err = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, createPackageLocalActivity, w.logger, w.pkgsvc, &createPackageLocalActivityParams{
				Key:    req.Key,
				Status: status,
			}).Get(activityOpts, &tinfo.req.PackageID)
		} else {
			// TODO: investigate better way to reset the package_.
			err = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, updatePackageLocalActivity, w.logger, w.pkgsvc, &updatePackageLocalActivityParams{
				PackageID: req.PackageID,
				Key:       req.Key,
				SIPID:     "",
				StoredAt:  temporalsdk_workflow.Now(ctx).UTC(),
				Status:    status,
			}).Get(activityOpts, nil)
		}

		if err != nil {
			return fmt.Errorf("error persisting package: %v", err)
		}
	}

	// Ensure that the status of the package and the preservation action is always updated when this
	// workflow function returns.
	defer func() {
		// Mark as failed unless it completed successfully or it was abandoned.
		if status != package_.StatusDone && status != package_.StatusAbandoned {
			status = package_.StatusError
		}

		// Use disconnected context so it also runs after cancellation.
		dctx, _ := temporalsdk_workflow.NewDisconnectedContext(ctx)
		activityOpts := withLocalActivityOpts(dctx)
		_ = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, updatePackageLocalActivity, w.logger, w.pkgsvc, &updatePackageLocalActivityParams{
			PackageID: tinfo.req.PackageID,
			Key:       tinfo.req.Key,
			SIPID:     tinfo.SIPID,
			StoredAt:  tinfo.StoredAt,
			Status:    status,
		}).Get(activityOpts, nil)

		if paStatus != package_.ActionStatusDone {
			paStatus = package_.ActionStatusError
		}

		_ = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, completePreservationActionLocalActivity, w.pkgsvc, &completePreservationActionLocalActivityParams{
			PreservationActionID: tinfo.PreservationActionID,
			Status:               paStatus,
			CompletedAt:          temporalsdk_workflow.Now(dctx).UTC(),
		}).Get(activityOpts, nil)
	}()

	// Activities running within a session.
	{
		var sessErr error
		var taskQueue string
		maxAttempts := 5

		if w.useArchivematica {
			taskQueue = temporal.AmWorkerTaskQueue
		} else {
			taskQueue = temporal.A3mWorkerTaskQueue
		}

		activityOpts := temporalsdk_workflow.WithActivityOptions(ctx, temporalsdk_workflow.ActivityOptions{
			StartToCloseTimeout: time.Minute,
			TaskQueue:           taskQueue,
		})
		for attempt := 1; attempt <= maxAttempts; attempt++ {
			sessCtx, err := temporalsdk_workflow.CreateSession(activityOpts, &temporalsdk_workflow.SessionOptions{
				CreationTimeout:  forever,
				ExecutionTimeout: forever,
			})
			if err != nil {
				return fmt.Errorf("error creating session: %v", err)
			}

			sessErr = w.SessionHandler(sessCtx, attempt, tinfo)

			// We want to retry the session if it has been canceled as a result
			// of losing the worker but not otherwise. This scenario seems to be
			// identifiable when we have an error but the root context has not
			// been canceled.
			if sessErr != nil && (errors.Is(sessErr, temporalsdk_workflow.ErrSessionFailed) || temporalsdk_temporal.IsCanceledError(sessErr)) {
				// Root context canceled, hence workflow canceled.
				if ctx.Err() == temporalsdk_workflow.ErrCanceled {
					return nil
				}

				logger.Error(
					"Session failed, will retry shortly (10s)...",
					"err", ctx.Err(),
					"attemptFailed", attempt,
					"attemptsLeft", maxAttempts-attempt,
				)

				_ = temporalsdk_workflow.Sleep(ctx, time.Second*10)

				continue
			}
			break
		}

		if sessErr != nil {
			return sessErr
		}

		status = package_.StatusDone

		paStatus = package_.ActionStatusDone
	}

	// Schedule deletion of the original in the watched data source.
	{
		if status == package_.StatusDone {
			if tinfo.req.RetentionPeriod != nil {
				err := temporalsdk_workflow.NewTimer(ctx, *tinfo.req.RetentionPeriod).Get(ctx, nil)
				if err != nil {
					logger.Warn("Retention policy timer failed", "err", err.Error())
				} else {
					activityOpts := withActivityOptsForRequest(ctx)
					_ = temporalsdk_workflow.ExecuteActivity(activityOpts, activities.DeleteOriginalActivityName, tinfo.req.WatcherName, tinfo.req.Key).Get(activityOpts, nil)
				}
			} else if tinfo.req.CompletedDir != "" {
				activityOpts := withActivityOptsForLocalAction(ctx)
				_ = temporalsdk_workflow.ExecuteActivity(activityOpts, activities.DisposeOriginalActivityName, tinfo.req.WatcherName, tinfo.req.CompletedDir, tinfo.req.Key).Get(activityOpts, nil)
			}
		}
	}

	logger.Info(
		"Workflow completed successfully!",
		"packageID", tinfo.req.PackageID,
		"watcher", tinfo.req.WatcherName,
		"key", tinfo.req.Key,
		"status", status.String(),
	)

	return nil
}

// SessionHandler runs activities that belong to the same session.
func (w *ProcessingWorkflow) SessionHandler(sessCtx temporalsdk_workflow.Context, attempt int, tinfo *TransferInfo) error {
	defer temporalsdk_workflow.CompleteSession(sessCtx)

	packageStartedAt := temporalsdk_workflow.Now(sessCtx).UTC()

	// Set in-progress status.
	{
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setStatusInProgressLocalActivity, w.pkgsvc, tinfo.req.PackageID, packageStartedAt).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	// Persist the preservation action for creating the AIP.
	{
		{
			var preservationActionType package_.PreservationActionType
			if tinfo.req.AutoApproveAIP {
				preservationActionType = package_.ActionTypeCreateAIP
			} else {
				preservationActionType = package_.ActionTypeCreateAndReviewAIP
			}

			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, createPreservationActionLocalActivity, w.pkgsvc, &createPreservationActionLocalActivityParams{
				WorkflowID: temporalsdk_workflow.GetInfo(ctx).WorkflowExecution.ID,
				Type:       preservationActionType,
				Status:     package_.ActionStatusInProgress,
				StartedAt:  packageStartedAt,
				PackageID:  tinfo.req.PackageID,
			}).Get(ctx, &tinfo.PreservationActionID)
			if err != nil {
				return err
			}
		}
	}

	// Download.
	{
		if tinfo.req.WatcherName != "" && !tinfo.req.IsDir {
			// TODO: even if TempFile is defined, we should confirm that the file is
			// locally available in disk, just in case we're in the context of a
			// session retry where a different worker is doing the work. In that
			// case, the activity would be executed again.
			if tinfo.TempFile == "" {
				activityOpts := withActivityOptsForLongLivedRequest(sessCtx)
				err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.DownloadActivityName, tinfo.req.WatcherName, tinfo.req.Key).Get(activityOpts, &tinfo.TempFile)
				if err != nil {
					return err
				}
			}
		}
	}

	// SFA-Preprocessing activities only are meant to be used with Archivematica.
	if w.useArchivematica {
		preProcCtx := temporalsdk_workflow.WithActivityOptions(sessCtx, temporalsdk_workflow.ActivityOptions{
			StartToCloseTimeout: time.Second * 5,
			RetryPolicy: &temporalsdk_temporal.RetryPolicy{
				InitialInterval:    time.Second,
				BackoffCoefficient: 2,
				MaximumInterval:    time.Second * 10,
				MaximumAttempts:    2,
				NonRetryableErrorTypes: []string{
					"TemporalTimeout:StartToClose",
				},
			},
		})

		// Extract package.
		var extractPackageRes sfa_activities.ExtractPackageResult
		err := workflow.ExecuteActivity(preProcCtx, sfa_activities.ExtractPackageName, &sfa_activities.ExtractPackageParams{
			Path: tinfo.TempFile,
			Key:  tinfo.req.Key,
		}).Get(sessCtx, &extractPackageRes)
		if err != nil {
			return err
		}
		w.cleanUpPath(extractPackageRes.Path)

		PreProcessingErr := func() error {
			// Validate SIP structure.
			var checkStructureRes sfa_activities.CheckSipStructureResult
			err = workflow.ExecuteActivity(preProcCtx, sfa_activities.CheckSipStructureName, &sfa_activities.CheckSipStructureParams{SipPath: extractPackageRes.Path}).Get(sessCtx, &checkStructureRes)
			if err != nil {
				return err
			}

			var allowedFileFormats sfa_activities.AllowedFileFormatsResult
			err = workflow.ExecuteActivity(preProcCtx, sfa_activities.AllowedFileFormatsName, &sfa_activities.AllowedFileFormatsParams{SipPath: extractPackageRes.Path}).Get(sessCtx, &allowedFileFormats)
			if err != nil {
				return err
			}

			// Validate metadata.xsd.
			var metadataValidation sfa_activities.MetadataValidationResult
			err = workflow.ExecuteActivity(preProcCtx, sfa_activities.MetadataValidationName, &sfa_activities.MetadataValidationParams{SipPath: extractPackageRes.Path}).Get(sessCtx, &metadataValidation)
			if err != nil {
				return err
			}

			// Repackage SFA Sip into a Bag.
			var sipCreation sfa_activities.SipCreationResult
			err = workflow.ExecuteActivity(preProcCtx, sfa_activities.SipCreationName, &sfa_activities.SipCreationParams{SipPath: extractPackageRes.Path}).Get(sessCtx, &sipCreation)
			if err != nil {
				return err
			}
			w.cleanUpPath(sipCreation.NewSipPath)

			// We do this so that the code above only stops when a non-bussines error is found.
			if !allowedFileFormats.Ok {
				return sfa_activities.ErrIlegalFileFormat
			}
			if !checkStructureRes.Ok {
				return sfa_activities.ErrInvaliSipStructure
			}
			tinfo.TempFile = sipCreation.NewSipPath
			tinfo.req.IsDir = true

			return nil
		}()
		if PreProcessingErr != nil {
			var sendToFailedRes sfa_activities.SendToFailedBucketResult
			err = workflow.ExecuteActivity(preProcCtx, sfa_activities.SendToFailedBucketName, &sfa_activities.SendToFailedBucketParams{
				FailureType: sfa_activities.FailureTransfer,
				Path:        tinfo.TempFile,
				Key:         tinfo.req.Key,
			}).Get(sessCtx, &sendToFailedRes)
			if err != nil {
				return err
			}
			return PreProcessingErr
		}
	}

	{
		var err error
		if w.useArchivematica {
			err = w.transferAM(sessCtx, tinfo)
		} else {
			err = w.transferA3m(sessCtx, tinfo)
		}
		if err != nil {
			return err
		}
	}

	// Persist SIPID.
	{
		activityOpts := withLocalActivityOpts(sessCtx)
		_ = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, updatePackageLocalActivity, w.logger, w.pkgsvc, &updatePackageLocalActivityParams{
			PackageID: tinfo.req.PackageID,
			Key:       tinfo.req.Key,
			SIPID:     tinfo.SIPID,
			StoredAt:  tinfo.StoredAt,
			Status:    package_.StatusInProgress,
		}).Get(activityOpts, nil)
	}

	// For the Archivematica workflow AIP creation, review, and storage will be
	// handled entirely by Archivematica and the AM Storage Service.
	if w.useArchivematica {
		return nil
	}

	// Identifier of the preservation task for upload to sips bucket.
	var uploadPreservationTaskID uint

	// Add preservation task for upload to review bucket.
	if !tinfo.req.AutoApproveAIP {
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, createPreservationTaskLocalActivity, w.pkgsvc, &createPreservationTaskLocalActivityParams{
			TaskID:               uuid.NewString(),
			Name:                 "Move AIP",
			Status:               package_.TaskStatusInProgress,
			StartedAt:            temporalsdk_workflow.Now(sessCtx).UTC(),
			Note:                 "Moving to review bucket",
			PreservationActionID: tinfo.PreservationActionID,
		}).Get(ctx, &uploadPreservationTaskID)
		if err != nil {
			return err
		}
	}

	// Upload AIP to MinIO.
	{
		activityOpts := temporalsdk_workflow.WithActivityOptions(sessCtx, temporalsdk_workflow.ActivityOptions{
			StartToCloseTimeout: time.Hour * 24,
			RetryPolicy: &temporalsdk_temporal.RetryPolicy{
				InitialInterval:    time.Second,
				BackoffCoefficient: 2,
				MaximumAttempts:    3,
			},
		})
		err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.UploadActivityName, &activities.UploadActivityParams{
			AIPPath: tinfo.AIPPath,
			AIPID:   tinfo.SIPID,
			Name:    tinfo.req.Key,
		}).Get(activityOpts, nil)
		if err != nil {
			return err
		}
	}

	// Complete preservation task for upload to review bucket.
	if !tinfo.req.AutoApproveAIP {
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, completePreservationTaskLocalActivity, w.pkgsvc, &completePreservationTaskLocalActivityParams{
			ID:          uploadPreservationTaskID,
			Status:      package_.TaskStatusDone,
			CompletedAt: temporalsdk_workflow.Now(sessCtx).UTC(),
			Note:        ref.New("Moved to review bucket"),
		}).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	var reviewResult *package_.ReviewPerformedSignal

	// Identifier of the preservation task for package review
	var reviewPreservationTaskID uint

	if tinfo.req.AutoApproveAIP {
		reviewResult = &package_.ReviewPerformedSignal{
			Accepted:   true,
			LocationID: tinfo.req.DefaultPermanentLocationID,
		}
	} else {
		// Set package to pending status.
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setStatusLocalActivity, w.pkgsvc, tinfo.req.PackageID, package_.StatusPending).Get(ctx, nil)
			if err != nil {
				return err
			}
		}

		// Set preservation action to pending status.
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setPreservationActionStatusLocalActivity, w.pkgsvc, tinfo.PreservationActionID, package_.ActionStatusPending).Get(ctx, nil)
			if err != nil {
				return err
			}
		}

		// Add preservation task for package review
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, createPreservationTaskLocalActivity, w.pkgsvc, &createPreservationTaskLocalActivityParams{
				TaskID:               uuid.NewString(),
				Name:                 "Review AIP",
				Status:               package_.TaskStatusPending,
				StartedAt:            temporalsdk_workflow.Now(sessCtx).UTC(),
				Note:                 "Awaiting user decision",
				PreservationActionID: tinfo.PreservationActionID,
			}).Get(ctx, &reviewPreservationTaskID)
			if err != nil {
				return err
			}
		}

		reviewResult = w.waitForReview(sessCtx)

		// Set package to in progress status.
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setStatusLocalActivity, w.pkgsvc, tinfo.req.PackageID, package_.StatusInProgress).Get(ctx, nil)
			if err != nil {
				return err
			}
		}

		// Set preservation action to in progress status.
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setPreservationActionStatusLocalActivity, w.pkgsvc, tinfo.PreservationActionID, package_.ActionStatusInProgress).Get(ctx, nil)
			if err != nil {
				return err
			}
		}
	}

	reviewCompletedAt := temporalsdk_workflow.Now(sessCtx).UTC()

	if reviewResult.Accepted {
		// Record package confirmation in review preservation task
		if !tinfo.req.AutoApproveAIP {
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, completePreservationTaskLocalActivity, w.pkgsvc, &completePreservationTaskLocalActivityParams{
				ID:          reviewPreservationTaskID,
				Status:      package_.TaskStatusDone,
				CompletedAt: reviewCompletedAt,
				Note:        ref.New("Reviewed and accepted"),
			}).Get(ctx, nil)
			if err != nil {
				return err
			}
		}

		// Identifier of the preservation task for permanent storage move.
		var movePreservationTaskID uint

		// Add preservation task for permanent storage move.
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, createPreservationTaskLocalActivity, w.pkgsvc, &createPreservationTaskLocalActivityParams{
				TaskID:               uuid.NewString(),
				Name:                 "Move AIP",
				Status:               package_.TaskStatusInProgress,
				StartedAt:            temporalsdk_workflow.Now(sessCtx).UTC(),
				Note:                 "Moving to permanent storage",
				PreservationActionID: tinfo.PreservationActionID,
			}).Get(ctx, &movePreservationTaskID)
			if err != nil {
				return err
			}
		}

		// Move package to permanent storage
		{
			activityOpts := withActivityOptsForRequest(sessCtx)
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.MoveToPermanentStorageActivityName, &activities.MoveToPermanentStorageActivityParams{
				AIPID:      tinfo.SIPID,
				LocationID: *reviewResult.LocationID,
			}).Get(activityOpts, nil)
			if err != nil {
				return err
			}
		}

		// Poll package move to permanent storage
		{
			activityOpts := withActivityOptsForLongLivedRequest(sessCtx)
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.PollMoveToPermanentStorageActivityName, &activities.PollMoveToPermanentStorageActivityParams{
				AIPID: tinfo.SIPID,
			}).Get(activityOpts, nil)
			if err != nil {
				return err
			}
		}

		// Complete preservation task for permanent storage move.
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, completePreservationTaskLocalActivity, w.pkgsvc, &completePreservationTaskLocalActivityParams{
				ID:          movePreservationTaskID,
				Status:      package_.TaskStatusDone,
				CompletedAt: temporalsdk_workflow.Now(sessCtx).UTC(),
				Note:        ref.New(fmt.Sprintf("Moved to location %s", *reviewResult.LocationID)),
			}).Get(ctx, nil)
			if err != nil {
				return err
			}
		}

		// Set package location
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setLocationIDLocalActivity, w.pkgsvc, tinfo.req.PackageID, *reviewResult.LocationID).Get(ctx, nil)
			if err != nil {
				return err
			}
		}
	} else if !tinfo.req.AutoApproveAIP {
		// Record package rejection in review preservation task
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, completePreservationTaskLocalActivity, w.pkgsvc, &completePreservationTaskLocalActivityParams{
				ID:          reviewPreservationTaskID,
				Status:      package_.TaskStatusDone,
				CompletedAt: reviewCompletedAt,
				Note:        ref.New("Reviewed and rejected"),
			}).Get(ctx, nil)
			if err != nil {
				return err
			}
		}

		// Reject package
		{
			activityOpts := withActivityOptsForRequest(sessCtx)
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.RejectPackageActivityName, &activities.RejectPackageActivityParams{
				AIPID: tinfo.SIPID,
			}).Get(activityOpts, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *ProcessingWorkflow) waitForReview(ctx temporalsdk_workflow.Context) *package_.ReviewPerformedSignal {
	var review package_.ReviewPerformedSignal
	signalChan := temporalsdk_workflow.GetSignalChannel(ctx, package_.ReviewPerformedSignalName)
	selector := temporalsdk_workflow.NewSelector(ctx)
	selector.AddReceive(signalChan, func(channel temporalsdk_workflow.ReceiveChannel, more bool) {
		_ = channel.Receive(ctx, &review)
	})
	selector.Select(ctx)
	return &review
}

func (w *ProcessingWorkflow) transferA3m(sessCtx temporalsdk_workflow.Context, tinfo *TransferInfo) error {
	// Bundle.
	{
		if tinfo.Bundle == (activities.BundleActivityResult{}) {
			activityOpts := withActivityOptsForLongLivedRequest(sessCtx)
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.BundleActivityName, &activities.BundleActivityParams{
				WatcherName:      tinfo.req.WatcherName,
				TransferDir:      "/home/a3m/.local/share/a3m/share",
				Key:              tinfo.req.Key,
				IsDir:            tinfo.req.IsDir,
				TempFile:         tinfo.TempFile,
				StripTopLevelDir: tinfo.req.StripTopLevelDir,
			}).Get(activityOpts, &tinfo.Bundle)
			if err != nil {
				return err
			}
		}
		w.cleanUpPath(tinfo.Bundle.FullPathBeforeStrip)
	}

	// Delete local temporary files.
	defer func() {
		activityOpts := withActivityOptsForRequest(sessCtx)
		_ = temporalsdk_workflow.ExecuteActivity(activityOpts, activities.CleanUpActivityName, &activities.CleanUpActivityParams{Paths: w.cleanUpPaths}).Get(activityOpts, nil)
	}()

	activityOpts := temporalsdk_workflow.WithActivityOptions(sessCtx, temporalsdk_workflow.ActivityOptions{
		StartToCloseTimeout: time.Hour * 24,
		HeartbeatTimeout:    time.Second * 5,
		RetryPolicy: &temporalsdk_temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	})

	params := &a3m.CreateAIPActivityParams{
		Name:                 tinfo.Name(),
		Path:                 tinfo.Bundle.FullPath,
		PreservationActionID: tinfo.PreservationActionID,
	}

	result := a3m.CreateAIPActivityResult{}
	err := temporalsdk_workflow.ExecuteActivity(activityOpts, a3m.CreateAIPActivityName, params).Get(sessCtx, &result)

	tinfo.SIPID = result.UUID
	tinfo.AIPPath = result.Path
	tinfo.StoredAt = temporalsdk_workflow.Now(sessCtx).UTC()

	return err
}

func (w *ProcessingWorkflow) transferAM(sessCtx temporalsdk_workflow.Context, tinfo *TransferInfo) error {
	var err error

	activityOpts := withActivityOptsForLongLivedRequest(sessCtx)

	// Zip transfer.
	var zip activities.ZipActivityResult
	err = temporalsdk_workflow.ExecuteActivity(
		activityOpts,
		activities.ZipActivityName,
		&activities.ZipActivityParams{SourceDir: tinfo.TempFile},
	).Get(activityOpts, &zip)
	if err != nil {
		return err
	}
	w.cleanUpPath(zip.Path) // Delete when workflow completes.
	zipPath := zip.Path

	defer func() {
		if err != nil {
			var sendToFailedRes sfa_activities.SendToFailedBucketResult
			bucketErr := workflow.ExecuteActivity(activityOpts, sfa_activities.SendToFailedBucketName, &sfa_activities.SendToFailedBucketParams{
				FailureType: sfa_activities.FailureSIP,
				Path:        zipPath,
				Key:         tinfo.req.Key,
			}).Get(sessCtx, &sendToFailedRes)
			errors.Join(err, bucketErr)
		}
	}()

	uploadResult := am.UploadTransferActivityResult{}
	err = temporalsdk_workflow.ExecuteActivity(
		activityOpts,
		am.UploadTransferActivityName,
		&am.UploadTransferActivityParams{SourcePath: zip.Path},
	).Get(activityOpts, &uploadResult)
	if err != nil {
		return err
	}

	activityOpts = temporalsdk_workflow.WithActivityOptions(sessCtx, temporalsdk_workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy: &temporalsdk_temporal.RetryPolicy{
			InitialInterval:    time.Second * 10,
			BackoffCoefficient: 1.0,
			MaximumAttempts:    3,
		},
	})

	result := am.StartTransferActivityResult{}
	err = temporalsdk_workflow.ExecuteActivity(
		activityOpts,
		am.StartTransferActivityName,
		&am.StartTransferActivityParams{
			Name: tinfo.Name(),
			Path: uploadResult.RemotePath,
		},
	).Get(activityOpts, &result)
	if err != nil {
		return err
	}

	var pollResult am.PollTransferActivityResult
	pollOpts := temporalsdk_workflow.WithActivityOptions(sessCtx, temporalsdk_workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy: &temporalsdk_temporal.RetryPolicy{
			InitialInterval:    time.Second * 10,
			BackoffCoefficient: 1.0,
		},
	})
	err = temporalsdk_workflow.ExecuteActivity(pollOpts, am.PollTransferActivityName, am.PollTransferActivityParams{
		TransferID: result.UUID,
	}).Get(pollOpts, &pollResult)
	if err != nil {
		return err
	}

	tinfo.SIPID = pollResult.SIPID
	tinfo.AIPPath = pollResult.Path
	tinfo.StoredAt = temporalsdk_workflow.Now(sessCtx).UTC()

	return nil
}

func (w *ProcessingWorkflow) cleanUpPath(path string) {
	if path == "" {
		return
	}

	w.cleanUpPaths = append(w.cleanUpPaths, path)
}
