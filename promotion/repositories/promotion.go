package repositories

type Promotion struct {
	ID              int
	PurchaseMin     int
	DiscountPercent int
}

func New() Promotion {
	return Promotion{
		ID:              1,
		PurchaseMin:     100,
		DiscountPercent: 10,
	}
}

func (p Promotion) Get() (Promotion, error) {
	return p, nil
}
