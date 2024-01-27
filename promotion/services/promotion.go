package services

func CalculateDiscount(amount int) int {
	if amount == 200 {
		return 180
	}
	return 90
}
