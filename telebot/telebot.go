package telebot

import (
	"fmt"
	"log"

	"github.com/captv89/yIntel/api"
	"github.com/captv89/yIntel/models"
	operation "github.com/captv89/yIntel/ops"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var config *models.Config

func StartBot(c *models.Config) {
	// Assign config to global variable
	config = c

	// log.Println("Received Config",config)

	log.Println("Starting bot... TelegramToken: ", config.TelegramToken)
	bot, err := tgbot.NewBotAPI(config.TelegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	// vUserId, _ := strconv.ParseInt(os.Getenv("V_ID"), 10, 64)
	// sUserId, _ := strconv.ParseInt(os.Getenv("S_ID"), 10, 64)
	// vishalUserId, _ := strconv.ParseInt(os.Getenv("VISHAL_ID"), 10, 64)

	validUsers := make(map[int64]bool)

	for i := 0; i < len(config.AllowedIds); i++ {
		log.Println("AllowedId: ", config.AllowedIds[i])
		validUsers[config.AllowedIds[i]] = true
	}

	// Allowed users for the bot
	// validUsers := map[int64]bool{vUserId: true, sUserId: true, vishalUserId: true}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbot.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			log.Println("[DEBUG] Update.Message is nil")
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			log.Println("[DEBUG] Update.Message.IsCommand() is false")
			continue
		}

		if !validUsers[update.Message.From.ID] {
			log.Println("[DEBUG] Update.Message.From.ID != 479110172")
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbot.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand:\n /sayhi : Just a simple Hi! \n /status : Check bot's health status. \n /topten : Top 10 news from hackernews.\n /random : Random news from hackernews. \n /euro : Checks current exchange rate EUR-INR. \n /usd : Checks current exchange rate EUR-INR.\n /joke : Random joke from Joke API.\n /chuck : Random joke from Chuck Norris API. \n /dadjoke : Random dad jokes. \n /meme : Random memes."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		case "topten":
			sendTopTenNews(bot, update.Message.Chat.ID)
		case "random":
			sendRandomNews(bot, update.Message.Chat.ID)
		case "euro":
			sendExchangeRate(bot, update.Message.Chat.ID, "eur")
		case "usd":
			sendExchangeRate(bot, update.Message.Chat.ID, "usd")
		case "joke":
			sendJokes(bot, update.Message.Chat.ID)
		case "chuck":
			sendChuckJokes(bot, update.Message.Chat.ID)
		case "dadjoke":
			sendDadJokes(bot, update.Message.Chat.ID)
		case "meme":
			sendMeme(bot, update.Message.Chat.ID)
		default:
			msg.Text = "I don't know that command"
		}

		log.Println(msg.Text)

		if msg.Text != "" {
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		} else {
			continue
		}
	}
}

//  Bot helper functions

func sendTopTenNews(bot *tgbot.BotAPI, chatID int64) {
	hackerNews := operation.Top10Stories()
	for _, story := range hackerNews {
		msg := tgbot.NewMessage(chatID, story)
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}

func sendRandomNews(bot *tgbot.BotAPI, chatID int64) {
	randomNews := operation.RandomStory()
	msg := tgbot.NewMessage(chatID, randomNews)
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func sendExchangeRate(bot *tgbot.BotAPI, chatID int64, currency string) {
	exchangeRate := api.GetExchangeRate(currency)
	exchangeMsg := fmt.Sprintf("From: %s \nOn: %s \nINR (₹): %f", currency, exchangeRate.Date, exchangeRate.INR)
	msg := tgbot.NewMessage(chatID, exchangeMsg)
	_, err := bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}
}

func sendJokes(bot *tgbot.BotAPI, chatID int64) {
	joke, status := api.Jokes()

	if status == true {
		jokeMsg := fmt.Sprintf("Category: %s\nType: %s\nJoke: %s", joke.Category, joke.JokeType, joke.Joke)

		msg := tgbot.NewMessage(chatID, jokeMsg)
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}

	} else {
		msg := tgbot.NewMessage(chatID, "Sorry, I couldn't find any jokes. Try again later.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}

func sendChuckJokes(bot *tgbot.BotAPI, chatID int64) {
	joke, status := api.Chuck()

	if status == true {
		jokeMsg := fmt.Sprintf("Joke: %s", joke.Value)

		msg := tgbot.NewMessage(chatID, jokeMsg)
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}

	} else {
		msg := tgbot.NewMessage(chatID, "Sorry, I couldn't find any jokes. Try again later.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}

func sendDadJokes(bot *tgbot.BotAPI, chatID int64) {
	joke, status := api.DadJokes()

	if status == true {
		jokeMsg := fmt.Sprintf("Joke: %s", joke.Joke)

		msg := tgbot.NewMessage(chatID, jokeMsg)
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}

	} else {
		msg := tgbot.NewMessage(chatID, "Sorry, I couldn't find any jokes. Try again later.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}

func sendMeme(bot *tgbot.BotAPI, chatID int64) {
	meme, status := api.Meme()

	if status == true {
		memeMsg := fmt.Sprintf("%s by %s", meme.Title, meme.Author)
		memeUrl := fmt.Sprintf("%s", meme.URL)
		msg := tgbot.NewPhoto(chatID, tgbot.FileURL(memeUrl))
		msg.Caption = memeMsg
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}

	} else {
		msg := tgbot.NewMessage(chatID, "Sorry, I couldn't find any memes. Try again later.")
		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}
