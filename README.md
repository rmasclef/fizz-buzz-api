# fizz-buzz-api

structured implementation of a fizz-buzz REST API

> This implementation ~~might be~~ **is** overkill for the needs

Features :
- Configuration
- GELF formatted logs
- Decoupled code (well ... the logger is coupled for now)
- Unit & Integration tests
- Domain Validation without reflection
- Request Content-type filtering
- Tracing support (correlation-ID)
- Graceful shutdown
- DI without reflection (required not to have a too big main)
- Docker support
- Profiling (pprof)

It only allows user to make JSON POST requests

## usage

### PROD (from docker image)

Docker image not pushed yet

### DEV (from code)
`make run HTTP_PORT=8080`
or
`go run ./cmd/main.go -HTTP_ADDR=:8080`

## example

## CURL 
```
curl --location --request POST 'localhost:8080/fizz-buzz' \
--header 'Content-Type: application/json' \
--data-raw '{
	"int1": 3,
	"int2": 4,
	"limit": 20,
	"str1": "fizz",
	"str2": "buzz"
}
'
```

## Go
```
package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "localhost:8080/fizz-buzz"
  method := "POST"
  payload := strings.NewReader(`{"int1": 3, "int2": 4, "limit": 20, "str1": "fizz", "str2": "buzz"}`)

  req, err := http.NewRequest(method, url, payload)
  if err != nil {
    fmt.Println(err)
  }
  req.Header.Add("Content-Type", "application/json")

  client := &http.Client {}
  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)

  fmt.Println(string(body))
}
```

the two above code snippets will return the following response body:

`["1","2","fizz","buzz","5","fizz","7","buzz","fizz","10","11","fizzbuzz","13","14","fizz","buzz","17","fizz","19","buzz"]`

## metrics and logs

On this version, we made the choice not to expose any metrics.
You can use a sideCar in order to have RED metrics such as "request per second", "error rate" ...
And you can also use a sideCar to determine which request has been made most of the time

Logs are directly sent to stdout in GELF format, we suggest you use a sidecar (fluentd, filebeats ...) in order to aggregate them in a log service (graylog, logstash, ...)

Don't hesitate to take a look at https://12factor.net/ :D

# Code structure
## cmd
This is the starting point of the Application, you will find the `main.go` file.

This folder contains the different available binaries/entry-points.

At this time we only have one.
If it were to change, we will create a subfolder for each binary (i.e the API, a consumer, an indexer ...).

## config

This folder contains all the configuration structures that will be use by our APP.

At this time we only have one
If it were to change, we will create a file for each binary (i.e the ServerConfig, the StorageConfig, the consumerConfig ...).

## internal

This folder contains all the non-domain code (di, HTTP implementation of our API ...)

If you want to add some storage implementation (ES, Redis, S3 etc ...) this goes here

> All packages inside this folder are not exported as go-modules (i.e you can not import this code outside of the project)

## pkg

This folder contains all the `domain` code.

All code that goes here is totally decoupled from implementations (i.e you should never see any ES/Mongo/HTTP import here)

You define the business/domain models here.

> All packages defined in this folder are accessible from the outside of the project
> 
> (i.e you can import them in another project using go-modules)
>
> If you want to make a gRPC version of the API, you can reuse the domain and "only implement" the gRPC part.

# TODO

The code is still not really decoupled but I tried to make it as small and concise as possible for this kind of architecture.

CI is not implemented yet.

Error messages are not complete
