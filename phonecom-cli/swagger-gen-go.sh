#!/usr/bin/env bash

scriptDir=`pwd`
swaggerFilePath=$scriptDir/phonecom-apiv4.swagger.json

echo "SDK dir path: $scriptDir"

mvn org.apache.maven.plugins:maven-dependency-plugin:2.10:get -DremoteRepositories=central -Dartifact=io.swagger:swagger-codegen-cli:LATEST -Dpackaging=jar -Ddest=$scriptDir/swagger-codegen.jar

cd $scriptDir

dir=$scriptDir/src/github.com/phonedotcom/API-SDK-go
rm -rf $dir

echo "Building SDK: go..."
java -jar $scriptDir/swagger-codegen.jar generate -i $swaggerFilePath -l go -o $dir

sed -i '/\/\/ to determine the Content-Type header/c\\tclearEmptyParams(localVarQueryParams)\n\n\t\/\/ to determine the Content-Type header' $dir/*_api.go
sed -i '/case "ssv":/c\\tcase "ssv", "multi":' $dir/api_client.go

# Patch media api
sed -i 's/json string/jsonParam string/' $dir/media_api.go
sed -i 's/(json, "")/(jsonParam, "")/' $dir/media_api.go
sed -i 's/@param json/@param jsonParam/' $dir/media_api.go

echo 'package swagger

func clearEmptyParams(paramMap map[string][]string) {

	for key, value := range paramMap {
		if (len(value) == 1 && value[0] == "") {
			delete(paramMap, key)
		}
	}
}' > $dir/util.go
