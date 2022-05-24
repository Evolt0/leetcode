package remove_outer_parentheses

func removeOuterParentheses(S string) string {
	result := make([]byte, 0, len(S))
	count := 0
	for i := range S {
		if S[i] == '(' {
			if count > 0 {
				result = append(result, S[i])
			}
			count++
		} else {
			count--
			if count > 0 {
				result = append(result, S[i])
			}
		}
	}
	return string(result)
}
