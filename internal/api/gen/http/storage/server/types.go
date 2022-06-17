// Code generated by goa v3.7.6, DO NOT EDIT.
//
// storage HTTP server types
//
// Command:
// $ goa-v3.7.6 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package server

import (
	storage "github.com/artefactual-labs/enduro/internal/api/gen/storage"
	goa "goa.design/goa/v3/pkg"
)

// UpdateRequestBody is the type of the "storage" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
}

// SubmitResponseBody is the type of the "storage" service "submit" endpoint
// HTTP response body.
type SubmitResponseBody struct {
	URL        string `form:"url" json:"url" xml:"url"`
	WorkflowID string `form:"workflow_id" json:"workflow_id" xml:"workflow_id"`
}

// UpdateResponseBody is the type of the "storage" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	OK bool `form:"ok" json:"ok" xml:"ok"`
}

// SubmitNotAvailableResponseBody is the type of the "storage" service "submit"
// endpoint HTTP response body for the "not_available" error.
type SubmitNotAvailableResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// SubmitNotValidResponseBody is the type of the "storage" service "submit"
// endpoint HTTP response body for the "not_valid" error.
type SubmitNotValidResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateNotAvailableResponseBody is the type of the "storage" service "update"
// endpoint HTTP response body for the "not_available" error.
type UpdateNotAvailableResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateNotValidResponseBody is the type of the "storage" service "update"
// endpoint HTTP response body for the "not_valid" error.
type UpdateNotValidResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewSubmitResponseBody builds the HTTP response body from the result of the
// "submit" endpoint of the "storage" service.
func NewSubmitResponseBody(res *storage.SubmitResult) *SubmitResponseBody {
	body := &SubmitResponseBody{
		URL:        res.URL,
		WorkflowID: res.WorkflowID,
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "storage" service.
func NewUpdateResponseBody(res *storage.UpdateResult) *UpdateResponseBody {
	body := &UpdateResponseBody{
		OK: res.OK,
	}
	return body
}

// NewSubmitNotAvailableResponseBody builds the HTTP response body from the
// result of the "submit" endpoint of the "storage" service.
func NewSubmitNotAvailableResponseBody(res *goa.ServiceError) *SubmitNotAvailableResponseBody {
	body := &SubmitNotAvailableResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewSubmitNotValidResponseBody builds the HTTP response body from the result
// of the "submit" endpoint of the "storage" service.
func NewSubmitNotValidResponseBody(res *goa.ServiceError) *SubmitNotValidResponseBody {
	body := &SubmitNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateNotAvailableResponseBody builds the HTTP response body from the
// result of the "update" endpoint of the "storage" service.
func NewUpdateNotAvailableResponseBody(res *goa.ServiceError) *UpdateNotAvailableResponseBody {
	body := &UpdateNotAvailableResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateNotValidResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "storage" service.
func NewUpdateNotValidResponseBody(res *goa.ServiceError) *UpdateNotValidResponseBody {
	body := &UpdateNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdatePayload builds a storage service update endpoint payload.
func NewUpdatePayload(body *UpdateRequestBody) *storage.UpdatePayload {
	v := &storage.UpdatePayload{
		WorkflowID: *body.WorkflowID,
	}

	return v
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.WorkflowID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("workflow_id", "body"))
	}
	return
}