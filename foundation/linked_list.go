package foundation

// LinkedListNode represents singly-linked list data structure.
type LinkedListNode struct {
	Element interface{}
	Next    *LinkedListNode
}

// Len calculates the length of the linked list.
func (list LinkedListNode) Len() (l int) {
	l = 1
	for curr := &list; curr.Next != nil; curr = curr.Next {
		l++
	}
	return
}
