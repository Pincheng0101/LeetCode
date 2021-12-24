package p0227

func calculate(s string) int {
	result, currentNumber, lastNumber := 0, 0, 0
	operation := '+'
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			currentNumber = currentNumber*10 + int(s[i]-'0')
		}
		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == len(s)-1 {
			if operation == '+' || operation == '-' {
				result += lastNumber
				lastNumber = currentNumber
				if operation == '-' {
					lastNumber = -lastNumber
				}
			} else if operation == '*' {
				lastNumber *= currentNumber
			} else if operation == '/' {
				lastNumber /= currentNumber
			}
			operation = rune(s[i])
			currentNumber = 0
		}
	}
	result += lastNumber
	return result
}
