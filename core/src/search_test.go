package core

import (
	"testing"

	"github.com/fayaz07/locator/common/models"
	"github.com/stretchr/testify/assert"
)

func TestSearchInts(t *testing.T) {
	assert := assert.New(t)

	a := []models.AsciiIndexModel{
		{Code: 1, Index: 1, Length: 1},
	}
	assert.Equal(0, SearchOnNames(a, 1, 1))
}
