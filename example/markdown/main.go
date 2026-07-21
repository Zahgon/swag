package main

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag/example/markdown/api"
	_ "github.com/swaggo/swag/example/markdown/docs"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/admin/user/", api.ListUsers).Methods("GET")
	router.HandleFunc("/admin/user/{id}", api.GetUser).Methods("GET")
	router.HandleFunc("/admin/user/", api.AddUser).Methods("POST")
	router.HandleFunc("/admin/user/{id}", api.UpdateUser).Methods("PUT")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", router)
}
