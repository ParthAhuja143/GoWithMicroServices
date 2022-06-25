package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/ParthAhuja143/GoWithMicroServices/data"
	"github.com/ParthAhuja143/GoWithMicroServices/errors"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l* log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	switch r.Method {
		case http.MethodGet:
			p.getProducts(rw, r)

		case http.MethodPost:
			p.addProduct(rw, r)

		case http.MethodPut:
			p.updateProduct(rw, r)	

		case http.MethodDelete:
			p.deleteProduct(rw, r)
		default:
			rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p* Products) getProducts(rw http.ResponseWriter, h *http.Request){
	p.l.Println("Handle GET Products")

	productList := data.GetProducts()

	err := productList.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p* Products) addProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST Products")

	product := &data.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, errors.ErrUnmarshal.Error(), http.StatusBadRequest)
	}

	newProductList := data.AddProduct(product)

	newProductList.ToJSON(rw)

	rw.WriteHeader(http.StatusOK)
}

func (p *Products) updateProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle PUT Products")

	id, err := getIDFromURI(rw, r)

	if err != nil  {
		if err == errors.ErrInvalidURI {
			http.Error(rw, errors.ErrInvalidURI.Error() , http.StatusBadRequest)
		}

		if id == -1 {
			http.Error(rw, errors.ErrProductNotFound.Error(), http.StatusBadRequest)
		}

		return
	}
	
	product := &data.Product{}

	err = product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, errors.ErrUnmarshal.Error(), http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, product)

	if err != nil {
		http.Error(rw, errors.ErrProductNotFound.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (p *Products) deleteProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle DELETE Products")

	id, err := getIDFromURI(rw, r)

	if err != nil {
		if err == errors.ErrInvalidURI {
			http.Error(rw, errors.ErrInvalidURI.Error() , http.StatusBadRequest)
		}

		if id == -1 {
			http.Error(rw, errors.ErrProductNotFound.Error(), http.StatusBadRequest)
		}

		return
	}

	err = data.DeleteProduct(id)

	if err != nil {
		http.Error(rw, errors.ErrProductNotFound.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func getIDFromURI(rw http.ResponseWriter, r *http.Request) (int, error){
	regex := regexp.MustCompile(`/([0-9]+)`)
	path := r.URL.Path

	uidArr := regex.FindAllStringSubmatch(path, -1)
	if len(uidArr) != 1 {
		return -1, errors.ErrInvalidURI
	}

	uid := uidArr[0][1]
	id, err := strconv.Atoi(uid)

	if err != nil {
		return -1, errors.ErrInvalidURI
	}

	return id, nil
}