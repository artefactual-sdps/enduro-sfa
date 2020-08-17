// Code generated by goa v3.2.3, DO NOT EDIT.
//
// batch HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	batch "github.com/artefactual-labs/enduro/internal/api/gen/batch"
	goahttp "goa.design/goa/v3/http"
)

// BuildSubmitRequest instantiates a HTTP request object with method and path
// set to call the "batch" service "submit" endpoint
func (c *Client) BuildSubmitRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SubmitBatchPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("batch", "submit", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSubmitRequest returns an encoder for requests sent to the batch submit
// server.
func EncodeSubmitRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*batch.SubmitPayload)
		if !ok {
			return goahttp.ErrInvalidType("batch", "submit", "*batch.SubmitPayload", v)
		}
		body := NewSubmitRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("batch", "submit", err)
		}
		return nil
	}
}

// DecodeSubmitResponse returns a decoder for responses returned by the batch
// submit endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeSubmitResponse may return the following errors:
//	- "not_available" (type *goa.ServiceError): http.StatusConflict
//	- "not_valid" (type *goa.ServiceError): http.StatusBadRequest
//	- error: internal error
func DecodeSubmitResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusAccepted:
			var (
				body SubmitResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("batch", "submit", err)
			}
			err = ValidateSubmitResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("batch", "submit", err)
			}
			res := NewSubmitBatchResultAccepted(&body)
			return res, nil
		case http.StatusConflict:
			var (
				body SubmitNotAvailableResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("batch", "submit", err)
			}
			err = ValidateSubmitNotAvailableResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("batch", "submit", err)
			}
			return nil, NewSubmitNotAvailable(&body)
		case http.StatusBadRequest:
			var (
				body SubmitNotValidResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("batch", "submit", err)
			}
			err = ValidateSubmitNotValidResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("batch", "submit", err)
			}
			return nil, NewSubmitNotValid(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("batch", "submit", resp.StatusCode, string(body))
		}
	}
}

// BuildStatusRequest instantiates a HTTP request object with method and path
// set to call the "batch" service "status" endpoint
func (c *Client) BuildStatusRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StatusBatchPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("batch", "status", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeStatusResponse returns a decoder for responses returned by the batch
// status endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeStatusResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body StatusResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("batch", "status", err)
			}
			err = ValidateStatusResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("batch", "status", err)
			}
			res := NewStatusBatchStatusResultOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("batch", "status", resp.StatusCode, string(body))
		}
	}
}
