package config

import (
	"io/ioutil"
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

func Parse() {
	if hclText, err := ioutil.ReadFile("./kubeq.conf"); err == nil {
		var out wombat
		log.Println("Parsing Config\n", string(hclText))
		err = hcl.Decode(&out, string(hclText))
		log.Printf("Redis Connect: %v\n", out.ID)
		log.Printf("  %+v\n", out.Queues[0])
	} else {
		log.Println(err)
	}
}
