package captcha

func LeftOperandFactory(pattern, operand int) Operand {
	if pattern == 1 {
		return &IntOperand{operand}
	}
	return &StringOperand{operand}
}
