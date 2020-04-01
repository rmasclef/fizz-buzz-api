FROM golang:1.14-alpine3.11 AS builder
WORKDIR /project
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
COPY . .
# @TODO add AppVersion flag
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o fizz-buzz-api ./cmd/...

FROM scratch
COPY --from=builder /project/fizz-buzz-api .
ENTRYPOINT ["/fizz-buzz-api"]
