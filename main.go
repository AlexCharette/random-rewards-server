package main

import (
	"log"

	api "random-rewards/pkg/api"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	app := api.App{}
	app.Initialize()
	app.Run()
}
