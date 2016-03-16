package foundation

type Queue []interface{}

func (queue Queue) Len() (l int) {
	l = len(queue)
	return
}

func (queue *Queue) Enqueue(element interface{}) {
	*queue = append(*queue, element)
}
