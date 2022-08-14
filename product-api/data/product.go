package data

import (
	"encoding/json"
	"io"
	"regexp"
	"time"

	"github.com/ParthAhuja143/GoWithMicroServices/product-api/errors"
	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
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

// Used in middleware in request context, we can use strings as well but preffered method is to use types
type RequestBodyProduct struct{}

func (p *Product) Validate() errors.HTTPError {
	validate := validator.New()
	// custom validation for sku
	validate.RegisterValidation("sku", validateSKU)

	err := validate.Struct(p)

	if err != nil {
		return errors.ErrValidatingProduct
	}

	return errors.NoErr

}

func validateSKU(fieldLevel validator.FieldLevel) bool {
	regex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := regex.FindAllString(fieldLevel.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}

func GetProducts() Products {
	return ProductList
}

func AddProduct(p *Product) Products {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
	return ProductList
}

func UpdateProduct(id int, p *Product) errors.HTTPError {
	_, pos, err := findProductByID(id)

	if err != errors.NoErr {
		return err
	}

	p.ID = id
	ProductList[pos] = p

	return errors.NoErr
}

func DeleteProduct(id int) errors.HTTPError {
	err := findProductByIDAndDelete(id)

	return err
}

func (p *Product) FromJSON(r io.Reader) errors.HTTPError {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(p)
	if err == nil {
		return errors.NoErr
	}

	return errors.ErrUnmarshal
}

func (p *Products) ToJSON(w io.Writer) errors.HTTPError {
	encoder := json.NewEncoder(w)

	err := encoder.Encode(p)

	if err == nil {
		return errors.NoErr
	}

	return errors.ErrMarshal
}

func getNextID() int {
	lastProduct := ProductList[len(ProductList)-1]
	return lastProduct.ID + 1
}

func findProductByID(id int) (*Product, int, errors.HTTPError) {

	for i, p := range ProductList {
		if p.ID == id {
			return p, i, errors.NoErr
		}
	}

	return nil, -1, errors.ErrProductNotFound
}

func findProductByIDAndDelete(id int) errors.HTTPError {
	newArr := make([]*Product, 0)
	var IDFound bool = false
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
	return errors.NoErr
}
