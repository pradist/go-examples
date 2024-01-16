package captcha

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringOperandWhenOperandIs1_StringOperandShouldReturnOne(t *testing.T) {
	expected := "One"

	o := StringOperand{1}
	actual := o.ToString()

	assert.Equal(t, expected, actual)
}

func TestStringOperandWhenOperandIs5_StringOperandShouldReturnFive(t *testing.T) {
	expected := "Five"

	o := StringOperand{5}
	actual := o.ToString()

	assert.Equal(t, expected, actual)
}

func TestStringOperandWhenOperandIs9_StringOperandShouldReturnNine(t *testing.T) {
	expected := "Nine"

	o := StringOperand{9}
	actual := o.ToString()

	assert.Equal(t, expected, actual)
}
