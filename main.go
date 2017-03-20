package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-redis/redis"
)

const schedulerName = "kubeq"

func doRedisStuff() (result string, err error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	result, err = client.Ping().Result()
	client.Close()
	return
}

func monitorJobQueue(done chan struct{}, wg *sync.WaitGroup) {
	foo, _ := doRedisStuff()
	log.Printf("doRedisStuff says: %s\n", foo)
	for {
		select {
		case <-done:
			wg.Done()
			log.Println("monitorJobQueue done")
			return
		}
	}
}

func main() {
	log.Println("Starting kubeq scheduler . . .")

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
