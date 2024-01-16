package captcha_test

import (
	"go-examples/captcha"
	"testing"

	"github.com/stretchr/testify/assert"
)

var placeholder = 1

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

func TestCaptcha_WhenPatternIs2_LeftOperandIs1_OperatorIs1_RightOperandIs1_StringShouldReturnOnePlus1(t *testing.T) {
	captcha := captcha.New(2, 1, 1, 1)

	expected := "One + 1"

	actual := captcha.Captcha()

	assert.Equal(t, expected, actual)
}

func TestCaptcha_WhenPatternIs2_LeftOperandIs2_OperatorIs3_RightOperandIs5_StringShouldReturnTwoPlus5(t *testing.T) {
	captcha := captcha.New(2, 2, 3, 5)

	expected := "Two / 5"

	actual := captcha.Captcha()

	assert.Equal(t, expected, actual)
}
