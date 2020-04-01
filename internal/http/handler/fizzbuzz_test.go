package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rmasclef/fizz_buzz_api/internal/http/handler"
	"github.com/rmasclef/fizz_buzz_api/pkg/fizz-buzz"
)

// =====================================================================================================================
// ============================================= INTEGRATION tests =====================================================
// =====================================================================================================================

func TestMainWithJSONRequest(t *testing.T) {
	jsonReq, _ := json.Marshal(&fizzbuzz.Request{Str1: "fizz", Str2: "buzz", Int1: 2, Int2: 3, Limit: 10})
	request := httptest.NewRequest(http.MethodPost, "/fizz-buzz", bytes.NewReader(jsonReq))
	response := httptest.NewRecorder()

	h := handler.FizzBuzz(&fizzbuzz.JSONTransformer{}, fizzbuzz.RequestValidator, fizzbuzz.FizzBuzzController)
	h.ServeHTTP(response, request)

	check := assert.New(t)
	check.Equal(http.StatusOK, response.Code, response.Body.String())
	check.JSONEq(`["1","fizz","buzz","fizz","5","fizzbuzz","7","fizz","buzz","fizz"]`, response.Body.String())
}

func TestMainWithJSONRequestAndMissingParams(t *testing.T) {
	jsonReq, _ := json.Marshal(&fizzbuzz.Request{Str1: "fizz", Int1: 2, Int2: 3, Limit: 10})
	request := httptest.NewRequest(http.MethodPost, "/fizz-buzz", bytes.NewReader(jsonReq))
	response := httptest.NewRecorder()

	h := handler.FizzBuzz(&fizzbuzz.JSONTransformer{}, fizzbuzz.RequestValidator, fizzbuzz.FizzBuzzController)
	h.ServeHTTP(response, request)

	check := assert.New(t)
	check.Equal(http.StatusOK, response.Code, response.Body.String())
	check.JSONEq(`["1","fizz","3","fizz","5","fizz","7","fizz","9","fizz"]`, response.Body.String())
}
