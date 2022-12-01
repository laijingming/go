package persist

import (
	"fmt"
	"strings"
)

func Save() chan interface{} {
	out := make(chan interface{})
	itemNum := 0
	go func() {
		for {
			item := <-out
			strArr := strings.Split(item.(string), ":")
			fmt.Println(strArr)
			fmt.Printf("Got %d item:%v\n", itemNum, item)
			itemNum++
		}
	}()
	return out
}
