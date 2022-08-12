// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /products
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta

package handlers

import (
	"log"

	"github.com/ParthAhuja143/GoWithMicroServices/data"
)

// A list of products returned in response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in:body
	Body []data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts create a new httpHandler with given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

/*func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
			p.GetProducts(rw, r)

		case http.MethodPost:
			p.AddProduct(rw, r)

		case http.MethodPut:
			p.UpdateProduct(rw, r)

		case http.MethodDelete:
			p.DeleteProduct(rw, r)
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}*/
