package commands

import (
	"log"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-fantasy-bot/internal/messages"
)

func (h *Handler) nextRace() error {
	circuit, err := h.F1.CurrentCircuit()
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	message, err := messages.MakeCircuit(circuit)
	if err != nil {
		return err
	}

	_, _, err = h.Slack.PostMessage(
		h.Command.ChannelID,
		*message,
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Printf("Slack post failed: %s", err.Error())
	}

	return err
}
