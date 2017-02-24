package main

import (
  "github.com/urfave/cli"
  "testing"
	"strconv"
)

var commandFlag = "-command"
var errorNotNullMessage = "expected no error from Run, got %s"

func createCliWithJsonIn(endpoint string, path string) (error, map[string] interface{}) {

  app := cli.NewApp()

  defaultCommand = endpoint
  defaultInput = path

  app.Flags = getCliFlags()
  configPath = "../../config.xml"
  var json (map[string] interface{})
  var err error

  app.Action = func(c *cli.Context) error {
    err, json = execute(c)
    return err
  }

  app.Run([]string{commandFlag, endpoint})

  return err, json
}

func createCli(endpoint string) (error, map[string] interface{}) {

  app := cli.NewApp()

  defaultCommand = endpoint

  app.Flags = getCliFlags()
  configPath = "../../config.xml"
  var json (map[string] interface{})
  var err error

  app.Action = func(c *cli.Context) error {
    err, json = execute(c)
    return err
  }

  app.Run([]string{commandFlag, endpoint})

  return err, json
}

func createGetCliStringId(endpoint string, id string) (error, map[string] interface{}) {

	if (id == "") {
		return nil, nil
	}

	var json (map[string] interface{})
	var err error

	app := cli.NewApp()

	defaultCommand = endpoint
	defaultId = id

	app.Flags = getCliFlags()
	configPath = "../../config.xml"

	app.Action = func(c *cli.Context) error {
		err, json = execute(c)
		return err
	}

	app.Run([]string{commandFlag, endpoint})

	return err, json
}

func createGetOrRemoveCli(endpoint string, id int) (error, map[string] interface{}) {

  if (id <= 0) {
    return nil, nil
  }

  var json (map[string] interface{})
  var err error

  app := cli.NewApp()

  defaultCommand = endpoint
  defaultId = strconv.Itoa(id)

  app.Flags = getCliFlags()
  configPath = "../../config.xml"

  app.Action = func(c *cli.Context) error {
    err, json = execute(c)
    return err
  }

  app.Run([]string{commandFlag, endpoint})

  return err, json
}

func assertErrorNotNull(t *testing.T, err error) {

  if err != nil {
    t.Fatalf(errorNotNullMessage, err)
  }
}

func getFirstIdString(json map[string] interface{}) string {

	items := json["items"].([]interface{})
	if (len(items) > 0) {
		firstItem := items[0].(map[string] interface{})
		firstId := firstItem["id"].(string)
		return firstId
	}

	return ""
}

func getFirstId(json map[string] interface{}) int {

  items := json["items"].([]interface{})
  if (len(items) > 0) {
    firstItem := items[0].(map[string] interface{})
    firstId := firstItem["id"].(float64)
    return int(firstId)
  }

  return 0
}

func getId(json map[string] interface{}) int {

  id := json["id"].(float64)
  return int(id)
}