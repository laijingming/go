package main

import (
	persist2 "crawler/persist"
	"crawler_distributed/persist"
	"crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Println("must specify port")
		return
	}
	log.Fatal(Server(fmt.Sprintf(":%d", *port), "user"))
}

func Server(host string, index string) error {

	return rpcsupport.ServerRpc(host, &persist.ItemSaveServer{
		Client: persist2.InitElastic(),
		Index:  index,
	})

}
