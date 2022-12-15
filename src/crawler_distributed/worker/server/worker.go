package main

import (
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
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
	log.Fatal(Server(fmt.Sprintf(":%d", *port)))
}

func Server(host string) error {
	return rpcsupport.ServerRpc(host, &worker.CrawlerServer{})

}
