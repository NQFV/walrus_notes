package main

import (
	"cli/app"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	pass, exists := os.LookupEnv("PASSWORD")
	if !exists {
		log.Fatal("Пароля нет")
	}
	jopa := app.NewApp(pass)
	jopa.RunApp()
}
