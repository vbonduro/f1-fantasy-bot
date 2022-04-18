package messages

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-ergast-api-go/pkg/ergast"
	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
)

func MakeCircuit(race *ergast.Race, circuit *f1fantasy.CircuitInfo) (*slack.MsgOption, error) {
	headerText := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*%s* %s\n", circuit.Name,
		MakeFlagEmoji(circuit.CountryIso)), false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	imageText := slack.NewTextBlockObject(PLAIN_TEXT, circuit.ShortName, false, false)
	imageSection := slack.NewImageBlock(*circuit.CircuitImage.Url, "test", "image-block", imageText)

	practice1Timestamp, err := makeTimeString(race.FirstPracticeStart)
	if err != nil {
		return nil, err
	}
	sprintTimestamp, err := makeTimeString(race.SprintRaceStart)
	if err != nil {
		return nil, err
	}
	practice2Timestamp, err := makeTimeString(race.SecondPracticeStart)
	if err != nil {
		return nil, err
	}
	practice3Timestamp, err := makeTimeString(race.ThirdPracticeStart)
	if err != nil {
		return nil, err
	}
	qualifyingTimestamp, err := makeTimeString(race.QualifyingStart)
	if err != nil {
		return nil, err
	}
	raceTimestamp, err := makeTimeString(race.RaceStart)
	if err != nil {
		return nil, err
	}

	practice1 := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Practice 1:*\n%s", practice1Timestamp), false, false)
	practice2 := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Practice 2:*\n%s", practice2Timestamp), false, false)
	practice3 := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Practice 3:*\n%s", practice3Timestamp), false, false)
	sprint := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Sprint:*\n%s", sprintTimestamp), false, false)
	qualifying := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Qualifying:*\n%s", qualifyingTimestamp), false, false)
	raceSched := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Race:*\n%s", raceTimestamp), false, false)
	totalLaps := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Laps:*\n%s", circuit.TotalLaps), false, false)
	lapDistance := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Lap Distance:*\n%s km", circuit.Length), false, false)
	lapRecord := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Lap Record:*\n%s", circuit.LapRecord), false, false)

	fieldSlice := make([]*slack.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, practice1)
	if practice3Timestamp == "NA" {
		fieldSlice = append(fieldSlice, qualifying)
	}
	fieldSlice = append(fieldSlice, practice2)
	if practice3Timestamp == "NA" {
		fieldSlice = append(fieldSlice, sprint)
	} else {
		fieldSlice = append(fieldSlice, practice3)
		fieldSlice = append(fieldSlice, qualifying)
	}
	fieldSlice = append(fieldSlice, raceSched)
	fieldSlice = append(fieldSlice, totalLaps)
	fieldSlice = append(fieldSlice, lapDistance)
	fieldSlice = append(fieldSlice, lapRecord)

	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

	message := slack.MsgOptionBlocks(headerSection, slack.NewDividerBlock(), imageSection, slack.NewDividerBlock(), fieldsSection)
	return &message, nil
}

func makeTimeString(eventTime time.Time) (string, error) {
	emptyTime := time.Time{}
	if eventTime == emptyTime {
		return "NA", nil
	}

	est, err := time.LoadLocation("Canada/Eastern")
	if err != nil {
		return "", err
	}

	cet, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		return "", err
	}

	timestamp := makeTimestamp(eventTime.In(est), "EST") + "\n" + makeTimestamp(eventTime.In(cet), "CET")
	return timestamp, nil
}

func makeTimestamp(day time.Time, tz string) string {
	return day.Format("2006-01-02 15:04:05") + " " + tz
}
