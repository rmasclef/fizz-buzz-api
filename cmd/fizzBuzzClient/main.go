package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/rmasclef/fizz_buzz_api/pkg/model"
)

func main() {
	req := &model.FizzBuzzRequest{
		Int1:                 3,
		Int2:                 5,
		Limit:                20,
		Str1:                 "fizz",
		Str2:                 "buzz",
	}
	b, err := proto.Marshal(req)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	httpResponse, err := http.Post("http://localhost:8080/fizz-buzz", "application/x-binary", bytes.NewReader(b))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatalf("Unable to read bytes from responses : %v", err)
	}

	resp := &model.FizzBuzzResponse{}
	_ = proto.Unmarshal(respBytes, resp)

	fmt.Printf("Response from API is : %v\n", resp.GetValues())
}
