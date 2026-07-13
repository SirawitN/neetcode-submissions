func evalRPN(tokens []string) int {
	stack := make([]string, len(tokens))
	stackIdx := -1
	var tmp, a, b int

	for _, t := range tokens {
		// fmt.Println(stack)
		switch t {
			case "+", "-", "*", "/":
				a, _ = strconv.Atoi(stack[stackIdx-1])
				b, _ = strconv.Atoi(stack[stackIdx])
				// fmt.Println(a,b,t)
				stackIdx -= 1

				if t=="+" {
					tmp = a+b
				} else if t=="-"{
					tmp = a-b
				} else if t=="*"{
					tmp = a*b
				} else {
					if b!=0{
						tmp = a/b
					}
				}
				stack[stackIdx] = strconv.Itoa(tmp)
			default:
				stackIdx += 1
				stack[stackIdx] = t
		}
	}

	// fmt.Println(stack)
	result, _ := strconv.Atoi(stack[stackIdx])
	return result
}
