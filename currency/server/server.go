package server

import (
	"context"

	protos "github.com/christianpickettcode/go_microservices_tut/currency/protos/currency"
	hclog "github.com/hashicorp/go-hclog"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	log hclog.Logger
}

type UnimplementedCurrencyServer struct {
}

// NewCurrency creates a new Currency server
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate
// for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}

