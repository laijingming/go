package main

import (
	"config"
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
	"fmt"
	"log"
)

func main() {
	log.Fatal(Server(fmt.Sprintf(":%d", config.WorkerPort0)))
}

func Server(host string) error {
	return rpcsupport.ServerRpc(host, &worker.CrawlerServer{})

}
