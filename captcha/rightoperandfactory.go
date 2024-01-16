package captcha

func RightOperandFactory(pattern, operand int) Operand {
	if pattern == 1 {
		return &StringOperand{operand}
	}
	return &IntOperand{operand}
}
