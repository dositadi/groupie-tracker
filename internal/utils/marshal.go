package utils

import "encoding/json"

func MarshalObject(object any) []byte {
	out, err := json.Marshal(object)
	if err == nil {
		return out
	}
	return nil
}
