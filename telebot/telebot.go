package telebot

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/captv89/yIntel/api"
	operation "github.com/captv89/yIntel/ops"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot() {
	bot, err := tgbot.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	vUserId, _ := strconv.ParseInt(os.Getenv("V_ID"), 10, 64)
	sUserId, _ := strconv.ParseInt(os.Getenv("S_ID"), 10, 64)

	log.Printf("Authorized user ids %d, %d", vUserId, sUserId)

	// Allowed users for the bot
	validUsers := map[int64]bool{vUserId: true, sUserId: true}

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
			msg.Text = "I understand:\n /sayhi : Just a simple Hi! \n /status : Check bot's health status. \n /topten : Top 10 news from hackernews.\n /random : Random news from hackernews. \n /euro : Checks current exchange rate EUR-INR. \n /usd : Checks current exchange rate EUR-INR."
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
