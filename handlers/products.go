package handlers

import (
	"log"
	"net/http"
	"product-api/data"
)

// Products is a handler for our coffee shop's products
type Products struct {
	l *log.Logger
}

// NewProducts initializes a new Products instance
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP contains HTTP processing logic
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts lists the products
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Can't marshal JSON", http.StatusInternalServerError)
	}
}

// addProduct adds a product to the product db
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	p.l.Printf("Product: %#v", prod)
	data.AddProduct(prod)
}
