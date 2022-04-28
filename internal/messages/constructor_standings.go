package messages

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-ergast-api-go/pkg/ergast"
	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
)

func MakeConstructorStandings(standings []ergast.ConstructorStanding, leader *f1fantasy.Player) (*slack.MsgOption, error) {
	headerText := slack.NewTextBlockObject(MARKDOWN, "*Constructor Standings* :racing_car:\n", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	var fastestConstructorImage *slack.ImageBlock
	if leader != nil {
		fastestConstructorImage = slack.NewImageBlock(*leader.ProfileImage.Url, "real fast car bud", "image-block", nil)
	}

	drivers := ":first_place_medal: " + makeConstructorName(standings[0]) + "\n"
	drivers += ":second_place_medal: " + makeConstructorName(standings[1]) + "\n"
	drivers += ":third_place_medal: " + makeConstructorName(standings[2]) + "\n"

	points := standings[0].Points + "\n"
	points += standings[1].Points + "\n"
	points += standings[2].Points + "\n"

	for index, entry := range standings[3:] {
		drivers += fmt.Sprintf("%d %s\n", index+4, makeConstructorName(entry))
		points += entry.Points + "\n"
	}
	driversText := slack.NewTextBlockObject(MARKDOWN, drivers, false, false)
	pointsText := slack.NewTextBlockObject(MARKDOWN, points, false, false)

	fields := make([]*slack.TextBlockObject, 0)
	fields = append(fields, driversText)
	fields = append(fields, pointsText)

	standingsSection := slack.NewSectionBlock(nil, fields, nil)

	var message slack.MsgOption
	if fastestConstructorImage != nil {
		message = slack.MsgOptionBlocks(headerSection, slack.NewDividerBlock(), fastestConstructorImage, slack.NewDividerBlock(), standingsSection)
	} else {
		message = slack.MsgOptionBlocks(headerSection, slack.NewDividerBlock(), standingsSection)
	}
	return &message, nil
}

func makeConstructorName(constructor ergast.ConstructorStanding) string {
	countryIso := NationalityToCountryIso(constructor.ConstructorInfo.Nationality)
	flagEmoji := ""
	if len(countryIso) > 0 {
		flagEmoji = MakeFlagEmoji(countryIso)
	}

	return constructor.ConstructorInfo.Name + " " + flagEmoji
}
