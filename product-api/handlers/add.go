package handlers

import (
	"net/http"

	"github.com/ParthAhuja143/GoWithMicroServices/data"
	"github.com/ParthAhuja143/GoWithMicroServices/errors"
)

//swagger:route POST /products addProduct
//Adds a product to the list
//responses:
//200: noResponse

//AddProducts returns the products from the data store

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	//get product from context
	product := r.Context().Value(data.RequestBodyProduct{}).(data.Product)
	newProductList := data.AddProduct(&product)

	err := newProductList.ToJSON(rw)

	if err != errors.NoErr {
		http.Error(rw, err.Title(), err.StatusCode())
		return
	}

	rw.WriteHeader(http.StatusOK)
}
