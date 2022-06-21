// Code generated by goa v3.7.6, DO NOT EDIT.
//
// storage client HTTP transport
//
// Command:
// $ goa-v3.7.6 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the storage service endpoint HTTP clients.
type Client struct {
	// Submit Doer is the HTTP client used to make requests to the submit endpoint.
	SubmitDoer goahttp.Doer

	// Update Doer is the HTTP client used to make requests to the update endpoint.
	UpdateDoer goahttp.Doer

	// Download Doer is the HTTP client used to make requests to the download
	// endpoint.
	DownloadDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the storage service servers.
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
		UpdateDoer:          doer,
		DownloadDoer:        doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Submit returns an endpoint that makes HTTP requests to the storage service
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
			return nil, goahttp.ErrRequestError("storage", "submit", err)
		}
		return decodeResponse(resp)
	}
}

// Update returns an endpoint that makes HTTP requests to the storage service
// update server.
func (c *Client) Update() goa.Endpoint {
	var (
		decodeResponse = DecodeUpdateResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("storage", "update", err)
		}
		return decodeResponse(resp)
	}
}

// Download returns an endpoint that makes HTTP requests to the storage service
// download server.
func (c *Client) Download() goa.Endpoint {
	var (
		decodeResponse = DecodeDownloadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDownloadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DownloadDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("storage", "download", err)
		}
		return decodeResponse(resp)
	}
}
