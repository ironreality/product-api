package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product represents an item in our coffee shop
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreateOn    string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeleteOn    string  `json:"-"`
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.43,
		SKU:         "abc323",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Americano",
		Description: "A little milk in coffee",
		SKU:         "fhg048",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}

// GetProducts returns the product list
func GetProducts() Products {
	return productList
}

// Products represent a list of product
type Products []*Product

// ToJSON serializes the product list to JSON
// and writes it to stream
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// AddProduct add a new product to the product db (i.e. our list)
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// FromJSON reads a product spec from r
func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}
