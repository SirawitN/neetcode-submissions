func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	var idx, tmp int

	for _, token := range tokens {
		if token != "+" && token != "-" && token != "*" && token != "/" {
			stack[idx], _ = strconv.Atoi(token)
			idx += 1
		} else {

			switch token {
				case "+":
					tmp = stack[idx-2] + stack[idx-1]
				case "-":
					tmp = stack[idx-2] - stack[idx-1]
				case "*":
					tmp = stack[idx-2] * stack[idx-1]
				default:
					tmp = stack[idx-2] / stack[idx-1]
			}

			idx -= 2
			stack[idx] = tmp
			idx += 1
		}
	}

	return stack[idx-1]
}
