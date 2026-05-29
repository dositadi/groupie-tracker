package helper

import "encoding/json"

/* type TransformTypes interface {
	any | map[string]any | map[int]any | map[bool]any
} */

func Marshal[T any](input T) []byte {
	json, err := json.Marshal(input)
	if err != nil {
		panic("Could not marshal JSON data.")
	}
	return json
}
