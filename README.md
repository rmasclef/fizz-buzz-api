# fizz-buzz-api

straight-forward implementation of a fizz-buzz REST API

It uses echo web framework in order not having to rewrite request binding, validation or HTTP middlewares.

It only allows user to make JSON POST requests

## usage

### prod

`docker run --rm --expose=8080 -p 8080:8080 rmasclef/fizz-buzz-api-go:v0.1.0`

### dev
`make run HTTP_PORT=8080`
or
`go run main.go HTTP_PORT=8080`

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

on this version, we made the choice not to expose any domain metrics.
you can use a sideCar in order to have RED metrics such as "request per second", "error rate" ...

logs are directly sent to stdout, we suggest you use a sidecar (fluentd, filebeats ...) in order to aggregate them in a log service (graylog, logstash, ...)
