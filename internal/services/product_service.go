package services

import (
	"golang-api/internal/entities"
	"golang-api/internal/repositories"
)

// GetAllProducts returns all products
func GetAllProducts() entities.Response {
	response := repositories.GetProducts()

	return response
}

// GetProductByID returns a product by its id if found, otherwise error
func GetProductByID(id string) entities.Response {
	response := repositories.GetProduct(id)

	return response
}
