// Code generated by goa v3.1.2, DO NOT EDIT.
//
// pipeline HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package server

import (
	"context"
	"net/http"

	pipeline "github.com/artefactual-labs/enduro/internal/api/gen/pipeline"
	pipelineviews "github.com/artefactual-labs/enduro/internal/api/gen/pipeline/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeListResponse returns an encoder for responses returned by the pipeline
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]*pipeline.EnduroStoredPipeline)
		enc := encoder(ctx, w)
		body := NewListResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the pipeline list
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			name *string
		)
		nameRaw := r.URL.Query().Get("name")
		if nameRaw != "" {
			name = &nameRaw
		}
		payload := NewListPayload(name)

		return payload, nil
	}
}

// EncodeShowResponse returns an encoder for responses returned by the pipeline
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*pipelineviews.EnduroStoredPipeline)
		enc := encoder(ctx, w)
		body := NewShowResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the pipeline show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  string
			err error

			params = mux.Vars(r)
		)
		id = params["id"]
		err = goa.MergeErrors(err, goa.ValidateFormat("id", id, goa.FormatUUID))

		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(id)

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show pipeline
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*pipeline.NotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalPipelineEnduroStoredPipelineToEnduroStoredPipelineResponse builds a
// value of type *EnduroStoredPipelineResponse from a value of type
// *pipeline.EnduroStoredPipeline.
func marshalPipelineEnduroStoredPipelineToEnduroStoredPipelineResponse(v *pipeline.EnduroStoredPipeline) *EnduroStoredPipelineResponse {
	res := &EnduroStoredPipelineResponse{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}
