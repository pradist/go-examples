package captcha

var stringNumbers = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
var stringOperators = []string{"+", "-", "/"}

type Captcha struct {
	pattern      int
	leftOperand  int
	operator     int
	rightOperand int
}

func New(pattern, leftOperand, operator, rightOperand int) *Captcha {
	return &Captcha{pattern, leftOperand, operator, rightOperand}
}

func (c *Captcha) Operator() string {
	return stringOperators[c.operator-1]
}

func (c *Captcha) Captcha() string {
	leftOperand := LeftOperandFactory(c.pattern, c.leftOperand)
	rightOperand := RightOperandFactory(c.pattern, c.rightOperand)
	return leftOperand.ToString() + " " + c.Operator() + " " + rightOperand.ToString()
}
