package main

import (
	"fmt"
	"testing"

	"github.com/christianpickettcode/go_microservices_tut/sdk/client"
	"github.com/christianpickettcode/go_microservices_tut/sdk/client/products"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
	t.Fail()
}