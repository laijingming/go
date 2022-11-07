package common

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type PostStruct struct {
	url      string
	paramStr string
	method   string
	header   http.Header
}

func (ps PostStruct) basePost() (*http.Response, error) {
	req, _ := http.NewRequest(ps.method, ps.url, strings.NewReader(ps.paramStr))
	req.Header = ps.header
	//req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	return client.Do(req)
}

//HttpPostUnmarshal json.Unmarshal进行解码
//执行时常90.18ms
func (ps PostStruct) httpPostUnmarshal() (interface{}, error) {
	var result interface{}
	resp, _ := ps.basePost()
	//此处request是http请求得到的json格式数据-》然后转化为【】byte格式数据
	resJson, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		data := []byte(string(resJson))
		err = json.Unmarshal(data, &result)
	}
	return result, err
}

//HttpPostNewDecoder json.NewDecoder解码
//执行时常58.8713ms
func (ps PostStruct) httpPostNewDecoder() (interface{}, error) {
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
