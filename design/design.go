package design

// https://goa.design/reference/goa/design/apidsl
// https://developer.service.hmrc.gov.uk/api-documentation/docs/authorisation/user-restricted-endpoints

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("hmrc", func() {
	Title("Shim for HRMRC oAuth API")
	Description("Provides known address for HMRC oAuth API to reply to.")
	Server("mtdServer", func() {
		Host("production", func() {
			Description("Deployed in AWS")
			URI("https://v1.hmrc.awltux.trade/mtd")
			Variable("version", String, "API version", func() {
				Default("v1")
			})
		})
		Host("development", func() {
			Description("Development hosts.")
			URI("http://localhost/mtd")
		})
	})

})

var StatePayload = Type("StatePayload", func() {
	Field(1, "state", String, "AES1 digest string to identify client")
})

var CodePayload = Type("CodePayload", func() {
	Field(1, "state", String, "AES1 digest string to identify client")
	Field(2, "code", String, "HMRC string to identify client")
})


var ErrorPayload = Type("ErrorPayload", func() {
	Field(1, "state", String, "AES1 digest string to identify client")
})

var _ = Service("mtd", func() {
	HTTP(func() {
		Path("/mtd")
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
			Response(StatusConflict)
			// Key already registered by another client
			Response(StatusUnauthorized)
		})

	})

	Method("retrieve", func() {
		Description("Store key that will store oauth token")
		Payload(StatePayload)
		HTTP(func() {
			GET("/{state}")
			Params(func() {
				Param("state", String, "Key submitted to oAuth call; normally AES1 digest")
			})
			// Token has been returned
			Response(StatusOK)
			// Key exists, but no token added yet
			Response(StatusNoContent)
			// Key was registered by a different IP
			Response(StatusUnauthorized)
			// No key by that name
			Response(StatusNotFound)
		})

	})

	Method("hmrc_callback", func() {
		Description("Authentication code response")
		Payload(CodePayload)
		HTTP(func() {
			POST("/")
			Params(func() {
				// These are used only for success condition
				Param("code", String, "Authorization code from HMRC; times out in 10 mins")
				Param("state", String, "Key submitted by client to oAuth call; normally AES1 digest")
				// These are used only for error condition
				Param("error", String, "access_denied")
				Param("error_description", String, "URL encoded error description")
				Param("error_code", String, "HMRC code for the error")
			})
			// Token added successfully
			Response(StatusOK)
			// No Key by that name; may have timed out
			Response(StatusNotFound)
		})

	})
})
