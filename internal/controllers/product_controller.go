package controllers

import (
	"encoding/json"
	"golang-api/internal/services"
	"net/http"

	"github.com/gorilla/mux"
)

// GetProducts returns all products
func GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	var response = services.GetAllProducts()

	if !response.Success || response.Data == nil {
		http.NotFound(responseWriter, request)
		return
	}

	json.NewEncoder(responseWriter).Encode(response.Data)
}

// GetProduct returns a product by id
func GetProduct(responseWriter http.ResponseWriter, request *http.Request) {
	ID := mux.Vars(request)["id"]
	var response = services.GetProductByID(ID)

	if !response.Success || response.Data == nil {
		http.NotFound(responseWriter, request)
		return
	}

	json.NewEncoder(responseWriter).Encode(response.Data)
}
