package commands

import (
	"log"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
	"github.com/vbonduro/f1-fantasy-bot/internal/messages"
)

func (h *Handler) driverStats(startingIndex int) error {
	f1FantasyApi := f1fantasy.NewApi()
	players, err := f1FantasyApi.GetPlayers()
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	drivers := filterDrivers(players.PlayerList)

	message, err := messages.MakeDriverStats(drivers, startingIndex)
	if err != nil {
		return err
	}

	_, err = h.Slack.PostEphemeral(
		h.Command.ChannelID,
		h.Command.UserID,
		*message,
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Printf("Slack post failed: %s", err.Error())
	}

	return err
}

func filterDrivers(players []f1fantasy.Player) []f1fantasy.Player {
	var drivers []f1fantasy.Player

	for _, player := range players {
		if !player.IsConstructor {
			drivers = append(drivers, player)
		}
	}

	return drivers
}
