package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rmasclef/fizz_buzz_api/pkg/fizz-buzz"
)

func TestFizzBuzzController(t *testing.T) {
	tests := []struct {
		name             string
		request          *fizzbuzz.Request
		expectedResponse *fizzbuzz.Response
	}{
		{
			name:             "test without data",
			request:          &fizzbuzz.Request{},
			expectedResponse: &fizzbuzz.Response{},
		},
		{
			name: "test with nominal data",
			request: &fizzbuzz.Request{
				Int1:  3,
				Int2:  5,
				Limit: 20,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedResponse: &fizzbuzz.Response{
				"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz",
			},
		},
		{
			name: "test with limit 0",
			request: &fizzbuzz.Request{
				Int1:  3,
				Int2:  5,
				Limit: 0,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedResponse: &fizzbuzz.Response{},
		},
		{
			name: "test with inverted int1 int2",
			request: &fizzbuzz.Request{
				Int1:  5,
				Int2:  3,
				Limit: 20,
				Str1:  "fizz",
				Str2:  "buzz",
			},
			expectedResponse: &fizzbuzz.Response{
				"1", "2", "buzz", "4", "fizz", "buzz", "7", "8", "buzz", "fizz", "11", "buzz", "13", "14", "fizzbuzz", "16", "17", "buzz", "19", "fizz",
			},
		},
		{
			name: "test with spaces",
			request: &fizzbuzz.Request{
				Int1:  3,
				Int2:  5,
				Limit: 20,
				Str1:  "fizz ",
				Str2:  " buzz",
			},
			expectedResponse: &fizzbuzz.Response{
				"1", "2", "fizz ", "4", " buzz", "fizz ", "7", "8", "fizz ", " buzz", "11", "fizz ", "13", "14", "fizz  buzz", "16", "17", "fizz ", "19", " buzz",
			},
		},
		{
			name: "test with special chars",
			request: &fizzbuzz.Request{
				Int1:  3,
				Int2:  5,
				Limit: 20,
				Str1:  "fizz&é\"'è§",
				Str2:  " #(~?-_`buzz",
			},
			expectedResponse: &fizzbuzz.Response{
				"1", "2", "fizz&é\"'è§", "4", " #(~?-_`buzz", "fizz&é\"'è§", "7", "8", "fizz&é\"'è§", " #(~?-_`buzz", "11", "fizz&é\"'è§", "13", "14", "fizz&é\"'è§ #(~?-_`buzz", "16", "17", "fizz&é\"'è§", "19", " #(~?-_`buzz",
			},
		},
		{
			name: "test with special chars",
			request: &fizzbuzz.Request{
				Int1:  3,
				Int2:  5,
				Limit: 20,
				Str1:  "fizz&é\"'è§",
				Str2:  " #(~?-_`buzz",
			},
			expectedResponse: &fizzbuzz.Response{
				"1", "2", "fizz&é\"'è§", "4", " #(~?-_`buzz", "fizz&é\"'è§", "7", "8", "fizz&é\"'è§", " #(~?-_`buzz", "11", "fizz&é\"'è§", "13", "14", "fizz&é\"'è§ #(~?-_`buzz", "16", "17", "fizz&é\"'è§", "19", " #(~?-_`buzz",
			},
		},
		{
			name: "test with empty str2",
			request: &fizzbuzz.Request{
				Int1:  3,
				Int2:  5,
				Limit: 20,
				Str1:  "fizz",
				Str2:  "",
			},
			expectedResponse: &fizzbuzz.Response{
				"1", "2", "fizz", "4", "5", "fizz", "7", "8", "fizz", "10", "11", "fizz", "13", "14", "fizz", "16", "17", "fizz", "19", "20",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResponse := fizzbuzz.FizzBuzzController(tt.request)
			// compare fizzbuzz responses
			assert.Equal(t, tt.expectedResponse, actualResponse)
			// check fizzbuzz response size
			assert.Len(t, *actualResponse, int(tt.request.Limit))
		})
	}
}
