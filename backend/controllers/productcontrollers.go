package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dimo92/devops-challenge-01/database"
	"github.com/dimo92/devops-challenge-01/entity"

	"github.com/gorilla/mux"
)

//GetAllProducts get all products data
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []entity.Product
	database.Connector.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

//GetProductBySKU returns product with specific SKU
func GetProductBySKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["SKU"]

	var product entity.Product
	database.Connector.First(&product, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

//CreateProduct creates product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var product entity.Product
	json.Unmarshal(requestBody, &product)

	database.Connector.Create(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

//UpdateProductBySKU updates product with respective SKU
func UpdateProductBySKU(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var product entity.Product
	json.Unmarshal(requestBody, &product)
	database.Connector.Save(&product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

//DeleteProductBySKU delete's product with specific SKU
func DeleteProductBySKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["SKU"]

	var product entity.Product
	SKU, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("SKU = ?", SKU).Delete(&product)
	w.WriteHeader(http.StatusNoContent)
}
