package captcha_test

import (
	"go-examples/captcha"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntOperand_WhenOperandIs1_ShouldReturn1(t *testing.T) {
	operand := captcha.IntOperand{Operand: 1}
	assert.Equal(t, "1", operand.ToString())
}

func TestIntOperand_WhenOperandIs5_ShouldReturn5(t *testing.T) {
	operand := captcha.IntOperand{Operand: 2}
	assert.Equal(t, "2", operand.ToString())
}

func TestIntOperand_WhenOperandIs9_ShouldReturn9(t *testing.T) {
	operand := captcha.IntOperand{Operand: 9}
	assert.Equal(t, "9", operand.ToString())
}
