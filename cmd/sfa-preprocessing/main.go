package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"

	"github.com/artefactual-sdps/enduro/internal/sfa/activities"
	"github.com/artefactual-sdps/enduro/internal/sfa/workflows"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	taskQueue := "sfa-preprocessing"
	temporalClient, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "default",
	})
	if err != nil {
		log.Fatal(err.Error() + "Failed to reach temporal client")
	}
	w := worker.New(temporalClient, taskQueue, worker.Options{})

	// Register workflows.
	w.RegisterWorkflowWithOptions(
		workflows.NewSFAWorkflow().Execute,
		workflow.RegisterOptions{Name: workflows.SFAWorkflowName},
	)

	// Register activities.
	w.RegisterActivityWithOptions(
		activities.NewCheckSipStructure().Execute,
		activity.RegisterOptions{Name: activities.CheckSipStructureName},
	)

	w.RegisterActivityWithOptions(
		activities.NewAllowedFileFormatsActivity().Execute,
		activity.RegisterOptions{Name: activities.AllowedFileFormatsName},
	)

	w.RegisterActivityWithOptions(
		activities.NewMetadataValidationActivity().Execute,
		activity.RegisterOptions{Name: activities.MetadataValidationActivityName},
	)

	w.RegisterActivityWithOptions(
		activities.NewSipCreationActivity().Execute,
		activity.RegisterOptions{Name: activities.SipCreationActivityName},
	)

	{
		workflowParams := &workflows.SFAWorkflowParams{SipDir: "./testdata/SIP_20111020_BFB_v60"}
		workflowOpts := client.StartWorkflowOptions{TaskQueue: taskQueue}
		wf, err := temporalClient.ExecuteWorkflow(ctx, workflowOpts, workflows.SFAWorkflowName, workflowParams)
		if err != nil {
			log.Fatal(err, "Failed to start workflow: ", workflows.SFAWorkflowName)
		}
		log.Print("Workflow started", "Workflow ID", wf.GetID())
	}

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatal(err, "Failed to start worker")
	}
	log.Print("Worker running")
}
