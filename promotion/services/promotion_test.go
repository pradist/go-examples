package services_test

import (
	"errors"
	"go-examples/promotion/services"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionService_WhenAmountIs0_ShouldReturnError(t *testing.T) {
	expected := errors.New("amount is less than zero")

	_, actual := services.CalculateDiscount(0)

	assert.Equal(t, expected, actual)
}

func TestPromotionService_WhenAmountIs100_ShouldPay90(t *testing.T) {
	expected := 90

	actual, _ := services.CalculateDiscount(100)

	assert.Equal(t, expected, actual)
}

func TestPromotionService_WhenAmountIs200_ShouldPay180(t *testing.T) {
	expected := 180

	actual, _ := services.CalculateDiscount(200)

	assert.Equal(t, expected, actual)
}

func TestPromotionService_WhenAmountIs300_ShouldPay270(t *testing.T) {
	expected := 270

	actual, _ := services.CalculateDiscount(300)

	assert.Equal(t, expected, actual)
}
