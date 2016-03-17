package foundation

import "errors"

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

// Pop pops element from the end (top) of the given stack.
func (stack *Stack) Pop() (element interface{}, err error) {
	s := *stack
	l := s.Len()
	if l == 0 {
		err = errors.New("Unable to pop element from empty stack.")
		return
	}
	element = s[l-1]
	*stack = s[:l-1]
	return
}

// Top returns top element of the stack.
func (stack Stack) Top() (element interface{}, err error) {
	l := stack.Len()
	if l == 0 {
		err = errors.New("There is no top element in empty stack.")
		return
	}
	element = stack[l-1]
	return
}
