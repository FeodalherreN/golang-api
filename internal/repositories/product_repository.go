package repositories

import (
	"fmt"
	"log"

	"golang-api/internal/database"
	"golang-api/internal/entities"
)

// GetProducts returns products from database
func GetProducts() entities.Response {
	var response entities.Response
	response.Success = false

	dbInstance, dbError := database.Connect()

	if dbError != nil {
		panic(dbError.Error())
	}

	query := "SELECT id, title, description FROM products"
	rows, queryError := dbInstance.Query(query)
	if queryError != nil {
		fmt.Println(queryError.Error())
		return response
	}

	var products []entities.Product
	for rows.Next() {
		var product entities.Product

		if err := rows.Scan(&product.ID, &product.Title, &product.Description); err != nil {
			log.Println(err.Error())
		}

		products = append(products, product)
	}

	defer dbInstance.Close()
	response.Data = products
	response.Success = true

	return response
}

// GetProduct returns product by given id
func GetProduct(id string) entities.Response {
	var response entities.Response
	response.Success = false

	dbInstance, dbError := database.Connect()

	if dbError != nil {
		panic(dbError.Error())
	}
	defer dbInstance.Close()

	var product entities.Product
	query := "SELECT id, title, description FROM products WHERE id = $1"
	err := dbInstance.QueryRow(query, id).Scan(&product.ID, &product.Title, &product.Description)
	if err != nil {
		fmt.Println(err.Error())
		return response
	}

	response.Data = product
	response.Success = true

	return response
}
