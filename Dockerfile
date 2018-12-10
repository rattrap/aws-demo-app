FROM golang:1.11.2 AS builder
RUN go version

COPY . /go/src/github.com/rattra/aws-demo-app/
WORKDIR /go/src/github.com/rattra/aws-demo-app/
RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/rattra/aws-demo-app/app .

EXPOSE 8123
ENTRYPOINT ["./app"]