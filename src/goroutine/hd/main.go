package main

import (
	"fmt"
	"lib/common"
	"net/http"
	"sync"
	"time"
)

const url = "http://dushitime.test.commpad.cn/cmd_t.php?tpass=abcddd123&uid=1000465&sevid=1"

func main() {
	ps := common.PostStruct{
		Url:      url,
		ParamStr: `{"huodong2":{"hd3063Create":{"name":"队1"}}}`,
		//ParamStr: `{"huodong2":{"hd3063Create":{"name":"测试"}}}`,
		//ParamStr: `{"huodong2":{"hd3063Info":{"name":"测试"}}}`,
		Method: http.MethodPost,
		Header: map[string][]string{
			"Content-Type": {"application/json"},
		},
	}
	workerNum, allNum := 10, 10
	var tSince common.TimeStruct
	tSince.CalculateTime()
	result := syncSelect(ps, workerNum, allNum)
	fmt.Println("syncWaitGroup耗时：", tSince.SliceTime(), " 开启协程数量：", workerNum, " 总工作量：", allNum, " 当前成功接收次数：", len(result))
	for k, v := range result {
		fmt.Println(k, v)
	}
	fmt.Println(len(result))
	time.Sleep(time.Second * 10)
}

//method1 传统同步机制：sync.WaitGroup
func syncWaitGroup(ps common.PostStruct, workerNum int, allNum int) map[int]interface{} {
	result := make(map[int]interface{})
	var wg sync.WaitGroup
	wg.Add(workerNum)
	worker := common.CreateWorker(common.ChanStruct{
		CFun: func(cs chan interface{}) {
			for i := 0; i < allNum/workerNum; i++ {
				res, _ := ps.HttpPostUnmarshal()
				cs <- res
			}
			wg.Done()
		},
		WNum: workerNum,
	})
	successNum := 0
	go func() {
		for w := range worker {
			result[successNum] = w
			successNum++
		}
	}()
	wg.Wait()
	return result
}

//method2 select
func syncSelect(ps common.PostStruct, workerNum int, allNum int) map[int]interface{} {
	result := make(map[int]interface{})
	worker := common.CreateWorker(common.ChanStruct{
		CFun: func(cs chan interface{}) {
			for {
				res, _ := ps.HttpPostUnmarshal()
				cs <- res
			}
		},
		WNum: workerNum,
	})
	successNum := 0
	stopTime := time.After(10 * time.Second) //超过10秒结束
	for {
		isDone := false
		select {
		case result[successNum] = <-worker:
			successNum++
			if successNum >= allNum {
				isDone = true //结束
			}
		case <-stopTime:
			isDone = true //结束
		}
		if isDone {
			break
		}
	}
	return result
}
