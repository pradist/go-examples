package captcha

import "strconv"

var stringNumbers = []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

type Captcha struct {
	pattern      int
	leftOperand  int
	operator     int
	rightOperand int
}

func New(pattern, leftOperand, operator, rightOperand int) *Captcha {
	return &Captcha{pattern, leftOperand, operator, rightOperand}
}

func (c *Captcha) LeftOperand() string {
	if c.pattern == 1 {
		return strconv.Itoa(c.leftOperand)
	}
	return stringNumbers[c.leftOperand]
}

func (c *Captcha) RightOperand() string {
	if c.pattern == 1 {
		return stringNumbers[c.rightOperand]
	}
	return strconv.Itoa(c.leftOperand)
}
