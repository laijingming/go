package main

import (
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
)

const baseUrl = "/list/"

// todo:做一个显示文件的server
func main() {
	//rwt := filelisting.Rwt{HeaderPath: "/list/"}
	//http.HandleFunc("/", global.ErrWrapper(rwt.ResponseWriter))
	//err := http.ListenAndServe(":8888", nil)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	showFileNew()
	openServer()
}

//region 通过服务器展示文件1
func showFile() {
	http.HandleFunc(baseUrl, func(writer http.ResponseWriter, request *http.Request) {
		url := "." + request.URL.Path
		file, err := os.Open(url)
		if err != nil {
			//writer.Write([]byte("路径为：" + request.URL.Path + "，的文件不存在。"))
			http.Error(writer,
				err.Error(),
				http.StatusInternalServerError)
			return
		}

		defer file.Close()

		all, err := ioutil.ReadAll(file)
		if err != nil {
			//在http server 里面panic不会被宕掉
			panic(err)
		}
		writer.Write(all)
	})
}

//endregion

//region 通过服务器展示文件2

func showFileNew() {
	sfh := sFH{baseUrl}
	http.HandleFunc(baseUrl, errorFileHandle(sfh.showFileHandle))
}

//定义函数错误类型
type tFileHandle func(writer http.ResponseWriter, request *http.Request) error
type sFH struct {
	headUrl string
}
type userMessage interface {
	error
	Message() string
}
type errMsg string

func (e errMsg) Error() string {
	return e.Message()
}

func (e errMsg) Message() string {
	return string(e)
}

//统一错误处理
func errorFileHandle(tfh tFileHandle) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("panic：", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}
		}()
		err := tfh(writer, request)
		if err != nil {
			if userErr, ok := err.(userMessage); ok {
				log.Println("came", userErr.Message())
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			log.Println("error handling request：", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code),
				code)
		}
	}
}

func (sfh sFH) showFileHandle(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, sfh.headUrl) != 0 {
		return errMsg("url need " + sfh.headUrl)
	}
	url := request.URL.Path[len(sfh.headUrl):]
	file, err := os.Open(url)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}

//endregion

//开放服务器，监听：8888
func openServer() {
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
		return
	}
}
