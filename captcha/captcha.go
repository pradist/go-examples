package captcha

type Captcha struct {
	pattern      int
	leftOperand  int
	operator     int
	rightOperand int
}

func New(pattern, leftOperand, operator, rightOperand int) *Captcha {
	return &Captcha{pattern, leftOperand, operator, rightOperand}
}

func (c *Captcha) Captcha() string {
	leftOperand := LeftOperandFactory(c.pattern, c.leftOperand)
	operator := OperatorFactory(c.operator)
	rightOperand := RightOperandFactory(c.pattern, c.rightOperand)
	return leftOperand.ToString() + " " + operator.ToString() + " " + rightOperand.ToString()
}
