package decider_test

import (
	"go-examples/decider"
	"testing"
)

func TestDecider_WhenCalled9_ShouldReturn9(t *testing.T) {
	decider := decider.Decider{
		RandImpl: func(x int) int { return 9 },
	}

	expected := 9

	actual := decider.Decide(9)

	if expected != actual {
		t.Errorf("expect %d but it got %d", expected, actual)
	}
}

func TestDecideWithRange_WhenCalled1and10_ShouldReturn5(t *testing.T) {
	decider := decider.Decider{
		RandImpl: func(x int) int { return 4 },
	}

	expected := 5

	actual := decider.DecideWithRange(1, 10)

	if expected != actual {
		t.Errorf("expect %d but it got %d", expected, actual)
	}
}
