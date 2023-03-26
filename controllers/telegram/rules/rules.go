package rules

import (
	tele "gopkg.in/telebot.v3"
)

func Reply(user *tele.User, message string) string {
	return "Username : " + user.Username
}
