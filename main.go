package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/smo921/kubeq/config"
	"github.com/smo921/kubeq/queue/redis"
)

const schedulerName = "kubeq"

func monitorJobQueue(qconf config.Queue, done chan struct{}, wg *sync.WaitGroup) {
	jobs, errc := redis.Read(qconf)
	for {
		select {
		case job := <-jobs:
			log.Printf("JOB:\n\tImage: '%v'\n\tArgs: '%v'\n", job.Image, job.Arguments)
		case err := <-errc:
			log.Println("ERROR:", err)
		case <-done:
			wg.Done()
			log.Println("monitorJobQueue done")
			return
		}
	}
}

func main() {
	var hclText []byte
	var err error

	log.Println("Starting kubeq scheduler . . .")
	if hclText, err = ioutil.ReadFile("./kubeq.conf"); err != nil {
		log.Fatal(err)
	}

	conf, err := config.Parse(string(hclText))
	log.Printf("Redis Connect: %v\n", conf.ID)
	log.Printf("  %+v\n", conf.Queues[0])

	doneChan := make(chan struct{})
	var wg sync.WaitGroup

	for _, qconf := range conf.Queues {
		wg.Add(1)
		go monitorJobQueue(*qconf, doneChan, &wg)
	}

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
