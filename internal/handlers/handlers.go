package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	mdw "guolmal/internal/middlewares"
	"guolmal/internal/routers"
	"log"
	"net/http"
	"os"
)

func Handlers() {
	router := mux.NewRouter()

	// add prefix for routes with "apiCreate"
	apiCreate := router.PathPrefix("/api/v1").Subrouter()

	// route for getting all the products
	router.HandleFunc("/", printHelloWorld).Methods(http.MethodGet)

	// return a list of all products in the database matching the query
	apiCreate.HandleFunc("/products", mdw.CheckDB(routers.SearchProducts)).Methods(http.MethodGet)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func printHelloWorld(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))
}
