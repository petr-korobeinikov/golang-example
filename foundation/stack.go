package foundation

// Stack represents a "stack" structure.
type Stack []interface{}

// Len returns length of the given stack.
func (stack Stack) Len() (l int) {
	l = len(stack)
	return
}

// Push appends element to the end of the given stack.
func (stack *Stack) Push(element interface{}) {
	*stack = append(*stack, element)
}
