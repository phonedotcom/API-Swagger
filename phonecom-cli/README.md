## Scope
This document describes the CLI for the Phone.com API.  
Programming language: Golang (tested with versions 1.7.5 and 1.8)  
Tool for describing the API: Swagger

## Swagger definition & CLI
The swagger definition that describes the Phone.com API was developed in an earlier stage.  

The Swagger definition describes the API endpoints defined at [https://apidocs.phone.com](https://apidocs.phone.com) from 01.01.2017.  

The CLI tool uses the generated Swagger client for the Go programming language.  

## Swagger Go Generated Code Workarounds
A bug was encountered for the generated Go client while developing the CLI.
The Phone.com expects non-empty parameters, but Swagger Go generator adds parameters to the request even if the value is empty.

#####Workaround:
This function was added:
```go
func clearEmptyParams(paramMap map[string][]string) {

    for key, value := range paramMap {
        if (len(value) == 1 && value[0] == "") {
            delete(paramMap, key)
        }
    }
}
```
Which is called in each generated API endpoint after the list of parameters is created.

## Code organization
There are 2 main Go packages that are used for the CLI operation:
* phonecom-go-sdk
* phonecom

The **phonecom-go-sdk** package contains all the generated swagger code along with the documentation in a form of README file and \*.md files.  
The **phonecom** package contains the code that adapts the generated Swagger code for the CLI.

## Libraries
There were 4 Go libraries that were used while developing the code. This libraries can be installed with go in the following way:
```bash
go get -u github.com/urfave/cli
go get -u github.com/go-resty/resty
go get -u github.com/yukithm/json2csv/cmd/json2csv
go get -u github.com/stretchr/testify
``` 


## JSON Configuration
In order the Phone.com API to be invoked, a JSON configuration file should be provided. The configuration file is used to provide the OAuth2 Authentication information. It is defined with the following format:  
```json
{
  "config": [
    {
      "type" : "main",
      "apiKeyPrefix" : "Bearer",
      "apiKey" : "Place your API key here",
      "accountId" : 1315091
    }
  ]
}
```
In the current version the configuration file contains information about the authentication.

Configuration fields:

* BaseApiPath - `string` The base API path. The default is [https://api.phone.com/v4](https://api.phone.com/v4)
* type - `string` The type of configuration. Planned for future extensions for defining multiple configurations. The value should be set to "main".
* apiKeyPrefix - `string` The prefix for the API key. Default: _Bearer_
* apiKey - `string` The API key
* accountId - `integer` The user account id

You can also define the API key token as a CLI parameter by specifying it with the -ak flag.

## Flags
This section describes the flags that are used in the CLI. For each flag there is a short version and a long version for the flags. They can be used interchangeably. An excerpt from the code is given below which provides the definition of the flags.
```go
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
  Name: "verbose-mode, vrm",
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
  Usage: "The filter value",
},
cli.StringFlag{
  Name: "sortType, st",
  Usage: "Type of sort",
},
cli.StringFlag{
  Name: "sortValue, sv",
  Usage: "The sort value",
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
  Usage: "Full list flag. If set, aditional information about the list is given",
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

## Usage
The CLI command to list all accounts is defined in the following way:
```
phonecom -c list-accounts -i ./src/test/jsonin/listAccounts.json
```

In addition to this, you can use the flags defined in the previous section. In order jus to test the invoked command you can use the "dryRun" option:
```
phonecom -c list-accounts -dr
```

The **Get** command can be invoked in two ways:
```
phonecom -c get-account -i ./src/test/jsonin/getAccounts.json
```
where the account id to be get is defined in the input json and:
```
phonecom -c get-account -id 1234
```
where the account id is provided directly through the console.  

For the List, Get and Delete calls the input JSONs are optional. For the Create and Replace calls the input JSON files are required.

## Filters and sorts

The fitlers and sorts are specific types of parameters.
There are 2 flags that describe each filter and sort.

For filters: 
* -ft (filter type) and -fv (filter value)

For sorts:
* -st (sort type) and -sv (sort value)

For example if you want to filter accounts by id grater than 20:
```
phonecom -c list-accounts -ft id -fv gt:20
```
Similarly, if you want to sort accounts by id in descending order:
```
phonecom -c list-accounts -st id -sv desc
```

## The full list flag

The default response for all ListGet API calls is a JSON response containing a list of all items of that list.

If you want to display the filters, sorts, limit, offset and total number of items in a wrapper object you need to add the -fl flag.
```
phonecom -c list-accounts -fl
```

## The send sms flags

There are 4 arguments that can be used with the create SMS command:
* from (flag: -from)
* to (flag: -to)
* text (flag: -text)
* extension id (flag: -id)

The command to send sms is of the following form:
```
phonecom -c create-account-sms -from {number1} -to {number2} -text {text} -id {extension id} 
```

## List of commands

The following commands can be invoked from the API:

```
create-account-calls
create-account-device
create-account-extension
create-account-menu
create-account-phone-number
create-account-queue
create-account-route
create-account-sms
create-account-subaccount
create-account-trunk
create-account-extension-contact
create-account-extension-contact-group

delete-account-menu
delete-account-queue
delete-account-route
delete-account-trunk
delete-account-extension-contact
delete-account-extension-contact-group

get-account
get-account-application
get-account-call-log
get-account-device
get-account-express-service-code
get-account-extension
get-account-media
get-account-menu
get-account-phone-number
get-account-queue
get-account-route
get-account-schedule
get-account-sms
get-account-trunk
get-account-extension-contact
get-account-extension-contact-group

list-accounts
list-account-applications
list-account-call-logs
list-account-devices
list-account-express-service-codes
list-account-extensions
list-account-media
list-account-menus
list-account-phone-numbers
list-account-queues
list-account-routes
list-account-schedules
list-account-sms
list-account-subaccounts
list-account-trunks
list-account-extension-caller-ids
list-account-extension-contacts
list-account-extension-contact-groups
list-available-phone-numbers
list-available-phone-number-regions

replace-account-device
replace-account-extension
replace-account-menu
replace-account-phone-number
replace-account-queue
replace-account-route
replace-account-trunk
replace-account-extension-contact
replace-account-extension-contact-group

 (default: "https://api.phone.com/v4/ping")
```

## Integration tests
The code contains integration tests located in the "phonecom" package, along with sample json files that are located in the src/test/jsonin directory.  
These files are only used for testing, and they are not necessary for the operation of the CLI:
* allCreateGetReplaceDelete_test.go
* allListGetEndpoints_test.go
* testUtil.go

The tests are standard go tests and can be run with the go test command.

## Error handling & display
The CLI handles and displays usage errors and http/network errors. The network errors messages are taken from the go-resty library. The usage messages are taken from the urfave-cli library. The HTTP errors are taken from the Phone.com API response. Additionally, with the -vrm flag additional information about the error might be provided.

## Deliverables
The code for the CLI along with the tests and the sample input JSONs will be pushed to Phone.com code repository. Additionally, all the changes made during the process of finding workarounds for overcoming the swagger bugs will be provided.
