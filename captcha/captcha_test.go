package captcha_test

import (
	"go-examples/captcha"
	"testing"
)

func TestCaptchaLeftOperand_WhenPattern1andInputIs1_SouldReturn1(t *testing.T) {
	expected := "1"

	actual := captcha.LeftOperand(1, 1)

	if expected != actual {
		t.Errorf("expect %s but it got %s", expected, actual)
	}
}
