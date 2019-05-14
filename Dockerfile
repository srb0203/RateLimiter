############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go build -o /go/bin/ratelimiter
############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
COPY --from=builder /go/bin/ratelimiter /go/bin/ratelimiter 

# Run the hello binary.
ENTRYPOINT ["/go/bin/ratelimiter"]