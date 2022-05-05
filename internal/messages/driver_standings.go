package messages

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-ergast-api-go/pkg/ergast"
	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
)

func MakeDriverStandings(standings []ergast.DriverStanding, leader *f1fantasy.Player) (*slack.MsgOption, error) {
	headerText := slack.NewTextBlockObject(MARKDOWN, "*Driver Standings* :checkered_flag:\n", false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	podiumDrivers := ":first_place_medal: " + makeDriverName(standings[0]) + "\n"
	podiumDrivers += ":second_place_medal: " + makeDriverName(standings[1]) + "\n"
	podiumDrivers += ":third_place_medal: " + makeDriverName(standings[2]) + "\n"

	podiumPoints := makeDriverScore(standings[0]) + "\n"
	podiumPoints += makeDriverScore(standings[1]) + "\n"
	podiumPoints += makeDriverScore(standings[2]) + "\n"

	podiumTitleText := slack.NewTextBlockObject(MARKDOWN, "Top 3 Drivers :medal:", false, false)
	podiumDriversText := slack.NewTextBlockObject(MARKDOWN, podiumDrivers, false, false)
	podiumPointsText := slack.NewTextBlockObject(MARKDOWN, podiumPoints, false, false)

	podiumFields := make([]*slack.TextBlockObject, 0)
	podiumFields = append(podiumFields, podiumDriversText)
	podiumFields = append(podiumFields, podiumPointsText)

	var podiumSection *slack.SectionBlock
	if leader != nil {
		firstPlaceDude := slack.NewImageBlockElement(leader.HeadshotImages.Profile, "picture of first place duder")
		podiumSection = slack.NewSectionBlock(podiumTitleText, podiumFields, slack.NewAccessory(firstPlaceDude))
	} else {
		podiumSection = slack.NewSectionBlock(podiumTitleText, podiumFields, nil)
	}

	drivers := ""
	points := ""
	for index, entry := range standings[3:] {
		drivers += fmt.Sprintf("%d %s\n", index+4, makeDriverName(entry))
		points += makeDriverScore(entry) + "\n"
	}
	driversText := slack.NewTextBlockObject(MARKDOWN, drivers, false, false)
	pointsText := slack.NewTextBlockObject(MARKDOWN, points, false, false)

	fields := make([]*slack.TextBlockObject, 0)
	fields = append(fields, driversText)
	fields = append(fields, pointsText)

	theRestSection := slack.NewSectionBlock(nil, fields, nil)

	message := slack.MsgOptionBlocks(headerSection, slack.NewDividerBlock(), podiumSection, slack.NewDividerBlock(), theRestSection)
	return &message, nil
}

func makeDriverName(driver ergast.DriverStanding) string {
	countryIso := NationalityToCountryIso(driver.DriverInfo.Nationality)
	flagEmoji := ""
	if len(countryIso) > 0 {
		flagEmoji = MakeFlagEmoji(countryIso)
	}

	return driver.DriverInfo.FirstName + " " + driver.DriverInfo.LastName + " " + flagEmoji
}

func makeDriverScore(driver ergast.DriverStanding) string {
	return driver.Points
}
