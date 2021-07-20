package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:"Benson",
		Price: 5,
		SKU:"abc-def-ghi",
	}

	err := p.Validator()
	if err != nil {
		t.Fatal(err)
	}
}