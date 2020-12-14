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
		p.GetProducts(rw, r)
		return
	}
	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// GetProducts lists the products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Can't marshal JSON", http.StatusInternalServerError)
	}
}
