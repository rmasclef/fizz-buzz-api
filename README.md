# fizz-buzz-api
implementation of a fizz-buzz REST ~~ful~~ API

## simple version
https://github.com/rmasclef/fizz-buzz-api/pull/1
https://github.com/rmasclef/fizz-buzz-api/tree/simple-version

It can be even smaller if we don't want content-type filtering, graceful shutdow and metrics handler
Uses `echo` framework to do things fast (request binding, validation, logs, error recovering ...)

## structured version

https://github.com/rmasclef/fizz-buzz-api/pull/2
https://github.com/rmasclef/fizz-buzz-api/tree/pure-go-version

Structured (over-ingineered) version of fizz-buzz API
Less understandable but more `prod-ready` and extensible ...

## bonus - protobuf version

https://github.com/rmasclef/fizz-buzz-api/tree/protobuf-version

Request binding and validation is done by protobuf model
Theorically faster than a JSON API (useless for our case)
