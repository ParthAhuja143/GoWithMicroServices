package handlers

import (
	"net/http"

	"github.com/ParthAhuja143/GoWithMicroServices/data"
	"github.com/ParthAhuja143/GoWithMicroServices/errors"
	"github.com/ParthAhuja143/GoWithMicroServices/utils"
)

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE Products")

	id, err := utils.GetIDFromURI(rw, r)

	if err != errors.NoErr {
		if err == errors.ErrInvalidURI {
			http.Error(rw, errors.ErrInvalidURI.Title(), errors.ErrInvalidURI.StatusCode())
		}

		if id == -1 {
			http.Error(rw, errors.ErrProductNotFound.Title(), errors.ErrProductNotFound.StatusCode())
		}

		return
	}

	err = data.DeleteProduct(id)

	if err != errors.NoErr {
		http.Error(rw, errors.ErrProductNotFound.Title(), errors.ErrProductNotFound.StatusCode())
		return
	}

	rw.WriteHeader(http.StatusOK)
}