package main

import (
	"encoding/json"
	"github.com/urfave/cli"
	"strconv"
)

var commandFlag = "-command"

func createCliConfig() CliConfig {

	var cliConfig CliConfig
	cliConfig.Path = "../../config.json"

	return cliConfig
}

func createCliWithJsonIn(endpoint string, path string) (error, map[string]interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name:  inputLong,
			Value: path,
		},
	}

	return doCreateCli(endpoint, flags)
}

func doCreateCli(endpoint string, flags []cli.Flag) (error, map[string]interface{}) {

	app := cli.NewApp()

	app.Flags = flags

	var response (map[string]interface{})
	var err error

	app.Action = func(c *cli.Context) error {
		err, response = execute(c, createCliConfig())
		return err
	}

	app.Run([]string{commandFlag, endpoint})

	return err, response
}

func createCreateSmsCliWithJsonIn(endpoint string, path string, from string, to string, text string) (error, map[string]interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name:  inputLong,
			Value: path,
		},
		cli.StringFlag{
			Name:  fromLong,
			Value: from,
		},
		cli.StringFlag{
			Name:  toLong,
			Value: to,
		},
		cli.StringFlag{
			Name:  textLong,
			Value: text,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createCreateTrunkCliWithJsonIn(endpoint string, path string, trunkName string, trunkUri string, trunkConcurrentCalls int, trunkMaxMinutes int) (error, map[string]interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name:  inputLong,
			Value: path,
		},
		cli.StringFlag{
			Name:  trunkNameLong,
			Value: trunkName,
		},
		cli.StringFlag{
			Name:  trunkUriLong,
			Value: trunkUri,
		},
		cli.IntFlag{
			Name:  maxConcurrentCallsLong,
			Value: trunkConcurrentCalls,
		},
		cli.IntFlag{
			Name:  maxMinutesPerMonthLong,
			Value: trunkMaxMinutes,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createReplaceCliWithJsonIn(endpoint string, path string, id int) (error, map[string]interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name:  inputLong,
			Value: path,
		},
		cli.StringFlag{
			Name:  idLong,
			Value: strconv.Itoa(id),
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createCliWithIdIdSecondary(endpoint string, id int, idSecondary int) (error, map[string]interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name:  idLong,
			Value: strconv.Itoa(id),
		},
		cli.StringFlag{
			Name:  idSecondaryLong,
			Value: strconv.Itoa(idSecondary),
		},
	}

	return doCreateCli(endpoint, flags)
}

func createReplaceContactCliWithJsonIn(endpoint string, path string, id int, idSecondary int) (error, map[string]interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name:  inputLong,
			Value: path,
		},
		cli.StringFlag{
			Name:  idLong,
			Value: strconv.Itoa(id),
		},
		cli.StringFlag{
			Name:  idSecondaryLong,
			Value: strconv.Itoa(idSecondary),
		},
	}

	return doCreateCli(endpoint, flags)
}

func createReplaceTrunkCliWithJsonIn(endpoint string, path string, trunkName string, trunkUri string, trunkConcurrentCalls int, trunkMaxMinutes int, id int) (error, map[string]interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name:  idLong,
			Value: strconv.Itoa(id),
		},
		cli.StringFlag{
			Name:  inputLong,
			Value: path,
		},
		cli.StringFlag{
			Name:  trunkNameLong,
			Value: trunkName,
		},
		cli.StringFlag{
			Name:  trunkUriLong,
			Value: trunkUri,
		},
		cli.IntFlag{
			Name:  maxConcurrentCallsLong,
			Value: trunkConcurrentCalls,
		},
		cli.IntFlag{
			Name:  maxMinutesPerMonthLong,
			Value: trunkMaxMinutes,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createCliListAvailablePhoneNumbers(endpoint string, filtersPhoneNumber []string, filtersCountryCode []string, filtersNpa []string, filtersNxx []string, filtersXxxx []string, filtersCity []string, filtersProvince []string, filtersCountry []string, filtersPrice []string, filtersCategory []string, sortInternal string, sortPrice string, sortPhoneNumber string) (error, map[string]interface{}) {

	filterValues := filtersPhoneNumber[0] + ";" + filtersCountryCode[0] + ";" + filtersNpa[0] + ";" + filtersNxx[0] + ";" + filtersXxxx[0] + ";" + filtersCity[0] + ";" + filtersProvince[0] + ";" + filtersCountry[0] + ";" + filtersPrice[0] // + ";" + filtersCategory[0]
	sortValues := sortInternal + ";" + sortPrice + ";" + sortPhoneNumber
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "phone_number;country_code;npa;nxx;xxxx;city;province;country;price", //;category",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "internal;price;phone_number",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createCliListAvailablePhoneNumberRegions(endpoint string, filtersCountryCode []string, filtersNpa []string, filtersNxx []string, filtersIsTollFree []string, filtersCity []string, filtersProvincePostalCode []string, filtersCountryPostalCode []string, sortCountryCode string, sortNpa string, sortNxx string, sortIsTollFree string, sortCity string, sortProvincePostalCode string, sortCountryPostalCode string) (error, map[string]interface{}) {

	filterValues := filtersCountryCode[0] + ";" + filtersNpa[0] + ";" + filtersNxx[0] + ";" + filtersIsTollFree[0] + ";" + filtersCity[0] + ";" + filtersProvincePostalCode[0] + ";" + filtersCountryPostalCode[0]
	sortValues := sortCountryCode + ";" + sortNpa + ";" + sortNxx + ";" + sortIsTollFree + ";" + sortCity + ";" + sortProvincePostalCode + ";" + sortCountryPostalCode
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "country_code;npa;nxx;is_toll_free;city;province_postal_code;country_postal_code",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "country_code;npa;nxx;is_toll_free;city;province_postal_code;country_postal_code",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createCliSortAvailablePhoneNumberRegions(endpoint string, sortCountryCode string, sortNpa string, sortNxx string, sortIsTollFree string, sortCity string, sortProvincePostalCode string, sortCountryPostalCode string) (error, map[string]interface{}) {

	sortValues := sortCountryCode + ";" + sortNpa + ";" + sortNxx + ";" + sortIsTollFree + ";" + sortCity + ";" + sortProvincePostalCode + ";" + sortCountryPostalCode
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "country_code;npa;nxx;is_toll_free;city;province_postal_code;country_postal_code",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortAccountsSubaccountsCli(endpoint string, filtersId []string, sortId string) (error, map[string]interface{}) {

	filterValues := filtersId[0]
	sortValues := sortId
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "id",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortExpressServiceCodesCli(endpoint string, filtersId []string) (error, map[string]interface{}) {

	filterValues := filtersId[0]
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortExtensionsCli(endpoint string, filtersId []string, filtersExtension []string, filtersName []string, sortId string, sortExtension string, sortName string) (error, map[string]interface{}) {

	filterValues := filtersId[0] + ";" + filtersExtension[0] + ";" + filtersName[0]
	sortValues := sortId + ";" + sortExtension + ";" + sortName
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id;extension;name",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "id;extension;name",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortCallerIdsCliWithId(endpoint string, id int, filtersNumber []string, filtersName []string, sortNumber string, sortName string) (error, map[string]interface{}) {

	if id <= 0 {
		return nil, nil
	}

	filterValues := filtersNumber[0] + ";" + filtersName[0]
	sortValues := sortNumber + ";" + sortName
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name:  idLong,
			Value: strconv.Itoa(id),
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "number;name",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "number;name",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortContactsCliWithId(endpoint string, id int, filtersId []string, filtersGroupId []string, filtersUpdatedAt []string, sortId string, sortUpdatedAt string) (error, map[string]interface{}) {

	if id <= 0 {
		return nil, nil
	}

	filterValues := filtersId[0] + ";" + filtersGroupId[0] + ";" + filtersUpdatedAt[0]
	sortValues := sortId + ";" + sortUpdatedAt
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.StringFlag{
			Name:  idLong,
			Value: strconv.Itoa(id),
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id;group_id;updated_at",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "id;updated_at",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli(endpoint string, filtersId []string, filtersName []string, sortId string, sortName string) (error, map[string]interface{}) {

	filterValues := filtersId[0] + ";" + filtersName[0]
	sortValues := sortId + ";" + sortName
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id;name",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "id;name",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli2(endpoint string, filtersId []string /*filtersName []string, */, sortId string /*, sortName string*/) (error, map[string]interface{}) {

	filterValues := filtersId[0] // + ";" + filtersName[0]
	sortValues := sortId         // + ";" + sortName
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id", //;name",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "id", //;name",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli3(endpoint string, filtersId []string, filtersName []string, sortId string /*, sortName string*/) (error, map[string]interface{}) {

	filterValues := filtersId[0] + ";" + filtersName[0]
	sortValues := sortId // + ";" + sortName
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id;name",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "id", //;name",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortSmsCli(endpoint string, filtersId []string, filtersDirection string, filtersFrom string, sortId string, sortCreatedAt string) (error, map[string]interface{}) {

	filterValues := filtersId[0] + ";" + filtersDirection + ";" + filtersFrom
	sortValues := sortId + ";" + sortCreatedAt
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id;direction;from",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "id;created_at",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortPhoneNumbersCli(endpoint string, filtersId []string, filtersName []string /*filtersPhoneNumber []string, */, sortId string, sortName string /*, sortPhoneNumber string*/) (error, map[string]interface{}) {

	filterValues := filtersId[0] + ";" + filtersName[0] // + ";" + filtersPhoneNumber[0]
	sortValues := sortId + ";" + sortName               // + ";" + sortPhoneNumber
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id;name", //;phone_number",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "id;name", //;phone_number",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortGroupsCli(endpoint string, id int, filtersId []string, filtersName []string, sortId string, sortName string) (error, map[string]interface{}) {

	if id <= 0 {
		return nil, nil
	}

	filterValues := filtersId[0] + ";" + filtersName[0]
	sortValues := sortId + ";" + sortName
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.StringFlag{
			Name:  idLong,
			Value: strconv.Itoa(id),
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id;name",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "id;name",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createFilterSortCallLogsCli(endpoint string, filtersId []string, filtersStartTime []string, filtersCreatedAt string /*filtersDirection string, */, filtersCalledNumber string, filtersType string, filtersExtension []string, sortId string, sortStartTime string, sortCreatedAt string) (error, map[string]interface{}) {

	filterValues := filtersId[0] + ";" + filtersStartTime[0] + ";" + filtersCreatedAt + ";" /*+ filtersDirection+ ";" */ + filtersCalledNumber + ";" + filtersType + ";" + filtersExtension[0]
	sortValues := sortId + ";" + sortStartTime + ";" + sortCreatedAt
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 25,
		},
		cli.StringFlag{
			Name:  filtersTypeLong,
			Value: "id;start_time;created_at;" /*direction;*/ + "called_number;type;extension",
		},
		cli.StringFlag{
			Name:  filtersValueLong,
			Value: filterValues,
		},
		cli.StringFlag{
			Name:  sortTypeLong,
			Value: "id;start_time;created_at",
		},
		cli.StringFlag{
			Name:  sortValueLong,
			Value: sortValues,
		},
		cli.BoolTFlag{
			Name: fullListLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createCli(endpoint string) (error, map[string]interface{}) {

	flags := []cli.Flag{

		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 5,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createGetCliStringId(endpoint string, id string) (error, map[string]interface{}) {

	if id == "" {
		return nil, nil
	}

	flags := []cli.Flag{

		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name:  idLong,
			Value: id,
		},
		cli.BoolTFlag{
			Name: verboseLong,
		},
	}

	return doCreateCli(endpoint, flags)
}

func createCliWithId(endpoint string, id int) (error, map[string]interface{}) {

	if id <= 0 {
		return nil, nil
	}

	flags := []cli.Flag{

		cli.StringFlag{
			Name:  commandLong,
			Value: endpoint,
		},
		cli.StringFlag{
			Name:  idLong,
			Value: strconv.Itoa(id),
		},
		cli.IntFlag{
			Name:  limitLong,
			Value: 5,
		},
	}

	return doCreateCli(endpoint, flags)
}

func getFirstIdString(json map[string]interface{}) string {

	items := json["items"].([]interface{})
	if len(items) > 0 {
		firstItem := items[0].(map[string]interface{})
		firstId := firstItem["id"].(string)
		return firstId
	}

	return ""
}

func getFirstId(jsonObject map[string]interface{}) int {

	items := jsonObject["items"].([]interface{})
	if len(items) > 0 {
		firstItem := items[0].(map[string]interface{})
		firstId := firstItem["id"].(json.Number)
		idToReturn, _ := json.Number.Int64(firstId)
		return int(idToReturn)
	}

	return 0
}

func getFilters(jsonObject map[string]interface{}) map[string]interface{} {

	filters := jsonObject["filters"].(map[string]interface{})

	return filters
}

func getSorts(jsonObject map[string]interface{}) map[string]interface{} {

	sorts := jsonObject["sort"].(map[string]interface{})

	return sorts
}

func getFirstAvailablePhoneNumber(json map[string]interface{}) string {

	items := json["items"].([]interface{})
	if len(items) > 0 {
		firstItem := items[0].(map[string]interface{})
		firstNumber := firstItem["phone_number"].(string)
		return string(firstNumber)
	}

	return ""
}

func getId(jsonObject map[string]interface{}) int {

	id := jsonObject["id"].(json.Number)
	idToReturn, _ := json.Number.Int64(id)
	return int(idToReturn)
}

func getCallId(jsonObject map[string]interface{}) string {

	id := jsonObject["id"].(string)
	return id
}

func getName(json map[string]interface{}) string {

	name := json["name"].(string)
	return name
}
