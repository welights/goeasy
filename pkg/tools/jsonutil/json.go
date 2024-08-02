// Package jsonutil uses json-iterator to get better performance than standard encoding/json package.
// And it's 100% compatible with encoding/json.
//
// 字符串在编码为JSON字符串时会被强制转换为有效的UTF-8，为了防止一些浏览器在JSON输出误解以为是HTML，
// “<”，“>”，“&”这类字符会被进行转义，如果不想被转义，就使用Encoder，并且SetEscapeHTML(false)即可。
package jsonutil

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

var defaultStrategy = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(v interface{}) ([]byte, error) {
	return defaultStrategy.Marshal(v)
}

func MarshalIndent(v interface{}) ([]byte, error) {
	return defaultStrategy.MarshalIndent(v, "", "  ")
}

func Unmarshal(data []byte, v interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}

func MarshalToString(v interface{}) (string, error) {
	return defaultStrategy.MarshalToString(v)
}

func UnmarshalFromString(data string, v interface{}) error {
	return defaultStrategy.UnmarshalFromString(data, v)
}

func NewEncoder(w io.Writer) *jsoniter.Encoder {
	return jsoniter.NewEncoder(w)
}

func NewDecoder(r io.Reader) *jsoniter.Decoder {
	return jsoniter.NewDecoder(r)
}
