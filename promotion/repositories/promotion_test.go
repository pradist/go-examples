package repositories_test

import (
	"go-examples/promotion/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionGet_ShouldReturnIDIs10(t *testing.T) {
	expected := 1

	p := repositories.New()
	actual := p.Get().ID

	assert.Equal(t, expected, actual)
}

func TestPromotionGet_ShouldReturnDiscountPercentIs10(t *testing.T) {
	expected := 10

	p := repositories.New()
	actual := p.Get().DiscountPercent

	assert.Equal(t, expected, actual)
}

func TestPromotionGet_ShouldReturnPurchaseMinIs100(t *testing.T) {
	expected := 100

	p := repositories.New()
	actual := p.Get().PurchaseMin

	assert.Equal(t, expected, actual)
}
