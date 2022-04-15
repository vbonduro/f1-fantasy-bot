package messages

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/vbonduro/f1-fantasy-api-go/pkg/f1fantasy"
)

func MakeLeaderboard(leaderboard *f1fantasy.LeagueLeaderboard) slack.MsgOption {
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

	return slack.MsgOptionBlocks(headerSection, slack.NewDividerBlock(), podiumSection, slack.NewDividerBlock(),
		theRestSection)
}

func makePlacementString(placement string, entry f1fantasy.LeaderboardEntry) string {
	return fmt.Sprintf("%s: `%-30s %-30s [%.2f]` %s\n", placement, entry.TeamName, entry.UserName, entry.Score,
		MakeFlagEmoji(entry.Country))
}
