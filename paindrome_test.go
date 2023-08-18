package interview_qs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, IsPalindrome("civic"))
	assert.True(t, IsPalindrome("rotor"))
	assert.True(t, IsPalindrome("Race car"))
	assert.True(t, IsPalindrome("Was it a car or a cat I saw"))
	assert.True(t, IsPalindrome("Never odd or even"))

	assert.False(t, IsPalindrome("assert"))
	assert.False(t, IsPalindrome("test"))
	assert.False(t, IsPalindrome("The sky is blue"))
	assert.False(t, IsPalindrome("The quick brown fox jumps over the lazy dog"))
}
