package main

import (
	persist2 "crawler/persist"
	"crawler_distributed/persist"
	"crawler_distributed/rpcsupport"
	"log"
)

func main() {
	log.Fatal(Server(":1234", "user"))
}

func Server(host string, index string) error {

	return rpcsupport.ServerRpc(host, &persist.ItemSaveServer{
		Client: persist2.InitElastic(),
		Index:  index,
	})

}
