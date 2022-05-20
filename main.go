package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
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

func slashCommandHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

// todo: refactor and make nicer
func handleBlockAction(action *slack.BlockAction, user slack.User, channel slack.Channel) {
	if strings.HasPrefix(action.ActionID, "/f1") {
		command := slackutil.SlashCommand{}
		command.Command = "/f1"
		command.Text = action.ActionID[4:]
		command.ChannelID = channel.ID
		command.UserID = user.ID
		err := CommandHandler.Handle(command)
		if err != nil {
			panic(err)
		}
	}
	log.Printf("Block action: " + action.ActionID)
}

func modalEventHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	uri, _ := url.ParseQuery(request.Body)
	payload := uri.Get("payload")

	var action slack.InteractionCallback
	err := json.Unmarshal([]byte(payload), &action)
	if err != nil {
		log.Printf("Could not parse action response JSON: %v", err)
	}

	if action.ActionCallback.BlockActions != nil {
		for _, blockAction := range action.ActionCallback.BlockActions {
			handleBlockAction(blockAction, action.User, action.Channel)
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.Path == "/slash" {
		return slashCommandHandler(request)
	} else if request.Path == "/modal" {
		return modalEventHandler(request)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
	}, nil
}

func main() {
	lambda.Start(handler)
}
