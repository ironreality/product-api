package handlers

import (
	"log"
	"net/http"
	"product-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// Products is a handler for our coffee shop's products
type Products struct {
	l *log.Logger
}

// NewProducts initializes a new Products instance
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts lists the products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Can't marshal JSON", http.StatusInternalServerError)
	}
}

// UpdateProduct replaces a product with a new one
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable convert id to integer number", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product", id)

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	p.l.Printf("Product: %#v", prod)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Internal error ocurred", http.StatusInternalServerError)
		return
	}
}

// AddProduct adds a product to the product db
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	p.l.Printf("Product: %#v", prod)
	data.AddProduct(prod)
}
