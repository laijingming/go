package main

import (
	"config"
	persist2 "crawler/persist"
	"crawler_distributed/persist"
	"crawler_distributed/rpcsupport"
	"fmt"
	"log"
)

func main() {
	log.Fatal(Server(fmt.Sprintf(":%d", config.ItemSaver0), "user"))
}

func Server(host string, index string) error {

	return rpcsupport.ServerRpc(host, &persist.ItemSaveServer{
		Client: persist2.InitElastic(),
		Index:  index,
	})

}
