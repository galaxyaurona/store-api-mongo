package main

import (
	"log"
	"net/http"

	"github.com/galaxyaurona/store-api-mongo/store"
	"github.com/joho/godotenv"
)

// GET ALL

// GET ONE

// CREATE ONE

// UPDATE

// DELETE

// Main

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := store.NewRouter()
	http.ListenAndServe(":8080", router)
}
