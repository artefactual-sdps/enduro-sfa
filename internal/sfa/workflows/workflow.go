package workflows

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"github.com/artefactual-sdps/enduro/internal/sfa/activities"
)

const SFAWorkflowName = "sfa-preprocessing"

type SFAWorkflow struct{}

func NewSFAWorkflow() *SFAWorkflow {
	return &SFAWorkflow{}
}

type SFAWorkflowParams struct {
	SipDir string
}

type SFAWorkflowResult struct{}

func (w *SFAWorkflow) Execute(ctx workflow.Context, params *SFAWorkflowParams) (*SFAWorkflowResult, error) {
	var err error
	res := &SFAWorkflowResult{}
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 60,
		HeartbeatTimeout:    time.Second * 30,
		WaitForCancellation: true,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	})

	var checkStructureRes activities.CheckSipStructureResult
	err = workflow.ExecuteActivity(ctx, activities.CheckSipStructureName, &activities.CheckSipStructureParams{
		SipPath: params.SipDir,
	}).Get(ctx, &checkStructureRes)
	if err != nil {
		return nil, err
	}

	var allowedFileFormats activities.AllowedFileFormatsResult
	err = workflow.ExecuteActivity(ctx, activities.AllowedFileFormatsName, &activities.AllowedFileFormatsParams{
		SipPath: params.SipDir,
	}).Get(ctx, &allowedFileFormats)
	if err != nil {
		return nil, err
	}

	var metadataValidation activities.MetadataValidationResult
	err = workflow.ExecuteActivity(ctx, activities.MetadataValidationActivityName, &activities.MetadataValidationParams{
		SipPath: params.SipDir,
	}).Get(ctx, &metadataValidation)
	if err != nil {
		return nil, err
	}

	var sipCreation activities.SipCreationResult
	err = workflow.ExecuteActivity(ctx, activities.SipCreationActivityName, &activities.SipCreationParams{
		SipPath: params.SipDir,
	}).Get(ctx, &sipCreation)
	if err != nil {
		return nil, err
	}

	// I do this for now so that the code above only stops when a non-bussiness error is found.
	if !allowedFileFormats.Ok {
		return nil, activities.ErrIlegalFileFormat
	}
	if !checkStructureRes.Ok {
		return nil, activities.ErrInvaliSipStructure
	}
	return res, nil
}
