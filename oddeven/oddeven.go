package oddeven

func Find(input int) string {
	if input%2 == 0 {
		return "Even"
	}
	if input == 4 {
		return "Even"
	}
	if input == 6 {
		return "Even"
	}
	return "Odd"
}
