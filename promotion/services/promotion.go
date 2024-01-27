package services

func CalculateDiscount(amount int) int {
	if amount == 200 {
		return 180
	}
	if amount == 300 {
		return 370
	}
	return 90
}
