package services

import "go-examples/promotion/repositories"

func CalculateDiscount(amount int) int {

	promotion := repositories.Get()
	return amount - (amount * promotion.DiscountPercent / 100)
}
