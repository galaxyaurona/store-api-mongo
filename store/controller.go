package store

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	products := c.Repository.GetProducts()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if products != nil {
		data, _ := json.Marshal(products)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

	return
}

func (c *Controller) AddProduct(w http.ResponseWriter, r *http.Request) {
	newProduct := Product{}
	rBody := r.Body
	fmt.Printf("request body %+v", rBody)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {

		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	product := c.Repository.AddProduct(newProduct)
	if (product == Product{}) {
		w.WriteHeader(500)
		w.Write([]byte("error adding product"))
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(product)
	return
}
