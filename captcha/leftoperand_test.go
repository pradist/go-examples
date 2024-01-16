package captcha_test

import (
	"go-examples/captcha"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftOperandFactoryWhenPatternIs1_operandShouldBeIntegerOperand(t *testing.T) {
	expected := &captcha.IntOperand{}

	actual := captcha.LeftOperandFactory(1, 1)

	assert.IsType(t, expected, actual)
}

func TestLeftOperandFactoryWhenPatternIs2_operandShouldBeStringOperand(t *testing.T) {
	expected := &captcha.StringOperand{}

	actual := captcha.LeftOperandFactory(2, 1)

	assert.IsType(t, expected, actual)
}
