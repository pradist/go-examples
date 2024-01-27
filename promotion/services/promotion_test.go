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

func TestPromotionService_WhenAmountIs200_ShouldPay180(t *testing.T) {
	expected := 180

	actual := services.CalculateDiscount(200)

	assert.Equal(t, expected, actual)
}

func TestPromotionService_WhenAmountIs300_ShouldPay370(t *testing.T) {
	expected := 370

	actual := services.CalculateDiscount(300)

	assert.Equal(t, expected, actual)
}
