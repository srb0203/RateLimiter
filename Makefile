all: build

build:
	go build -o ratelimiter

test:
	go test ./... -v

clean:
	go clean

image:
	docker build -t cirocosta/l7 .