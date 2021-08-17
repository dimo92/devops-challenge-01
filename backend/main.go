package main

import (
	"log"
	"net/http"

	"github.com/dimo92/devops-challenge-01/controllers"
	"github.com/dimo92/devops-challenge-01/database"
	"github.com/dimo92/devops-challenge-01/entity"

	"github.com/gorilla/mux"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/signin", controllers.SignIn).Methods("POST")
	router.HandleFunc("/create", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/get/{sku}", controllers.GetProductBySKU).Methods("GET")
	router.HandleFunc("/update/{sku}", controllers.UpdateProductBySKU).Methods("PUT")
	router.HandleFunc("/delete/{sku}", controllers.DeleteProductBySKU).Methods("DELETE")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "db:3306",
			User:       "db_user_catalog",
			Password:   "db_user_pass_catalog",
			DB:         "catalog",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	database.Migrate(&entity.Product{}, &entity.User{})
}
