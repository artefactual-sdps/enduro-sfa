// Code generated by goa v3.7.10, DO NOT EDIT.
//
// package HTTP server encoders and decoders
//
// Command:
// $ goa-v3.7.10 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"

	package_ "github.com/artefactual-labs/enduro/internal/api/gen/package_"
	package_views "github.com/artefactual-labs/enduro/internal/api/gen/package_/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the package
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*package_.ListResult)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the package list
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			name                *string
			aipID               *string
			earliestCreatedTime *string
			latestCreatedTime   *string
			location            *string
			status              *string
			cursor              *string
			err                 error
		)
		nameRaw := r.URL.Query().Get("name")
		if nameRaw != "" {
			name = &nameRaw
		}
		aipIDRaw := r.URL.Query().Get("aip_id")
		if aipIDRaw != "" {
			aipID = &aipIDRaw
		}
		if aipID != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("aipID", *aipID, goa.FormatUUID))
		}
		earliestCreatedTimeRaw := r.URL.Query().Get("earliest_created_time")
		if earliestCreatedTimeRaw != "" {
			earliestCreatedTime = &earliestCreatedTimeRaw
		}
		if earliestCreatedTime != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("earliestCreatedTime", *earliestCreatedTime, goa.FormatDateTime))
		}
		latestCreatedTimeRaw := r.URL.Query().Get("latest_created_time")
		if latestCreatedTimeRaw != "" {
			latestCreatedTime = &latestCreatedTimeRaw
		}
		if latestCreatedTime != nil {
			err = goa.MergeErrors(err, goa.ValidateFormat("latestCreatedTime", *latestCreatedTime, goa.FormatDateTime))
		}
		locationRaw := r.URL.Query().Get("location")
		if locationRaw != "" {
			location = &locationRaw
		}
		statusRaw := r.URL.Query().Get("status")
		if statusRaw != "" {
			status = &statusRaw
		}
		if status != nil {
			if !(*status == "new" || *status == "in progress" || *status == "done" || *status == "error" || *status == "unknown" || *status == "queued" || *status == "pending" || *status == "abandoned") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("status", *status, []interface{}{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
			}
		}
		cursorRaw := r.URL.Query().Get("cursor")
		if cursorRaw != "" {
			cursor = &cursorRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewListPayload(name, aipID, earliestCreatedTime, latestCreatedTime, location, status, cursor)

		return payload, nil
	}
}

// EncodeShowResponse returns an encoder for responses returned by the package
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*package_views.EnduroStoredPackage)
		enc := encoder(ctx, w)
		body := NewShowResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the package show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(id)

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show package
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			var res *package_.PackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteResponse returns an encoder for responses returned by the
// package delete endpoint.
func EncodeDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteRequest returns a decoder for requests sent to the package
// delete endpoint.
func DecodeDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeletePayload(id)

		return payload, nil
	}
}

// EncodeDeleteError returns an encoder for errors returned by the delete
// package endpoint.
func EncodeDeleteError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			var res *package_.PackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCancelResponse returns an encoder for responses returned by the
// package cancel endpoint.
func EncodeCancelResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeCancelRequest returns a decoder for requests sent to the package
// cancel endpoint.
func DecodeCancelRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewCancelPayload(id)

		return payload, nil
	}
}

// EncodeCancelError returns an encoder for errors returned by the cancel
// package endpoint.
func EncodeCancelError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_running":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCancelNotRunningResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "not_found":
			var res *package_.PackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCancelNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRetryResponse returns an encoder for responses returned by the package
// retry endpoint.
func EncodeRetryResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeRetryRequest returns a decoder for requests sent to the package retry
// endpoint.
func DecodeRetryRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewRetryPayload(id)

		return payload, nil
	}
}

// EncodeRetryError returns an encoder for errors returned by the retry package
// endpoint.
func EncodeRetryError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_running":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRetryNotRunningResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "not_found":
			var res *package_.PackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRetryNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeWorkflowResponse returns an encoder for responses returned by the
// package workflow endpoint.
func EncodeWorkflowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*package_views.EnduroPackageWorkflowStatus)
		enc := encoder(ctx, w)
		body := NewWorkflowResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeWorkflowRequest returns a decoder for requests sent to the package
// workflow endpoint.
func DecodeWorkflowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewWorkflowPayload(id)

		return payload, nil
	}
}

// EncodeWorkflowError returns an encoder for errors returned by the workflow
// package endpoint.
func EncodeWorkflowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			var res *package_.PackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewWorkflowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeBulkResponse returns an encoder for responses returned by the package
// bulk endpoint.
func EncodeBulkResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*package_.BulkResult)
		enc := encoder(ctx, w)
		body := NewBulkResponseBody(res)
		w.WriteHeader(http.StatusAccepted)
		return enc.Encode(body)
	}
}

// DecodeBulkRequest returns a decoder for requests sent to the package bulk
// endpoint.
func DecodeBulkRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body BulkRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateBulkRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewBulkPayload(&body)

		return payload, nil
	}
}

// EncodeBulkError returns an encoder for errors returned by the bulk package
// endpoint.
func EncodeBulkError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_available":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewBulkNotAvailableResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "not_valid":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewBulkNotValidResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeBulkStatusResponse returns an encoder for responses returned by the
// package bulk_status endpoint.
func EncodeBulkStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*package_.BulkStatusResult)
		enc := encoder(ctx, w)
		body := NewBulkStatusResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeConfirmResponse returns an encoder for responses returned by the
// package confirm endpoint.
func EncodeConfirmResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeConfirmRequest returns a decoder for requests sent to the package
// confirm endpoint.
func DecodeConfirmRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body ConfirmRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateConfirmRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id uint

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewConfirmPayload(&body, id)

		return payload, nil
	}
}

// EncodeConfirmError returns an encoder for errors returned by the confirm
// package endpoint.
func EncodeConfirmError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_available":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewConfirmNotAvailableResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "not_valid":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewConfirmNotValidResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "not_found":
			var res *package_.PackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewConfirmNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRejectResponse returns an encoder for responses returned by the
// package reject endpoint.
func EncodeRejectResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeRejectRequest returns a decoder for requests sent to the package
// reject endpoint.
func DecodeRejectRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewRejectPayload(id)

		return payload, nil
	}
}

// EncodeRejectError returns an encoder for errors returned by the reject
// package endpoint.
func EncodeRejectError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_available":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRejectNotAvailableResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "not_valid":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRejectNotValidResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "not_found":
			var res *package_.PackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRejectNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeMoveResponse returns an encoder for responses returned by the package
// move endpoint.
func EncodeMoveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeMoveRequest returns a decoder for requests sent to the package move
// endpoint.
func DecodeMoveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MoveRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMoveRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id uint

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewMovePayload(&body, id)

		return payload, nil
	}
}

// EncodeMoveError returns an encoder for errors returned by the move package
// endpoint.
func EncodeMoveError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_available":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMoveNotAvailableResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "not_valid":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMoveNotValidResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "not_found":
			var res *package_.PackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMoveNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeMoveStatusResponse returns an encoder for responses returned by the
// package move_status endpoint.
func EncodeMoveStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*package_.MoveStatusResult)
		enc := encoder(ctx, w)
		body := NewMoveStatusResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeMoveStatusRequest returns a decoder for requests sent to the package
// move_status endpoint.
func DecodeMoveStatusRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  uint
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseUint(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "unsigned integer"))
			}
			id = uint(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewMoveStatusPayload(id)

		return payload, nil
	}
}

// EncodeMoveStatusError returns an encoder for errors returned by the
// move_status package endpoint.
func EncodeMoveStatusError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "failed_dependency":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMoveStatusFailedDependencyResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusFailedDependency)
			return enc.Encode(body)
		case "not_found":
			var res *package_.PackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMoveStatusNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalPackageViewsEnduroMonitorPingEventViewToEnduroMonitorPingEventResponseBody
// builds a value of type *EnduroMonitorPingEventResponseBody from a value of
// type *package_views.EnduroMonitorPingEventView.
func marshalPackageViewsEnduroMonitorPingEventViewToEnduroMonitorPingEventResponseBody(v *package_views.EnduroMonitorPingEventView) *EnduroMonitorPingEventResponseBody {
	if v == nil {
		return nil
	}
	res := &EnduroMonitorPingEventResponseBody{
		Message: v.Message,
	}

	return res
}

// marshalPackageViewsEnduroPackageCreatedEventViewToEnduroPackageCreatedEventResponseBody
// builds a value of type *EnduroPackageCreatedEventResponseBody from a value
// of type *package_views.EnduroPackageCreatedEventView.
func marshalPackageViewsEnduroPackageCreatedEventViewToEnduroPackageCreatedEventResponseBody(v *package_views.EnduroPackageCreatedEventView) *EnduroPackageCreatedEventResponseBody {
	if v == nil {
		return nil
	}
	res := &EnduroPackageCreatedEventResponseBody{
		ID: *v.ID,
	}
	if v.Item != nil {
		res.Item = marshalPackageViewsEnduroStoredPackageViewToEnduroStoredPackageResponseBody(v.Item)
	}

	return res
}

// marshalPackageViewsEnduroStoredPackageViewToEnduroStoredPackageResponseBody
// builds a value of type *EnduroStoredPackageResponseBody from a value of type
// *package_views.EnduroStoredPackageView.
func marshalPackageViewsEnduroStoredPackageViewToEnduroStoredPackageResponseBody(v *package_views.EnduroStoredPackageView) *EnduroStoredPackageResponseBody {
	res := &EnduroStoredPackageResponseBody{
		ID:          *v.ID,
		Name:        v.Name,
		Location:    v.Location,
		Status:      *v.Status,
		WorkflowID:  v.WorkflowID,
		RunID:       v.RunID,
		AipID:       v.AipID,
		CreatedAt:   *v.CreatedAt,
		StartedAt:   v.StartedAt,
		CompletedAt: v.CompletedAt,
	}

	return res
}

// marshalPackageViewsEnduroPackageDeletedEventViewToEnduroPackageDeletedEventResponseBody
// builds a value of type *EnduroPackageDeletedEventResponseBody from a value
// of type *package_views.EnduroPackageDeletedEventView.
func marshalPackageViewsEnduroPackageDeletedEventViewToEnduroPackageDeletedEventResponseBody(v *package_views.EnduroPackageDeletedEventView) *EnduroPackageDeletedEventResponseBody {
	if v == nil {
		return nil
	}
	res := &EnduroPackageDeletedEventResponseBody{
		ID: *v.ID,
	}

	return res
}

// marshalPackageViewsEnduroPackageUpdatedEventViewToEnduroPackageUpdatedEventResponseBody
// builds a value of type *EnduroPackageUpdatedEventResponseBody from a value
// of type *package_views.EnduroPackageUpdatedEventView.
func marshalPackageViewsEnduroPackageUpdatedEventViewToEnduroPackageUpdatedEventResponseBody(v *package_views.EnduroPackageUpdatedEventView) *EnduroPackageUpdatedEventResponseBody {
	if v == nil {
		return nil
	}
	res := &EnduroPackageUpdatedEventResponseBody{
		ID: *v.ID,
	}
	if v.Item != nil {
		res.Item = marshalPackageViewsEnduroStoredPackageViewToEnduroStoredPackageResponseBody(v.Item)
	}

	return res
}

// marshalPackageViewsEnduroPackageStatusUpdatedEventViewToEnduroPackageStatusUpdatedEventResponseBody
// builds a value of type *EnduroPackageStatusUpdatedEventResponseBody from a
// value of type *package_views.EnduroPackageStatusUpdatedEventView.
func marshalPackageViewsEnduroPackageStatusUpdatedEventViewToEnduroPackageStatusUpdatedEventResponseBody(v *package_views.EnduroPackageStatusUpdatedEventView) *EnduroPackageStatusUpdatedEventResponseBody {
	if v == nil {
		return nil
	}
	res := &EnduroPackageStatusUpdatedEventResponseBody{
		ID:     *v.ID,
		Status: *v.Status,
	}

	return res
}

// marshalPackageViewsEnduroPackageLocationUpdatedEventViewToEnduroPackageLocationUpdatedEventResponseBody
// builds a value of type *EnduroPackageLocationUpdatedEventResponseBody from a
// value of type *package_views.EnduroPackageLocationUpdatedEventView.
func marshalPackageViewsEnduroPackageLocationUpdatedEventViewToEnduroPackageLocationUpdatedEventResponseBody(v *package_views.EnduroPackageLocationUpdatedEventView) *EnduroPackageLocationUpdatedEventResponseBody {
	if v == nil {
		return nil
	}
	res := &EnduroPackageLocationUpdatedEventResponseBody{
		ID:       *v.ID,
		Location: *v.Location,
	}

	return res
}

// marshalPackageEnduroStoredPackageToEnduroStoredPackageResponseBody builds a
// value of type *EnduroStoredPackageResponseBody from a value of type
// *package_.EnduroStoredPackage.
func marshalPackageEnduroStoredPackageToEnduroStoredPackageResponseBody(v *package_.EnduroStoredPackage) *EnduroStoredPackageResponseBody {
	res := &EnduroStoredPackageResponseBody{
		ID:          v.ID,
		Name:        v.Name,
		Location:    v.Location,
		Status:      v.Status,
		WorkflowID:  v.WorkflowID,
		RunID:       v.RunID,
		AipID:       v.AipID,
		CreatedAt:   v.CreatedAt,
		StartedAt:   v.StartedAt,
		CompletedAt: v.CompletedAt,
	}

	return res
}

// marshalPackageViewsEnduroPackageWorkflowHistoryViewToEnduroPackageWorkflowHistoryResponseBody
// builds a value of type *EnduroPackageWorkflowHistoryResponseBody from a
// value of type *package_views.EnduroPackageWorkflowHistoryView.
func marshalPackageViewsEnduroPackageWorkflowHistoryViewToEnduroPackageWorkflowHistoryResponseBody(v *package_views.EnduroPackageWorkflowHistoryView) *EnduroPackageWorkflowHistoryResponseBody {
	if v == nil {
		return nil
	}
	res := &EnduroPackageWorkflowHistoryResponseBody{
		ID:      v.ID,
		Type:    v.Type,
		Details: v.Details,
	}

	return res
}
