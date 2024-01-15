package decider_test

import (
	"go-examples/decider"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecider_WhenCalled9_ShouldReturn9(t *testing.T) {
	decider := decider.Decider{
		RandImpl: func(x int) int { return 9 },
	}

	expected := 9

	actual := decider.Decide(9)

	assert.Equal(t, expected, actual)
}

func TestDecideWithRange_WhenCalled1and10_ShouldReturn5(t *testing.T) {
	decider := decider.Decider{
		RandImpl: func(x int) int { return 4 },
	}

	expected := 5

	actual := decider.DecideWithRange(1, 10)

	assert.Equal(t, expected, actual)
}
