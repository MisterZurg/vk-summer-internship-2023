package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const Greet = "Pick resource you want"

const AboutText = "This is product task for an VK internship as Go-developer."

var MainMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("CodeWars", "codewars")),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("LeetCode", "leetcode")),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("CodeForces", "codeforces")),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("About", "about")),
)

var AboutMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("View source", "source")),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Back", "main")),
)

var CodeWarsMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Show User", "cwUser"),
		tgbotapi.NewInlineKeyboardButtonData("Show Problem", "cwProblem"),
	),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Back", "main")),
)

var LeetCodeMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Show User", "lcUser"),
		tgbotapi.NewInlineKeyboardButtonData("Show Problem", "lcProblem"),
	),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Back", "main")),
)

var CodeForcesMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Show User", "cfUser")),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Back", "main")),
)
