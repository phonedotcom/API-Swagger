package main

import (
	"encoding/json"
  "strings"
)

func validateJson(jsonString string) map[string] interface{} {

	var js map[string] interface{}
  d := json.NewDecoder(strings.NewReader(jsonString))
  d.UseNumber()
	d.Decode(&js)

	if (js != nil) {
		return js
	}

	return nil
}
