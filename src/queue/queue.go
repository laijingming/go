package queue

type Queue []interface{}

// Push 插入队列
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}
