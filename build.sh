#!/usr/bin/env bash

echo "Updating libraries..."
export GOPATH=`pwd`

go get -u github.com/urfave/cli
go get -u github.com/go-resty/resty
go get -u github.com/yukithm/json2csv/cmd/json2csv
go get -u github.com/stretchr/testify
go get -u github.com/igorsimevski/phonecom-goclient

# Directories
cliSwaggerDir=src/phonecom-go-sdk
goSourceDir=SDKs/go-client-generated

echo "Removing old phonecom Go SDK"
rm -r $cliSwaggerDir

echo "Copying and modifying new Go SDK"

cp -r $goSourceDir $cliSwaggerDir
sed -i '/\/\/ to determine the Content-Type header/c\\tclearEmptyParams(localVarQueryParams)\n\n\t\/\/ to determine the Content-Type header' $cliSwaggerDir/*_api.go
sed -i '/case "ssv":/c\\tcase "ssv", "multi":' $cliSwaggerDir/api_client.go

echo 'package swagger


func clearEmptyParams(paramMap map[string][]string) {

	for key, value := range paramMap {
		if (len(value) == 1 && value[0] == "") {
			delete(paramMap, key)
		}
	}
}' > $cliSwaggerDir/util.go

echo "Generating executables for Windows, Linux and Mac"

GOARCH=amd64 GOOS=darwin go build -o phonecom-mac phonecom
GOARCH=amd64 GOOS=windows go build -o phonecom-windows phonecom
GOARCH=amd64 GOOS=linux go build -o phonecom-linux phonecom

echo "Done."
