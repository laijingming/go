- channel
    - 处理goroutine与goroutine之前的通讯
- buffered channel
- range
- 不要通过共享内存来通信，通过通信来共享内存

### 使用channel来等待goroutine结束
- 简单使用
```
func main() {
	chanDemo2()
}

func chanDemo2() {
	c := make(chan int) //初始化chan
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	for i := 0; i < 10; i++ {
		c <- i
	}
	time.Sleep(time.Microsecond)
}
```

```
package main

import (
	"fmt"
	"sync"
)

type woker struct {
	in   chan int
	done func()
}

//func createWorker(id int) chan int {
//func createWorker(id int) <-chan int {//只允从chan获得数据
//func createWorker(id int) chan<- int { //只允发数据给chan
func createWorker(id int, wg *sync.WaitGroup) woker { //只允发数据给chan
	w := woker{
		make(chan int),
		func() {
			wg.Done()
		},
	}
	go doWorker(id, w)
	return w
}

func doWorker(id int, w woker) {
	for n := range w.in {
		fmt.Printf("doWorker %d,received %c\n", id, n)
		w.done()
	}
}

func chanDemo() {
	var w sync.WaitGroup
	var workers [10]woker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &w)
	}
	w.Add(20)
	for i, w := range workers {
		w.in <- 'a' + i
	}

	for i, w := range workers {
		w.in <- 'A' + i
	}
	w.Wait()
}

func main() {
	chanDemo()
	//bufferedChannel()
}

```
### 使用channel来实现树的遍历
```main.go
package main

import (
	"fmt"
	tree "study/tree/functional"
)

func main() {
	var t tree.TN
	t = tree.TN{Value: 3}
	t.Left = &tree.TN{}
	t.Right = &tree.TN{Value: 5}
	t.Right.Left = &tree.TN{Value: 4}
	t.Left.Right = &tree.TN{Value: 22}
	t.Traverse()
	tCount := 0
	t.TraverseFunc(func(tn *tree.TN) {
		tCount++
	})
	fmt.Printf("t count:%d", tCount)
	//var mt myTreeNode
	//mt.node = &t
	//mt.MyTraverse()

	channel := t.TraverseWithChannel()
	maxNode := 0
	for c := range channel {
		if maxNode < c.Value {
			maxNode = c.Value
		}
	}
	fmt.Println("\nMax node value：", maxNode)

}

//扩展别人类型
type myTreeNode struct {
	node *tree.TN
}

func (myNt *myTreeNode) MyTraverse() {
	if myNt.node == nil || myNt.node == nil {
		return
	}
	left := myTreeNode{myNt.node.Left}
	right := myTreeNode{myNt.node.Right}
	left.MyTraverse()
	right.MyTraverse()
	myNt.node.Print()
}

```
```node.go
package tree

func (n *TN) Traverse() {
	n.TraverseFunc(func(tn *TN) {
		tn.Print()
	})
}

func (n *TN) TraverseFunc(f func(*TN)) {
	if n == nil {
		return
	}
	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}

func (n *TN) TraverseWithChannel() chan *TN {
	c := make(chan *TN)
	go func() {
		n.TraverseFunc(func(n *TN) {
			c <- n
		})
		close(c)
	}()
	return c
}

```
### select进行调度
```
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d,received %d\n", id, n)
	}
}
func createWorker(id int) chan<- int { //只允发数据给chan
	c := make(chan int)
	go worker(id, c)
	return c
}
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			//随机1500毫秒
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Microsecond)
			//time.Sleep(1 * time.Second)
			out <- i
			i++
		}
	}()
	return out
}
func main() {
	//c1, c2 := generator(), generator()
	//for  {
	//	select {
	//	case n := <-c1:
	//		fmt.Println("Received from c1:",n)
	//	case n := <-c2:
	//		fmt.Println("Received from c1:",n)
	//
	//	}
	//}
	c1, c2 := generator(), generator()
	var valuesInt []int
	w := createWorker(1)
	n := 0
	stopTime := time.After(10 * time.Second)
	tk := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(valuesInt) > 1 {
			activeWorker = w
			activeValue = valuesInt[0]
		}
		select {
		case n = <-c1:
			valuesInt = append(valuesInt, n)
		case n = <-c2:
			valuesInt = append(valuesInt, n)
		//case <-time.After(time.Microsecond * 800):
		//	fmt.Println("timeout")
		case activeWorker <- activeValue:
			valuesInt = valuesInt[1:]
		case <-tk:
			fmt.Println("valuesInt len=", len(valuesInt))
		case <-stopTime:
			fmt.Println("stop run")
			return //结束

		}
	}
}

```
### 传统同步机制
- WaitGroup
- Mutex
- Cond
```
package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("save increment")
	func() {//defer控制在函数体
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Microsecond)
	fmt.Println(a.get())
}

```