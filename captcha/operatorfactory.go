package captcha

func OperatorFactory(operator int) Operator {
	return &StringOperator{Operator: operator}
}
