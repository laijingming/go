package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type PostStruct struct {
	Url      string
	ParamStr string
	Method   string
	Header   http.Header
}

func (ps PostStruct) basePost() (*http.Response, error) {

	req, _ := http.NewRequest(ps.Method, ps.Url, strings.NewReader(ps.ParamStr))
	req.Header = ps.Header
	//req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	return client.Do(req)
}

//HttpPostUnmarshal json.Unmarshal进行解码
//执行时常 63.225648ms
func (ps PostStruct) HttpPostUnmarshal() (interface{}, error) {
	var result interface{}
	resp, _ := ps.basePost()
	//此处request是http请求得到的json格式数据-》然后转化为【】byte格式数据
	resByte, err := ioutil.ReadAll(resp.Body)
	resJson := string(resByte)
	index := strings.Index(resJson, "<pre>")
	if index != -1 {
		resJson = resJson[0:index]
	}
	if err == nil {
		data := []byte(resJson)
		err = json.Unmarshal(data, &result)
	}
	return result, err
}

//HttpPostNewDecoder json.NewDecoder解码
//执行时常63.116348ms
func (ps PostStruct) HttpPostNewDecoder() (interface{}, error) {
	resp, _ := ps.basePost()
	var result interface{}
	err := json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

/*
二、区别：
1、json.NewDecoder是从一个流里面直接进行解码，代码精干；
2、json.Unmarshal是从已存在与内存中的json进行解码；
3、相对于解码，json.NewEncoder进行大JSON的编码比json.marshal性能高，因为内部使用pool。

三、场景应用：
1、json.NewDecoder用于http连接与socket连接的读取与写入，或者文件读取；
2、json.Unmarshal用于直接是byte的输入。
*/

type TimeStruct struct {
	Time time.Time
}

//CalculateTime 保存时间
func (t *TimeStruct) CalculateTime() {
	t.Time = time.Now()
}

//SliceTime 计算时间差
func (t *TimeStruct) SliceTime() time.Duration {
	since := time.Since(t.Time)
	t.Time = time.Now()
	return since
}

//PrintSliceTime 输出时间差
func (t *TimeStruct) PrintSliceTime() {
	fmt.Println(t.SliceTime())
}

type ChanStruct struct {
	CFun func(chan interface{})
	WNum int
}

//CreateWorker 创建协程
func CreateWorker(cs ChanStruct) chan interface{} {
	c := make(chan interface{})
	for i := cs.WNum; i > 0; i-- {
		go cs.CFun(c)
	}
	return c
}
