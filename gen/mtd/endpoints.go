// Code generated by goa v3.0.2, DO NOT EDIT.
//
// mtd endpoints
//
// Command:
// $ goa gen github.com/awltux/hmrc_oauth/design

package mtd

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "mtd" service endpoints.
type Endpoints struct {
	Register     goa.Endpoint
	Retrieve     goa.Endpoint
	HmrcCallback goa.Endpoint
}

// NewEndpoints wraps the methods of the "mtd" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Register:     NewRegisterEndpoint(s),
		Retrieve:     NewRetrieveEndpoint(s),
		HmrcCallback: NewHmrcCallbackEndpoint(s),
	}
}

// Use applies the given middleware to all the "mtd" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Register = m(e.Register)
	e.Retrieve = m(e.Retrieve)
	e.HmrcCallback = m(e.HmrcCallback)
}

// NewRegisterEndpoint returns an endpoint function that calls the method
// "register" of service "mtd".
func NewRegisterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StatePayload)
		return nil, s.Register(ctx, p)
	}
}

// NewRetrieveEndpoint returns an endpoint function that calls the method
// "retrieve" of service "mtd".
func NewRetrieveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*StatePayload)
		return s.Retrieve(ctx, p)
	}
}

// NewHmrcCallbackEndpoint returns an endpoint function that calls the method
// "hmrc_callback" of service "mtd".
func NewHmrcCallbackEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CodePayload)
		return nil, s.HmrcCallback(ctx, p)
	}
}
