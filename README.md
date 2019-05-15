# RateLimiter

HTTP API Rate limiter in Go.

# Description

Ratelimiter limits the number of calls to an HTTP API. It uses a credit based system. Each IP address is assigned a specified number of credits which corresponds to the number of requests it can make in a given amount of time. The credit reduces by one on each request. The credits are replenished after the given amount of time has elapased. For example if the app allows for 10 requests per hour then the credits will be replenished after an hour.

# Build and Run

Use the makefile to build the ratelimiter executable:

```bash
make
```

Run the app using the following:

```bash
./ratelimiter
```
