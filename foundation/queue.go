package foundation

// Queue represents a "queue" structure.
type Queue []interface{}

// Len returns length of the given queue.
func (queue Queue) Len() (l int) {
	l = len(queue)
	return
}

// Enqueue appends element to the queue.
func (queue *Queue) Enqueue(element interface{}) {
	*queue = append(*queue, element)
}
