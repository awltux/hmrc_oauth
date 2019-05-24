package hmrc

import (
	"context"
	"log"

	mtd "github.com/awltux/hmrc_oauth/gen/mtd"
)

// mtd service example implementation.
// The example methods log the requests and return zero values.
type mtdsrvc struct {
	logger *log.Logger
}

// NewMtd returns the mtd service implementation.
func NewMtd(logger *log.Logger) mtd.Service {
	return &mtdsrvc{logger}
}

// Store key that will store oauth token
func (s *mtdsrvc) Register(ctx context.Context, p *mtd.StatePayload) (err error) {
	s.logger.Print("mtd.register")
	return
}

// Store key that will store oauth token
func (s *mtdsrvc) Retrieve(ctx context.Context, p *mtd.StatePayload) (err error) {
	s.logger.Print("mtd.retrieve")
	return
}

// Authentication code response
func (s *mtdsrvc) HmrcCallback(ctx context.Context, p *mtd.CodePayload) (err error) {
	s.logger.Print("mtd.hmrc_callback")
	return
}
