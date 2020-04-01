package fizzbuzz

import (
	"encoding/json"
	"errors"
)

type Transformer interface {
	FromBytes(requestBody []byte) (*Request, error)
	ToBytes(fbr *Response) ([]byte, error)
}

// =====================================================================================================================
// =================================================== JSON ============================================================
// =====================================================================================================================
type JSONTransformer struct{}

// transform an HTTP Request body into a fizzbuzz.Request
func (t *JSONTransformer) FromBytes(requestBody []byte) (*Request, error) {
	if requestBody == nil {
		return nil, errors.New("cannot JSON unmarshal an empty body")
	}
	// deserialize the request body into a domain fizzbuzz.Request
	fbr := new(Request)
	err := json.Unmarshal(requestBody, fbr)
	if err != nil {
		return nil, err
	}
	return fbr, nil
}

func (t *JSONTransformer) ToBytes(fbr *Response) ([]byte, error) {
	return json.Marshal(fbr)
}

// HERE YOU MAY IMPLEMENT OTHER FORMAT (XML, PROTOBUF ...)
