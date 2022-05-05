package commands

import (
	"errors"
	"os"
	"strconv"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-fantasy-bot/internal/slackutil"
)

const (
	F1_USER     = "F1_USER"
	F1_PASSWORD = "F1_PASSWORD"
	SLACK_OAUTH = "SLACK_OAUTH"
	F1_LEAGUE   = "F1_LEAGUE"
)

type Handler struct {
	//F1       *f1fantasy.AuthenticatedApi
	Slack    *slack.Client
	Command  slackutil.SlashCommand
	LeagueId int
}

func (h *Handler) Init() error {
	h.Slack = slack.New(os.Getenv(SLACK_OAUTH))
	// todo: Authenticated API currently broken =(
	// f1, err := f1fantasy.NewAuthenticatedApi(os.Getenv(F1_USER), os.Getenv(F1_PASSWORD))
	// if err != nil {
	// 	return err
	// }
	h.LeagueId, _ = strconv.Atoi(os.Getenv(F1_LEAGUE))
	return nil
}

func (h *Handler) Handle(command slackutil.SlashCommand) error {
	if command.Command != "/f1" {
		return errors.New("Invalid Command: " + command.Command)
	}

	// if h.F1.Expired() {
	// 	log.Printf("F1 Session Expired! Renew...")
	// 	f1, err := f1fantasy.NewAuthenticatedApi(os.Getenv(F1_USER), os.Getenv(F1_PASSWORD))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	h.F1 = f1
	// }
	h.Command = command

	if len(command.Text) == 0 {
		return h.help()
	}

	switch command.Text {
	case "fantasy leaderboard":
		return h.fantasyLeaderboard()
	case "next race":
		return h.nextRace()
	case "driver standings":
		return h.driverStandings()
	case "constructor standings":
		return h.constructorStandings()
	}

	return errors.New("Invalid Command: " + command.Text)
}
