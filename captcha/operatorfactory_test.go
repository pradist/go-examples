package captcha_test

import (
	"go-examples/captcha"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorFactory_ShouldBeStringOperator(t *testing.T) {
	expected := &captcha.StringOperator{1}

	actual := captcha.OperatorFactory(1)

	assert.Equal(t, expected, actual)
}
