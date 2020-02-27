FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/simple-go-clean-arch

WORKDIR /go/src/github.com/moemoe89/simple-go-clean-arch

COPY . /go/src/github.com/moemoe89/simple-go-clean-arch

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/goose -env=docker up && /go/bin/simple-go-clean-arch
