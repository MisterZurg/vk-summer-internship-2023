package main

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"

	"vk-summer-internship-2023/internal/config"
	"vk-summer-internship-2023/internal/handler"
	"vk-summer-internship-2023/internal/service"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := config.New()
	if err != nil {
		logger.Fatal().Err(err).Msg("Configuration error")
	}

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPIToken)
	if err != nil {
		logger.Fatal().Err(err).Msg(fmt.Sprintf("%w", err))
	}

	bot.Debug = true

	logger.Info().Msg(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	srv := service.New(&logger)
	h := handler.New(&logger, srv)

	for update := range updates {
		if update.Message != nil {
			h.HandleUpdate(bot, update)
		} else if update.CallbackQuery != nil {
			callback := update.CallbackQuery
			h.HandleCallBack(bot, *callback)
		}
	}
}
