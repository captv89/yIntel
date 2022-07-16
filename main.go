package main

import (
	"log"

	"github.com/captv89/yIntel/telebot"
	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// hackerNews := hackernews.Top10Stories()
	// fmt.Println(hackerNews)

	telebot.StartBot()

}