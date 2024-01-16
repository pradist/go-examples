package captcha

type StringOperand struct {
	operand int
}

func (o *StringOperand) ToString() string {
	return stringNumbers[o.operand-1]
}
