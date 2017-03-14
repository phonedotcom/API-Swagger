#!/bin/bash


scriptDir=`pwd`
sdkDir=$1
swaggerFilePath=$scriptDir/phonecom.json

if [[ $# -lt 2 ]] ; then
  echo 'Usage: First argument: path to SDK, Second Argument: Path to Java SDK'
  exit 0
fi

echo "SDK dir path: $sdkDir"

mvn org.apache.maven.plugins:maven-dependency-plugin:2.10:get -DremoteRepositories=central -Dartifact=io.swagger:swagger-codegen-cli:LATEST -Dpackaging=jar -Ddest=$scriptDir/swagger-codegen.jar

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

find $sdkDir/*-client/* ! -iregex '(.git)' | xargs rm -r
cd $sdkDir

for sdk in "${sdks[@]}"
do
  dir=$sdkDir/$sdk-client

  if [ $sdk != "java" ]
  then
    java -jar $scriptDir/swagger-codegen.jar generate -i $swaggerFilePath -l $sdk -o $dir
  fi

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

  elif [ $sdk == "java" ]
  then
    echo "Copying Java SDK..."
    cp -r $2/* $dir
  fi

  #zip -r $dir.zip $dir
done
