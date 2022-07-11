// Code generated by goa v3.7.10, DO NOT EDIT.
//
// batch client HTTP transport
//
// Command:
// $ goa-v3.7.10 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the batch service endpoint HTTP clients.
type Client struct {
	// Submit Doer is the HTTP client used to make requests to the submit endpoint.
	SubmitDoer goahttp.Doer

	// Status Doer is the HTTP client used to make requests to the status endpoint.
	StatusDoer goahttp.Doer

	// Hints Doer is the HTTP client used to make requests to the hints endpoint.
	HintsDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the batch service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		SubmitDoer:          doer,
		StatusDoer:          doer,
		HintsDoer:           doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Submit returns an endpoint that makes HTTP requests to the batch service
// submit server.
func (c *Client) Submit() goa.Endpoint {
	var (
		encodeRequest  = EncodeSubmitRequest(c.encoder)
		decodeResponse = DecodeSubmitResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSubmitRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SubmitDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("batch", "submit", err)
		}
		return decodeResponse(resp)
	}
}

// Status returns an endpoint that makes HTTP requests to the batch service
// status server.
func (c *Client) Status() goa.Endpoint {
	var (
		decodeResponse = DecodeStatusResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildStatusRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.StatusDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("batch", "status", err)
		}
		return decodeResponse(resp)
	}
}

// Hints returns an endpoint that makes HTTP requests to the batch service
// hints server.
func (c *Client) Hints() goa.Endpoint {
	var (
		decodeResponse = DecodeHintsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildHintsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.HintsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("batch", "hints", err)
		}
		return decodeResponse(resp)
	}
}
