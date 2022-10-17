package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Rwt struct {
	HeaderPath string
}

type userMessage string

func (e userMessage) Error() string {
	return string(e)
}

func (e userMessage) Message() string {
	return string(e)
}

func (rwt Rwt) ResponseWriter(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, rwt.HeaderPath) != 0 && request.URL.Path != "/favicon.ico" {
		return userMessage("url need " + rwt.HeaderPath)
	}
	path := request.URL.Path[len(rwt.HeaderPath):]
	file, err := os.Open(path)
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

func TestUserMessage() error {
	return userMessage("user message")
}
