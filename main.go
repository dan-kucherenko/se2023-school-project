package main

import (
	"github.com/dan-kucherenko/se-school-project/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// load .env file for the future
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	router.ActivateRouter()
}
