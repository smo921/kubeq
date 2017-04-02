package redis

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/smo921/kubeq/config"
	"github.com/smo921/kubeq/job"
)

// Read pulls jobs from the redis queue and sends them to the job channel
func Read(conf config.Queue) (<-chan job.Job, <-chan error) {
	client := redis.NewClient(&redis.Options{
		Addr:        conf.Address(),
		Password:    conf.Password, // no password set
		DB:          conf.Database, // use default DB
		DialTimeout: conf.Timeout * time.Second,
	})

	jobc := make(chan job.Job)
	errc := make(chan error, 1)

	go func() {
		for {
			result, err := client.LPop(conf.Name).Result()
			if err == redis.Nil {
				log.Println("No jobs found.")
				time.Sleep(10 * time.Second)
				continue
			} else if err != nil {
				errc <- err
				time.Sleep(10 * time.Second)
				continue
			}

			var job job.Job
			err = json.Unmarshal([]byte(result), &job)
			if err != nil {
				log.Printf("ERROR: %s\nUnable to create job from: %v", err, result)
				time.Sleep(10 * time.Second)
				continue
			}
			jobc <- job
		}
		client.Close()
		return
	}()

	return jobc, errc
}
