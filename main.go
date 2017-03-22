package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hashicorp/hcl"
	"github.com/smo921/kubeq/queue"
)

const schedulerName = "kubeq"

func monitorJobQueue(done chan struct{}, wg *sync.WaitGroup) {
	foo, err := queue.DoRedisStuff()
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	} else {
		log.Printf("DoRedisStuff says: '%s'\n", foo)
	}
	for {
		select {
		case <-done:
			wg.Done()
			log.Println("monitorJobQueue done")
			return
		}
	}
}

func parseConfig() {
	if hclText, err := ioutil.ReadFile("./kubeq.conf"); err == nil {
		var out map[string]string
		err = hcl.Decode(&out, string(hclText))
		log.Printf("Redis Connect: %s:%s\n", out["redis_host"], out["redis_port"])
	} else {
		log.Println(err)
	}
}

func main() {
	log.Println("Starting kubeq scheduler . . .")

	parseConfig()
	doneChan := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go monitorJobQueue(doneChan, &wg)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-sigChan:
			log.Println("Shutdown signal received . . .")
			os.Exit(0)
		}
	}
}
