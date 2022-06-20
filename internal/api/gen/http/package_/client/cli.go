// Code generated by goa v3.7.6, DO NOT EDIT.
//
// package HTTP client CLI support package
//
// Command:
// $ goa-v3.7.6 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	package_ "github.com/artefactual-labs/enduro/internal/api/gen/package_"
	goa "goa.design/goa/v3/pkg"
)

// BuildListPayload builds the payload for the package list endpoint from CLI
// flags.
func BuildListPayload(package_ListName string, package_ListAipID string, package_ListEarliestCreatedTime string, package_ListLatestCreatedTime string, package_ListStatus string, package_ListCursor string) (*package_.ListPayload, error) {
	var err error
	var name *string
	{
		if package_ListName != "" {
			name = &package_ListName
		}
	}
	var aipID *string
	{
		if package_ListAipID != "" {
			aipID = &package_ListAipID
			err = goa.MergeErrors(err, goa.ValidateFormat("aipID", *aipID, goa.FormatUUID))
			if err != nil {
				return nil, err
			}
		}
	}
	var earliestCreatedTime *string
	{
		if package_ListEarliestCreatedTime != "" {
			earliestCreatedTime = &package_ListEarliestCreatedTime
			err = goa.MergeErrors(err, goa.ValidateFormat("earliestCreatedTime", *earliestCreatedTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var latestCreatedTime *string
	{
		if package_ListLatestCreatedTime != "" {
			latestCreatedTime = &package_ListLatestCreatedTime
			err = goa.MergeErrors(err, goa.ValidateFormat("latestCreatedTime", *latestCreatedTime, goa.FormatDateTime))
			if err != nil {
				return nil, err
			}
		}
	}
	var status *string
	{
		if package_ListStatus != "" {
			status = &package_ListStatus
			if !(*status == "new" || *status == "in progress" || *status == "done" || *status == "error" || *status == "unknown" || *status == "queued" || *status == "pending" || *status == "abandoned") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("status", *status, []interface{}{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var cursor *string
	{
		if package_ListCursor != "" {
			cursor = &package_ListCursor
		}
	}
	v := &package_.ListPayload{}
	v.Name = name
	v.AipID = aipID
	v.EarliestCreatedTime = earliestCreatedTime
	v.LatestCreatedTime = latestCreatedTime
	v.Status = status
	v.Cursor = cursor

	return v, nil
}

// BuildShowPayload builds the payload for the package show endpoint from CLI
// flags.
func BuildShowPayload(package_ShowID string) (*package_.ShowPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_ShowID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &package_.ShowPayload{}
	v.ID = id

	return v, nil
}

// BuildDeletePayload builds the payload for the package delete endpoint from
// CLI flags.
func BuildDeletePayload(package_DeleteID string) (*package_.DeletePayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_DeleteID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &package_.DeletePayload{}
	v.ID = id

	return v, nil
}

// BuildCancelPayload builds the payload for the package cancel endpoint from
// CLI flags.
func BuildCancelPayload(package_CancelID string) (*package_.CancelPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_CancelID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &package_.CancelPayload{}
	v.ID = id

	return v, nil
}

// BuildRetryPayload builds the payload for the package retry endpoint from CLI
// flags.
func BuildRetryPayload(package_RetryID string) (*package_.RetryPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_RetryID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &package_.RetryPayload{}
	v.ID = id

	return v, nil
}

// BuildWorkflowPayload builds the payload for the package workflow endpoint
// from CLI flags.
func BuildWorkflowPayload(package_WorkflowID string) (*package_.WorkflowPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_WorkflowID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &package_.WorkflowPayload{}
	v.ID = id

	return v, nil
}

// BuildDownloadPayload builds the payload for the package download endpoint
// from CLI flags.
func BuildDownloadPayload(package_DownloadID string) (*package_.DownloadPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_DownloadID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &package_.DownloadPayload{}
	v.ID = id

	return v, nil
}

// BuildBulkPayload builds the payload for the package bulk endpoint from CLI
// flags.
func BuildBulkPayload(package_BulkBody string) (*package_.BulkPayload, error) {
	var err error
	var body BulkRequestBody
	{
		err = json.Unmarshal([]byte(package_BulkBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"operation\": \"retry\",\n      \"size\": 11073631979459909334,\n      \"status\": \"abandoned\"\n   }'")
		}
		if !(body.Operation == "retry" || body.Operation == "cancel" || body.Operation == "abandon") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.operation", body.Operation, []interface{}{"retry", "cancel", "abandon"}))
		}
		if !(body.Status == "new" || body.Status == "in progress" || body.Status == "done" || body.Status == "error" || body.Status == "unknown" || body.Status == "queued" || body.Status == "pending" || body.Status == "abandoned") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", body.Status, []interface{}{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &package_.BulkPayload{
		Operation: body.Operation,
		Status:    body.Status,
		Size:      body.Size,
	}
	{
		var zero uint
		if v.Size == zero {
			v.Size = 100
		}
	}

	return v, nil
}

// BuildPreservationActionsPayload builds the payload for the package
// preservation-actions endpoint from CLI flags.
func BuildPreservationActionsPayload(package_PreservationActionsID string) (*package_.PreservationActionsPayload, error) {
	var err error
	var id uint
	{
		var v uint64
		v, err = strconv.ParseUint(package_PreservationActionsID, 10, strconv.IntSize)
		id = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT")
		}
	}
	v := &package_.PreservationActionsPayload{}
	v.ID = id

	return v, nil
}
