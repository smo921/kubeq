package config

import (
	"log"
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	input := `
  queue "test" {
    type = "test_type"
    host = "localhost"
    port = 1234
  }
`
	expect := config{
		Queues: []*queue{
			&queue{
				Driver: "test_type",
				Host:   "localhost",
				Port:   1234,
			},
		},
	}
	out, err := Parse(input)
	if err != nil {
		log.Println(err)
	}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("bad: %#v !== %#v", out, expect)
	}
}
