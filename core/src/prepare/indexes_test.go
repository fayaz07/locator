package prepare

import (
	"testing"

	"github.com/fayaz07/locator/core/src/models"
	"github.com/stretchr/testify/assert"
)

func TestGenSearchIndexes(t *testing.T) {
	assert := assert.New(t)

	asciiIndexes := []models.AsciiIndexModel{
		{Code: 99, Index: 1, Length: 2},
		{Code: 100, Index: 1, Length: 2},
		{Code: 100, Index: 4, Length: 5},
		{Code: 101, Index: 6, Length: 7},
		{Code: 101, Index: 8, Length: 9},
		{Code: 101, Index: 10, Length: 11},
		{Code: 102, Index: 12, Length: 13},
		{Code: 102, Index: 14, Length: 15},
		{Code: 102, Index: 16, Length: 17},
		{Code: 102, Index: 18, Length: 19},
		{Code: 102, Index: 20, Length: 21},
		{Code: 103, Index: 20, Length: 21},
	}

	expected := []models.SearchIndexModel{
		{Code: 99, AsciiStart: 1, AsciiEnd: 1},
		{Code: 100, AsciiStart: 2, AsciiEnd: 3},
		{Code: 101, AsciiStart: 4, AsciiEnd: 6},
		{Code: 102, AsciiStart: 7, AsciiEnd: 11},
		{Code: 103, AsciiStart: 12, AsciiEnd: 12},
	}

	result := PrepareData{}.generateSearchIndexes(asciiIndexes)

	assert.Equal(expected, result)
}
