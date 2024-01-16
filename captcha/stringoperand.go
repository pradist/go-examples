package captcha

var stringNumbers = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

type StringOperand struct {
	Operand int
}

func (o *StringOperand) ToString() string {
	return stringNumbers[o.Operand-1]
}
