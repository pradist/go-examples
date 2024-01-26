package repositories_test

import (
	"go-examples/promotion/repositories"
	"testing"
)

func TestPromotion_ShouldDiscountPercentIs10(t *testing.T) {
	expected := 10

	actual := repositories.Get()

	if expected != actual {
		t.Errorf("expected %v but it got %v", expected, actual)
	}

}
