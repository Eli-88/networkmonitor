package parser

type Encoder interface {
	Marshal(v any) ([]byte, error)
}

type Decoder interface {
	Unmarshal(data []byte, v any) error
}

type Parser interface {
	Encoder
	Decoder
}
