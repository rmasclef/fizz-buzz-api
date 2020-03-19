package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"

	"github.com/rmasclef/fizz_buzz_api/pkg/controller"
	"github.com/rmasclef/fizz_buzz_api/pkg/model"
)

func FizzBuzzHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Printf("Content Length Received : %v\n", req.ContentLength)

	// create domain request from HTTP request
	request := &model.FizzBuzzRequest{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}
	err = proto.Unmarshal(data, request)
	// err = jsonpb.Unmarshal(data, request)
	if err != nil {
		log.Fatalf("Unable to transform HTTP request : %v", err)
	}

	// call domain controller
	results := controller.FizzBuzz(request)

	// write domain response into HTTP response
	response, err := proto.Marshal(results)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	_, err = resp.Write(response)
	if err != nil {
		log.Fatalf("Unable to write data into HTTP response : %v", err)
	}
}
