all: build

build:
	go build -o ratelimiter

test:
	go test ./... -v

clean:
	go clean

image:
	docker build -t ratelimiter .

image-run:
	docker run -p 8050:8050 ratelimiter