# Scope
This document describes the CLI for the Phone.com API.  
Programming language: Golang (tested with versions 1.7.5 and 1.8)  
Tool for describing the API: Swagger

# Swagger definition & CLI
The swagger definition that describes the Phone.com API was developed in an earlier stage.  

The Swagger definition describes the API endpoints defined at [https://apidocs.phone.com](https://apidocs.phone.com) from 01.01.2017.  

The CLI tool uses the generated Swagger client for the Go programming language.  

There were 2 known bugs encountered for the generated Go client while developing the CLI. One is the duplication of the collectionFormat variable:
```golang
var collectionFormat = "multi"
```

To fix this bug the generated code had to be modified so that the duplicated definitions were removed.  

Other manual intervention that was needed to be done was to check if the filter (and similar parameters are empty, so that the generated lines of type:
```golang
localVarQueryParams.Add(
    "filters[id]",
     a.Configuration.APIClient.ParameterToString(filtersId, collectionFormat))
```
were modified with the included check:
```golang
if len(filtersId) > 0 {
 localVarQueryParams.Add(
    "filters[id]", 
     a.Configuration.APIClient.ParameterToString(filtersId, collectionFormat))
}
```
# Code organization
The there are 2 main Go packages that are used for the CLI operation:
* phonecom-go-sdk
* phonecom

The **phonecom-go-sdk** package contains all the generated swagger code along with the documentation in a form of README file and \*.md files.  
The **phonecom** package contains the code that adapts the generated Swagger code for the CLI.

# Libraries
There were 3 Go libraries that were used while developing the code. This libraries can be installed with go in the following way:
* go get github.com/urfave/cli
* go get -u github.com/go-resty/resty
* go get github.com/yukithm/json2csv/cmd/json2csv

# XML Configuration
In order the Phone.com API to be invoked, an XML configuration file should be provided. The configuration file is used to provide the OAuth2 Authentication information. It is defined with the following format:  
```xml
<?xml version="1.0" encoding="UTF-8" ?>
<Data>
    <Config>
        <Type>main</Type>
        <ApiKeyPrefix>Bearer</ApiKeyPrefix>
        <ApiKey>The API Key for Phone.com</ApiKey>
    </Config>
</Data>
```
In the current version the configuration file contains information about the authentication. This can be extended in future versions if needed.

# Flags
This section describes the flags that are used in the CLI. For each flag there is a short version and a long version for the flags. They can be used interchangeably. An excerpt from the code is given below which provides the definition of the flags.
```golang
cli.StringFlag{
  Name: "command, c",
  Value: defaultCommand,
  Usage: "Phone.com API command that you want to execute",
},
cli.StringFlag{
  Name: "id",
  Value: defaultId,
  Usage: "ID of entity you want to operate",
},
cli.IntFlag{
  Name: "id-secondary, is",
  Value: 0,
  Usage: "Secondary ID of entity you want to operate",
},
cli.IntFlag{
  Name: "limit, l",
  Value: 5,
  Usage: "Upper limit of results you want to fetch",
},
cli.IntFlag{
  Name: "offset, o",
  Value: 0,
  Usage: "Offset of results you want to fetch",
},
cli.IntFlag{
  Name: "account, a",
  Value: 1315091,
  Usage: "Phone.com API account to use in API calls",
},
cli.BoolFlag{
  Name: "dryrun, d",
  Usage: "Print the expected action without executing the API command",
},
cli.StringFlag{
  Name: "input, i",
  Value: defaultInput,
  Usage: "Specify the path to the JSON file for making the API call",
},
cli.BoolFlag{
  Name: "verbose, vr",
  Usage: "Activate verbose mode",
},
cli.StringFlag{
  Name: "contact",
  Usage: "Path to the local JSON descriptor to a contact object",
},
cli.StringFlag{
  Name: "billing-contact",
  Usage: "Path to the local JSON descriptor to a billing contact object",
},
cli.StringFlag{
  Name: "from",
  Value: defaultFrom,
  Usage: "The phone number of the SMS sender",
},
cli.StringFlag{
  Name: "to",
  Value: defaultTo,
  Usage: "The phone number of the SMS receiver",
},
cli.StringFlag{
  Name: "text",
  Value: defaultText,
  Usage: "SMS body",
},
cli.StringFlag{
  Name: "name",
  Usage: "Trunk name",
},
cli.StringFlag{
  Name: "uri",
  Usage: "Trunk uri",
},
cli.IntFlag{
  Name: "max-concurrent-calls",
  Value: 60,
  Usage: "Maximum concurrent calls for the trunk",
},
cli.IntFlag{
  Name: "max-minutes-per-month",
  Value: 800,
  Usage: "Maximum of minutes per month for the trunk",
},
cli.StringFlag{
  Name: "filtersType, ft",
  Usage: "Type of filter",
},
cli.StringFlag{
  Name: "filtersValue, fv",
  Usage: "Type of filter",
},
cli.StringFlag{
  Name: "sortType, st",
  Usage: "Type of filter",
},
cli.StringFlag{
  Name: "sortValue, sv",
  Usage: "Type of filter",
},
cli.StringFlag{
  Name: "samplein, si",
  Usage: "Generate sample input json",
},
cli.StringFlag{
  Name: "sampleout, so",
  Usage: "Generate sample output json",
},
cli.BoolFlag{
  Name: "fullList, fl",
  Usage: "Generate sample output json",
},
cli.StringFlag{
  Name: "inputFormat, if",
  Usage: "Input format. Can be 'json' or 'xml'. Default is json",
},
cli.StringFlag{
  Name: "outputFormat, of",
  Usage: "Output format. Can be 'json' or 'csv'. Default is json",
},
cli.StringFlag{
  Name: "api-key, ak",
  Usage: "The API key for Phone.com",
},
cli.StringFlag{
  Name: "api-key-prefix, akp",
  Usage: "The API key prefix for Phone.com",
},
```

# Usage
The CLI command to list all accounts is defined in the following way:
```
phonecom -c list-accounts -i ./src/test/jsonin/listAccounts.json
```

In addition to this, you can use the flags defined in the previous section. In order jus to test the invoked command you can use the ÒdryRunÓ option:
```
phonecom -c list-accounts -dr
```

The **Get** command can be invoked in two ways:
```
phonecom -c get-accounts -i ./src/test/jsonin/getAccounts.json
```
where the account id to be get is defined in the input json and:
```
phonecom -c get-accounts -id 1234
```
where the account id is provided directly through the console.  

For the List, Get and Delete calls the input JSONs are optional. For the Create and Replace calls the input JSON files are required.

# Integration tests
The code contains integration tests located in the ÒphonecomÓ package, along with sample json files that are located in the src/test/jsonin directory.  
These files are only used for testing, and they are not necessary for the operation of the CLI:
* allCreateGetReplaceDelete_test.go
* allListGetEndpoints_test.go
* testUtil.go

The tests are standard go tests and can be run with the go test command.

# Error handling & display
The CLI handles and displays usage errors and http/network errors. The network errors messages are taken from the go-resty library. The usage messages are taken from the urfave-cli library. The HTTP errors are taken from the Phone.com API response. Additionally, with the -vr flag additional information about the error might be provided.

# Deliverables
The code for the CLI along with the tests and the sample input JSONs will be pushed to Phone.com code repository. Additionally, all the changes made during the process of finding workarounds for overcoming the swagger bugs will be provided.
