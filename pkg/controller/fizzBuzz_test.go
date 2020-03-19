package controller_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rmasclef/fizz_buzz_api/pkg/controller"
	"github.com/rmasclef/fizz_buzz_api/pkg/model"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		want *model.FizzBuzzRequest
		got  *model.FizzBuzzResponse
	}{
		{
			name: "test without data",
			want: &model.FizzBuzzRequest{},
			got: &model.FizzBuzzResponse{},
		},
		{
			name: "test with nominal data",
			want: &model.FizzBuzzRequest{
				Int1:                 3,
				Int2:                 5,
				Limit:                20,
				Str1:                 "fizz",
				Str2:                 "buzz",
			},
			got: &model.FizzBuzzResponse{
				Values: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"},
			},
		},
		{
			name: "test with limit 0",
			want: &model.FizzBuzzRequest{
				Int1:                 3,
				Int2:                 5,
				Limit:                0,
				Str1:                 "fizz",
				Str2:                 "buzz",
			},
			got: &model.FizzBuzzResponse{},
		},
		{
			name: "test with negative int1 int2",
			want: &model.FizzBuzzRequest{
				Int1:                 -3,
				Int2:                 -5,
				Limit:                20,
				Str1:                 "fizz",
				Str2:                 "buzz",
			},
			got: &model.FizzBuzzResponse{
				Values: []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz"},
			},
		},
		{
			name: "test with inverted int1 int2",
			want: &model.FizzBuzzRequest{
				Int1:                 5,
				Int2:                 3,
				Limit:                20,
				Str1:                 "fizz",
				Str2:                 "buzz",
			},
			got: &model.FizzBuzzResponse{
				Values: []string{"1", "2", "buzz", "4", "fizz", "buzz", "7", "8", "buzz", "fizz", "11", "buzz", "13", "14", "fizzbuzz", "16", "17", "buzz", "19", "fizz"},
			},
		},
		{
			name: "test with negative limit",
			want: &model.FizzBuzzRequest{
				Int1:                 3,
				Int2:                 5,
				Limit:                -20,
				Str1:                 "fizz",
				Str2:                 "buzz",
			},
			got: &model.FizzBuzzResponse{},
		},
		{
			name: "test with spaces",
			want: &model.FizzBuzzRequest{
				Int1:                 3,
				Int2:                 5,
				Limit:                20,
				Str1:                 "fizz ",
				Str2:                 " buzz",
			},
			got: &model.FizzBuzzResponse{
				Values: []string{"1", "2", "fizz ", "4", " buzz", "fizz ", "7", "8", "fizz ", " buzz", "11", "fizz ", "13", "14", "fizz  buzz", "16", "17", "fizz ", "19", " buzz"},
			},
		},
		{
			name: "test with special chars",
			want: &model.FizzBuzzRequest{
				Int1:                 3,
				Int2:                 5,
				Limit:                20,
				Str1:                 "fizz&é\"'è§",
				Str2:                 " #(~?-_`buzz",
			},
			got: &model.FizzBuzzResponse{
				Values: []string{"1", "2", "fizz&é\"'è§", "4", " #(~?-_`buzz", "fizz&é\"'è§", "7", "8", "fizz&é\"'è§", " #(~?-_`buzz", "11", "fizz&é\"'è§", "13", "14", "fizz&é\"'è§ #(~?-_`buzz", "16", "17", "fizz&é\"'è§", "19", " #(~?-_`buzz"},
			},
		},
		{
			name: "test with special chars",
			want: &model.FizzBuzzRequest{
				Int1:                 3,
				Int2:                 5,
				Limit:                20,
				Str1:                 "fizz&é\"'è§",
				Str2:                 " #(~?-_`buzz",
			},
			got: &model.FizzBuzzResponse{
				Values: []string{"1", "2", "fizz&é\"'è§", "4", " #(~?-_`buzz", "fizz&é\"'è§", "7", "8", "fizz&é\"'è§", " #(~?-_`buzz", "11", "fizz&é\"'è§", "13", "14", "fizz&é\"'è§ #(~?-_`buzz", "16", "17", "fizz&é\"'è§", "19", " #(~?-_`buzz"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// compare fizzbuzz responses
			assert.Equal(t, controller.FizzBuzz(tt.want), tt.got)
			// check result size only when limit is not < 0
			if tt.want.Limit >= 0 {
				assert.Len(t, controller.FizzBuzz(tt.want).Values, int(tt.want.Limit))
			}
		})
	}
}
