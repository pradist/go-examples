package captcha

func LeftOperand(pattern, left int) string {
	if left == 2 {
		return "2"
	}
	if left == 3 {
		return "3"
	}
	if left == 4 {
		return "4"
	}
	return "1"
}
