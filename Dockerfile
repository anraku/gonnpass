FROM golang:1.13.0-alpine3.10 as builder
RUN apk --no-cache add git
WORKDIR /go/src/github.com/anraku/gonnpass
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY go.mod go.sum ./
RUN GO111MODULE=on go mod download
COPY . ./
RUN go build -o /app

FROM alpine:3.10 as executor
COPY --from=builder /app  /app 
CMD ["/app "]