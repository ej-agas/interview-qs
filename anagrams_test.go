package interview_qs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	wordList := []string{
		"acer",
		"acre",
		"care",
		"racer",
	}

	assert.Equal(t, []string{"acer", "acre", "care"}, FindAnagrams("race", wordList))

	wordList2 := []string{
		"coronavirus",
		"carnivorous",
	}

	assert.Equal(t, wordList2, FindAnagrams("coronavirus", wordList2))
}
