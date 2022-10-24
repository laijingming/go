package queue

import "fmt"

func ExampleQueue_Push() {
	queue := Queue{1}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push("abc")
	queue.Pop()
	fmt.Println(queue)

	// Output:
	// [1 2 3]
}
