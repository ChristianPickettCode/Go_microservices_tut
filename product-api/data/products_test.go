package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "p1",
		Price: 1.99,
		SKU:   "abc-sdf-adfas",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
