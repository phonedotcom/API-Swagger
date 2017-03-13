#!/bin/bash

swaggerFilePath=../phonecom.json

mvn org.apache.maven.plugins:maven-dependency-plugin:2.10:get -DremoteRepositories=central -Dartifact=io.swagger:swagger-codegen-cli:LATEST -Dpackaging=jar -Ddest=swagger-codegen.jar

declare -a sdks=(
  "go"
  "android"
  "csharp"
  "java"
  "objc"
  "php"
  "python"
  "ruby"
  "swift3"
  "typescript-node"
)

sdkDir=SDKs

rm -r $sdkDir
mkdir $sdkDir
cd $sdkDir

for sdk in "${sdks[@]}"
do
  dir=$sdk-client
  java -jar ../swagger-codegen.jar generate -i $swaggerFilePath -l $sdk -o $dir

echo "Building SDK: $sdk..."

  if [ $sdk == "go" ]
  then
    echo "Bug fixing go SDK..."

    sed -i '/\/\/ to determine the Content-Type header/c\\tclearEmptyParams(localVarQueryParams)\n\n\t\/\/ to determine the Content-Type header' $dir/*_api.go
    sed -i '/case "ssv":/c\\tcase "ssv", "multi":' $dir/api_client.go

echo 'package swagger

func clearEmptyParams(paramMap map[string][]string) {

	for key, value := range paramMap {
		if (len(value) == 1 && value[0] == "") {
			delete(paramMap, key)
		}
	}
}' > $dir/util.go
  fi

  zip -r $dir.zip $dir
done
