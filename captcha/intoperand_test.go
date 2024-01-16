package captcha_test

import (
	"go-examples/captcha"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntOperandWhenOperandIs1_IntOperandShouldReturn1(t *testing.T) {
	expected := "1"

	o := captcha.IntOperand{1}
	actual := o.ToString()

	assert.Equal(t, expected, actual)
}

func TestIntOperandWhenOperandIs5_IntOperandShouldReturn5(t *testing.T) {
	expected := "5"

	o := captcha.IntOperand{5}
	actual := o.ToString()

	assert.Equal(t, expected, actual)
}

func TestIntOperandWhenOperandIs9_IntOperandShouldReturn9(t *testing.T) {
	expected := "9"

	o := captcha.IntOperand{9}
	actual := o.ToString()

	assert.Equal(t, expected, actual)
}
