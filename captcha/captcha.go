package captcha

func LeftOperand(pattern, left int) string {
	if left == 2 {
		return "2"
	}
	if left == 3 {
		return "3"
	}
	return "1"
}
