package messages

import "strings"

const (
	MARKDOWN   = "mrkdwn"
	PLAIN_TEXT = "plain_text"
)

func MakeFlagEmoji(countryIso string) string {
	flag := ""
	if len(countryIso) > 0 {
		flag = ":flag-" + strings.ToLower(countryIso) + ":"
	}
	return flag
}
