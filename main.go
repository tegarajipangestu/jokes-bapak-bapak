package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"strings"
	"os"
	"io"
	"encoding/csv"
	"math/rand"
	"time"
)

type Joke struct {
		id string
		puns string
		tags []string
}

func main() {
	bot, err := tgbotapi.NewBotAPI("323650569:AAH0miXDpgJJQoFOJ9Mr2HQ8QMLP281Iq1w")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		jokes := randomizeJokes()

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, jokes)

		bot.Send(msg)
	}
}

func fetchJokes() ([]Joke) {
  f, err := os.Open("jokes-bapak-bapak.csv")
  if err != nil {
      // return nil, err
  }
  defer f.Close()

  csvr := csv.NewReader(f)

  var result []Joke
  for {
    row, err := csvr.Read()
    if err != nil {
        if err == io.EOF {
            err = nil
						return result
        }
    }
		if row != nil {
			joke := Joke{id: row[0], puns: row[1], tags: strings.Split(row[2], ",")}
			result = append(result, joke)
		}

  }
	return result
}

func randomizeJokes() (string){
	jokes := fetchJokes()
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	return jokes[r.Intn(len(jokes)-1)].puns
}
