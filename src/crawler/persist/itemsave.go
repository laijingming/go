package persist

import "fmt"

func Save() chan interface{} {
	out := make(chan interface{})
	itemNum := 0
	go func() {
		for {
			item := <-out
			fmt.Printf("Got %d item:%v\n", itemNum, item)
			itemNum++
		}
	}()
	return out
}
