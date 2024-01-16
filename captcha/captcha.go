package captcha

import (
	"fmt"
	"math/rand"
)

type Captcha struct {
	pattern      int
	leftOperand  int
	operator     int
	rightOperand int
}

func New(pattern, leftOperand, operator, rightOperand int) *Captcha {
	fmt.Println(pattern, leftOperand, operator, rightOperand)
	return &Captcha{pattern, leftOperand, operator, rightOperand}
}

func Create() *Captcha {
	return New(
		rand.Intn(2)+1,
		rand.Intn(2)+1,
		rand.Intn(2)+1,
		rand.Intn(2)+1)
}

func (c *Captcha) Captcha() string {
	leftOperand := LeftOperandFactory(c.pattern, c.leftOperand)
	operator := OperatorFactory(c.operator)
	rightOperand := RightOperandFactory(c.pattern, c.rightOperand)
	return leftOperand.ToString() + " " + operator.ToString() + " " + rightOperand.ToString()
}
