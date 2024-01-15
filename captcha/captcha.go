package captcha

import "strconv"

var numbers = []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

func LeftOperand(pattern, left int) string {
	if pattern == 2 {
		return numbers[left]
	}
	return strconv.Itoa(left)
}

func RightOperand(pattern, right int) string {
	if pattern == 2 {
		return "1"
	}
	return numbers[right]
}
