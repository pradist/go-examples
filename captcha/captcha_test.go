package captcha_test

import (
	"go-examples/captcha"
	"testing"

	"github.com/stretchr/testify/assert"
)

var placeholder = 1

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs1_LeftOperandShouldReturn1(t *testing.T) {
	captcha := captcha.New(1, 1, placeholder, placeholder)

	expected := "1"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs2_LeftOperandShouldReturn2(t *testing.T) {
	captcha := captcha.New(1, 2, placeholder, placeholder)

	expected := "2"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs3_LeftOperandShouldReturn3(t *testing.T) {
	captcha := captcha.New(1, 3, placeholder, placeholder)

	expected := "3"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs4_LeftOperandShouldReturn4(t *testing.T) {
	captcha := captcha.New(1, 4, placeholder, placeholder)

	expected := "4"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andInputIs1_LeftOperandShouldReturnOne(t *testing.T) {
	captcha := captcha.New(2, 1, placeholder, placeholder)

	expected := "One"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andInputIs2_LeftOperandShouldReturnTwo(t *testing.T) {
	captcha := captcha.New(2, 2, placeholder, placeholder)

	expected := "Two"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andInputIs3_LeftOperandShouldReturnThree(t *testing.T) {
	captcha := captcha.New(2, 3, placeholder, placeholder)

	expected := "Three"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andInputIs4_LeftOperandShouldReturnFour(t *testing.T) {
	captcha := captcha.New(2, 4, placeholder, placeholder)

	expected := "Four"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaRightOperand_WhenPatternIs1andInputIs1_RightOperandShouldReturnOne(t *testing.T) {
	captcha := captcha.New(1, 1, placeholder, placeholder)

	expected := "One"

	actual := captcha.RightOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaRightOperand_WhenPatternIs1andInputIs2_RightOperandShouldReturnTwo(t *testing.T) {
	captcha := captcha.New(1, placeholder, placeholder, 2)

	expected := "Two"

	actual := captcha.RightOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaRightOperand_WhenPatternIs2andInput1_RightOperandShouldReturn1(t *testing.T) {
	captcha := captcha.New(2, 1, placeholder, placeholder)

	expected := "1"

	actual := captcha.RightOperand()

	assert.Equal(t, expected, actual)
}

func TestCatchaOperator_WhenInputIs1_OperatorShouldReturnPlus(t *testing.T) {
	captcha := captcha.New(placeholder, placeholder, 1, placeholder)

	expected := "+"

	actual := captcha.Operator()

	assert.Equal(t, expected, actual)
}

func TestCatchaOperator_WhenInputIs2_OperatorShouldReturnMinus(t *testing.T) {
	captcha := captcha.New(placeholder, placeholder, 2, placeholder)

	expected := "-"

	actual := captcha.Operator()

	assert.Equal(t, expected, actual)
}
