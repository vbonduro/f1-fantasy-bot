package commands

import (
	"log"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-ergast-api-go/pkg/ergast"
	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
	"github.com/vbonduro/f1-fantasy-bot/internal/messages"
)

func (h *Handler) nextRace() error {
	race, err := ergast.NextRace()
	if err != nil {
		log.Printf("%s\n", err.Error())
		return err
	}

	f1FantasyApi := f1fantasy.NewApi()
	circuits, err := f1FantasyApi.GetCircuits()
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	message, err := messages.MakeCircuit(race, &circuits[race.RoundNumber-1].Info)
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
