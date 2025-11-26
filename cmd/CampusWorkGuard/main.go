package main

import (
	"CampusWorkGuardBackend/internal/router"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
