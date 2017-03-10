package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListAccounts(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listAccounts)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getAccount, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListAccounts(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1315091"
	idSlice = append(idSlice, expectedId)

	sortId := "asc"

	err, response := createFilterSortAccountsSubaccountsCli(listAccounts, idSlice, sortId)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
}

func TestFilterWithSpaceInParamValue(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "str1 str2"
	idSlice = append(idSlice, expectedId)

	err, response := createFilterSortExpressServiceCodesCli(listAccounts, idSlice)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
}

func TestFilterWithHashInParamValue(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "str1#str2"
	idSlice = append(idSlice, expectedId)

	err, response := createFilterSortExpressServiceCodesCli(listAccounts, idSlice)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
}

func TestListApplications(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listApplications)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getApplication, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListApplications(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	//nameSlice := make([]string, 0)
	//expectedName := "internal_only"
	//nameSlice = append(nameSlice, expectedName)

	sortId := "asc"
	//sortName := "asc"

	err, response := createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli2(listApplications, idSlice /*nameSlice, */, sortId /*, sortName*/)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	//name := filters["name"].(string)
	//assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	//assert.Equal(t, sortName, sorts["name"])
}

func TestListCallLogs(t *testing.T) {

	var err error

	err, _ = createCli(listCallLogs)
	assert.NoError(t, err)
}

func TestFilterSortListCallLogs(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	startTimeSlice := make([]string, 0)
	expectedStartTime := "internal_only"
	startTimeSlice = append(startTimeSlice, expectedStartTime)
	expectedCreatedAt := "internal_only"
	//expectedDirection := "internal_only"
	expectedCalledNumber := "internal_only"
	expectedType := "internal_only"
	extensionSlice := make([]string, 0)
	expectedExtension := "internal_only"
	extensionSlice = append(extensionSlice, expectedExtension)

	sortId := "asc"
	sortStartTime := "asc"
	sortCreatedAt := "asc"

	err, response := createFilterSortCallLogsCli(listCallLogs, idSlice, startTimeSlice, expectedCreatedAt /*expectedDirection, */, expectedCalledNumber, expectedType, extensionSlice, sortId, sortStartTime, sortCreatedAt)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	startTime := filters["start_time"].(string)
	assert.Equal(t, expectedStartTime, startTime)
	assert.Equal(t, expectedCreatedAt, filters["created_at"])
	//assert.Equal(t, expectedDirection, filters["direction"])
	assert.Equal(t, expectedCalledNumber, filters["called_number"])
	assert.Equal(t, expectedType, filters["type"])
	extension := filters["extension"].(string)
	assert.Equal(t, expectedExtension, extension)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortStartTime, sorts["start_time"])
	assert.Equal(t, sortCreatedAt, sorts["created_at"])
}

func TestListDevices(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listDevices)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getDevice, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListDevices(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	nameSlice := make([]string, 0)
	expectedName := "internal_only"
	nameSlice = append(nameSlice, expectedName)

	sortId := "asc"
	sortName := "asc"

	err, response := createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli(listDevices, idSlice, nameSlice, sortId, sortName)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortName, sorts["name"])
}

func TestExpressServiceCodes(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listExpressServiceCodes)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getExpressServiceCode, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortExpressServiceCodes(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "324202"
	idSlice = append(idSlice, expectedId)

	err, response := createFilterSortExpressServiceCodesCli(listExpressServiceCodes, idSlice)
	assert.NoError(t, err)
	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
}

func TestListExtensions(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listExtensions)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getExtension, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListExtensions(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1767963"
	idSlice = append(idSlice, expectedId)
	extensionSlice := make([]string, 0)
	expectedExtension := "111"
	extensionSlice = append(extensionSlice, expectedExtension)
	nameSlice := make([]string, 0)
	expectedName := "ws3mo6pehbk5"
	nameSlice = append(nameSlice, expectedName)

	sortId := "asc"
	sortExtension := "asc"
	sortName := "asc"

	err, response := createFilterSortExtensionsCli(listExtensions, idSlice, extensionSlice, nameSlice, sortId, sortExtension, sortName)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	extension := filters["extension"].(string)
	assert.Equal(t, expectedExtension, extension)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortExtension, sorts["extension"])
	assert.Equal(t, sortName, sorts["name"])
}

func TestListCallerIds(t *testing.T) {

	var err error
	var json map[string]interface{}

	err, json = createCli(listExtensions)
	assert.NoError(t, err)
	extensionId := getFirstId(json)

	err, _ = createCliWithId(listCallerIds, extensionId)
	assert.NoError(t, err)
}

func TestFilterSortListCallerIds(t *testing.T) {

	numberSlice := make([]string, 0)
	expectedNumber := "ws3mo6pehbk5"
	numberSlice = append(numberSlice, expectedNumber)
	nameSlice := make([]string, 0)
	expectedName := "ws3mo6pehbk5"
	nameSlice = append(nameSlice, expectedName)

	sortNumber := "asc"
	sortName := "asc"

	var err error
	var json map[string]interface{}

	err, json = createCli(listExtensions)
	assert.NoError(t, err)
	extensionId := getFirstId(json)

	err, response := createFilterSortCallerIdsCliWithId(listCallerIds, extensionId, numberSlice, nameSlice, sortNumber, sortName)
	assert.NoError(t, err)

	filters := getFilters(response)
	number := filters["number"].(string)
	assert.Equal(t, expectedNumber, number)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortNumber, sorts["number"])
	assert.Equal(t, sortName, sorts["name"])
}

func TestListContacts(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listExtensions)
	assert.NoError(t, err)
	extensionId := getFirstId(json)

	err, json = createCliWithId(listContacts, extensionId)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithIdIdSecondary(getContact, extensionId, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListContacts(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "ws3mo6pehbk5"
	idSlice = append(idSlice, expectedId)
	groupIdSlice := make([]string, 0)
	expectedGroupId := "ws3mo6pehbk5"
	groupIdSlice = append(groupIdSlice, expectedGroupId)
	updatedAtSlice := make([]string, 0)
	expectedUpdatedAt := "ws3mo6pehbk5"
	updatedAtSlice = append(updatedAtSlice, expectedUpdatedAt)

	sortId := "asc"
	sortUpdatedAt := "asc"

	var json map[string]interface{}
	var err error

	err, json = createCli(listExtensions)
	assert.NoError(t, err)
	extensionId := getFirstId(json)

	err, response := createFilterSortContactsCliWithId(listContacts, extensionId, idSlice, groupIdSlice, updatedAtSlice, sortId, sortUpdatedAt)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	groupId := filters["group_id"].(string)
	assert.Equal(t, expectedGroupId, groupId)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortUpdatedAt, sorts["updated_at"])
}

func TestListGroups(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listExtensions)
	assert.NoError(t, err)
	extensionId := getFirstId(json)

	err, json = createCliWithId(listGroups, extensionId)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithIdIdSecondary(getGroup, extensionId, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListGroups(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	nameSlice := make([]string, 0)
	expectedName := "internal_only"
	nameSlice = append(nameSlice, expectedName)

	sortId := "asc"
	sortName := "asc"

	var json map[string]interface{}
	var err error

	err, json = createCli(listExtensions)
	assert.NoError(t, err)
	extensionId := getFirstId(json)

	err, response := createFilterSortGroupsCli(listGroups, extensionId, idSlice, nameSlice, sortId, sortName)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortName, sorts["name"])
}

func TestListMedia(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listMedia)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getRecording, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListMedia(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	nameSlice := make([]string, 0)
	expectedName := "internal_only"
	nameSlice = append(nameSlice, expectedName)

	sortId := "asc"
	sortName := "asc"

	err, response := createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli(listMedia, idSlice, nameSlice, sortId, sortName)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortName, sorts["name"])
}

func TestListMenus(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listMenus)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getMenu, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListMenus(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	nameSlice := make([]string, 0)
	expectedName := "internal_only"
	nameSlice = append(nameSlice, expectedName)

	sortId := "asc"
	sortName := "asc"

	err, response := createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli(listMenus, idSlice, nameSlice, sortId, sortName)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortName, sorts["name"])
}

func TestListPhoneNumbers(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listPhoneNumbers)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getPhoneNumber, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListPhoneNumbers(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	nameSlice := make([]string, 0)
	expectedName := "internal_only"
	nameSlice = append(nameSlice, expectedName)
	//phoneNumberSlice := make([]string, 0)
	//expectedPhoneNumber := "internal_only"
	//phoneNumberSlice = append(phoneNumberSlice, expectedPhoneNumber)

	sortId := "asc"
	sortName := "asc"
	//sortPhoneNumber := "asc"

	err, response := createFilterSortPhoneNumbersCli(listPhoneNumbers, idSlice, nameSlice /*phoneNumberSlice, */, sortId, sortName /*, sortPhoneNumber*/)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)
	//phoneNumber := filters["phone_number"].(string)
	//assert.Equal(t, expectedPhoneNumber, phoneNumber)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortName, sorts["name"])
	//assert.Equal(t, sortPhoneNumber, sorts["phone_number"])
}

func TestListQueues(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listQueues)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getQueue, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListQueues(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	nameSlice := make([]string, 0)
	expectedName := "internal_only"
	nameSlice = append(nameSlice, expectedName)

	sortId := "asc"
	sortName := "asc"

	err, response := createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli(listQueues, idSlice, nameSlice, sortId, sortName)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortName, sorts["name"])
}

func TestListRoutes(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listRoutes)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getRoute, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListRoutes(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	nameSlice := make([]string, 0)
	expectedName := "internal_only"
	nameSlice = append(nameSlice, expectedName)

	sortId := "asc"
	sortName := "asc"

	err, response := createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli(listRoutes, idSlice, nameSlice, sortId, sortName)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortName, sorts["name"])
}

func TestListSchedules(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listSchedules)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getSchedule, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListSchedules(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	nameSlice := make([]string, 0)
	expectedName := "internal_only"
	nameSlice = append(nameSlice, expectedName)

	sortId := "asc"
	//sortName := "asc"

	err, response := createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli3(listSchedules, idSlice, nameSlice, sortId /*, sortName*/)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	//assert.Equal(t, sortName, sorts["name"])
}

func TestListSms(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listSms)
	assert.NoError(t, err)

	firstId := getFirstIdString(json)
	err, json = createGetCliStringId(getSms, firstId)

	assert.NoError(t, err)
}

func TestFilterSortListSms(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	expectedDirection := "in"
	expectedFrom := "+16309624775"

	sortId := "asc"
	sortCreatedAt := "asc"

	err, response := createFilterSortSmsCli(listSms, idSlice, expectedDirection, expectedFrom, sortId, sortCreatedAt)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	assert.Equal(t, expectedDirection, filters["direction"])
	assert.Equal(t, expectedFrom, filters["from"])

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortCreatedAt, sorts["created_at"])
}

func TestListSubaccounts(t *testing.T) {

	var err error

	err, _ = createCli(listSubaccounts)
	assert.NoError(t, err)
}

func TestFilterSortListSubaccounts(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1315091"
	idSlice = append(idSlice, expectedId)

	sortId := "asc"

	err, response := createFilterSortAccountsSubaccountsCli(listSubaccounts, idSlice, sortId)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
}

func TestListTrunks(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listTrunks)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	err, json = createCliWithId(getTrunk, int(firstId))

	assert.NoError(t, err)
}

func TestFilterSortListTrunks(t *testing.T) {

	idSlice := make([]string, 0)
	expectedId := "1356077"
	idSlice = append(idSlice, expectedId)
	nameSlice := make([]string, 0)
	expectedName := "internal_only"
	nameSlice = append(nameSlice, expectedName)

	sortId := "asc"
	sortName := "asc"

	err, response := createFilterSortApplicationsDevicesMediaMenusQueusRoutesSchedulesTrunksCli(listTrunks, idSlice, nameSlice, sortId, sortName)
	assert.NoError(t, err)

	filters := getFilters(response)
	id := filters["id"].(string)
	assert.Equal(t, expectedId, id)
	name := filters["name"].(string)
	assert.Equal(t, expectedName, name)

	sorts := getSorts(response)
	assert.Equal(t, sortId, sorts["id"])
	assert.Equal(t, sortName, sorts["name"])
}

func TestListAvailablePhoneNumbers(t *testing.T) {

	var err error

	err, _ = createCli(listAvailablePhoneNumbers)
	assert.NoError(t, err)
}

func TestListAvailablePhoneNumberRegions(t *testing.T) {

	var err error

	err, _ = createCli(listAvailablePhoneNumberRegions)
	assert.NoError(t, err)
}

func TestFilterSortListAvailablePhoneNumbers(t *testing.T) {

	phoneNumberSlice := make([]string, 0)
	expectedPhoneNumber := "+12234687"
	phoneNumberSlice = append(phoneNumberSlice, expectedPhoneNumber)
	countryCodeSlice := make([]string, 0)
	expectedCountryCode := "1"
	countryCodeSlice = append(countryCodeSlice, expectedCountryCode)
	npaSlice := make([]string, 0)
	expectedNpa := "gt:10"
	npaSlice = append(npaSlice, expectedNpa)
	nxxSlice := make([]string, 0)
	expectedNxx := "831"
	nxxSlice = append(nxxSlice, expectedNxx)
	xxxxSlice := make([]string, 0)
	expectedXxxx := "7863"
	xxxxSlice = append(xxxxSlice, expectedXxxx)
	citySlice := make([]string, 0)
	expectedCity := "Martinsville"
	citySlice = append(citySlice, expectedCity)
	provinceSlice := make([]string, 0)
	expectedProvince := "IN"
	provinceSlice = append(provinceSlice, expectedProvince)
	countrySlice := make([]string, 0)
	expectedCountry := "US"
	countrySlice = append(countrySlice, expectedCountry)
	priceSlice := make([]string, 0)
	expectedPrice := "0"
	priceSlice = append(priceSlice, expectedPrice)
	categorySlice := make([]string, 0) // check filter
	expectedCategory := "10"
	categorySlice = append(categorySlice, expectedCategory)

	sortInternal := "asc"
	sortPrice := "asc"
	sortPhoneNumber := "desc"

	err, response := createCliListAvailablePhoneNumbers(listAvailablePhoneNumbers, phoneNumberSlice, countryCodeSlice, npaSlice, nxxSlice, xxxxSlice, citySlice, provinceSlice, countrySlice, priceSlice, categorySlice, sortInternal, sortPrice, sortPhoneNumber)
	assert.NoError(t, err)

	filters := getFilters(response)
	phoneNumber := filters["phone_number"].(string)
	assert.Equal(t, expectedPhoneNumber, phoneNumber)
	countryCode := filters["country_code"].(string)
	assert.Equal(t, expectedCountryCode, countryCode)
	npa := filters["npa"].(string)
	assert.Equal(t, expectedNpa, npa)
	nxx := filters["nxx"].(string)
	assert.Equal(t, expectedNxx, nxx)
	xxxx := filters["xxxx"].(string)
	assert.Equal(t, expectedXxxx, xxxx)
	city := filters["city"].(string)
	assert.Equal(t, expectedCity, city)
	province := filters["province"].(string)
	assert.Equal(t, expectedProvince, province)
	country := filters["country"].(string)
	assert.Equal(t, expectedCountry, country)
	price := filters["price"].(string)
	assert.Equal(t, expectedPrice, price)
	//category := filters["category"].(string)
	//assert.Equal(t, expectedCategory, category)

	sorts := getSorts(response)
	assert.Equal(t, sortInternal, sorts["internal"])
	assert.Equal(t, sortPrice, sorts["price"])
	assert.Equal(t, sortPhoneNumber, sorts["phone_number"])
}

func TestFilterSortListAvailablePhoneNumberRegions(t *testing.T) {

	countryCodeSlice := make([]string, 0)
	expectedCountryCode := "1"
	countryCodeSlice = append(countryCodeSlice, expectedCountryCode)
	npaSlice := make([]string, 0)
	expectedNpa := "gt:10"
	npaSlice = append(npaSlice, expectedNpa)
	nxxSlice := make([]string, 0)
	expectedNxx := "249"
	nxxSlice = append(nxxSlice, expectedNxx)
	isTollFreeSlice := make([]string, 0)
	expectedIsTollFree := "true"
	isTollFreeSlice = append(isTollFreeSlice, expectedIsTollFree)
	citySlice := make([]string, 0)
	expectedCity := "Prilep"
	citySlice = append(citySlice, expectedCity)
	provincePostalCodeSlice := make([]string, 0)
	expectedProvincePostalCode := "2345"
	provincePostalCodeSlice = append(provincePostalCodeSlice, expectedProvincePostalCode)
	countryPostalCodeSlice := make([]string, 0)
	expectedCountryPostalCode := "23453"
	countryPostalCodeSlice = append(countryPostalCodeSlice, expectedCountryPostalCode)

	sortCountryCode := "asc"
	sortNpa := "asc"
	sortNxx := "desc"
	sortIsTollFree := "asc"
	sortCity := "desc"
	sortProvincePostalCode := "asc"
	sortCountryPostalCode := "desc"

	err, response := createCliListAvailablePhoneNumberRegions(listAvailablePhoneNumberRegions, countryCodeSlice, npaSlice, nxxSlice, isTollFreeSlice, citySlice, provincePostalCodeSlice, countryPostalCodeSlice, sortCountryCode, sortNpa, sortNxx, sortIsTollFree, sortCity, sortProvincePostalCode, sortCountryPostalCode)
	assert.NoError(t, err)

	filters := getFilters(response)
	countryCode := filters["country_code"].(string)
	assert.Equal(t, expectedCountryCode, countryCode)
	npa := filters["npa"].(string)
	assert.Equal(t, expectedNpa, npa)
	nxx := filters["nxx"].(string)
	assert.Equal(t, expectedNxx, nxx)
	isTollFree := filters["is_toll_free"].(string)
	assert.Equal(t, expectedIsTollFree, isTollFree)
	city := filters["city"].(string)
	assert.Equal(t, expectedCity, city)
	provincePostalCode := filters["province_postal_code"].(string)
	assert.Equal(t, expectedProvincePostalCode, provincePostalCode)
	countryPostalCode := filters["country_postal_code"].(string)
	assert.Equal(t, expectedCountryPostalCode, countryPostalCode)

	sorts := getSorts(response)
	assert.Equal(t, sortCountryCode, sorts["country_code"])
	assert.Equal(t, sortNpa, sorts["npa"])
	assert.Equal(t, sortNxx, sorts["nxx"])
	assert.Equal(t, sortIsTollFree, sorts["is_toll_free"])
	assert.Equal(t, sortCity, sorts["city"])
	assert.Equal(t, sortProvincePostalCode, sorts["province_postal_code"])
	assert.Equal(t, sortCountryPostalCode, sorts["country_postal_code"])
}

func TestSortListAvailablePhoneNumberRegions(t *testing.T) {

	sortCountryCode := "asc"
	sortNpa := "desc"
	sortNxx := "desc"
	sortIsTollFree := "asc"
	sortCity := "desc"
	sortProvincePostalCode := "asc"
	sortCountryPostalCode := "desc"

	err, response := createCliSortAvailablePhoneNumberRegions(listAvailablePhoneNumberRegions, sortCountryCode, sortNpa, sortNxx, sortIsTollFree, sortCity, sortProvincePostalCode, sortCountryPostalCode)
	assert.NoError(t, err)

	sorts := getSorts(response)
	assert.Equal(t, sortCountryCode, sorts["country_code"])
	assert.Equal(t, sortNpa, sorts["npa"])
	assert.Equal(t, sortNxx, sorts["nxx"])
	assert.Equal(t, sortIsTollFree, sorts["is_toll_free"])
	assert.Equal(t, sortCity, sorts["city"])
	assert.Equal(t, sortProvincePostalCode, sorts["province_postal_code"])
	assert.Equal(t, sortCountryPostalCode, sorts["country_postal_code"])
}
