package captcha_test

import (
	"go-examples/captcha"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringOperator_WhenOperatorIs1_ShouldReturnPlus(t *testing.T) {
	expected := "+"

	o := captcha.StringOperator{1}
	actual := o.ToString()

	assert.Equal(t, expected, actual)
}

func TestStringOperator_WhenOperatorIs2_ShouldReturnMinus(t *testing.T) {
	expected := "-"

	o := captcha.StringOperator{2}
	actual := o.ToString()

	assert.Equal(t, expected, actual)
}

func TestStringOperand_WhenOperandIs3_ShouldReturnDivide(t *testing.T) {
	expected := "/"

	o := captcha.StringOperator{3}
	actual := o.ToString()

	assert.Equal(t, expected, actual)
}
