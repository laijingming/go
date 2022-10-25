package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	//TODO implement me
	return fmt.Sprintf("implement me：%s", r.Contents)
}

func (r *Retriever) Get(url string) string {
	return r.Contents
}
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = url
	return "ok"
}
