package strings

import (
	"strings"

	"github.com/fayaz07/locator/core/src/models"
)

const (
	bracket_open  = "("
	bracket_close = ")"
	specialChars  = "[]{}()<>?/\\|!@#$%^&*_-+=~`.:;,'\""
)

func SanitizePlaceName(place models.LocationModel) string {
	optionalName := GetStringInBrackets(place.PlaceName)
	if optionalName != "" && (place.City == optionalName ||
		place.Block == optionalName ||
		place.State == optionalName) {
		// trim optional name from place name
		place.PlaceName = TrimStringInBrackets(place.PlaceName)
	}
	place.PlaceName = removeSpecialChars(place.PlaceName)
	place.PlaceName = strings.Replace(place.PlaceName, "  ", " ", -1)
	return strings.TrimSpace(place.PlaceName)
}

func removeSpecialChars(s string) string {
	s = strings.Replace(s, ".", " ", -1)
	s = strings.Replace(s, "-", " ", -1)
	for _, c := range specialChars {
		s = strings.Replace(s, string(c), "", -1)
	}
	return strings.TrimSpace(s)
}

func GetStringInBrackets(s string) string {
	index := strings.Index(s, bracket_open)
	if index != -1 {
		s = strings.Split(s, bracket_open)[1]
		s = strings.Split(s, bracket_close)[0]
		return s
	}
	return ""
}

func TrimStringInBrackets(s string) string {
	index := strings.Index(s, bracket_open)
	if index != -1 {
		endIndex := strings.Index(s, bracket_close)
		if endIndex != -1 {
			toRemove := s[index : endIndex+1]
			s = strings.Replace(s, toRemove, "", -1)
		}
		return strings.TrimSpace(s)
	}
	return strings.TrimSpace(s)
}

func ConvertToAscii(s string) int {
	ascii := 0
	s = strings.ToValidUTF8(strings.ToUpper(s), "")
	for _, c := range s {
		ascii += int(c)
	}
	return ascii
}
