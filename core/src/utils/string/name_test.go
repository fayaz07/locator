package strings

import (
	"testing"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/stretchr/testify/assert"
)

func TestSanitizePlaceName(t *testing.T) {
	assert := assert.New(t)

	// IN,382016,(Gandhinagar) Sector 16,Gujarat,Gandhi Nagar,Gandhinagar,23.0976,72.8913,1,72.8913,23.0976
	place := models.LocationModel{
		CountryCode: "IN",
		PostalCode:  "382016",
		PlaceName:   "(Gandhinagar) Sector 16",
		State:       "Gujarat",
		City:        "Gandhi Nagar",
		Block:       "Gandhinagar",
		Latitude:    23.0976,
		Longitude:   72.8913,
		Accuracy:    1,
		Coordinates: models.CoordinatesModel{
			Lon: 72.8913,
			Lat: 23.0976,
		},
	}

	assert.Equal("Sector 16", SanitizePlaceName(place))
}

func TestGetStringInBrackets(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Gandhinagar", GetStringInBrackets("(Gandhinagar) Sector 16"))
	assert.Equal("Gandhinagar", GetStringInBrackets("Sector 16 (Gandhinagar)"))
	assert.Equal("Gandhinagar", GetStringInBrackets("Sector 16 (Gandhinagar) Road 2"))
	assert.Equal("Gandhinagar", GetStringInBrackets("Sector 16 (Gandhinagar) Road 2 (Gandhinagar)"))
}

func TestTrimStringInBrackets(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Sector 16", TrimStringInBrackets("(Gandhinagar) Sector 16"))
	assert.Equal("Sector 16", TrimStringInBrackets("Sector 16 (Gandhinagar)"))
	assert.Equal("Sector 16 Road 2", TrimStringInBrackets("Sector 16 Road 2 (Gandhinagar)"))
	assert.Equal("Sector 16 Road 2", TrimStringInBrackets("Sector 16 Road 2 (Gandhinagar) (Gandhinagar)"))
}

func TestRemoveSpecialChars(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Sector 16", removeSpecialChars("Sector 16"))
	assert.Equal("Sector 16 Road 2", removeSpecialChars("Sector 16 Road 2"))
	assert.Equal("Sector 16 Road 2", removeSpecialChars("Sector 16 Road 2 \"/;:.,<>?[]{}\\|!@#$%^&*()_+-='"))
	assert.Equal("Sector 16 Road 2", removeSpecialChars("Sector.16 Road 2 \"/;:.,<>?[]{}\\|!@#$%^&*()_+-='"))
	assert.Equal("Sector 16 Road 2", removeSpecialChars("Sector-16 Road 2 \"/;:.,<>?[]{}\\|!@#$%^&*()_+-='"))
}
