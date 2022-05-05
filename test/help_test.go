package test

import (
	"os"
	"testing"

	"github.com/vbonduro/f1-fantasy-bot/internal/commands"
	"github.com/vbonduro/f1-fantasy-bot/internal/slackutil"
)

func TestHelp(t *testing.T) {
	handler := commands.Handler{}
	err := handler.Init()
	if err != nil {
		panic(err)
	}
	command := slackutil.SlashCommand{}
	command.Command = "/f1"
	command.ChannelID = os.Getenv("CHANNEL_ID")
	command.UserID = os.Getenv("USER_ID")
	err = handler.Handle(command)
	if err != nil {
		panic(err)
	}
}
