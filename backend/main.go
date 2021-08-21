package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dimo92/devops-challenge-01/controllers"
	"github.com/dimo92/devops-challenge-01/database"
	"github.com/dimo92/devops-challenge-01/entity"

	"github.com/gorilla/mux"
)

var db_name = os.Getenv("DB_NAME")
var db_user = os.Getenv("DB_USER")
var db_password = os.Getenv("DB_PASSWORD")
var db_host = os.Getenv("DB_HOST")
var db_port = os.Getenv("DB_PORT")

func main() {
	fmt.Println(db_name, " - ", db_user, " - ", db_password, " - ", db_host, " - ", db_port, " - ")
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
			Name:     db_name,
			User:     db_user,
			Password: db_password,
			Host:     db_host,
			Port:     db_port,
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}

	database.Migrate(&entity.Product{}, &entity.User{})
}
