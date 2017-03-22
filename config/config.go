package config

import (
	"log"

	"github.com/hashicorp/hcl"
)

type wombat struct {
	ID     string `hcl:"-"`
	Queues []*foo `hcl:"queue,expand"`
}

type foo struct {
	Name   string `hcl:,key`
	Driver string `hcl:"type"`
	Host   string `hcl:"host"`
	Port   int    `hcl:"port"`
}

func Parse(input string) (out wombat, err error) {
	log.Println("Parsing Config\n", input)
	err = hcl.Decode(&out, input)
	log.Printf("Redis Connect: %v\n", out.ID)
	log.Printf("  %+v\n", out.Queues[0])
	return
}
