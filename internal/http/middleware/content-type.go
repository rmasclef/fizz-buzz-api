package middleware

import (
	"fmt"
	"net/http"
)

func ContentTypeFilterer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		ct := req.Header.Get("Content-Type")
		if !isJSONRequest(ct) {
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write([]byte(fmt.Sprintf("This API only allows 'application/json' requests (provided: %s).", ct)))
			return
		}
		next.ServeHTTP(writer, req)
	})
}

func isJSONRequest(ct string) bool {
	switch ct {
	case "application/json", "application/json; charset=UTF-8":
		return true
	default:
		return false
	}
}
