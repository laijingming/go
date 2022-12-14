package main

import (
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
	"log"
)

func main() {
	log.Fatal(Server(":9201"))
}

func Server(host string) error {
	return rpcsupport.ServerRpc(host, &worker.CrawlerServer{})

}
