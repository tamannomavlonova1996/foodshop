package main

import (
	"Food_shop/db"
	"Food_shop/handlers"
	_ "database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"net/http"
)

func main() {
	db, err := db.InitDB()
	if err == nil {
		fmt.Errorf("DB conn error: ", err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", handlers.GetProducts).Methods("Get")
	router.HandleFunc("/products/{id}", handlers.GetProductsbyID).Methods("Get")
	router.HandleFunc("/products/{id}", handlers.UpdatePruducts).Methods("PUT")
	router.HandleFunc("/products/{id}", handlers.DeletePruducts).Methods("DELETE")

	http.ListenAndServe(":8282", router)

}
