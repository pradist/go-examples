package captcha_test

import (
	"go-examples/captcha"
	"testing"

	"github.com/stretchr/testify/assert"
)

var placeholder = 1

func TestCaptchaLeftOperand_WhenPatternIs1andLeftOperandIs1_LeftOperandShouldReturn1(t *testing.T) {
	captcha := captcha.New(1, 1, placeholder, placeholder)

	expected := "1"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs1andLeftOperandIs2_LeftOperandShouldReturn2(t *testing.T) {
	captcha := captcha.New(1, 2, placeholder, placeholder)

	expected := "2"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs1andLeftOperandIs3_LeftOperandShouldReturn3(t *testing.T) {
	captcha := captcha.New(1, 3, placeholder, placeholder)

	expected := "3"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs1andLeftOperandIs4_LeftOperandShouldReturn4(t *testing.T) {
	captcha := captcha.New(1, 4, placeholder, placeholder)

	expected := "4"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andLeftOperandIs1_LeftOperandShouldReturnOne(t *testing.T) {
	captcha := captcha.New(2, 1, placeholder, placeholder)

	expected := "One"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andLeftOperandIs2_LeftOperandShouldReturnTwo(t *testing.T) {
	captcha := captcha.New(2, 2, placeholder, placeholder)

	expected := "Two"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andLeftOperandIs3_LeftOperandShouldReturnThree(t *testing.T) {
	captcha := captcha.New(2, 3, placeholder, placeholder)

	expected := "Three"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andLeftOperandIs4_LeftOperandShouldReturnFour(t *testing.T) {
	captcha := captcha.New(2, 4, placeholder, placeholder)

	expected := "Four"

	actual := captcha.LeftOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaRightOperand_WhenPatternIs1andRightOperandIs1_RightOperandShouldReturnOne(t *testing.T) {
	captcha := captcha.New(1, 1, placeholder, placeholder)

	expected := "One"

	actual := captcha.RightOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaRightOperand_WhenPatternIs1andRightOperandIs2_RightOperandShouldReturnTwo(t *testing.T) {
	captcha := captcha.New(1, placeholder, placeholder, 2)

	expected := "Two"

	actual := captcha.RightOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaRightOperand_WhenPatternIs2andRightOperand1_RightOperandShouldReturn1(t *testing.T) {
	captcha := captcha.New(2, 1, placeholder, placeholder)

	expected := "1"

	actual := captcha.RightOperand()

	assert.Equal(t, expected, actual)
}

func TestCaptchaOperator_WhenOperatorIs1_OperatorShouldReturnPlus(t *testing.T) {
	captcha := captcha.New(placeholder, placeholder, 1, placeholder)

	expected := "+"

	actual := captcha.Operator()

	assert.Equal(t, expected, actual)
}

func TestCaptchaOperator_WhenOperatorIs2_OperatorShouldReturnMinus(t *testing.T) {
	captcha := captcha.New(placeholder, placeholder, 2, placeholder)

	expected := "-"

	actual := captcha.Operator()

	assert.Equal(t, expected, actual)
}

func TestCaptcha_WhenPatternIs1_LeftOperandIs1_OperatorIs1_RightOperandIs1_StringShouldReturn1PlusOne(t *testing.T) {
	captcha := captcha.New(1, 1, 1, 1)

	expected := "1 + One"

	actual := captcha.Captcha()

	assert.Equal(t, expected, actual)
}

func TestCaptcha_WhenPatternIs1_LeftOperandIs1_OperatorIs1_RightOperandIs2_StringShouldReturn1PlusNine(t *testing.T) {
	captcha := captcha.New(1, 1, 1, 9)

	expected := "1 + Nine"

	actual := captcha.Captcha()

	assert.Equal(t, expected, actual)
}

func TestCaptcha_WhenPatternIs2_LeftOperandIs1_OperatorIs1_RightOperandIs1_StringShouldReturnOnePlusOne(t *testing.T) {
	captcha := captcha.New(2, 1, 1, 1)

	expected := "One + 1"

	actual := captcha.Captcha()

	assert.Equal(t, expected, actual)
}
