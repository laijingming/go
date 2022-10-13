package queue

type Queue []int

//插入队列
func (q *Queue) Push(v int){
	*q = append(*q,v)
}