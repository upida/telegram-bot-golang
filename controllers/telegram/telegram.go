package telegram

import (
	"bot/controllers/telegram/rules"
	"bot/helper/json"
	"bot/helper/logger"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func Telegram(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		logError, _ := json.Encode(err)
		logger.Write("ERROR [godotenv.Load] : " + logError)
	}

	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		logError, _ := json.Encode(err)
		logger.Write("ERROR [tele.NewBot] : " + logError)
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		var (
			user = c.Sender()
			text = c.Text()
		)

		logUser, err := json.Encode(user)
		if err != nil {
			logError, _ := json.Encode(err)
			logger.Write("ERROR [json.Encode(user)] : " + logError)
		}
		logger.Write("User : " + logUser)
		logger.Write("Text : " + text)

		message := rules.Reply(user, text)
		return c.Send(message)
	})

	b.Start()
}
