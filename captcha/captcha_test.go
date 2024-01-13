package captcha_test

import (
	"go-examples/captcha"
	"testing"
)

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs1_ShouldReturn1(t *testing.T) {
	expected := "1"

	actual := captcha.LeftOperand(1, 1)

	if expected != actual {
		t.Errorf("expect %s but it got %s", expected, actual)
	}
}

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs2_ShouldReturn2(t *testing.T) {
	expected := "2"

	actual := captcha.LeftOperand(1, 2)

	if expected != actual {
		t.Errorf("expect %s but it got %s", expected, actual)
	}
}

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs3_ShouldReturn3(t *testing.T) {
	expected := "3"

	actual := captcha.LeftOperand(1, 3)

	if expected != actual {
		t.Errorf("expect %s but it got %s", expected, actual)
	}
}

func TestCaptchaLeftOperand_WhenPatternIs1andInputIs4_ShouldReturn4(t *testing.T) {
	expected := "4"

	actual := captcha.LeftOperand(1, 4)

	if expected != actual {
		t.Errorf("expect %s but it got %s", expected, actual)
	}
}
