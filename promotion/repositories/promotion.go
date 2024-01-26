package repositories

type Promotion struct {
	ID              int
	PurchaseMin     int
	DiscountPercent int
}

func Get() Promotion {
	return Promotion{
		ID:              1,
		PurchaseMin:     100,
		DiscountPercent: 10,
	}
}
