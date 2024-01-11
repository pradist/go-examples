package oddeven

func Find(input int) string {
	if input%2 == 0 {
		return "Even"
	}
	return "Odd"
}
