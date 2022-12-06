package persist

import (
	"crawler/model"
	"encoding/json"
	"fmt"
	"testing"
)

func TestItemSaver(t *testing.T) {
	es := InitElastic()
	user := model.User{
		Name: "小强",
		Id:   "1",
	}
	id := es.save("user", user.Id, user)
	result := es.get("user", id)
	//反序列化
	var getUser model.User
	err := json.Unmarshal(result.Source, &getUser)
	if err != nil {
		return
	}
	fmt.Println(getUser)
	if getUser != user {
		t.Errorf("Got getUser:%v,user:%v", getUser, user)
	}
}
