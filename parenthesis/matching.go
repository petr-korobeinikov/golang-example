package parenthesis

// IsMatching check parenthesis order in given string.
// Returns true, if parenthesis are matching, and false otherwise.
func IsMatching(s string) bool {
	l := len(s)
	if l == 0 {
		return true
	}

	pairs := make(map[rune]rune)
	pairs['('] = ')'
	pairs['{'] = '}'
	pairs['['] = ']'

	isOpening := func(c rune) bool {
		_, found := pairs[c]
		return found
	}

	isClosing := func(c rune) bool {
		for k := range pairs {
			if c == pairs[k] {
				return true
			}
		}
		return false
	}

	stack := []rune{}
	for _, c := range s {
		if isOpening(c) {
			stack = append(stack, c)
		} else if isClosing(c) {
			if len(stack) == 0 {
				return false
			}

			var o rune
			o, stack = stack[len(stack)-1], stack[:len(stack)-1]

			if c != pairs[o] {
				return false
			}
		}
	}

	return len(stack) == 0
}
