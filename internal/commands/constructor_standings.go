package commands

import (
	"errors"
	"log"
	"strings"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-ergast-api-go/pkg/ergast"
	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
	"github.com/vbonduro/f1-fantasy-bot/internal/messages"
)

func (h *Handler) constructorStandings() error {
	standings, err := ergast.ConstructorStandings()
	if err != nil {
		log.Printf("%s\n", err.Error())
		return err
	}

	f1FantasyApi := f1fantasy.NewApi()
	players, err := f1FantasyApi.GetPlayers()
	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	leader, err := findConstructor(standings[0], players.PlayerList)
	if err != nil {
		log.Printf("%s", err.Error())
	}

	message, err := messages.MakeConstructorStandings(standings, leader)
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

func findConstructor(constructor ergast.ConstructorStanding, players []f1fantasy.Player) (*f1fantasy.Player, error) {
	for _, player := range players {
		if strings.ToLower(player.FirstName) == strings.ToLower(constructor.ConstructorInfo.Name) {
			return &player, nil
		}
	}
	return nil, errors.New("Unable to find constructor " + constructor.ConstructorInfo.Name)
}
