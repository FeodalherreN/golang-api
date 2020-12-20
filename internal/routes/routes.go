package routes

import (
	"fmt"
	"log"
	"net/http"

	"golang-api/internal/controllers"
	"golang-api/internal/middleware"

	"github.com/gorilla/mux"
)

// HandleRequests creates an instance of mux and set routes that points towards controllers
func HandleRequests() {
	port := "8080"
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(middleware.HeaderMiddleware)
	setRoutes(myRouter)
	printEndpoints(myRouter)
	fmt.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), myRouter))
}

func setRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts)
	router.HandleFunc("/api/product/{id}", controllers.GetProduct)
}

func printEndpoints(router *mux.Router) {
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})
}
