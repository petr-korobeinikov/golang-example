package foundation

// LinkedList represents signly-linked list data structure.
type LinkedList struct {
	Element interface{}
	Next    *LinkedList
}

// Len calculates the length of LinkedList.
func (list LinkedList) Len() (l int) {
	l = 1
	for curr := &list; curr.Next != nil; curr = curr.Next {
		l++
	}
	return
}
