package utils

import "encoding/json"

func ToJson(v any) ([]byte, error) {
	return json.Marshal(v)
}

func FromJson(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
