# This Dockerfile builds an image for a client_golang example.
#
# Use as (from the root for the client_golang repository):
#    docker build -f examples/$name/Dockerfile -t prometheus/golang-example-$name .

# Builder image, where we build the example.
FROM golang:1.9.0 AS builder
WORKDIR /go/src/github.com/ben-st/go-mux-example
COPY . .
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux go build -o mux

# Final image.
FROM scratch
LABEL maintainer "Benjamin Stein <info@diffus.org>"
COPY --from=builder /go/src/github.com/ben-st/go-mux-example/mux .
EXPOSE 8080
ENTRYPOINT ["/mux"]