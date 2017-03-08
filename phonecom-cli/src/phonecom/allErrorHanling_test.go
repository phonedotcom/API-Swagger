package main

import (
	"testing"
	"github.com/urfave/cli"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestNoArguments(t *testing.T) {

	err, response := createCliNoArguments()

	assert.NoError(t, err)

	datetime := response["datetime"].(string)
	assert.NotEmpty(t, datetime)
}

func createCliNoArguments() (error, map[string] interface{}) {

	app := cli.NewApp()

	app.Flags = []cli.Flag{cli.StringFlag{
		Name: commandLong,
		Value: defaultCommand,
	},}

	var response (map[string] interface{})
	var err error

	app.Action = func(c *cli.Context) error {
		err, response = execute(c, createCliConfig())
		return err
	}

	app.Run([]string{""})

	return err, response
}

func TestFlagNonExistingEndpoint(t *testing.T) {

	errorCommand := "ERRRRRRRRRRRRRorrrr"
	err, _ := createCli(errorCommand)

	assert.EqualError(t, err, fmt.Sprintf(invalidCommand, errorCommand, getAllCommands()))
}

func TestFlagEmptyCommand(t *testing.T) {

	errorCommand := ""
	err, _ := createCli(errorCommand)

	assert.EqualError(t, err, fmt.Sprintf(invalidCommand, errorCommand, getAllCommands()))
}

func TestFlagBadJsonIn(t *testing.T) {

	path := "../test/errorhandling/badJson.json"
	err, _ := createCliWithJsonIn(listAccounts, path)

	assert.EqualError(t, err, fmt.Sprintf(couldNotUnmarshal, path))
}

func TestFlagNonExistingFile(t *testing.T) {

	path := "non-existingfile"
	err, _ := createCliWithJsonIn(listAccounts, path)

	assert.EqualError(t, err, fmt.Sprintf(couldNotReadFile, path))
}
