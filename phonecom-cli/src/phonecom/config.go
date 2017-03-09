package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"phonecom-go-sdk"
)

type Query struct {
	ConfigList []CliConfig `xml:"Config"`
}

type CliConfig struct {
	BaseApiPath  string
	ApiKeyPrefix string
	ApiKey       string
	Type         string
	AccountId    int32
	Path         string
}

func (c *CliConfig) getConfig() CliConfig {

	configPath := c.getConfigPath()
	xmlFile, err := os.Open(configPath)

	var noConfig CliConfig

	if err != nil {
		fmt.Println("Could not read config.xml", err)
		return noConfig
	}

	defer xmlFile.Close()

	content, _ := ioutil.ReadAll(xmlFile)

	var q Query
	xml.Unmarshal(content, &q)

	for _, config := range q.ConfigList {
		if config.Type == "main" {
			return config
		}
	}

	return noConfig
}

func (c *CliConfig) createSwaggerConfig(xmlConfig CliConfig) *swagger.Configuration {

	var baseApiPath string = xmlConfig.BaseApiPath
	var apiKeyPrefix string = xmlConfig.ApiKeyPrefix
	var apiKey string = xmlConfig.ApiKey

	var swaggerConfig = swagger.NewConfiguration()

	if len(baseApiPath) > 0 {
		swaggerConfig.BasePath = baseApiPath
	}

	swaggerConfig.APIKeyPrefix["Authorization"] = apiKeyPrefix
	swaggerConfig.APIKey["Authorization"] = apiKey

	return swaggerConfig
}

func (c *CliConfig) createOrReadCliConfig(param CliParams) (CliConfig, error) {

	var cliConfig CliConfig = c.getConfig()

	if cliConfig.ApiKey == "" {
		if param.apiKey == "" {
			return cliConfig, errors.New("No API key provided. Please provide Phone.com API key via -ak CLI flag of via config.xml")
		} else {
			cliConfig.ApiKey = param.apiKey
		}
	}

	if cliConfig.ApiKeyPrefix == "" {
		if param.apiKeyPrefix == "" {
			cliConfig.ApiKeyPrefix = "Bearer"
		} else {
			cliConfig.ApiKey = param.apiKeyPrefix
		}
	}

	if cliConfig.AccountId == 0 {
		if param.accountId == 0 {
			return cliConfig, errors.New("No account id provided. Please provide account id via -a CLI flag or via config.xml")
		} else {
			cliConfig.AccountId = param.accountId
		}
	}

	return cliConfig, nil
}

func (c *CliConfig) getConfigPath() string {

	if c.Path != "" {
		return c.Path
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	return dir + "/config.xml"
}
