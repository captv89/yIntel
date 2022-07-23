package main

import (
	"log"
	"os"

	"github.com/captv89/yIntel/models"
	"github.com/captv89/yIntel/telebot"
	yaml "gopkg.in/yaml.v2"
)

var config models.Config

func init() {

	// Load config.yaml file
	log.Println("Loading config.yml...")
	file, err := os.Open("config.yml")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	// Init Decoder
	d := yaml.NewDecoder(file)

	// Decode config.yaml file
	if err := d.Decode(&config); err != nil {
		log.Panic(err)
	}
	log.Println("Config loaded successfully!")

	// Set config to models.Config
	log.Println(config)
}

func main() {
	// Start bot
	telebot.StartBot(&config)
}
