package services

import "go-examples/promotion/repositories"

func CalculateDiscount(amount int) int {
	p := repositories.New()
	promotion := p.Get()

	return amount - (amount * promotion.DiscountPercent / 100)
}
