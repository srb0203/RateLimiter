# RateLimiter

HTTP API Rate limiter in Go.

# Description

Ratelimiter limits the number of calls to an HTTP API. It uses a credit based system. Each IP address is assigned a specified number of credits which corresponds to the number of requests it can make in a given amount of time. The credit reduces by one on each request. The credits are replenished after the given amount of time has elapased. For example if the app allows for 10 requests per hour then the credits will be replenished after an hour.

Dockerfile is provided to build a small docker image than can be deployed anywhere.

# Usage

Define the number of requests allowed and the total time.

```golang
var numberOfRequests = 2 //total number of requests allowed
var timeLimit = 10.0     //time limit in seconds
```

# Build and Run

Use the makefile to build the ratelimiter executable:

```bash
make
```

Run the app using the following:

```bash
./ratelimiter
```

# Build docker image

Use the makefile to build the docker image:

```bash
make image
```

Run the docker container:

```bash
make image-run
```
