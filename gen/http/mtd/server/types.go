// Code generated by goa v3.0.2, DO NOT EDIT.
//
// mtd HTTP server types
//
// Command:
// $ goa gen github.com/awltux/hmrc_oauth/design

package server

import (
	mtd "github.com/awltux/hmrc_oauth/gen/mtd"
)

// NewRegisterStatePayload builds a mtd service register endpoint payload.
func NewRegisterStatePayload(state string) *mtd.StatePayload {
	return &mtd.StatePayload{
		State: &state,
	}
}

// NewRetrieveStatePayload builds a mtd service retrieve endpoint payload.
func NewRetrieveStatePayload(state string) *mtd.StatePayload {
	return &mtd.StatePayload{
		State: &state,
	}
}

// NewHmrcCallbackCodePayload builds a mtd service hmrc_callback endpoint
// payload.
func NewHmrcCallbackCodePayload(code *string, state *string, error *string, errorDescription *string, errorCode *string) *mtd.CodePayload {
	return &mtd.CodePayload{
		Code:             code,
		State:            state,
		Error:            error,
		ErrorDescription: errorDescription,
		ErrorCode:        errorCode,
	}
}
