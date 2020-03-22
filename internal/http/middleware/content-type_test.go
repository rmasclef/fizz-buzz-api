package middleware_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rmasclef/fizz_buzz_api/internal/http/middleware"
)

func TestContentTypeFilterer(t *testing.T) {
	tests := []struct {
		name                  string
		contentType           string
		handlerShouldBeCalled bool
	}{
		{
			name:                  "test with json",
			contentType:           "application/json",
			handlerShouldBeCalled: true,
		},
		{
			name:                  "test without value",
			contentType:           "",
			handlerShouldBeCalled: false,
		},
		{
			name:                  "test with invalid content-type",
			contentType:           "multipart/form-data",
			handlerShouldBeCalled: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "http://fake", nil)
			req.Header.Set("Content-Type", tt.contentType)

			recorder := httptest.NewRecorder()

			nextHandlerCalled := false
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				nextHandlerCalled = true
			})

			m := middleware.ContentTypeFilterer(handler)
			m.ServeHTTP(recorder, getRequest(tt.contentType))

			resp := recorder.Result()

			assert.Equal(t, tt.handlerShouldBeCalled, nextHandlerCalled)

			if !tt.handlerShouldBeCalled {
				body, _ := ioutil.ReadAll(resp.Body)

				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				assert.Equal(t, fmt.Sprintf("This API only allows 'application/json' requests (provided: %s).", tt.contentType), string(body))
				return
			}
			assert.Equal(t, http.StatusOK, resp.StatusCode)
		})
	}
}

func getRequest(ct string) *http.Request {
	req := httptest.NewRequest(http.MethodGet, "http://fake", nil)
	req.Header.Set("Content-Type", ct)
	return req
}
