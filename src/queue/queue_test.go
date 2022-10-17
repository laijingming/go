package queue

import (
	"fmt"
)

func ExampleQueue_Push() {
	queue := Queue{1}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue)

	// Output:
	// [1 1 2 3]
	// 2
	// 3
}
