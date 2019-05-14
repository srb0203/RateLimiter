# Build a binary
FROM golang:alpine AS builder
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
RUN go build -o /go/bin/ratelimiter

#Build a smaller image from previous binary
FROM alpine
# Copy the executable.
COPY --from=builder /go/bin/ratelimiter /go/bin/ratelimiter 
# Run the binary.
ENTRYPOINT ["/go/bin/ratelimiter"]