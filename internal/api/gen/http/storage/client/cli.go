// Code generated by goa v3.8.3, DO NOT EDIT.
//
// storage HTTP client CLI support package
//
// Command:
// $ goa-v3.8.3 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package client

import (
	"encoding/json"
	"fmt"

	storage "github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	goa "goa.design/goa/v3/pkg"
)

// BuildSubmitPayload builds the payload for the storage submit endpoint from
// CLI flags.
func BuildSubmitPayload(storageSubmitBody string, storageSubmitAipID string) (*storage.SubmitPayload, error) {
	var err error
	var body SubmitRequestBody
	{
		err = json.Unmarshal([]byte(storageSubmitBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Dolorem non voluptate voluptatem voluptate maxime facere.\"\n   }'")
		}
	}
	var aipID string
	{
		aipID = storageSubmitAipID
	}
	v := &storage.SubmitPayload{
		Name: body.Name,
	}
	v.AipID = aipID

	return v, nil
}

// BuildUpdatePayload builds the payload for the storage update endpoint from
// CLI flags.
func BuildUpdatePayload(storageUpdateAipID string) (*storage.UpdatePayload, error) {
	var aipID string
	{
		aipID = storageUpdateAipID
	}
	v := &storage.UpdatePayload{}
	v.AipID = aipID

	return v, nil
}

// BuildDownloadPayload builds the payload for the storage download endpoint
// from CLI flags.
func BuildDownloadPayload(storageDownloadAipID string) (*storage.DownloadPayload, error) {
	var aipID string
	{
		aipID = storageDownloadAipID
	}
	v := &storage.DownloadPayload{}
	v.AipID = aipID

	return v, nil
}

// BuildAddLocationPayload builds the payload for the storage add_location
// endpoint from CLI flags.
func BuildAddLocationPayload(storageAddLocationBody string) (*storage.AddLocationPayload, error) {
	var err error
	var body AddLocationRequestBody
	{
		err = json.Unmarshal([]byte(storageAddLocationBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"config\": {\n         \"Type\": \"s3\",\n         \"Value\": \"\\\"JSON\\\"\"\n      },\n      \"description\": \"Accusantium vitae.\",\n      \"name\": \"Sit ut quam praesentium odio assumenda fuga.\",\n      \"purpose\": \"aip_store\",\n      \"source\": \"unspecified\"\n   }'")
		}
		if !(body.Source == "unspecified" || body.Source == "minio") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.source", body.Source, []interface{}{"unspecified", "minio"}))
		}
		if !(body.Purpose == "unspecified" || body.Purpose == "aip_store") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.purpose", body.Purpose, []interface{}{"unspecified", "aip_store"}))
		}
		if body.Config != nil {
			if !(body.Config.Type == "s3") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.config.Type", body.Config.Type, []interface{}{"s3"}))
			}
		}
		if err != nil {
			return nil, err
		}
	}
	v := &storage.AddLocationPayload{
		Name:        body.Name,
		Description: body.Description,
		Source:      body.Source,
		Purpose:     body.Purpose,
	}
	if body.Config != nil {
		switch body.Config.Type {
		case "s3":
			var val *storage.S3Config
			json.Unmarshal([]byte(body.Config.Value), &val)
			v.Config = val
		}
	}

	return v, nil
}

// BuildMovePayload builds the payload for the storage move endpoint from CLI
// flags.
func BuildMovePayload(storageMoveBody string, storageMoveAipID string) (*storage.MovePayload, error) {
	var err error
	var body MoveRequestBody
	{
		err = json.Unmarshal([]byte(storageMoveBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"location_id\": \"Modi aut blanditiis.\"\n   }'")
		}
	}
	var aipID string
	{
		aipID = storageMoveAipID
	}
	v := &storage.MovePayload{
		LocationID: body.LocationID,
	}
	v.AipID = aipID

	return v, nil
}

// BuildMoveStatusPayload builds the payload for the storage move_status
// endpoint from CLI flags.
func BuildMoveStatusPayload(storageMoveStatusAipID string) (*storage.MoveStatusPayload, error) {
	var aipID string
	{
		aipID = storageMoveStatusAipID
	}
	v := &storage.MoveStatusPayload{}
	v.AipID = aipID

	return v, nil
}

// BuildRejectPayload builds the payload for the storage reject endpoint from
// CLI flags.
func BuildRejectPayload(storageRejectAipID string) (*storage.RejectPayload, error) {
	var aipID string
	{
		aipID = storageRejectAipID
	}
	v := &storage.RejectPayload{}
	v.AipID = aipID

	return v, nil
}

// BuildShowPayload builds the payload for the storage show endpoint from CLI
// flags.
func BuildShowPayload(storageShowAipID string) (*storage.ShowPayload, error) {
	var aipID string
	{
		aipID = storageShowAipID
	}
	v := &storage.ShowPayload{}
	v.AipID = aipID

	return v, nil
}

// BuildShowLocationPayload builds the payload for the storage show-location
// endpoint from CLI flags.
func BuildShowLocationPayload(storageShowLocationUUID string) (*storage.ShowLocationPayload, error) {
	var err error
	var uuid string
	{
		uuid = storageShowLocationUUID
		err = goa.MergeErrors(err, goa.ValidateFormat("uuid", uuid, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
	}
	v := &storage.ShowLocationPayload{}
	v.UUID = uuid

	return v, nil
}
