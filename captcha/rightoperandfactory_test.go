package captcha_test

import (
	"go-examples/captcha"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightOperandFactoryWhenPatternIs1_operandShouldBeStringOperand(t *testing.T) {
	expected := &captcha.StringOperand{}

	actual := captcha.RightOperandFactory(1, 1)

	assert.IsType(t, expected, actual)
}

func TestRightOperandFactoryWhenPatternIs2_operandShouldBeIntegerOperand(t *testing.T) {
	expected := &captcha.IntOperand{}

	actual := captcha.RightOperandFactory(2, 1)

	assert.IsType(t, expected, actual)
}
