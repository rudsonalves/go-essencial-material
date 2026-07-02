package textutil

func OnlyDigits(value string) string {
	numbers := []rune{}

	for _, r := range value {
		if r >= '0' && r <= '9' {
			numbers = append(numbers, r)
		}
	}

	return string(numbers)
}
