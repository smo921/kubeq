package config

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/hcl"
)

type config struct {
	ID     string   `hcl:"-"`
	Queues []*Queue `hcl:"queue,expand"`
}

type Queue struct {
	Name     string        `hcl:,key`
	Driver   string        `hcl:"type"`
	Host     string        `hcl:"host"`
	Port     int           `hcl:"port"`
	Database int           `hcl:"db"`
	Password string        `hcl:"password"`
	Timeout  time.Duration `hcl:timeout`
}

func Parse(input string) (out config, err error) {
	log.Println("Parsing Config\n", input)
	err = hcl.Decode(&out, input)
	if err == nil {
		for _, q := range out.Queues {
			if q.Timeout == 0 {
				log.Println("No timeout specified, defaulting to 2 seconds.")
				q.Timeout = 2 * time.Second
			}
		}
	}
	return
}

func (q *Queue) Address() string {
	return fmt.Sprintf("%s:%d", q.Host, q.Port)
}
