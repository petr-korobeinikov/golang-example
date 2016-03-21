package foundation

// LinkedListNode represents singly-linked list data structure.
type LinkedListNode struct {
	Element interface{}
	Next    *LinkedListNode
}

// Len calculates the length of the linked list.
func (list LinkedListNode) Len() (l int) {
	l = 0
	for curr := &list; curr != nil; curr = curr.Next {
		l++
	}
	return
}

func (list LinkedListNode) Last() (element *LinkedListNode) {
	for curr := &list; curr != nil; curr = curr.Next {
		element = curr
	}
	return
}
