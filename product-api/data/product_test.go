package data

import (
	"testing"

	"github.com/ParthAhuja143/GoWithMicroServices/product-api/errors"
)

func TestCheckValidation(t *testing.T) {
	p := &Product{}

	err := p.Validate()

	if err != errors.NoErr {
		t.Fatal(err)
	}
}
