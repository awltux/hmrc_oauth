// Code generated by goa v3.0.2, DO NOT EDIT.
//
// mtd HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/awltux/hmrc_oauth/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeRegisterResponse returns an encoder for responses returned by the mtd
// register endpoint.
func EncodeRegisterResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

// DecodeRegisterRequest returns a decoder for requests sent to the mtd
// register endpoint.
func DecodeRegisterRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			state string

			params = mux.Vars(r)
		)
		state = params["state"]
		payload := NewRegisterStatePayload(state)

		return payload, nil
	}
}

// EncodeRegisterError returns an encoder for errors returned by the register
// mtd endpoint.
func EncodeRegisterError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "key_length_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewRegisterKeyLengthErrorResponseBody(res)
			w.Header().Set("goa-error", "key_length_error")
			w.WriteHeader(http.StatusPreconditionFailed)
			return enc.Encode(body)
		case "key_already_exists":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewRegisterKeyAlreadyExistsResponseBody(res)
			w.Header().Set("goa-error", "key_already_exists")
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "key_ip_mismatch":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewRegisterKeyIPMismatchResponseBody(res)
			w.Header().Set("goa-error", "key_ip_mismatch")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRetrieveResponse returns an encoder for responses returned by the mtd
// retrieve endpoint.
func EncodeRetrieveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRetrieveRequest returns a decoder for requests sent to the mtd
// retrieve endpoint.
func DecodeRetrieveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			state string

			params = mux.Vars(r)
		)
		state = params["state"]
		payload := NewRetrieveStatePayload(state)

		return payload, nil
	}
}

// EncodeRetrieveError returns an encoder for errors returned by the retrieve
// mtd endpoint.
func EncodeRetrieveError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewRetrieveInvalidRequestResponseBody(res)
			w.Header().Set("goa-error", "invalid_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "key_has_no_token":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewRetrieveKeyHasNoTokenResponseBody(res)
			w.Header().Set("goa-error", "key_has_no_token")
			w.WriteHeader(http.StatusPartialContent)
			return enc.Encode(body)
		case "key_ip_mismatch":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewRetrieveKeyIPMismatchResponseBody(res)
			w.Header().Set("goa-error", "key_ip_mismatch")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "matching_key_not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewRetrieveMatchingKeyNotFoundResponseBody(res)
			w.Header().Set("goa-error", "matching_key_not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeHmrcCallbackResponse returns an encoder for responses returned by the
// mtd hmrc_callback endpoint.
func EncodeHmrcCallbackResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeHmrcCallbackRequest returns a decoder for requests sent to the mtd
// hmrc_callback endpoint.
func DecodeHmrcCallbackRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			code             *string
			state            *string
			error            *string
			errorDescription *string
			errorCode        *string
		)
		codeRaw := r.URL.Query().Get("code")
		if codeRaw != "" {
			code = &codeRaw
		}
		stateRaw := r.URL.Query().Get("state")
		if stateRaw != "" {
			state = &stateRaw
		}
		errorRaw := r.URL.Query().Get("error")
		if errorRaw != "" {
			error = &errorRaw
		}
		errorDescriptionRaw := r.URL.Query().Get("error_description")
		if errorDescriptionRaw != "" {
			errorDescription = &errorDescriptionRaw
		}
		errorCodeRaw := r.URL.Query().Get("error_code")
		if errorCodeRaw != "" {
			errorCode = &errorCodeRaw
		}
		payload := NewHmrcCallbackCodePayload(code, state, error, errorDescription, errorCode)

		return payload, nil
	}
}

// EncodeHmrcCallbackError returns an encoder for errors returned by the
// hmrc_callback mtd endpoint.
func EncodeHmrcCallbackError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "matching_key_not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewHmrcCallbackMatchingKeyNotFoundResponseBody(res)
			w.Header().Set("goa-error", "matching_key_not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "invalid_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewHmrcCallbackInvalidRequestResponseBody(res)
			w.Header().Set("goa-error", "invalid_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "key_length_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewHmrcCallbackKeyLengthErrorResponseBody(res)
			w.Header().Set("goa-error", "key_length_error")
			w.WriteHeader(http.StatusPreconditionFailed)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
