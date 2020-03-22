package fizzbuzz

type (
	Request struct {
		Str1  string `json:"str1"`
		Str2  string `json:"str2"`
		Int1  uint   `json:"int1"`
		Int2  uint   `json:"int2"`
		Limit uint   `json:"limit"`
	}
	Response []string
)
