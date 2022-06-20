package handlers

import (
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l* log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	productList := data.GetProducts()
}