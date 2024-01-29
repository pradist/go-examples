package services

import (
	"go-examples/promotion/repositories"

	"errors"
)

type PromotionService struct {
	repo repositories.Promotion
}

func New(repo repositories.Promotion) PromotionService {
	return PromotionService{
		repo: repo,
	}
}

func (p PromotionService) CalculateDiscount(amount int) (int, error) {
	if amount == 0 {
		return 0, errors.New("amount is less than zero")
	}

	r := repositories.NewPromotionRepository()
	promotion, _ := r.GetPromotion()

	return amount - (amount * promotion.DiscountPercent / 100), nil
}
