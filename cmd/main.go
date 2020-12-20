package main

import (
	"fmt"
	"golang-api/internal/routes"
)

func main() {
	fmt.Println("Starting RESTful products API...")
	routes.HandleRequests()
}
