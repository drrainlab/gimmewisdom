package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Wisdom struct {
	num    int
	author string
	quote  string
}

func (w *Wisdom) format() string {
	return "*" + w.quote + "*\n\n" + "_ – " + w.author + "_"
}

var wisdoms []Wisdom

func initWisdoms(filepath string) int {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Panic("\nCannot read wisdoms file!\n", err)
	}
	splitted := bytes.Split(data, []byte("\n\n"))
	wisdoms = make([]Wisdom, len(splitted))
	for i, unparsedQuote := range splitted {
		splitted := strings.Split(string(unparsedQuote), ". -")
		wisdoms[i] = Wisdom{num: i, author: splitted[1], quote: splitted[0]}
	}
	return len(wisdoms)
}

func main() {

	var (
		publicURL = os.Getenv("PUBLIC_URL")
		port      = os.Getenv("PORT")
		token     = os.Getenv("TOKEN")
	)

	initWisdoms("/app/wisdoms.txt")

	fmt.Println("Total wisdoms parsed: ", len(wisdoms))

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(publicURL))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)

	for update := range updates {

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		cmd := update.Message.Command()

		// Process command
		if len(cmd) > 0 {

			fmt.Printf("Команда: %s", cmd)

			switch cmd {
			case "gimmewisdom":

				rand.Seed(time.Now().UnixNano())
				n := rand.Intn(len(wisdoms))
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, wisdoms[n].format())
				msg.ParseMode = tgbotapi.ModeMarkdown
				bot.Send(msg)

			}
		}
	}
}
