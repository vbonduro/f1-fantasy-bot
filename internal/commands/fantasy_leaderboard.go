package commands

import (
	"log"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-fantasy-bot/internal/messages"
)

func (h *Handler) fantasyLeaderboard() error {
	leaderboard, err := h.F1.GetLeagueLeaderboard(h.LeagueId)
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}
	_, _, err = h.Slack.PostMessage(
		h.Command.ChannelID,
		messages.MakeLeaderboard(leaderboard),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Printf("Slack post failed: %s", err.Error())
	}

	return err
}
