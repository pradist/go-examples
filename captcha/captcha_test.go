package captcha_test

import (
	"go-examples/captcha"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs1_ShouldReturn1(t *testing.T) {
	expected := "1"

	actual := captcha.LeftOperand(1, 1)

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs2_ShouldReturn2(t *testing.T) {
	expected := "2"

	actual := captcha.LeftOperand(1, 2)

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs3_ShouldReturn3(t *testing.T) {
	expected := "3"

	actual := captcha.LeftOperand(1, 3)

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs4_ShouldReturn4(t *testing.T) {
	expected := "4"

	actual := captcha.LeftOperand(1, 4)

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andInputIs1_ShouldReturnOne(t *testing.T) {
	expected := "One"

	actual := captcha.LeftOperand(2, 1)

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andInputIs2_ShouldReturnTwo(t *testing.T) {
	expected := "Two"

	actual := captcha.LeftOperand(2, 2)

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andInputIs3_ShouldReturnThree(t *testing.T) {
	expected := "Three"

	actual := captcha.LeftOperand(2, 3)

	assert.Equal(t, expected, actual)
}

func TestCaptchaLeftOperand_WhenPatternIs2andInputIs4_ShouldReturnFour(t *testing.T) {
	expected := "Four"

	actual := captcha.LeftOperand(2, 4)

	assert.Equal(t, expected, actual)
}

func TestCaptchaRightOperand_WhenPatternIs1andInputIs1_ShouldReturnOne(t *testing.T) {
	expected := "One"

	actual := captcha.RightOperand(1, 1)

	assert.Equal(t, expected, actual)
}

func TestCaptchaRightOperand_WhenPatternIs1andInputIs2_ShouldReturnTwo(t *testing.T) {
	expected := "Two"

	actual := captcha.RightOperand(1, 2)

	assert.Equal(t, expected, actual)
}

func TestCaptchaRightOperand_WhenPatternIs2andInput1_ShouldReturn1(t *testing.T) {
	expected := "1"

	actual := captcha.RightOperand(2, 1)

	assert.Equal(t, expected, actual)
}
