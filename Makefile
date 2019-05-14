all: build

build:
	go build -o ratelimiter

test:
	go test ./... -v

image:
	docker build -t cirocosta/l7 .