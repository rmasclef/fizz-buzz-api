package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gol4ng/logger"

	"github.com/rmasclef/fizz_buzz_api/pkg/fizz-buzz"
)

// handle a fizzbuzz request
func FizzBuzz(t fizzbuzz.Transformer, v fizzbuzz.Validator, c fizzbuzz.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.FromContext(r.Context(), logger.NewNopLogger())

		// get request body
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			l.Error("read request body error", logger.Error("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			// @TODO add some more information
			_, _ = w.Write([]byte("read request body error"))
			return
		}

		// bind HTTP Request (form-data or JSON or whatever ...) to a fizzBuzzRequest
		fbr, err := t.FromBytes(reqBody)
		if err != nil {
			l.Error("transform request error", logger.Error("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			// @TODO add some more information
			_, _ = w.Write([]byte("transform error"))
			return
		}

		// validate fizzbuzz request
		if err = v(fbr); err != nil {
			l.Error("validation error", logger.Error("error", err))
			w.WriteHeader(http.StatusBadRequest)
			// @TODO add some more information
			_, _ = w.Write([]byte(fmt.Sprintf("validation error : %s", err)))
			return
		}

		// call fizzbuzz controller and transform the fizzbuzzResponse into []byte
		resp, err := t.ToBytes(c(fbr))
		if err != nil {
			l.Error("transform response error", logger.Error("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			// @TODO add some more information
			_, _ = w.Write([]byte("response transform error"))
			return
		}

		_, _ = w.Write(resp)
	}
}
