package foundation

import "errors"

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

// Dequeue dequeues element from the queue.
func (queue *Queue) Dequeue() (element interface{}, err error) {
	q := *queue
	l := q.Len()
	if l == 0 {
		err = errors.New("Can't dequeue element from an empty queue.")
		return
	}
	element = q[0]
	*queue = q[1:]
	return
}
