package messages

import (
	"fmt"
	"strconv"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
)

func MakeDriverStats(drivers []f1fantasy.Player, startingIndex int) (*slack.MsgOption, error) {
	const DRIVERS_PER_MESSAGE = 3

	endingIndex := Min(startingIndex+DRIVERS_PER_MESSAGE, len(drivers)-1)
	driversInMessage := drivers[startingIndex:endingIndex]

	var sections [DRIVERS_PER_MESSAGE]*slack.SectionBlock
	for index, driver := range driversInMessage {
		driverHeadshot := slack.NewImageBlockElement(driver.HeadshotImages.Profile, "quicky mcfastface")

		wins := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Wins:*\n%d", driver.DriverStats.Wins), false, false)
		podiums := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Podiums:*\n%d", driver.DriverStats.Podiums), false, false)
		poles := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Poles:*\n%d", driver.DriverStats.Poles), false, false)
		fastestLaps := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Fastest Laps:*\n%d", driver.DriverStats.FastestLaps), false, false)
		grandsPrix := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Grand Prix Entered:*\n%d", driver.DriverStats.TotalGrandPrix), false, false)
		championshipTitles := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*World Championships:*\n%d", driver.DriverStats.Titles), false, false)
		bestFinishes := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Best Finishes:*\n%s", driver.DriverStats.HighestRaceFinished), false, false)

		fieldSlice := make([]*slack.TextBlockObject, 0)
		fieldSlice = append(fieldSlice, wins)
		fieldSlice = append(fieldSlice, podiums)
		fieldSlice = append(fieldSlice, poles)
		fieldSlice = append(fieldSlice, fastestLaps)
		fieldSlice = append(fieldSlice, grandsPrix)
		fieldSlice = append(fieldSlice, championshipTitles)
		fieldSlice = append(fieldSlice, bestFinishes)

		driverName := slack.NewTextBlockObject(MARKDOWN, "*"+driver.FirstName+" "+driver.LastName+"*  "+MakeFlagEmoji(*driver.CountryIso), false, false)
		sections[index] = slack.NewSectionBlock(driverName, fieldSlice, slack.NewAccessory(driverHeadshot))
	}

	nextButtonText := slack.NewTextBlockObject(PLAIN_TEXT, "Next", false, false)
	nextButton := slack.NewButtonBlockElement("driver stats "+strconv.Itoa(endingIndex), "", nextButtonText)

	var buttonBlock *slack.ActionBlock
	buttonBlock = slack.NewActionBlock("", nextButton)
	if startingIndex == 0 {
		buttonBlock = slack.NewActionBlock("", nextButton)
	} else if endingIndex != len(drivers)-1 {
		prevButtonText := slack.NewTextBlockObject(PLAIN_TEXT, "Previous", false, false)
		prevButton := slack.NewButtonBlockElement("driver stats "+strconv.Itoa(startingIndex-DRIVERS_PER_MESSAGE), "", prevButtonText)
		buttonBlock = slack.NewActionBlock("", prevButton, nextButton)
	}

	// todo(vbonduro): It would be great if this could be appended :(
	//message := slack.MsgOptionBlocks(sections[0], slack.NewDividerBlock(), sections[1], slack.NewDividerBlock(), sections[2], slack.NewDividerBlock(), sections[3], slack.NewDividerBlock(), sections[4], slack.NewDividerBlock(), buttonBlock)
	message := slack.MsgOptionBlocks(sections[0], slack.NewDividerBlock(), sections[1], slack.NewDividerBlock(), sections[2], slack.NewDividerBlock(), buttonBlock)
	return &message, nil
}
