package main

import (
	"errorhandling/filelistingserver/filelisting"
	"errorhandling/filelistingserver/global"
	"net/http"
	_ "net/http/pprof"
)

// todo:做一个显示文件的server
func main() {
	rwt := filelisting.Rwt{HeaderPath: "/list/"}
	http.HandleFunc("/", global.ErrWrapper(rwt.ResponseWriter))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
		return
	}
}
