package main

import (
	"config"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai"
	itemsaver "crawler_distributed/persist/client"
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker/client"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverPort      = flag.Int("saver_host", 0, "the port for me to listen on")
	workerProcessHosts = flag.String("worker_hosts", "", "worker hosts(comma separated)")
)

func main() {
	flag.Parse()
	saver, err := itemsaver.ItemSaver(*itemSaverPort)
	if err != nil {
		panic(err)
	}
	processor, err := client.CreateProcessor(createProcessClientPool(strings.Split(*workerProcessHosts, ",")))
	if err != nil {
		panic(err)
	}
	e := &engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         saver,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url: config.ZAMainUrl,
		Parser: &engine.NewFunParser{
			Func: zhenai.ParserCityList,
			Name: "ParserCityList",
		},
	})

}

func createProcessClientPool(host []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range host {
		c, err := rpcsupport.NewClient(fmt.Sprintf(":%s", h))
		if err != nil {
			log.Printf("error connecting to port:%s", h)
			continue
		}
		log.Printf("connecting to port:%s", h)
		clients = append(clients, c)
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out

}
