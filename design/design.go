package design
// https://goa.design/reference/goa/design/apidsl
// https://developer.service.hmrc.gov.uk/api-documentation/docs/authorisation/user-restricted-endpoints

import (
  . "goa.design/goa/v3/dsl"
)

var _ = API("hmrc_oauth", func() {
        Title("Shim for HRMRC oAuth API")
        Description("Provides known address for HMRC oAuth API to reply to.")
        Scheme("https")
        Host("localhost")
})

var _ = Resource("hmrc_oauth", func() {
        BasePath("/hmrc_oauth")



        Action("register", func() {
                Description("Store key that will store oauth token")
                Routing(POST("/:state"))
                Params(func() {
                        Param("state", Integer, "Key submitted to oAuth call; normally AES1 digest")
                })

                // New key added
                Response(Created)
                // Key already exists
                Response(Conflict)
                // Key already registered by another client
                Response(Unauthorized)
        })

        Action("retrieve", func() {
                Description("Store key that will store oauth token")
                Routing(GET("/:state"))
                Params(func() {
                        Param("state", Integer, "Key submitted to oAuth call; normally AES1 digest")
                })

                // Token has been returned
                Response(OK)
                // Key exists, but no token added yet
                Response(NoContent)
                // Key was registered by a different IP
                Response(Unauthorized)
                // No key by that name
                Response(NotFound)

        })

        Action("hmrc_callback", func() {
                Description("Authentication code response")
                Routing(POST("/"))
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
                Response(OK)
                // No Key by that name; may have timed out
                Response(NotFound)
        })
})

