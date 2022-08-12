package handlers

import (
	"net/http"
	"strconv"

	"github.com/ParthAhuja143/GoWithMicroServices/data"
	"github.com/ParthAhuja143/GoWithMicroServices/errors"
	"github.com/gorilla/mux"
)

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Products")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, errors.ErrInvalidURI.Title(), errors.ErrInvalidURI.StatusCode())
	}

	product := r.Context().Value(data.RequestBodyProduct{}).(data.Product)

	httperr := data.UpdateProduct(id, &product)

	if httperr != errors.NoErr {
		http.Error(rw, errors.ErrProductNotFound.Title(), errors.ErrProductNotFound.StatusCode())
		return
	}

	rw.WriteHeader(http.StatusOK)
}
