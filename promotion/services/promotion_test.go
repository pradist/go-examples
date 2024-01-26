package services_test

import (
	"go-examples/promotion/services"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionService_WhenAmountIs100_ShouldPay90(t *testing.T) {
	expected := 90

	actual := services.CalculateDiscount(100)

	assert.Equal(t, expected, actual)
}
