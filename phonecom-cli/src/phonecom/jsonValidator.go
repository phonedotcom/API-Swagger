package main

import (
	"encoding/json"
)

func validateJson(jsonString string) map[string] interface{} {

	var js map[string] interface{}
	json.Unmarshal([]byte(jsonString), &js)

	if (js != nil) {
		return js
	}

	return nil
}
