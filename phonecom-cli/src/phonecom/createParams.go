package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"phonecom-go-sdk"
)

type ListParams struct {
	accountId int32
	limit     int32
	offset    int32
	fields    string
}

func getListParams(inputFile string) (error, ListParams) {

	err, data := readAndUnmarshal(inputFile)

	var params ListParams

	if err != nil || data == nil {
		return err, params
	}

	params.accountId = getField(data["account_id"])
	params.limit = getField(data["limit"])
	params.offset = getField(data["offset"])
	params.fields = getFieldString(data["fields"])

	return err, params
}

func getField(field interface{}) int32 {

	if field == nil {
		return -1
	}

	return int32(field.(float64))
}

func readAndUnmarshal(inputFile string) (error, map[string]interface{}) {

	file, e := ioutil.ReadFile(inputFile)

	if e != nil {
		return errors.New(fmt.Sprintf(couldNotReadFile, inputFile)), nil
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(file, &dat); err != nil {
		return errors.New(fmt.Sprintf(couldNotUnmarshal, inputFile)), nil
	}

	return nil, dat
}

func getFieldString(field interface{}) string {

	if field == nil {
		return ""
	}

	return field.(string)
}

type FilterParams struct {
	filtersId                 []string
	filtersExtension          []string
	filtersGroupId            []string
	filtersUpdatedAt          []string
	filtersPhoneNumber        []string
	filtersName               []string
	filtersNumber             []string
	filtersDirection          string
	filtersCalledNumber       string
	filtersType               string
	filtersCountryCode        []string
	filtersCountry            []string
	filtersNpa                []string
	filtersNxx                []string
	filtersXxxx               []string
	filtersCity               []string
	filtersProvince           []string
	filtersPrice              []string
	filtersCategory           []string
	filtersFrom               string
	filtersTo                 []string
	filtersIsTollFree         []string
	filtersProvincePostalCode []string
	filtersCountryPostalCode  []string
	filtersStartTime          []string
	filtersCreatedAt          string
}

func getFiltersParams(inputFile string) (
	error,
	FilterParams) {

	err, data := readAndUnmarshal(inputFile)

	var params FilterParams

	if err != nil || data == nil {
		return err, params
	}

	params.filtersId = createStringArray(data["filters[id]"])
	params.filtersExtension = createStringArray(data["filters[extension]"])
	params.filtersGroupId = createStringArray(data["filters[group_id]"])
	params.filtersUpdatedAt = createStringArray(data["filters[updated_at]"])
	params.filtersPhoneNumber = createStringArray(data["filters[phone_number]"])
	params.filtersName = createStringArray(data["filters[name]"])
	params.filtersNumber = createStringArray(data["filters[number]"])
	params.filtersDirection = getFieldString(data["filters[direction]"])
	params.filtersCalledNumber = getFieldString(data["filters[called_number]"])
	params.filtersType = getFieldString(data["filters[type]"])
	params.filtersCountryCode = createStringArray(data["filters[country_code]"])
	params.filtersCountry = createStringArray(data["filters[countries]"])
	params.filtersNpa = createStringArray(data["filters[npa]"])
	params.filtersNxx = createStringArray(data["filters[nxx]"])
	params.filtersXxxx = createStringArray(data["filters[xxxx]"])
	params.filtersCity = createStringArray(data["filters[city]"])
	params.filtersProvince = createStringArray(data["filters[province]"])
	params.filtersPrice = createStringArray(data["filters[price]"])
	params.filtersCategory = createStringArray(data["filters[category]"])
	params.filtersFrom = getFieldString(data["filters[from]"])
	params.filtersTo = createStringArray(data["filters[to]"])
	params.filtersIsTollFree = createStringArray(data["filters[is_toll_free]"])
	params.filtersProvincePostalCode = createStringArray(data["filters[province_postal_code]"])
	params.filtersCountryPostalCode = createStringArray(data["filters[country_postal_code]"])
	params.filtersStartTime = createStringArray(data["filters[start_time]"])
	params.filtersCreatedAt = getFieldString(data["filters[created_at]"])

	return err, params
}

func createStringArray(filter interface{}) []string {

	if filter == nil {
		return nil
	}

	filterArray := filter.([]interface{})
	str1 := filterArray[0].(string)

	if len(filterArray) == 1 && filterArray[0] == "" {
		return make([]string, 0)
	}

	return []string{str1}
}

type OtherParams struct {
	groupBy     []string
	extensionId int32
	callId      string
	deviceId    int32
	codeId      int32
	contactId   int32
	groupId     int32
	recordingId int32
	menuId      int32
	numberId    int32
	queueId     int32
	routeId     int32
	scheduleId  int32
	smsId       string
	trunkId     int32
}

func getOtherParams(inputFile string) (
	error,
	OtherParams) {

	err, data := readAndUnmarshal(inputFile)

	var params OtherParams

	if err != nil || data == nil {
		return err, params
	}

	params.groupBy = createStringArray(data["group_by"])
	params.extensionId = getField(data["extension_id"])
	params.callId = getFieldString(data["call_id"])
	params.deviceId = getField(data["device_id"])
	params.codeId = getField(data["code_id"])
	params.contactId = getField(data["contact_id"])
	params.groupId = getField(data["group_id"])
	params.recordingId = getField(data["recording_id"])
	params.menuId = getField(data["menu_id"])
	params.numberId = getField(data["number_id"])
	params.queueId = getField(data["queue_id"])
	params.routeId = getField(data["route_id"])
	params.scheduleId = getField(data["schedule_id"])
	params.smsId = getFieldString(data["sms_id"])
	params.trunkId = getField(data["trunk_id"])

	return err, params
}

type SortParams struct {
	sortId                 string
	sortPhoneNumber        string
	sortNumber             string
	sortName               string
	sortInternal           string
	sortPrice              string
	sortExtension          string
	sortUpdatedAt          string
	sortCreatedAt          string
	sortStartTime          string
	sortCountryCode        string
	sortNpa                string
	sortNxx                string
	sortIsTollFree         string
	sortCity               string
	sortProvincePostalCode string
	sortCountryPostalCode  string
}

func getSortParams(inputFile string) (
	error,
	SortParams) {

	err, data := readAndUnmarshal(inputFile)
	var params SortParams

	if err != nil || data == nil {
		return err, params
	}

	params.sortId = getFieldString(data["sort[id]"])
	params.sortPhoneNumber = getFieldString(data["sort[phone_number]"])
	params.sortNumber = getFieldString(data["sort[number]"])
	params.sortName = getFieldString(data["sort[name]"])
	params.sortInternal = getFieldString(data["sort[internal]"])
	params.sortPrice = getFieldString(data["sort[price]"])
	params.sortExtension = getFieldString(data["sort[extension]"])
	params.sortUpdatedAt = getFieldString(data["sort[updated_at]"])
	params.sortCreatedAt = getFieldString(data["sort[created_at]"])
	params.sortStartTime = getFieldString(data["sort[start_time]"])
	params.sortCountryCode = getFieldString(data["sort[country_code]"])
	params.sortNpa = getFieldString(data["sort[npa]"])
	params.sortNxx = getFieldString(data["sort[nxx]"])
	params.sortIsTollFree = getFieldString(data["sort[is_toll_free]"])
	params.sortCity = getFieldString(data["sort[city]"])
	params.sortProvincePostalCode = getFieldString(data["sort[province_postal_code]"])
	params.sortCountryPostalCode = getFieldString(data["sort[country_postal_code]"])

	return err, params
}

func createDeviceParams(inputFile string) swagger.CreateDeviceParams {

	var params swagger.CreateDeviceParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createCallParams(inputFile string) swagger.CreateCallParams {

	var params swagger.CreateCallParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func readAndUnmarshalFile(inputFile string, params interface{}) error {

	file, e := ioutil.ReadFile(inputFile)

	if e != nil {
		return e
	}

	fmt.Printf("%s\n", string(file))

	json.Unmarshal(file, params)

	return nil
}

func createMenuParams(inputFile string) swagger.CreateMenuParams {

	var params swagger.CreateMenuParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func replaceMenuParams(inputFile string) swagger.ReplaceMenuParams {

	var params swagger.ReplaceMenuParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createExtensionsParams(inputFile string) swagger.CreateExtensionParams {

	var params swagger.CreateExtensionParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createQueueParams(inputFile string) swagger.CreateQueueParams {

	var params swagger.CreateQueueParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createRouteParams(inputFile string) swagger.CreateRouteParams {

	var params swagger.CreateRouteParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createSmsParams(from string, to string, text string, extensionId int32) swagger.CreateSmsParams {

	var par swagger.CreateSmsParams
	par.From = from
	par.To = to
	par.Text = text
	par.ExtensionId = extensionId

	return par
}

func createSubaccountParams(inputFile string, contact string, billingContact string) swagger.CreateSubaccountParams {

	var params swagger.CreateSubaccountParams
	if contact != "" {
		readAndUnmarshalFile(inputFile, &params)
		var contactObject swagger.ContactSubaccount
		var billingContactObject swagger.ContactSubaccount
		readAndUnmarshalFile(contact, &contactObject)
		readAndUnmarshalFile(billingContact, &billingContactObject)
		params.Contact = contactObject
		params.BillingContact = billingContactObject
		return params

	}
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createTrunkParams(name string, uri string, max_concurrent_calls int32, max_minutes_per_month int32) swagger.CreateTrunkParams {

	var params swagger.CreateTrunkParams

	params.Name = name
	params.Uri = uri
	params.MaxConcurrentCalls = max_concurrent_calls
	params.MaxMinutesPerMonth = max_minutes_per_month

	return params
}

func replaceExtensionParams(inputFile string) swagger.ReplaceExtensionParams {

	var params swagger.ReplaceExtensionParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createContactParams(inputFile string) swagger.CreateContactParams {

	var params swagger.CreateContactParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createGroupParams(inputFile string) swagger.CreateGroupParams {

	var params swagger.CreateGroupParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func createPhoneNumberParams(inputFile string) swagger.CreatePhoneNumberParams {

	var params swagger.CreatePhoneNumberParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}

func replacePhoneNumberParams(inputFile string) swagger.ReplacePhoneNumberParams {

	var params swagger.ReplacePhoneNumberParams
	readAndUnmarshalFile(inputFile, &params)

	return params
}
