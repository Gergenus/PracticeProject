package main

import (
	"log"

	telegrambot "github.com/Gergenus/telegramBot"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = telegrambot.InitBot()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
