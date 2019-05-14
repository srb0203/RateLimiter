all: build

#build executable
build:
	go build -o ratelimiter

#run tests
test:
	go test ./... -v

#clean the project
clean:
	go clean

#make docker image
image:
	docker build -t ratelimiter .

#run docker container
image-run:
	docker run -p 8050:8050 ratelimiter

#get test coverage results
cover:
	go test -coverprofile=cover.out

#generate html file for coverage results
coverfile:
	go tool cover -html=cover.out -o cover.html