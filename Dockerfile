FROM golang:latest

RUN mkdir -p /go/src/github.com/moemoe89/simple-go-clean-arch

WORKDIR /go/src/github.com/moemoe89/simple-go-clean-arch

COPY . /go/src/github.com/moemoe89/simple-go-clean-arch

RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go get gopkg.in/go-playground/validator.v10
RUN mkdir -p /go/src/github.com/moemoe89/simple-go-clean-arch/vendor/github.com/go-playground/validator/v10
RUN cp -rf /go/src/gopkg.in/go-playground/validator.v10/* /go/src/github.com/moemoe89/simple-go-clean-arch/vendor/github.com/go-playground/validator/v10
RUN go install

ENTRYPOINT /go/bin/goose -env=docker up && /go/bin/simple-go-clean-arch
