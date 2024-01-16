package captcha

var stringOperators = []string{"+", "-", "/"}

type StringOperator struct {
	Operator int
}

func (o *StringOperator) ToString() string {
	return stringOperators[o.Operator-1]
}
