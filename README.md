# letsgo-mux
[![Build Status](https://travis-ci.org/letsgo-framework/letsgo-mux.svg?branch=master)](https://travis-ci.org/letsgo-framework/letsgo-mux)
[![Go Report Card](https://goreportcard.com/badge/github.com/letsgo-framework/letsgo-mux)](https://goreportcard.com/report/github.com/letsgo-framework/letsgo-mux)
[![Coverage Status](https://coveralls.io/repos/github/letsgo-framework/letsgo-mux/badge.svg?branch=master)](https://coveralls.io/github/letsgo-framework/letsgo-mux?branch=master)
[![Sourcegraph](https://sourcegraph.com/github.com/letsgo-framework/letsgo-mux/-/badge.svg)](https://sourcegraph.com/github.com/letsgo-framework/letsgo-mux?badge)
[![Join the chat at https://gitter.im/letsgo-framework/community](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/letsgo-framework/community)

## Go api starter
### Ingredients

- Go
- [mux] ( https://github.com/gorilla/mux )
- [mongodb] ( https://www.mongodb.com/ )
- [mongo-go-driver] ( https://github.com/mongodb/mongo-go-driver )
- [oauth2] ( https://github.com/golang/oauth2 )
- [check] ( https://godoc.org/gopkg.in/check.v1 )
- [godotenv] ( https://github.com/joho/godotenv )
- [cors] ( github.com/gin-contrib/cors )
***
### Directory Structure

By default, your project's structure will look like this:

- `/controllers`: contains the core code of your application.
- `/database`: contains mongo-go-driver connector.
- `/helpers`: contains helpers functions of your application.
- `/middlewares`: contains middlewares of your application.
- `/routes`: directory contains RESTful api routes of your application.
- `/tests`: contains tests of your application.
- `/types`: contains the types/structures of your application.
***
### Environment Configuration

letsGo uses `godotenv` for setting environment variables. The root directory of your application will contain a `.env.example` file.
copy and rename it to `.env` to set your environment variables.

You need to create a `.env.testing` file from `.env.example` for running tests.
***
### Setting up

- clone letsGo
- change package name in `go.mod` to your package name
- change the internal package (controllers, tests, helpers etc.) paths as per your requirement
- setup `.env` and `.env.testing`
- run `go mod download` to install dependencies

#### OR `letsgo-cli` can be used to setup new project

### install letsgo-cli
```
go get github.com/letsgo-framework/letsgo-cli
```


### Create a new project

```bash
letsgo-cli init -importPath=<import_namespace> -directory=<project_name> -router=<router>
```

- **letsgo-cli init -importPath=github.com -directory=myapp -router=gin**<br/>
  Generates a new project called **myapp** in your `GOPATH` inside `github.com` and installs the default plugins through the glide.
***
### Run : ```go run main.go```
***
### Build : ```go build```
***
### Test : ```go test tests/main_test.go```

### Coverall :
```
go test -v -coverpkg=./... -coverprofile=coverage.out ./...

goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN
```
***
### Authentication

letsgo uses Go OAuth2 (https://godoc.org/golang.org/x/oauth2) for authentication.
***

### Deploy into Docker

```
sudo docker run --rm -v "$PWD":/go/src/github.com/letsgo-framework/letsgo -w /go/src/github.com/letsgo-framework/letsgo iron/go:dev go build -o letsgo
```
```
sudo docker build -t sab94/letsgo .
```
```
sudo docker run --rm -p 8080:8080 sab94/letsgo
```
