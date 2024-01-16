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
