VERSION := $(shell cat VERSION)
BINARY = kubeq
DEPS := $(shell find . -type f -iname '*.go' | egrep -v _test.g-o$)
TESTS := $(shell find . -type f -iname *_test.go)

all: $(BINARY)

$(BINARY): $(DEPS)
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .

image: $(BINARY)
	docker build -t smo921/$(BINARY):$(VERSION) .
	docker push smo921/$(BINARY):$(VERSION)

test: $(TESTS)
	go test ./...

clean:
	rm -f $(BINARY)
