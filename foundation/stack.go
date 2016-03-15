package foundation

// Stack represents a "stack" structure.
type Stack []interface{}

// Len returns length of the given stack.
func (stack Stack) Len() (l int) {
	l = len(stack)
	return
}
