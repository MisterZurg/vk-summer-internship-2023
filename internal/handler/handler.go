package handler

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"

	"vk-summer-internship-2023/internal/service"
)

type Handler struct {
	logger  *zerolog.Logger
	service Service
}

type Service interface {
	Do(service.Params, time.Time) (string, error)
}

func New(logger *zerolog.Logger, srv *service.Service) *Handler {
	return &Handler{
		logger:  logger,
		service: srv,
	}
}

const (
	CODE_WARS = iota + 1
	LEET_CODE
	CODE_FORCES
)

var askUsername, askProblemTitle bool

var interviewPlatform int

var username, title string

func (h *Handler) HandleUpdate(bot *tgbotapi.BotAPI, upd tgbotapi.Update) {
	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, upd.Message.Text)
	if askUsername {
		username = upd.Message.Text
		msg = tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf("UPD: You've inputed %s", username))
		askUsername = false

		switch interviewPlatform {
		case CODE_WARS:
			info, _ := h.service.Do(service.Params{Cmd: "cwUser", Arg: username}, time.Now())
			msg = tgbotapi.NewMessage(upd.Message.Chat.ID, info)
		case LEET_CODE:
			info, _ := h.service.Do(service.Params{Cmd: "lcUser", Arg: username}, time.Now())
			msg = tgbotapi.NewMessage(upd.Message.Chat.ID, info)
		case CODE_FORCES:
			info, _ := h.service.Do(service.Params{Cmd: "cfUser", Arg: username}, time.Now())
			msg = tgbotapi.NewMessage(upd.Message.Chat.ID, info)
		}
		bot.Send(msg)
		return
	}

	if askProblemTitle {
		title = upd.Message.Text
		msg = tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf("UPD: You've inputed %s", username))
		askProblemTitle = false

		switch interviewPlatform {
		case CODE_WARS:
			info, _ := h.service.Do(service.Params{Cmd: "cwUser", Arg: username}, time.Now())
			msg = tgbotapi.NewMessage(upd.Message.Chat.ID, info)
		}
	}

	switch upd.Message.Text {
	case "start":
		msg = tgbotapi.NewMessage(upd.Message.Chat.ID, service.Greet)
		msg.ReplyMarkup = service.MainMenu

	default:
		msg = tgbotapi.NewMessage(upd.Message.Chat.ID, "I don't know that command, try input `start` to start worling")
	}
	bot.Send(msg)
}

func (h *Handler) HandleCallBack(bot *tgbotapi.BotAPI, cb tgbotapi.CallbackQuery) {
	msg := tgbotapi.NewMessage(cb.Message.Chat.ID, cb.Data)
	switch cb.Data {

	// Front Buttons
	case "about":
		msg = tgbotapi.NewMessage(cb.Message.Chat.ID, service.AboutText)
		msg.ReplyMarkup = service.AboutText
	case "codewars":
		msg.ReplyMarkup = service.CodeWarsMenu
	case "leetcode":
		msg.ReplyMarkup = service.LeetCodeMenu
	case "codeforces":
		msg.ReplyMarkup = service.CodeForcesMenu
	// Nested Buttons
	case "cwUser":
		interviewPlatform = CODE_WARS
		askUsername = true
		msg = tgbotapi.NewMessage(cb.Message.Chat.ID, "Enter CodeWars Username")

	case "cwProblem":
		interviewPlatform = CODE_WARS
		askProblemTitle = true
		msg = tgbotapi.NewMessage(cb.Message.Chat.ID, "Enter CodeWars problem")

	case "lcUser":
		interviewPlatform = LEET_CODE
		askUsername = true
		msg = tgbotapi.NewMessage(cb.Message.Chat.ID, "Enter LeetCode Username")

	case "cfUser":
		interviewPlatform = CODE_FORCES
		askUsername = true
		msg = tgbotapi.NewMessage(cb.Message.Chat.ID, "Enter CodeForces Handle")

	case "main":
		msg.ReplyMarkup = service.MainMenu
	}

	bot.Send(msg)
}
