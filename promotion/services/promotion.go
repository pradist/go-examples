package services

import (
	"errors"
	"go-examples/promotion/repositories"
)

func CalculateDiscount(amount int) (int, error) {

	if amount == 0 {
		return 0, errors.New("amount is less than zero")
	}

	p := repositories.New()

	promotion := p.Get()

	return amount - (amount * promotion.DiscountPercent / 100), nil
}
