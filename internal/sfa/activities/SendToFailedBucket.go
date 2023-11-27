package activities

import (
	"context"
	"os"

	"gocloud.dev/blob"
)

const SendToFailedBucketName = "send-to-failed-bucket"

type SendToFailedBucketActivity struct {
	failedBucket *blob.Bucket
}

func NewSendToFailedBuckeActivity(bucket *blob.Bucket) *SendToFailedBucketActivity {
	return &SendToFailedBucketActivity{failedBucket: bucket}
}

type SendToFailedBucketParams struct {
	Path string
	Key  string
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

	if err := sf.failedBucket.Upload(ctx, res.FailedKey, f, &blob.WriterOptions{
		ContentType: "application/octet-stream",
	}); err != nil {
		return nil, err
	}

	return res, nil
}
