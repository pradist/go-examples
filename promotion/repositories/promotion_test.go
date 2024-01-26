package repositories_test

import (
	"go-examples/promotion/repositories"
	"testing"
)

func TestPromotionGet_ShouldReturnDiscountPercentIs10(t *testing.T) {
	expected := 10

	actual := repositories.Get().DiscountPercent

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}

func TestPromotionGet_ShouldReturnPurchaseMinIs100(t *testing.T) {
	expected := 100

	actual := repositories.Get().PurchaseMin

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}
}
