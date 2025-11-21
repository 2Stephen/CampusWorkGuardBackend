package main

import (
	"CampusWorkGuardBackend/routers"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := routers.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
