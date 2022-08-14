package HTTPmiddleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ParthAhuja143/GoWithMicroServices/product-api/data"
	"github.com/ParthAhuja143/GoWithMicroServices/product-api/errors"
)

// Serializing/Marshalling the request body product and adding to context
func MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware validating request body")
		product := data.Product{}

		err := product.FromJSON(r.Body)

		if err != errors.NoErr {
			http.Error(rw, err.Title(), err.StatusCode())
			return
		}

		ValidationErr := product.Validate()

		if ValidationErr != errors.NoErr {
			http.Error(rw, ValidationErr.Title(), ValidationErr.StatusCode())
			return
		}

		/*
			ADDING PRODUCT TO THE CONTEXT

			data.RequestBodyProduct -> Key
			product -> value
		*/

		ctx := context.WithValue(r.Context(), data.RequestBodyProduct{}, product)
		r = r.WithContext(ctx)

		// call next handler
		fmt.Println("Middleware validated request body")
		next.ServeHTTP(rw, r)
	})
}
