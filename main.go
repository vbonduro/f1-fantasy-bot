package main

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vbonduro/f1-fantasy-bot/internal/commands"
	"github.com/vbonduro/f1-fantasy-bot/internal/slackutil"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	command := slackutil.MakeSlashCommand(request.Body)

	// todo(vbonduro): Should use slack secret instead, but not sure how just yet :(
	if command.Token != os.Getenv("SLACK_VERIFICATION_TOKEN") {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid token.",
			StatusCode: http.StatusOK,
		}, nil
	}

	if command.Command != "/f1" {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid command.",
			StatusCode: http.StatusOK,
		}, nil
	}

	switch command.Text {
	case "fantasy leaderboard":
		return commands.FantasyLeaderboard(command), nil
	}

	steiner := []string{"HE DOES NOT FOK SMASH MY DOOR!",
		"WE LOOK LIKE A BUNCH OF VANKERS!",
		"THIS IS WHY EVERYONE HATES YOU!",
		"WE COULD HAVE LOOKED LIKE ROCKSTARS!"}
	return events.APIGatewayProxyResponse{
		Body:       steiner[rand.Intn(len(steiner))],
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handler)
}
