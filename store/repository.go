package store

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

type Repository interface {
	AddProduct(product Product) Product
	GetProducts() Products
	GetProduct(id int64) Product
	UpdateProduct(id int64, product Product) Product
	DeleteProduct(id int64) bool
}

var err = godotenv.Load()

var SERVER = os.Getenv("DATABASE_URL")
var DBNAME = os.Getenv("DATABASE_NAME")
var COLLECTION = os.Getenv("DATABASE_COLLECTION")

type repository struct{}

func (repo repository) AddProduct(newProduct Product) Product {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	// don't let them insert new product
	newProduct.ID = time.Now().Unix()
	err = session.DB(DBNAME).C(COLLECTION).Insert(newProduct)
	if err != nil {
		log.Fatal(err)
		return Product{}
	}
	fmt.Println("Added New Product ID- ", newProduct.ID)

	return newProduct
}

func (repo repository) GetProducts() Products {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
		return Products{}
	}
	defer session.Close()

	products := Products{}
	collections := session.DB(DBNAME).C(COLLECTION).Find(nil)
	fmt.Printf("collections here %v", collections)
	if err = collections.All(&products); err != nil {
		fmt.Println("Failed to retreive result", err)
	}
	return products
}

func (repo repository) DeleteProduct(id int64) bool {
	return true
}

func (repo repository) UpdateProduct(id int64, updateProduct Product) Product {
	return Product{}
}
func (repo repository) GetProduct(id int64) Product {
	return Product{}
}
