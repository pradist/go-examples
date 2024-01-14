package captcha

import "strconv"

var number = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

func LeftOperand(pattern, left int) string {
	if pattern == 2 {
		if left == 2 {
			return "Two"
		}
		if left == 3 {
			return "Three"
		}
		if left == 4 {
			return "Four"
		}
		return number[left-1]
	}
	return strconv.Itoa(left)
}
