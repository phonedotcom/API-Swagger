package main

import (
	"fmt"
  "encoding/json"
)

func validateResponse(jsonObject map[string] interface{}) string {

	jsonError := jsonObject[responseError]

	if (jsonError != nil) {

		jsonError := jsonError.(map[string] interface{})

		return fmt.Sprintf(
			msgHttpError,
      json.Number.String(jsonError[responseHttpStatusCode].(json.Number)), jsonError[responseMessage].(string))
	}

	return ""
}
