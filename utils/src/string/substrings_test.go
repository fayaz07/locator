package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSubstrings(t *testing.T) {
	assert := assert.New(t)

	s := "apple"
	minSize := 2
	maxSize := 4
	expected := []string{"ap", "app", "appl", "pp", "ppl", "pple", "pl", "ple", "le"}

	assert.Equal(expected, GenerateSubstrings(s, minSize, maxSize, len(s)))

	s = "Nizamabad"
	minSize = 3
	maxSize = 7
	expected = []string{
		"Niz", "Niza", "Nizam", "Nizama", "Nizamab",
		"iza", "izam", "izama", "izamab", "izamaba",
		"zam", "zama", "zamab", "zamaba", "zamabad",
		"ama", "amab", "amaba", "amabad",
		"mab", "maba", "mabad",
		"aba", "abad", "bad",
	}
	assert.Equal(expected, GenerateSubstrings(s, minSize, maxSize, len(s)))

	s = "Nizamabad"
	minSize = 3
	maxSize = 7
	maxStartIndex := 3
	expected = []string{
		"Niz", "Niza", "Nizam", "Nizama", "Nizamab",
		"iza", "izam", "izama", "izamab", "izamaba",
		"zam", "zama", "zamab", "zamaba", "zamabad",
	}

	assert.Equal(expected, GenerateSubstrings(s, minSize, maxSize, maxStartIndex))
}
