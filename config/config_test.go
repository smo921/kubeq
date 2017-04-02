package config

import (
	"reflect"
	"testing"
	"time"
)

func TestParseConfig(t *testing.T) {
	input := `
queue "test" {
  type = "test_type"
  host = "localhost"
  port = 1234
}
queue "test2" {
	type = "test_type"
	host = "testhost"
	port = 7890
}
`
	expect := config{
		Queues: []*Queue{
			&Queue{
				Name:    "test",
				Driver:  "test_type",
				Host:    "localhost",
				Port:    1234,
				Timeout: 2 * time.Second,
			},
			&Queue{
				Name:    "test2",
				Driver:  "test_type",
				Host:    "testhost",
				Port:    7890,
				Timeout: 2 * time.Second,
			},
		},
	}
	out, err := Parse(input)
	if err != nil {
		t.Log(err)
	}
	t.Logf("%#v\n", out)
	for _, q := range out.Queues {
		t.Logf("%#v\n", q)
	}
	if !reflect.DeepEqual(out, expect) {
		t.Fatalf("bad: %#v !== %#v", out, expect)
	}
}
