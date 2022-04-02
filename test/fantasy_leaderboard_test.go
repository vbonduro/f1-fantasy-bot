package test

import (
	"os"
	"testing"

	"github.com/vbonduro/f1-fantasy-bot/internal/commands"
	"github.com/vbonduro/f1-fantasy-bot/internal/slackutil"
)

func TestFantasyLeaderboard(t *testing.T) {
	command := slackutil.SlashCommand{}
	command.ChannelID = os.Getenv("CHANNEL_ID")
	commands.FantasyLeaderboard(command)
}
