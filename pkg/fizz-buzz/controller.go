package fizzbuzz

import "strconv"

type Controller func(*Request) *Response

func FizzBuzzController(fizzBuzzReq *Request) *Response {
	fizzBuzzResp := make(Response, fizzBuzzReq.Limit)

	for i := uint(1); i <= fizzBuzzReq.Limit; i++ {
		res := ""
		if i%fizzBuzzReq.Int1 == 0 {
			res += fizzBuzzReq.Str1
		}
		if i%fizzBuzzReq.Int2 == 0 {
			res += fizzBuzzReq.Str2
		}
		if res != "" {
			fizzBuzzResp[i-1] = res
			continue
		}
		fizzBuzzResp[i-1] = strconv.FormatUint(uint64(i), 10)
	}
	return &fizzBuzzResp
}
