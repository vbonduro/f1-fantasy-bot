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

func NationalityToCountryIso(nationality string) string {
	var ISO_MAP = map[string]string{
		"Monegasque": "MC",
		"Italian":    "IT",
		"Dutch":      "NL",
		"Austrian":   "AT",
		"Mexican":    "MX",
		"British":    "GB",
		"German":     "DE",
		"Spanish":    "ES",
		"Finnish":    "FI",
		"Swiss":      "CH",
		"French":     "FR",
		"Danish":     "DK",
		"American":   "US",
		"Australian": "AU",
		"Japanese":   "JP",
		"Chinese":    "CN",
		"Thai":       "TH",
		"Canadian":   "CA",
	}
	return ISO_MAP[nationality]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
