FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/go-clean-arch-ratu

WORKDIR /go/src/github.com/moemoe89/go-clean-arch-ratu

COPY . /go/src/github.com/moemoe89/go-clean-arch-ratu

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go mod download
RUN go install

ENTRYPOINT /go/bin/goose -env=docker up && /go/bin/go-clean-arch-ratu
