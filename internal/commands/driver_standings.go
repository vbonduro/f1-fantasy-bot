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

func (h *Handler) driverStandings() error {
	standings, err := ergast.DriverStandings()
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

	leader, err := findPlayer(standings[0], players.PlayerList)
	if err != nil {
		log.Printf("%s", err.Error())
	}

	message, err := messages.MakeDriverStandings(standings, leader)
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

func findPlayer(driver ergast.DriverStanding, players []f1fantasy.Player) (*f1fantasy.Player, error) {
	for _, player := range players {
		if strings.ToLower(player.FirstName) == strings.ToLower(driver.DriverInfo.FirstName) &&
			strings.ToLower(player.LastName) == strings.ToLower(driver.DriverInfo.LastName) {
			return &player, nil
		}
	}
	return nil, errors.New("Unable to find driver " + driver.DriverInfo.FirstName + " " + driver.DriverInfo.LastName)
}
