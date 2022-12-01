package common

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type HttpStruct struct {
	Url      string
	ParamStr string
	Method   string
	Header   http.Header
}

func (ps *HttpStruct) HttpBase() (*http.Response, error) {
	req, _ := http.NewRequest(ps.Method, ps.Url, strings.NewReader(ps.ParamStr))
	req.Header = ps.Header
	//req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	return client.Do(req)
}

type PostStruct struct {
	Url      string
	ParamStr string
	Header   http.Header
}

func (ps PostStruct) basePost() (*http.Response, error) {

	req, _ := http.NewRequest(http.MethodPost, ps.Url, strings.NewReader(ps.ParamStr))
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

//HttpGet 基础get请求
func HttpGet(url string) []byte {
	if strings.Index(url, "https") == -1 {
		url = strings.Replace(url, "http", "https", 1)
	}
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	request.Header.Add("cookie", "sid=14f7b3a6-2efd-4bb9-bdee-b59d40a974f1; ec=rPr4ZFqN-1631082127157-a9662a3bbd7b3-666224740; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1668415549; hccesp_lttk=AAAAAgAAAAAAAAAFAAAAAQAAAAeBwwi0wpEfjAVEAd2OIgKgWhPcfOHSVjvgoSJ3HeqF0QAAAAAAAAAAAAAAQHuCWxyd3QUuIz9Koo5pWAoLEU7Bl5VDEtcn/hqylJ6rNGYfFCKkrqelJ3+DzLwb5vuN76fJIQecpuKhmqSOkOE=; FSSBBIl1UgzbN7NO=5cywccPaKuutAeBVB4yzslilgYzrjGijwFPBeG5K5SLfxf7M6mOKN8LPai0ox180DDmREARRX1.90xMSker1n1G; hcbm_tk=MjZvkSuwF9k49wWeq8hF3nE0yS5XhXRnME7swrXaN/PnRNl/GKtUgjby8sffl00Q0Jl/fGD6Bv3fCMFiJxqfvJPI8AW128i/EkRxTyRWA+1ka9cyNolYelpYji9NAXz2otN2fjKK77TscE4s5mt7L27r1QxOIepMEhXvYDiCCGEElymWW7R2/vnIk3KLbmGkCY0hQumPk5FTlI4yeawODwVTXh0GUJo/OydLuScdvCESbMS7kg6+wVeyj6CoTT+7Lggz75qVfLEDSnfszV3oQuqb9HSM+QmF5065nUJULdd+HD9zWqekwPDpW72aonLORdLz7g5whq38UWVgF5HR+WqftHFNXGQzWE1ZRwe43Xvc47axwaV80YeBzLitkrvwEDd2zoiP75awPzZt4/npScnK0O1CYbsPIGDKNqzWzQ==_RERXV4AmIa4VjY7e; FSSBBIl1UgzbN7NP=5361siDJ09ulqqqDlITZYQAfUuiIa1mgnI9._Mxt9DM2HHuMaVv6ZIZU1frhGMKwroXHWmWZGIyK54Rv9EcbaR_Ew0ZRy7ZQzCoMQXm0FTZGgY1z7solyY02yOvYB5qiYtrTm4ZcAX_lJoavpDB97ZtPDM2aFqldlKNN1ulVDOkHrfK2GlWRWqnNaHloHP1bXv1z3ehUH.913rc3xiY5qfnyIfW65qTtBmTWh0VnBhK5V_tl3rO8vRr2aC_si6D2J9; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1668584242; _efmdata=dyY8qPJvZQjO4CreY%2FTLLk42f4COphOi34kNpyobU12eSzK62Musuq0WGynAhTPLGv6S%2Bvm74cQIHTXIF2noCMNAlRyVcy3re%2BqrrfrDl1E%3D; _exid=Iv%2Fhiboc4nYawAIlaW8YDSlva7KCDCm3n2JQ8LJPsUn9baVJzsPsjysH7OrCFzy5MfPfmU1UW2%2FVhkiZV%2Bgeeg%3D%3D")
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	//DumpResponse头部一起返回
	//httputil.DumpResponse(response, true)
	if response.StatusCode != http.StatusOK {
		fmt.Printf("wrong status code:%d %s\n", response.StatusCode, url)
		return nil
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return bytes
}

const baseDir = "./src/"

func ReadFile(name string) []byte {
	file, err := os.ReadFile(baseDir + name)
	if err != nil {
		return nil
	}
	return file
}

func WriteFile(name string, contents []byte) {
	err := os.WriteFile(baseDir+name, contents, 0666)
	if err != nil {
		panic(err)
	}

}

type CrawlerRequestStruct struct {
	url       string
	ParserFun func([]byte) CrawlerParserResultStruct
}

type CrawlerParserResultStruct struct {
	requests []CrawlerRequestStruct
	items    []interface{}
}

func CrawlerRun(CRS CrawlerRequestStruct) {

}
