package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rmasclef/fizz_buzz_api/pkg/fizz-buzz"
)

func TestFizzbuzzJSONTransformer_FromBytes(t *testing.T) {
	tests := []struct {
		name                 string
		requestBody          []byte
		want                 *fizzbuzz.Request
		expectedErrorMessage string
	}{
		{
			name:                 "test with good data",
			requestBody:          []byte(`{"str1":"fizz","str2":"buzz","int1":2,"int2":3,"limit":10}`),
			want:                 &fizzbuzz.Request{Str1: "fizz", Str2: "buzz", Int1: 2, Int2: 3, Limit: 10},
			expectedErrorMessage: "",
		},
		{
			name:                 "test with partial data",
			requestBody:          []byte(`{"int1":2,"int2":3,"limit":10}`),
			want:                 &fizzbuzz.Request{Int1: 2, Int2: 3, Limit: 10},
			expectedErrorMessage: "",
		},
		{
			name:                 "test with wrong data",
			requestBody:          []byte(`test:invalid_json`),
			want:                 nil,
			expectedErrorMessage: "invalid character 'e' in literal true (expecting 'r')",
		},
		{
			name:                 "test without body",
			requestBody:          nil,
			want:                 nil,
			expectedErrorMessage: "cannot JSON unmarshal an empty body",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := &fizzbuzz.JSONTransformer{}
			got, err := tf.FromBytes(tt.requestBody)

			if tt.expectedErrorMessage != "" {
				assert.EqualError(t, err, tt.expectedErrorMessage)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFizzbuzzJSONTransformer_ToBytes(t *testing.T) {
	tests := []struct {
		name          string
		fbr           *fizzbuzz.Response
		want          []byte
		expectedError error
	}{
		{
			name:          "test with good data",
			fbr:           &fizzbuzz.Response{"1", "fizz", "buzz", "fizz", "5", "fizzbuzz"},
			want:          []byte(`["1","fizz","buzz","fizz","5","fizzbuzz"]`),
			expectedError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tf := &fizzbuzz.JSONTransformer{}
			got, err := tf.ToBytes(tt.fbr)

			if tt.expectedError != nil {
				assert.Equal(t, tt.expectedError, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
