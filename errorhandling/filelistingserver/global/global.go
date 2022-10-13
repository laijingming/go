package global

import (
	"log"
	"net/http"
	"os"
)

type AppHandler func(writer http.ResponseWriter, request *http.Request) error

type userMessage interface {
	error
	Message() string
}

func ErrWrapper(handler AppHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		// 错误统一处理
		if err != nil {
			//import "github.com/gpmgo/gopm/modules/log"
			//log.Warn("Error handling request: %s",
			//  err.Error())
			log.Printf("Error handling request: %s",
				err.Error())
			// Type Assertion
			if userErr, ok := err.(userMessage); ok {
				log.Println("came")
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			// 错误都写在writer里
			http.Error(writer,
				http.StatusText(code),
				code)
		}
	}
}
