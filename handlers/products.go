package handlers

import (
	"context"
	"fmt"
	"log"
	"microserce/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

// NewProducts creates a Products handler with a given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}



// getProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get Products")
	// fetch the products from the data store
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Could not Marshal data", http.StatusInternalServerError)
	}
}

// postProduct adds a new product to the data store
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars  := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w,"Can convert to interger", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle Put Product", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	// p.l.Printf("Prod %#v", prod)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
	}
}

type KeyProduct struct {}

func (p Products) MilddlewareProductValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Could not unmarshall json", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validator()
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not validate product: %s", err), http.StatusBadRequest)
			return
		}

		// agg the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		// Call the next handler,, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(w, req)
	})
}