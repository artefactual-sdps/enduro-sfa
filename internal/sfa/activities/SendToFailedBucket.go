package activities

import (
	"context"
	"os"

	"gocloud.dev/blob"
)

const SendToFailedBucketName = "send-to-failed-bucket"
const FailureSIP = "sip"
const FailureTransfer = "transfer"

type SendToFailedBucketActivity struct {
	failedTransferBucket *blob.Bucket
	failedSipBucket      *blob.Bucket
}

func NewSendToFailedBuckeActivity(transfer, sip *blob.Bucket) *SendToFailedBucketActivity {
	return &SendToFailedBucketActivity{
		failedTransferBucket: transfer,
		failedSipBucket:      sip,
	}
}

type SendToFailedBucketParams struct {
	FailureType string
	Path        string
	Key         string
}

type SendToFailedBucketResult struct {
	FailedKey string
}

func (sf *SendToFailedBucketActivity) Execute(ctx context.Context, params *SendToFailedBucketParams) (*SendToFailedBucketResult, error) {
	res := &SendToFailedBucketResult{}
	f, err := os.Open(params.Path)
	if err != nil {
		return nil, err
	}
	res.FailedKey = "Failed_" + params.Key

	switch params.FailureType {
	case FailureTransfer:
		err = sf.failedTransferBucket.Upload(ctx, res.FailedKey, f, &blob.WriterOptions{ContentType: "application/octet-stream"})
	case FailureSIP:
		err = sf.failedSipBucket.Upload(ctx, res.FailedKey, f, &blob.WriterOptions{ContentType: "application/octet-stream"})
	}
	if err != nil {
		return nil, err
	}

	return res, nil
}
