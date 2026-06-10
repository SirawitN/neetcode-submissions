func evalRPN(tokens []string) int {
	stack := make([]int, len(tokens))
	var idx, tmp int
	var leftVar, rightVar int

	for _, token := range tokens {
		if token != "+" && token != "-" && token != "*" && token != "/" {
			stack[idx], _ = strconv.Atoi(token)
			idx += 1
		} else {
			leftVar = stack[idx-2]
			rightVar = stack[idx-1]
			idx -= 2

			switch token {
				case "+":
					tmp = leftVar + rightVar
				case "-":
					tmp = leftVar - rightVar
				case "*":
					tmp = leftVar * rightVar
				default:
					tmp = leftVar / rightVar
			}

			stack[idx] = tmp
			idx += 1
		}
	}

	return stack[idx-1]
}
