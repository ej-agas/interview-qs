package interview_qs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoCrystalBalls(t *testing.T) {
	index5 := []bool{false, false, false, false, false, true, true, true, true, true, true}

	assert.Equal(t, "Crystal Balls break at index 5.", TwoCrystalBalls(index5))

	index3 := []bool{false, false, false, true, true, true, true, true}

	assert.Equal(t, "Crystal Balls break at index 3.", TwoCrystalBalls(index3))

	index0 := []bool{true, true, true, true, true, true, true}

	assert.Equal(t, "Crystal Balls break at index 0.", TwoCrystalBalls(index0))

	index6 := []bool{false, false, false, false, false, false, true}

	assert.Equal(t, "Crystal Balls break at index 6.", TwoCrystalBalls(index6))

	wontBreak := []bool{false, false, false}

	assert.Equal(t, "Crystal balls did not break at any point.", TwoCrystalBalls(wontBreak))

}
