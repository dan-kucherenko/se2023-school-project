package main

import (
	"github.com/dan-kucherenko/se-school-project/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnvFile()
	router.ActivateRouter()
}

// load .env file for future function calls
func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
