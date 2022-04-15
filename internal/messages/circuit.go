package messages

import (
	"fmt"
	"time"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
)

func MakeCircuit(circuit *f1fantasy.CircuitInfo) (*slack.MsgOption, error) {
	raceWeekend := fmt.Sprintf("%d-%02d-%02d", circuit.StartDay.Year(), circuit.StartDay.Month(),
		circuit.StartDay.Day())
	headerText := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*%s* %s\n\n*Race Weekend:* %s", circuit.Name,
		MakeFlagEmoji(circuit.CountryIso), raceWeekend), false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	imageText := slack.NewTextBlockObject(PLAIN_TEXT, circuit.ShortName, false, false)
	imageSection := slack.NewImageBlock(*circuit.CircuitImage.Url, "test", "image-block", imageText)

	practice1Timestamp, err := makeTimeString(circuit.StartDay, circuit.Practice1, circuit.GmtOffset)
	if err != nil {
		return nil, err
	}
	practice2Timestamp, err := makeTimeString(circuit.StartDay, circuit.Practice2, circuit.GmtOffset)
	if err != nil {
		return nil, err
	}
	practice3Timestamp, err := makeTimeString(circuit.StartDay, circuit.Practice3, circuit.GmtOffset)
	if err != nil {
		return nil, err
	}
	qualifyingTimestamp, err := makeTimeString(circuit.StartDay, circuit.Qualifying, circuit.GmtOffset)
	if err != nil {
		return nil, err
	}
	raceTimestamp, err := makeTimeString(circuit.StartDay, circuit.Race, circuit.GmtOffset)
	if err != nil {
		return nil, err
	}

	practice1 := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Practice 1:*\n%s", practice1Timestamp), false, false)
	practice2 := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Practice 2:*\n%s", practice2Timestamp), false, false)
	practice3 := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Practice 3:*\n%s", practice3Timestamp), false, false)
	qualifying := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Qualifying:*\n%s", qualifyingTimestamp), false, false)
	race := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Race:*\n%s", raceTimestamp), false, false)
	totalLaps := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Laps:*\n%s", circuit.TotalLaps), false, false)
	lapDistance := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Lap Distance:*\n%s km", circuit.Length), false, false)
	lapRecord := slack.NewTextBlockObject(MARKDOWN, fmt.Sprintf("*Lap Record:*\n%s", circuit.LapRecord), false, false)

	fieldSlice := make([]*slack.TextBlockObject, 0)
	fieldSlice = append(fieldSlice, practice1)
	fieldSlice = append(fieldSlice, practice2)
	fieldSlice = append(fieldSlice, practice3)
	fieldSlice = append(fieldSlice, qualifying)
	fieldSlice = append(fieldSlice, race)
	fieldSlice = append(fieldSlice, totalLaps)
	fieldSlice = append(fieldSlice, lapDistance)
	fieldSlice = append(fieldSlice, lapRecord)

	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

	message := slack.MsgOptionBlocks(headerSection, slack.NewDividerBlock(), imageSection, slack.NewDividerBlock(), fieldsSection)
	return &message, nil
}

func makeTimeString(raceStart time.Time, eventTime string, utcOffsetStr string) (string, error) {
	utcTime, err := makeUtcTime(raceStart, eventTime, utcOffsetStr)
	if err != nil {
		return "", err
	}

	est, err := time.LoadLocation("EST")
	if err != nil {
		return "", err
	}

	cet, err := time.LoadLocation("CET")
	if err != nil {
		return "", err
	}

	timestamp := makeTimestamp(utcTime.In(est), "EST") + "\n" + makeTimestamp(utcTime.In(cet), "CET")
	return timestamp, nil
}

func makeUtcTime(raceStart time.Time, eventTime string, utcOffsetStr string) (*time.Time, error) {
	var offsetHours int
	var offsetMinutes int
	_, err := fmt.Sscanf(utcOffsetStr, "%d:%d", &offsetHours, &offsetMinutes)
	if err != nil {
		return nil, err
	}
	// todo(vbonduro): Not sure why, but the offsets seem to be off by an hour :shrug:
	utcOffset := (offsetHours+1)*3600 + offsetMinutes*60

	var timeDay string
	var timeHour int
	var timeMinutes int
	_, err = fmt.Sscanf(eventTime, "%s %d:%d", &timeDay, &timeHour, &timeMinutes)
	if err != nil {
		return nil, err
	}

	switch timeDay {
	case "Fri":
		raceStart = raceStart.AddDate(0, 0, -1)
	case "Sun":
		raceStart = raceStart.AddDate(0, 0, 1)
	}

	loc := time.FixedZone("dangerzone", utcOffset)
	start := time.Date(raceStart.Year(), raceStart.Month(), raceStart.Day(), timeHour, timeMinutes, 0, 0, loc)

	return &start, nil
}

func makeTimestamp(day time.Time, tz string) string {
	return day.Format("2006-01-02 15:04:05") + " " + tz
}
