package main

import (
	"strconv"
	"fmt"
)

func validateResponse(json map[string] interface{}) string {

	jsonError := json[responseError]

	if (jsonError != nil) {

		jsonError := jsonError.(map[string] interface{})

		return fmt.Sprintf(
			msgHttpError,
			strconv.FormatFloat(jsonError[responseHttpStatusCode].(float64), 'f', -1, 64), jsonError[responseMessage].(string))
	}

	return ""
}
