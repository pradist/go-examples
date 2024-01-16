package captcha

import "strconv"

type IntOperand struct {
	Operand int
}

func (o *IntOperand) ToString() string {
	return strconv.Itoa(o.Operand)
}
