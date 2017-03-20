VERSION := $(shell cat VERSION)
BINARY = kubeq

all: main.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
	docker build -t smo921/$(BINARY):$(VERSION) .
	docker push smo921/$(BINARY):$(VERSION)

clean:
	rm -f $(BINARY)
