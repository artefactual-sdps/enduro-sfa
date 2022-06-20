// Code generated by goa v3.7.6, DO NOT EDIT.
//
// storage HTTP client types
//
// Command:
// $ goa-v3.7.6 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	storage "github.com/artefactual-labs/enduro/internal/api/gen/storage"
	goa "goa.design/goa/v3/pkg"
)

// SubmitRequestBody is the type of the "storage" service "submit" endpoint
// HTTP request body.
type SubmitRequestBody struct {
	AipID string `form:"aip_id" json:"aip_id" xml:"aip_id"`
	Name  string `form:"name" json:"name" xml:"name"`
}

// UpdateRequestBody is the type of the "storage" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	AipID string `form:"aip_id" json:"aip_id" xml:"aip_id"`
}

// SubmitResponseBody is the type of the "storage" service "submit" endpoint
// HTTP response body.
type SubmitResponseBody struct {
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// UpdateResponseBody is the type of the "storage" service "update" endpoint
// HTTP response body.
type UpdateResponseBody struct {
	OK *bool `form:"ok,omitempty" json:"ok,omitempty" xml:"ok,omitempty"`
}

// SubmitNotAvailableResponseBody is the type of the "storage" service "submit"
// endpoint HTTP response body for the "not_available" error.
type SubmitNotAvailableResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// SubmitNotValidResponseBody is the type of the "storage" service "submit"
// endpoint HTTP response body for the "not_valid" error.
type SubmitNotValidResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateNotAvailableResponseBody is the type of the "storage" service "update"
// endpoint HTTP response body for the "not_available" error.
type UpdateNotAvailableResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// UpdateNotValidResponseBody is the type of the "storage" service "update"
// endpoint HTTP response body for the "not_valid" error.
type UpdateNotValidResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewSubmitRequestBody builds the HTTP request body from the payload of the
// "submit" endpoint of the "storage" service.
func NewSubmitRequestBody(p *storage.SubmitPayload) *SubmitRequestBody {
	body := &SubmitRequestBody{
		AipID: p.AipID,
		Name:  p.Name,
	}
	return body
}

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "storage" service.
func NewUpdateRequestBody(p *storage.UpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		AipID: p.AipID,
	}
	return body
}

// NewSubmitResultAccepted builds a "storage" service "submit" endpoint result
// from a HTTP "Accepted" response.
func NewSubmitResultAccepted(body *SubmitResponseBody) *storage.SubmitResult {
	v := &storage.SubmitResult{
		URL: *body.URL,
	}

	return v
}

// NewSubmitNotAvailable builds a storage service submit endpoint not_available
// error.
func NewSubmitNotAvailable(body *SubmitNotAvailableResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewSubmitNotValid builds a storage service submit endpoint not_valid error.
func NewSubmitNotValid(body *SubmitNotValidResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateResultAccepted builds a "storage" service "update" endpoint result
// from a HTTP "Accepted" response.
func NewUpdateResultAccepted(body *UpdateResponseBody) *storage.UpdateResult {
	v := &storage.UpdateResult{
		OK: *body.OK,
	}

	return v
}

// NewUpdateNotAvailable builds a storage service update endpoint not_available
// error.
func NewUpdateNotAvailable(body *UpdateNotAvailableResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewUpdateNotValid builds a storage service update endpoint not_valid error.
func NewUpdateNotValid(body *UpdateNotValidResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateSubmitResponseBody runs the validations defined on SubmitResponseBody
func ValidateSubmitResponseBody(body *SubmitResponseBody) (err error) {
	if body.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "body"))
	}
	return
}

// ValidateUpdateResponseBody runs the validations defined on UpdateResponseBody
func ValidateUpdateResponseBody(body *UpdateResponseBody) (err error) {
	if body.OK == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ok", "body"))
	}
	return
}

// ValidateSubmitNotAvailableResponseBody runs the validations defined on
// submit_not_available_response_body
func ValidateSubmitNotAvailableResponseBody(body *SubmitNotAvailableResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateSubmitNotValidResponseBody runs the validations defined on
// submit_not_valid_response_body
func ValidateSubmitNotValidResponseBody(body *SubmitNotValidResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateNotAvailableResponseBody runs the validations defined on
// update_not_available_response_body
func ValidateUpdateNotAvailableResponseBody(body *UpdateNotAvailableResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateUpdateNotValidResponseBody runs the validations defined on
// update_not_valid_response_body
func ValidateUpdateNotValidResponseBody(body *UpdateNotValidResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
