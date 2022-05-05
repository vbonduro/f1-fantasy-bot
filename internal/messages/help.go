package messages

import (
	"github.com/slack-go/slack"
)

func MakeHelp() slack.MsgOption {
	headerText := slack.NewTextBlockObject(MARKDOWN, "Here's all of the cool stuff you can do with this app:", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	runCommandButtonText := slack.NewTextBlockObject(PLAIN_TEXT, "Run", false, false)

	fantasyLeaderboardText := slack.NewTextBlockObject(MARKDOWN,
		"*:trophy: F1 Fantasy Leaderboard*\nPost the current F1 Fantasy leaderboard to the channel", false, false)
	fantasyLeaderboardButton := slack.NewButtonBlockElement("/f1 fantasy leaderboard", "", runCommandButtonText)
	fantasyLeaderboardSection := slack.NewSectionBlock(fantasyLeaderboardText, nil, slack.NewAccessory(fantasyLeaderboardButton))

	nextRaceText := slack.NewTextBlockObject(MARKDOWN,
		"*:motorway: Next Race*\nPost information about the next race to the channel", false, false)
	nextRaceButton := slack.NewButtonBlockElement("/f1 next race", "", runCommandButtonText)
	nextRaceSection := slack.NewSectionBlock(nextRaceText, nil, slack.NewAccessory(nextRaceButton))

	driverStandingsText := slack.NewTextBlockObject(MARKDOWN,
		"*:checkered_flag: F1 Driver Standings*\nMessage me the driver standings for the season", false, false)
	driverStandingsButton := slack.NewButtonBlockElement("/f1 driver standings", "", runCommandButtonText)
	driverStandingsSection := slack.NewSectionBlock(driverStandingsText, nil, slack.NewAccessory(driverStandingsButton))

	constructorStandingsText := slack.NewTextBlockObject(MARKDOWN,
		"*:racing_car: F1 Constructor Standings*\nMessage me the constructor standings for the season", false, false)
	constructorStandingsButton := slack.NewButtonBlockElement("/f1 constructor standings", "", runCommandButtonText)
	constructorStandingsSection := slack.NewSectionBlock(constructorStandingsText, nil, slack.NewAccessory(constructorStandingsButton))

	return slack.MsgOptionBlocks(headerSection, slack.NewDividerBlock(), fantasyLeaderboardSection, nextRaceSection, driverStandingsSection,
		constructorStandingsSection)
}
