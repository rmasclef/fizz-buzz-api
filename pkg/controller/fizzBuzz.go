package controller

import (
	"strconv"

	"github.com/rmasclef/fizz_buzz_api/pkg/model"
)

func FizzBuzz(fizzBuzzReq *model.FizzBuzzRequest) *model.FizzBuzzResponse {
	results := &model.FizzBuzzResponse{}

	for i := 1; i <= int(fizzBuzzReq.Limit); i++ {
		res := ""
		if i%int(fizzBuzzReq.Int1) == 0 {
			res += fizzBuzzReq.Str1
		}
		if i%int(fizzBuzzReq.Int2) == 0 {
			res += fizzBuzzReq.Str2
		}
		if res != "" {
			results.Values = append(results.Values, res)
			continue
		}
		results.Values = append(results.Values, strconv.Itoa(i))
	}
	return results
}
