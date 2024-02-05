FROM golang:alpine3.19

RUN mkdir $GOPATH/src/service

ADD ./. $GOPATH/src/service

WORKDIR $GOPATH/src/service

RUN go mod tidy
RUN go build -o service ./cmd/service/main.go

EXPOSE 8001
CMD ./service