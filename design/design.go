package design

// https://goa.design/reference/goa/design/apidsl
// https://developer.service.hmrc.gov.uk/api-documentation/docs/authorisation/user-restricted-endpoints

import (
	. "github.com/awltux/goa"
)

var _ = API("hmrc", func() {
	Title("Shim for HRMRC oAuth API")
	Description("Provides known address for HMRC oAuth API to reply to.")
	Server("mtdServer", func() {
		Host("prod", func() {
			Description("Deployed in AWS")
			URI("https://hmrc.awltux.trade/v1/mtd")
			Variable("version", String, "API version", func() {
				Default("v1")
			})
		})
		Host("dev", func() {
			Description("Development hosts.")
			URI("http://localhost:8088/v1/mtd")
		})
	})

})

var StatePayload = Type("StatePayload", func() {
	Field(1, "state", String, "AES1 digest string to identify client")
})

// CodePayload is the code/state
var CodePayload = Type("CodePayload", func() {
	Field(1, "state", String, "AES1 digest string to identify client")
	Field(2, "code", String, "HMRC string to identify client")
	Field(3, "error", String, "Error String")
	Field(4, "error_description", String, "Error String")
	Field(5, "error_code", String, "Error String")
})

var _ = Service("mtd", func() {
	Error("key_length_error")
	Error("key_already_exists")
	Error("key_ip_mismatch")
	Error("key_has_no_token")
	Error("matching_key_not_found")
	Error("invalid_request")

	HTTP(func() {
		Path("/v1/mtd")
	})

	Method("register", func() {
		Description("Store key that will store oauth token")
		Payload(StatePayload)
		HTTP(func() {
			POST("/{state}")
			Params(func() {
				Param("state", String, "Key submitted to oAuth call; normally AES1 digest")
			})
			// New key added
			Response(StatusCreated)
			// Key already exists
			Response("key_length_error", StatusPreconditionFailed)
			Response("key_already_exists", StatusConflict)
			// Key already registered by another client
			Response("key_ip_mismatch", StatusUnauthorized)

		})

	})

	Method("retrieve", func() {
		Description("Store key that will store oauth token")
		Payload(StatePayload)
		Result(String, "Token from HMRC oAuth")
		HTTP(func() {
			GET("/{state}")
			Params(func() {
				Param("state", String, "Key submitted to oAuth call; normally AES1 digest")
			})
			// Token has been returned
			Response(StatusOK)
			Response("invalid_request", StatusBadRequest)
			// Key exists, but no token added yet
			Response("key_has_no_token", StatusPartialContent)
			// Key was registered by a different IP
			Response("key_ip_mismatch", StatusUnauthorized)
			// No key by that name
			Response("matching_key_not_found", StatusNotFound)

		})

	})

	Method("hmrc_callback", func() {
		Description("Authentication code response")
		Payload(CodePayload)
		HTTP(func() {
			POST("")
			Params(func() {
				// These are used only for success condition
				Param("code", String, "Authorization code from HMRC; times out in 10 mins")
				Param("state", String, "Key submitted by client to oAuth call; normally AES1 digest")
				Param("error", String, "access_denied")
				Param("error_description", String, "URL encoded error description")
				Param("error_code", String, "HMRC code for the error")
			})
			// Token added successfully
			Response(StatusOK)
			// No Key by that name; may have timed out
			Response("matching_key_not_found", StatusNotFound)
			Response("invalid_request", StatusBadRequest)
			Response("key_length_error", StatusPreconditionFailed)

		})

	})
})
