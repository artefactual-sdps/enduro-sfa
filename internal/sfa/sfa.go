package sfa

import (
	"context"
	"fmt"

	temporalsdk_activity "go.temporal.io/sdk/activity"
	temporalsdk_worker "go.temporal.io/sdk/worker"
	"gocloud.dev/blob"

	"github.com/artefactual-sdps/enduro/internal/config"
	"github.com/artefactual-sdps/enduro/internal/sfa/activities"
	"github.com/artefactual-sdps/enduro/internal/storage"
)

func RegisterActivities(ctx context.Context, w temporalsdk_worker.Worker, cfg config.Configuration) error {
	// Set-up failed transfers bucket.
	var ft *blob.Bucket
	{
		fl, err := storage.NewInternalLocation(&cfg.FailedTransfers)
		if err != nil {
			return fmt.Errorf("error setting up failed transfers location: %v", err)
		}
		ft, err = fl.OpenBucket(ctx)
		if err != nil {
			return fmt.Errorf("error getting failed transfers bucket: %v", err)
		}
	}

	// Set-up failed sip bucket.
	var fs *blob.Bucket
	{
		fl, err := storage.NewInternalLocation(&cfg.FailedSips)
		if err != nil {
			return fmt.Errorf("error setting up failed SIPs location: %v", err)
		}
		fs, err = fl.OpenBucket(ctx)
		if err != nil {
			return fmt.Errorf("error getting failed SIPs bucket: %v", err)
		}
	}

	// Register activities.
	w.RegisterActivityWithOptions(activities.NewExtractPackage().Execute, temporalsdk_activity.RegisterOptions{Name: activities.ExtractPackageName})
	w.RegisterActivityWithOptions(activities.NewCheckSipStructure().Execute, temporalsdk_activity.RegisterOptions{Name: activities.CheckSipStructureName})
	w.RegisterActivityWithOptions(activities.NewAllowedFileFormatsActivity().Execute, temporalsdk_activity.RegisterOptions{Name: activities.AllowedFileFormatsName})
	w.RegisterActivityWithOptions(activities.NewMetadataValidationActivity().Execute, temporalsdk_activity.RegisterOptions{Name: activities.MetadataValidationName})
	w.RegisterActivityWithOptions(activities.NewSipCreationActivity().Execute, temporalsdk_activity.RegisterOptions{Name: activities.SipCreationName})
	w.RegisterActivityWithOptions(activities.NewSendToFailedBuckeActivity(ft, fs).Execute, temporalsdk_activity.RegisterOptions{Name: activities.SendToFailedBucketName})

	return nil
}
