package remove_outer_parentheses

func removeOuterParentheses(s string) string {
	result := make([]byte, 0, len(s))
	count := 0
	for i := range s {
		if s[i] == '(' {
			if count > 0 {
				result = append(result, s[i])
			}
			count++
		} else {
			count--
			if count > 0 {
				result = append(result, s[i])
			}
		}
	}
	return string(result)
}
