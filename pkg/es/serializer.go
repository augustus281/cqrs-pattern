package es

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

var jsonIter = jsoniter.ConfigCompatibleWithStandardLibrary

type Serializer interface {
	SerializeEvent(aggregate Aggregate, event Event) (Event, error)
	DeserializeEvent(event Event) (any, error)
}

func Marshal(v any) ([]byte, error) {
	return jsonIter.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return jsonIter.Unmarshal(data, v)
}

func NewDecoder(r io.Reader) *jsoniter.Decoder {
	return jsonIter.NewDecoder(r)
}

func NewEncoder(w io.Writer) *jsoniter.Encoder {
	return jsonIter.NewEncoder(w)
}
