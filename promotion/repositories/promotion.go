package repositories

type Promotion struct {
	PurchaseMin     int
	DiscountPercent int
}

func Get() Promotion {
	return Promotion{
		PurchaseMin:     100,
		DiscountPercent: 10,
	}
}
