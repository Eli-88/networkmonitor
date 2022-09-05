package parser

import "encoding/json"

// interface compliance
var _ Parser = &jsonParser{}

type jsonParser struct{}

func MakeJsonParser() Parser {
	return &jsonParser{}
}

func (jsonParser) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonParser) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
