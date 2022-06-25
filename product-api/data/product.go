package data

import (
	"encoding/json"
	"io"
	"time"

	"github.com/ParthAhuja143/GoWithMicroServices/errors"
)

type Product struct {
	ID          int  	`json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32	`json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

var ProductList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Expresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "abc3224",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() Products{
	return ProductList
}

func AddProduct(p* Product) Products{
	p.ID = getNextID()
	ProductList = append(ProductList, p)
	return ProductList
}

func UpdateProduct(id int, p *Product) error{
	_, pos, err := findProductByID(id)

	if err != nil {
		return err
	}

	p.ID = id
	ProductList[pos] = p

	return nil
}

func DeleteProduct(id int) error {
	 err := findProductByIDAndDelete(id)

	 if err != nil {
		return err
	 }

	 return nil
}

func (p *Product) FromJSON(r io.Reader) error{
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error{
	encoder := json.NewEncoder(w)
	return encoder.Encode(p);
}

func getNextID() int{
	lastProduct := ProductList[len(ProductList)-1]
	return lastProduct.ID+1
}

func findProductByID(id int) (*Product, int, error) {

	for i, p := range ProductList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, errors.ErrProductNotFound
}

func findProductByIDAndDelete(id int) error {
	newArr := make([]*Product, 0)
	var IDFound bool = false;
	for _, p := range ProductList {
		if p.ID == id {
			IDFound = true
			continue
		}

		newArr = append(newArr, p)
	}
	if !IDFound {
		return errors.ErrProductNotFound
	}

	ProductList = newArr
	return nil
}