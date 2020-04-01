package fizzbuzz

import "errors"

type Validator func(*Request) error

func RequestValidator(r *Request) error {
	if r.Limit > 1024 {
		return errors.New("limit can not exceed 1024 values")
	}
	if r.Int1 == 0 || r.Int2 == 0 {
		return errors.New("int1 and int2 must be greater than 0")
	}
	return nil
}
