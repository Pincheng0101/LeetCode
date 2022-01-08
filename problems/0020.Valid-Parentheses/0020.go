package p0020

// 使用 Stack
func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			// Push Data to Stack
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			// Pop Data from Stack
			// 取出最後一個值
			end := stack[len(stack)-1]
			if (end == '(' && s[i] == ')') ||
				(end == '[' && s[i] == ']') ||
				(end == '{' && s[i] == '}') {
				// 刪除最後一個值
				stack = stack[:len(stack)-1]
				continue
			}
			return false
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
