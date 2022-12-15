package itemsaver

import (
	"crawler/model"
	"crawler_distributed/rpcsupport"
	"fmt"
)

func ItemSaver(host int) (chan model.User, error) {
	client, err := rpcsupport.NewClient(fmt.Sprintf(":%d", host))
	if err != nil {
		return nil, err
	}
	out := make(chan model.User)
	go func() {
		for {
			user := <-out
			var result string
			err = client.Call("ItemSaveServer.Save", user, &result)
			if result == "ok" {
				fmt.Printf("Got save #%s item:%v\n", user.Id, user)
			}

		}
	}()
	return out, nil
}
