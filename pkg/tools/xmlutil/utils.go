package xmlutil

import (
	"bytes"
	"encoding/xml"

	"golang.org/x/net/html/charset"
)

func Unmarshal(data []byte, v interface{}) error {
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	return decoder.Decode(v)
}

func Marshal(v interface{}) ([]byte, error) {
	var b bytes.Buffer
	if err := xml.NewEncoder(&b).Encode(v); err != nil {
		return []byte{}, err
	}
	return b.Bytes(), nil
}
