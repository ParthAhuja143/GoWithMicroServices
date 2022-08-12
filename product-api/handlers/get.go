package handlers

import (
	"net/http"

	"github.com/ParthAhuja143/GoWithMicroServices/data"
	"github.com/ParthAhuja143/GoWithMicroServices/errors"
)

//swagger:route GET /products listProducts
//Returns a list of products
//responses:
//200: productsResponse

//GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	p.l.Println("Handle GET Products")

	productList := data.GetProducts()

	err := productList.ToJSON(rw)

	if err != errors.NoErr {
		http.Error(rw, err.Title(), err.StatusCode())
	}
}