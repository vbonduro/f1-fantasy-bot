package commands

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
	"github.com/vbonduro/f1-fantasy-bot/internal/slackutil"
)

func FantasyLeaderboard(command slackutil.SlashCommand) events.APIGatewayProxyResponse {
	leagueId, _ := strconv.Atoi(os.Getenv("F1_LEAGUE"))
	user := os.Getenv("F1_USER")
	password := os.Getenv("F1_PASSWORD")
	api, err := f1fantasy.NewAuthenticatedApi(user, password)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "F1 Fantasy Authentication Failed!",
			StatusCode: http.StatusOK,
		}
	}
	leaderboard, err := api.GetLeagueLeaderboard(leagueId)
	if err != nil {
		log.Printf("%s", err.Error())
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Failed to get leaderboard: %s", err.Error()),
			StatusCode: http.StatusOK,
		}
	}

	slackApi := slack.New(os.Getenv("SLACK_OAUTH"))
	_, _, err = slackApi.PostMessage(
		command.ChannelID,
		makeLeaderboardMessage(leaderboard),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Printf("SLack post failed: %s", err.Error())
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}
}

func makeLeaderboardMessage(leaderboard *f1fantasy.LeagueLeaderboard) slack.MsgOption {
	headerText := slack.NewTextBlockObject("mrkdwn", "*Fantasy Leaderboard*", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	podiumString := "*Podium* :trophy:\n\n"
	podiumString += makePlacementString(":first_place_medal:", leaderboard.Leaderboard.Entries[0])
	podiumString += makePlacementString(":second_place_medal:", leaderboard.Leaderboard.Entries[1])
	podiumString += makePlacementString(":third_place_medal:", leaderboard.Leaderboard.Entries[2])
	podiumText := slack.NewTextBlockObject("mrkdwn", podiumString, false, false)
	podiumSection := slack.NewSectionBlock(podiumText, nil, nil)

	theRestString := ""
	for index, entry := range leaderboard.Leaderboard.Entries[3:] {
		theRestString += makePlacementString(fmt.Sprintf("%02d", index+4), entry)
	}
	theRestText := slack.NewTextBlockObject("mrkdwn", theRestString, false, false)
	theRestSection := slack.NewSectionBlock(theRestText, nil, nil)

	return slack.MsgOptionBlocks(headerSection, slack.NewDividerBlock(), podiumSection, slack.NewDividerBlock(), theRestSection)
}

func makePlacementString(placement string, entry f1fantasy.LeaderboardEntry) string {
	return fmt.Sprintf("%s: `%-30s %-30s [%.2f]` :flag-%s:\n", placement, entry.TeamName, entry.UserName, entry.Score, entry.Country)
}
