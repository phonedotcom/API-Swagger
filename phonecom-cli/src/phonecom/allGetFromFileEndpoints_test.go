package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var testJsonCreator SampleJsonCreator

const testJsonType = "json"
const testJsonExtension = "." + testJsonType

func TestGetAccount(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listAccounts)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	accountSample := GetAccountJson{int32(firstId)}

	filePath := "../../getAccount.json"
	testJsonCreator.marshalInput(accountSample, filePath, testJsonType)
	err, _ = createCliWithFile(getAccount, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetApplication(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listApplications)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	applicationSample := GetApplicationJson{int32(firstId)}

	filePath := "../../getApplication"
	testJsonCreator.marshalInput(applicationSample, filePath, testJsonType)
	err, _ = createCliWithFile(getApplication, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetCallLog(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listCallLogs)
	assert.NoError(t, err)

	firstId := getFirstIdString(json)
	callLogSample := GetCallLogJson{firstId}

	filePath := "../../getCallLog"
	testJsonCreator.marshalInput(callLogSample, filePath, testJsonType)
	err, _ = createCliWithFile(getCallLog, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetDevice(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listDevices)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	deviceSample := GetDeviceJson{int32(firstId)}

	filePath := "../../getDevice"
	testJsonCreator.marshalInput(deviceSample, filePath, testJsonType)
	err, _ = createCliWithFile(getDevice, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetExpressServiceCode(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listExpressServiceCodes)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	getExpressServiceCodeSample := GetExpressServiceCodeJson{int32(firstId)}

	filePath := "../../getExpressServiceCode"
	testJsonCreator.marshalInput(getExpressServiceCodeSample, filePath, testJsonType)
	err, _ = createCliWithFile(getExpressServiceCode, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetExtension(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listExtensions)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	extensionSample := GetExtensionJson{int32(firstId)}

	filePath := "../../getExtension"
	testJsonCreator.marshalInput(extensionSample, filePath, testJsonType)
	err, _ = createCliWithFile(getExtension, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetContact(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listExtensions)
	assert.NoError(t, err)
	extensionId := getFirstId(json)

	err, json = createCliWithId(listContacts, extensionId)
	assert.NoError(t, err)
	contactId := getFirstId(json)

	contactSample := GetContactJson{int32(extensionId), int32(contactId)}

	filePath := "../../getContact"
	testJsonCreator.marshalInput(contactSample, filePath, testJsonType)
	err, _ = createCliWithFile(getContact, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetGroup(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listExtensions)
	assert.NoError(t, err)
	extensionId := getFirstId(json)

	err, json = createCliWithId(listGroups, extensionId)
	assert.NoError(t, err)

	groupId := getFirstId(json)
	groupSample := GetGroupJson{int32(extensionId), int32(groupId)}

	filePath := "../../getGroup"
	testJsonCreator.marshalInput(groupSample, filePath, testJsonType)
	err, _ = createCliWithFile(getGroup, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetMedia(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listMedia)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	mediaSample := GetMediaJson{int32(firstId)}

	filePath := "../../getMedia"
	testJsonCreator.marshalInput(mediaSample, filePath, testJsonType)
	err, _ = createCliWithFile(getMedia, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetMenu(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listMenus)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	menuSample := GetMenuJson{int32(firstId)}

	filePath := "../../getMenu"
	testJsonCreator.marshalInput(menuSample, filePath, testJsonType)
	err, _ = createCliWithFile(getMenu, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetPhoneNumber(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listPhoneNumbers)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	phoneNumberSample := GetPhoneNumberJson{int32(firstId)}

	filePath := "../../getPhoneNumber"
	testJsonCreator.marshalInput(phoneNumberSample, filePath, testJsonType)
	err, _ = createCliWithFile(getPhoneNumber, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetQueue(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listQueues)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	queueSample := GetQueueJson{int32(firstId)}

	filePath := "../../getQueue"
	testJsonCreator.marshalInput(queueSample, filePath, testJsonType)
	err, _ = createCliWithFile(getQueue, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetRoute(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listRoutes)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	routeSample := GetRouteJson{int32(firstId)}

	filePath := "../../getRoute"
	testJsonCreator.marshalInput(routeSample, filePath, testJsonType)
	err, _ = createCliWithFile(getRoute, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetSchedule(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listSchedules)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	scheduleSample := GetScheduleJson{int32(firstId)}

	filePath := "../../getSchedule"
	testJsonCreator.marshalInput(scheduleSample, filePath, testJsonType)
	err, _ = createCliWithFile(getSchedule, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetSms(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listSms)
	assert.NoError(t, err)

	firstId := getFirstIdString(json)
	smsSample := GetSmsJson{firstId}

	filePath := "../../getSms"
	testJsonCreator.marshalInput(smsSample, filePath, testJsonType)
	err, _ = createCliWithFile(getSms, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}

func TestGetTrunk(t *testing.T) {

	var json map[string]interface{}
	var err error

	err, json = createCli(listTrunks)
	assert.NoError(t, err)

	firstId := getFirstId(json)
	trunkSample := GetTrunkJson{int32(firstId)}

	filePath := "../../getTrunk"
	testJsonCreator.marshalInput(trunkSample, filePath, testJsonType)
	err, _ = createCliWithFile(getTrunk, filePath+testJsonExtension)

	assert.NoError(t, err)
	os.Remove(filePath + testJsonExtension)
}
