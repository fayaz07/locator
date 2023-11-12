package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosestStrings(t *testing.T) {
	assert := assert.New(t)

	strings := []string{"New York City", "Los Angeles", "San Francisco", "Washington D.C.", "Franklin"}
	pattern := "New York"
	expected := []string{"New York City"}

	assert.Equal(expected, FindClosestStrings(strings, pattern))

	pattern = "Fra"
	expected = []string{"San Francisco", "Franklin"}
	assert.Equal(expected, FindClosestStrings(strings, pattern))

	pattern = "San"
	expected = []string{"San Francisco"}
	assert.Equal(expected, FindClosestStrings(strings, pattern))

	pattern = "San F"
	expected = []string{"San Francisco"}
	assert.Equal(expected, FindClosestStrings(strings, pattern))

	strings = []string{"apple", "banana", "cherry", "grape", "orange", "pear"}
	pattern = "xyz"
	assert.Empty(FindClosestStrings(strings, pattern))

	pattern = "a"
	expected = []string{"apple", "banana", "grape", "orange", "pear"}
	assert.Equal(expected, FindClosestStrings(strings, pattern))

	pattern = "ap"
	expected = []string{"apple", "grape"}
	assert.Equal(expected, FindClosestStrings(strings, pattern))
}
