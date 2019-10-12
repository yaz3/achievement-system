FROM golang:1.12

RUN mkdir -p /go/src/achievement-system
WORKDIR /go/src/achievement-system

ADD . /go/src/achievement-system
#RUN go mod init achievement-system
#RUN go mod vendor

#RUN GO111MODULE=on go mod download
RUN GO111MODULE=on go get ./...

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build

