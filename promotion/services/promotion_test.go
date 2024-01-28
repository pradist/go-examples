package services_test

import (
	"go-examples/promotion/repositories"
	"go-examples/promotion/services"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type promotionRepositoriesMock struct {
	MockGet func() repositories.Promotion
}

func (p *promotionRepositoriesMock) GetPromotion() repositories.Promotion {
	return p.MockGet()
}

func NewPromotionRepositoriesMock() *promotionRepositoriesMock {
	return &promotionRepositoriesMock{
		MockGet: func() repositories.Promotion {
			return repositories.Promotion{
				ID:              1,
				PurchaseMin:     100,
				DiscountPercent: 10,
			}
		},
	}
}

func TestPromotionService_WhenAmountIs0_ShouldReturnError(t *testing.T) {
	expected := errors.New("amount is less than zero")

	r := NewPromotionRepositoriesMock().GetPromotion()
	s := services.New(r)
	_, actual := s.CalculateDiscount(0)

	assert.Equal(t, expected, actual)
}

func TestPromotionService_WhenAmountIs100_ShouldPay90(t *testing.T) {
	expected := 90

	r := NewPromotionRepositoriesMock().GetPromotion()
	s := services.New(r)
	actual, _ := s.CalculateDiscount(100)

	assert.Equal(t, expected, actual)
}

func TestPromotionService_WhenAmountIs200_ShouldPay180(t *testing.T) {
	expected := 180

	r := NewPromotionRepositoriesMock().GetPromotion()
	s := services.New(r)
	actual, _ := s.CalculateDiscount(200)

	assert.Equal(t, expected, actual)
}

func TestPromotionService_WhenAmountIs300_ShouldPay270(t *testing.T) {
	expected := 270

	r := NewPromotionRepositoriesMock()
	s := services.New(r.GetPromotion())
	actual, _ := s.CalculateDiscount(300)

	assert.Equal(t, expected, actual)
}
