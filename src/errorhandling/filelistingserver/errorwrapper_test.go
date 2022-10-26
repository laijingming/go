package main

import (
	"errorhandling/filelistingserver/global"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type testUserMessage string

func (e testUserMessage) Error() string {
	return string(e)
}

func (e testUserMessage) Message() string {
	return string(e)
}

func errorPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}
func errorUserError(writer http.ResponseWriter, request *http.Request) error {
	return testUserMessage("user error")
}

var tests = []struct {
	h       global.AppHandler
	code    int
	message string
}{
	{errorPanic, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)},
	{errorUserError, http.StatusBadRequest, "user error"},
}

func TestErrWrapper(t *testing.T) {

	for _, test := range tests {
		f := global.ErrWrapper(test.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com" /*target*/, nil, /*body*/
		)
		f(response, request)
		verifyResponse(response.Result(), test.code, test.message, t)
	}
}
func TestErrWrapperInServer(t *testing.T) {
	for _, test := range tests {
		f := global.ErrWrapper(test.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, test.code, test.message, t)
	}
}

func verifyResponse(resp *http.Response,
	expectedCode int, expectedMsg string,
	t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode ||
		body != expectedMsg {
		t.Errorf("expect (%d, %s); "+
			"got (%d, %s)",
			expectedCode, expectedMsg,
			resp.StatusCode, body)
	}
}
