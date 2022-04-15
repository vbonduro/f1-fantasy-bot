package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/vbonduro/f1-fantasy-bot/internal/commands"
	"github.com/vbonduro/f1-fantasy-bot/internal/slackutil"
)

var CommandHandler commands.Handler

func init() {
	err := CommandHandler.Init()
	if err != nil {
		panic(err)
	}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	command := slackutil.MakeSlashCommand(request.Body)

	// todo(vbonduro): Should use slack secret instead, but not sure how just yet :(
	if command.Token != os.Getenv("SLACK_VERIFICATION_TOKEN") {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid token.",
			StatusCode: http.StatusOK,
		}, nil
	}

	err := CommandHandler.Handle(command)
	if err != nil {
		log.Printf("Command " + command.Text + " failed with error " + err.Error())
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error: %s", err.Error()),
			StatusCode: http.StatusOK,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handler)
}
