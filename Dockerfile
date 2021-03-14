FROM golang:1.16

WORKDIR /go/src/github.com/jeeeevs/rest-api-go
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod vendor
COPY . .
RUN go install
CMD ["rest-api-go"]