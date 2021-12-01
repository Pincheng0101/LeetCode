package p2011

func finalValueAfterOperations(operations []string) int {
	value := 0
	for _, op := range operations {
		switch op {
		case "++X", "X++":
			value++
		case "--X", "X--":
			value--
		}
	}
	return value
}
