[![Build Status](https://travis-ci.org/moemoe89/simple-go-clean-arch.svg?branch=master)](https://travis-ci.org/moemoe89/simple-go-clean-arch)
[![codecov](https://codecov.io/gh/moemoe89/simple-go-clean-arch/branch/master/graph/badge.svg)](https://codecov.io/gh/moemoe89/simple-go-clean-arch)
[![Go Report Card](https://goreportcard.com/badge/github.com/moemoe89/simple-go-clean-arch)](https://goreportcard.com/report/github.com/moemoe89/simple-go-clean-arch)

# SIMPLE-GO-CLEAN-ARCH #

Simple Go Clean Arch Using Golang (Gin Gonic Framework) as Programming Language, MariaDB as Database

## Directory structure
Your project directory structure should look like this
```
  + your_gopath/
  |
  +--+ src/github.com/moemoe89
  |  |
  |  +--+ simple-go-clean-arch/
  |     |
  |     +--+ main.go
  |        + api/
  |        + routers/
  |        + ... any other source code
  |
  +--+ bin/
  |  |
  |  +-- ... executable file
  |
  +--+ pkg/
     |
     +-- ... all dependency_library required

```

## Setup and Build

* Setup Golang <https://golang.org/>
* Setup MariaDB <https://www.mariadb.org/>
* Under `$GOPATH`, do the following command :
```
$ mkdir -p src/github.com/moemoe89
$ cd src/github.com/moemoe89
$ git clone <url>
$ mv <cloned directory> simple-go-clean-arch
```

## Running Migration
* Copy `config-sample.json` to `config.json` and changes the value based on your configurations.
* Create MySQL database for example named `simple_api` and do migration with `Goose` <https://bitbucket.org/liamstask/goose/>
* Change database configuration on dbconf.yml like `dialect` and `dsn` for each environtment
* Do the following command
```
$ cd $GOPATH/src/github.com/moemoe89/simple-go-clean-arch
$ goose -env=development up
```

## Documetation with Swagger
For open swagger access via browser
```
{{url}}/swagger/index.html
```
For updating swagger
```
$ swag init
```

## Running Application with Makefile
Make config file for local :
```
$ cp config-sample.json config-local.json
```
Build
```
$ cd $GOPATH/src/github.com/moemoe89/simple-go-clean-arch
$ make build
```
Run
```
$ cd $GOPATH/src/github.com/moemoe89/simple-go-clean-arch
$ make run
```
Stop
```
$ cd $GOPATH/src/github.com/moemoe89/simple-go-clean-arch
$ make stop
```
Make config file for docker :
```
$ cp config-sample.json config-docker.json
```
Docker Build
```
$ cd $GOPATH/src/github.com/moemoe89/simple-go-clean-arch
$ make docker-build
```
Docker Up
```
$ cd $GOPATH/src/github.com/moemoe89/simple-go-clean-arch
$ make docker-up
```
Docker Down
```
$ cd $GOPATH/src/github.com/moemoe89/simple-go-clean-arch
$ make docker-down
```

## How to Run with Docker
Make config file for docker :
```
$ cp config-sample.json config.json
```
Build
```
$ docker-compose build
```
Run
```
$ docker-compose up
```
Stop
```
$ docker-compose down
```

## How to Run Unit Test
Run
```
$ go test ./...
```
Run with cover
```
$ go test ./... -cover
```
Run with HTML output
```
$ go test ./... -coverprofile=c.out && go tool cover -html=c.out
```
