package oddeven_test

import (
	"go-examples/oddeven"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenInputIs1ShouldRetuenOdd(t *testing.T) {
	expected := "Odd"

	actual := oddeven.Find(1)

	assert.Equal(t, expected, actual)
}
