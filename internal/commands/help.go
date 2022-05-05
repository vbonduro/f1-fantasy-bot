package commands

import (
	"log"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-fantasy-bot/internal/messages"
)

func (h *Handler) help() error {
	_, err := h.Slack.PostEphemeral(
		h.Command.ChannelID,
		h.Command.UserID,
		messages.MakeHelp(),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Printf("Slack post failed: %s", err.Error())
	}

	return err
}
