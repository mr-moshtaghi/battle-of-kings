package telegram

import (
	"battle-of-kings/internal/service"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
	"time"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot
}

func NewTelegram(app *service.App, apikey string) (*Telegram, error) {
	pref := telebot.Settings{
		Token:  apikey,
		Poller: &telebot.LongPoller{Timeout: 60 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Error("couldn't connect to telegram servers")
		return nil, err
	}
	t := &Telegram{bot: bot, App: app}
	t.setupHandlers()
	return t, nil
}

func (t *Telegram) Start() {
	t.bot.Start()
}
