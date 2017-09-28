package main

import (
	"encoding/json"
	"github.com/phonedotcom/API-SDK-go"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateCall(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomName := randomString(12)
	callerPhoneNumber := "+12105912768"
	callerExtension := 1750618
	callerCallerId := "+12105988591"
	callerPrivate := "1"
	calleePhoneNumber := "+12105988591"
	calleeExtension := 1750618
	calleeCallerId := "+12105912768"
	calleePrivate := "1"
	CallParamsJson := swagger.CreateCallParams{callerPhoneNumber, int32(callerExtension), callerCallerId, callerPrivate, calleePhoneNumber, int32(calleeExtension), calleeCallerId, calleePrivate}
	fileName := "../test/jsonin/createCall" + randomName + ".json"
	b, err := json.Marshal(CallParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createCall, fileName)
	assert.NoError(t, err)

	id := getCallId(result)

	if id == "" {
		t.FailNow()
	}

	defer os.Remove(fileName)
}

func TestCreateDeleteDevice(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomName := randomString(12)
	DeviceParamsJson := CreateDeviceJson{randomName}
	fileName := "../test/jsonin/createDevice" + randomName + ".json"
	b, err := json.Marshal(DeviceParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createDevice, fileName)
	assert.NoError(t, err)

	id := getId(result)

	if id == 0 {
		t.FailNow()
	}

	createCliWithId(deleteDevice, id)
	defer os.Remove(fileName)
}

func TestListReplaceDevice(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listDevices)
	assert.NoError(t, err)

	firstId := getFirstId(result)

	assert.NoError(t, err)

	randomName := randomString(12)
	DeviceParamsJson := ReplaceDeviceJson{int32(firstId), randomName}
	fileName := "../test/jsonin/replaceDevice" + randomName + ".json"
	b, err := json.Marshal(DeviceParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	t.Log(firstId)
	err, result = createCliWithJsonInAndId(replaceDevice, fileName, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestCreateDeleteListener(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomName := randomString(12)
	callbacksSlice := make([]swagger.CallbackObject, 0)
	CallbackJsonObject := swagger.CallbackObject{"https://www.phone.com", "main", "POST", "123", "************", 10}
	callbacksSlice = append(callbacksSlice, CallbackJsonObject)
	ListenerParamsJson := swagger.CreateListenerParams{"callback", "call.new", callbacksSlice}
	fileName := "../test/jsonin/createListener" + randomName + ".json"
	b, err := json.Marshal(ListenerParamsJson )
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createListener, fileName)
	assert.NoError(t, err)

	id := getId(result)

	if id == 0 {
		t.FailNow()
	}

	createCliWithId(deleteListener, id)
	defer os.Remove(fileName)
}

func TestListReplaceListener(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listListeners)
	assert.NoError(t, err)

	firstId := getFirstId(result)

	assert.NoError(t, err)

	randomName := randomString(12)
	callbacksSlice := make([]swagger.CallbackObject, 0)
	CallbackJsonObject := swagger.CallbackObject{"https://www.phone.com", "main", "GET", "123", "************", 9}
	callbacksSlice = append(callbacksSlice, CallbackJsonObject)
	ListenerParamsJson := swagger.CreateListenerParams{"callback", "call.update", callbacksSlice}
	fileName := "../test/jsonin/replaceListener" + randomName + ".json"
	b, err := json.Marshal(ListenerParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	t.Log(firstId)
	err, result = createCliWithJsonInAndId(replaceListener, fileName, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestCreateExtension(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomName := randomString(12)
	randomNum := randomNumber(10, 9999999)
	stringEmailSlice := make([]string, 0)
	stringEmailSlice = append(stringEmailSlice, "asd@asd.com")

	ExtensionParamsJson := swagger.CreateExtensionParams{swagger.VoicemailInput{"0", "12345", swagger.GreetingInput{"alternate", swagger.MediaSummary{}, swagger.MediaSummary{}, "0"}, "wav", swagger.Notification{stringEmailSlice, "+18189640644"}, "human"}, swagger.CallNotifications{stringEmailSlice, "+12546551378"}, "+12019570328", "unlimited", int32(randomNum), "1", "The name", "The full name", "America/Los_Angeles", nil, "619", "1", "1"}

	fileName := "../test/jsonin/createExtension" + randomName + ".json"
	b, err := json.Marshal(ExtensionParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createExtension, fileName)
	assert.NoError(t, err)

	id := getId(result)
	if id == 0 {
		t.FailNow()
	}
	os.Remove(fileName)
}

func TestListReplaceExtension(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listExtensions)
	assert.NoError(t, err)

	firstId := getFirstId(result)

	randomName := randomString(12)
	stringEmailSlice := make([]string, 0)
	stringEmailSlice = append(stringEmailSlice, "asd@asd.com")
	ExtensionParamsJson := swagger.ReplaceExtensionParams{swagger.VoicemailInput{"0", "12345", swagger.GreetingInput{"alternate", swagger.MediaSummary{}, swagger.MediaSummary{}, "0"}, "wav", swagger.Notification{stringEmailSlice, "+18189640644"}, "human"}, swagger.CallNotifications{stringEmailSlice, "+12546551378"}, nil, randomName, "America/Los_Angeles", "1", 111, "1", "unlimited", "bobby McFerrin", "1", "private", "619", ""}
	fileName := "../test/jsonin/replaceExtension" + randomName + ".json"
	b, err := json.Marshal(ExtensionParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(replaceExtension, fileName, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestCreateDeleteContact(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listExtensions)
	assert.NoError(t, err)

	extensionId := getFirstId(result)

	err, result = createCli(listAvailablePhoneNumbers)
	assert.NoError(t, err)

	availablePhoneNumbers := getFirstNAvailablePhoneNumbers(result, 5)

	randomName := randomString(12)
	emails := make([]swagger.Email, 1)
	emails[0] = swagger.Email{"primary", randomName + "@asd.com"}

	addresses := make([]swagger.AddressListContacts, 1)
	addresses[0] = swagger.AddressListContacts{"home", randomName + " street and number", "The City", "The State", "12345", "The Country"}

	ContactParamsJson := swagger.CreateContactParams{"Geordi", "middle name", "last name", "prefix", "phoneticFirstName", "phoneticMiddleName", "phoneticLastName", "suffix", "nickname", "company", "department", "jobTitle", emails, availablePhoneNumbers, addresses, nil}
	fileName := "../test/jsonin/createContact" + randomName + ".json"
	b, err := json.Marshal(ContactParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(createContact, fileName, extensionId)
	assert.NoError(t, err)

	id := getId(result)
	if id == 0 {
		t.FailNow()
	}

	createCliWithIdIdSecondary(deleteContact, extensionId, id)
	os.Remove(fileName)
}

func TestCreateDeleteContactSixPhoneNumbers(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listExtensions)
	assert.NoError(t, err)

	extensionId := getFirstId(result)

	err, result = createCli(listAvailablePhoneNumbers)
	assert.NoError(t, err)

	availablePhoneNumbers := getFirstNAvailablePhoneNumbers(result, 6)

	randomName := randomString(12)
	emails := make([]swagger.Email, 1)
	emails[0] = swagger.Email{"primary", randomName + "@asd.com"}

	addresses := make([]swagger.AddressListContacts, 1)
	addresses[0] = swagger.AddressListContacts{"home", randomName + " street and number", "The City", "The State", "12345", "The Country"}

	ContactParamsJson := swagger.CreateContactParams{"Geordi", "middle name", "last name", "prefix", "phoneticFirstName", "phoneticMiddleName", "phoneticLastName", "suffix", "nickname", "company", "department", "jobTitle", emails, availablePhoneNumbers, addresses, nil}
	fileName := "../test/jsonin/createContact" + randomName + ".json"
	b, err := json.Marshal(ContactParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(createContact, fileName, extensionId	)
	assert.Error(t, err)

	os.Remove(fileName)
}

func TestListReplaceContact(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listExtensions)
	assert.NoError(t, err)

	extensionId := getFirstId(result)

	err, result = createCliWithId(listContacts, extensionId)
	assert.NoError(t, err)

	contactId := getFirstId(result)

	randomName := randomString(12)
	ContactParamsJson := swagger.CreateContactParams{randomName, "middle name", "last name", "prefix", "phoneticFirstName", "phoneticMiddleName", "phoneticLastName", "suffix", "nickname", "company", "department", "jobTitle", nil, nil, nil, nil}
	fileName := "../test/jsonin/replaceContact" + randomName + ".json"
	b, err := json.MarshalIndent(ContactParamsJson, "", "  ")
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createReplaceContactCliWithJsonIn(replaceContact, fileName, extensionId, contactId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestCreateDeleteGroup(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listExtensions)
	assert.NoError(t, err)

	extensionId := getFirstId(result)

	randomName := randomString(12)
	fileName := "../test/jsonin/createGroup" + randomName + ".json"
	GroupParamsJson := swagger.CreateGroupParams{"Ferengi Traders"}
	b, err := json.Marshal(GroupParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(createGroup, fileName, extensionId)
	assert.NoError(t, err)

	id := getId(result)
	if id == 0 {
		t.FailNow()
	}

	createCliWithIdIdSecondary(deleteGroup, extensionId, id)
	os.Remove(fileName)
}

func TestListReplaceGroup(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listExtensions)
	assert.NoError(t, err)

	extensionId := getFirstId(result)

	err, result = createCliWithId(listGroups, extensionId)
	assert.NoError(t, err)

	groupId := getFirstId(result)

	randomName := randomString(12)
	groupParamsJson := swagger.CreateGroupParams{randomName}
	fileName := "../test/jsonin/replaceGroup" + randomName + ".json"
	b, err := json.Marshal(groupParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createReplaceContactCliWithJsonIn(replaceGroup, fileName, extensionId, groupId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestCreateDeleteMenu(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomName := randomString(12)
	MenuParamsJson := CreateMenuJson{randomName, "1", 3}
	fileName := "../test/jsonin/createMenu" + randomName + ".json"
	b, err := json.Marshal(MenuParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createMenu, fileName)
	assert.NoError(t, err)

	id := getId(result)

	if id <= 0 {
		t.Fatal()
	}

	createCliWithId(deleteMenu, id)
	os.Remove(fileName)
}

func TestCreateDeleteOauthClientRedirectUri(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listOauthClients)
	assert.NoError(t, err)
	clientId := getFirstId(result)

	randomName := randomString(12)
	RedirectUriParamsJson := swagger.CreateRedirectUriParams{"https://www.google.com/redirect" + randomName}
	fileName := "../test/jsonin/createRedirectUri" + randomName + ".json"
	b, err := json.Marshal(RedirectUriParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(createOauthClientRedirectUri, fileName, clientId)
	assert.NoError(t, err)

	uriId := getId(result)

	if uriId <= 0 {
		t.Fatal()
	}

	createCliWithIdIdSecondary(deleteOauthClientRedirectUri, clientId, uriId)
	os.Remove(fileName)
}

func TestCreateDeletePaymentMethod(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomName := randomString(12)
	MenuParamsJson := swagger.CreatePaymentParams{randomName, "primary", "cc", "52d937f179874afe88531e12345f2be"}
	fileName := "../test/jsonin/createPayment" + randomName + ".json"
	b, err := json.Marshal(MenuParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createPaymentMethod, fileName)
	assert.NoError(t, err)

	id := getId(result)

	if id <= 0 {
		t.Fatal()
	}

	createCliWithId(deletePaymentMethod, id)
}

func TestListPatchPaymentMethod(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listPaymentMethods)
	assert.NoError(t, err)

	firstId := getFirstId(result)

	randomName := randomString(12)
	PaymentPatchJson := swagger.PatchPaymentParams{"primary"}
	fileName := "../test/jsonin/replacePayment" + randomName + ".json"
	b, err := json.Marshal(PaymentPatchJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(patchPaymentMethod, fileName, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestCreateDeleteMediaTts(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomName := randomString(12)
	MediaParamsJson := swagger.CreateMediaParams{randomName, "tts", "hold_music", "allison", randomString(100), "Y", 900, 100, "Notes aboute the media", "N"}
	fileName := "../test/jsonin/createMedia" + randomName + ".json"
	b, err := json.Marshal(MediaParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createMediaTts, fileName)
	assert.NoError(t, err)

	id := getId(result)

	if id <= 0 {
		t.Fatal()
	}

	createCliWithId(deleteMedia, id)
	os.Remove(fileName)
}

func TestListReplaceMenu(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listMenus)
	assert.NoError(t, err)

	firstId := getFirstId(result)

	randomName := randomString(12)
	MenuParamsJson := ReplaceMenuJson{int32(firstId), randomName, "0", 5}
	fileName := "../test/jsonin/replaceMenu" + randomName + ".json"
	b, err := json.Marshal(MenuParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(replaceMenu, fileName, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestListReplaceMedia(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listMedia)
	assert.NoError(t, err)

	firstId := getFirstId(result)

	randomName := randomString(12)
	MediaParamsJson := swagger.CreateMediaParams{randomName, "tts", "hold_music", "allison", randomString(100), "Y", 900, 100, "Notes aboute the media", "N"}
	fileName := "../test/jsonin/replaceMedia" + randomName + ".json"
	b, err := json.Marshal(MediaParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(replaceMediaTts, fileName, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestCreatePhoneNumber(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listExtensions)
	assert.NoError(t, err)

	firstExtensionId := getFirstId(result)

	err, result = createCli(listAvailablePhoneNumbers)
	assert.NoError(t, err)

	firstId := getFirstAvailablePhoneNumber(result)

	randomName := randomString(12)
	stringEmailSlice := make([]string, 0)
	stringEmailSlice = append(stringEmailSlice, "asd@asd.com")
	PhoneNumberParamsJson := swagger.CreatePhoneNumberParams{firstId, "", "Phone Name Now", "1", "1", swagger.CallerIdPhoneNumber{"nameOfOwner", "business"}, swagger.SmsForwardingParams{"extension", int32(firstExtensionId), 0}, swagger.CallNotifications{stringEmailSlice, "+12546551378"}}
	fileName := "../test/jsonin/createPhoneNumber" + randomName + ".json"
	b, err := json.Marshal(PhoneNumberParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createPhoneNumber, fileName)
	assert.NoError(t, err)

	id := getId(result)
	if id == 0 {
		t.FailNow()
	}
	os.Remove(fileName)
}

func TestListReplacePhoneNumber(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listExtensions)
	assert.NoError(t, err)

	firstExtensionId := getFirstId(result)

	err, result = createCli(listPhoneNumbers)
	assert.NoError(t, err)

	firstId := getFirstId(result)

	randomName := randomString(12)
	stringEmailSlice := make([]string, 0)
	stringEmailSlice = append(stringEmailSlice, "asd@asd.com")
	PhoneNumberParamsJson := swagger.ReplacePhoneNumberParams{"", "Robert", "0", "0", swagger.CallerIdPhoneNumber{"nameOfOwner", "business"}, swagger.SmsForwardingParams{"extension", int32(firstExtensionId), 0}, nil, swagger.CallNotifications{stringEmailSlice, "+12546551378"}}
	fileName := "../test/jsonin/replacePhoneNumber" + randomName + ".json"
	b, err := json.Marshal(PhoneNumberParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(replacePhoneNumber, fileName, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestCreateDeleteQueue(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomName := randomString(12)
	QueueParamsJson := CreateQueueJson{randomName, 60, "called_number", 10}
	fileName := "../test/jsonin/createQueue" + randomName + ".json"
	b, err := json.Marshal(QueueParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createQueue, fileName)
	assert.NoError(t, err)

	id := getId(result)
	if id == 0 {
		t.FailNow()
	}

	createCliWithId(deleteQueue, id)
	os.Remove(fileName)
}

func TestListReplaceQueue(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listQueues)
	assert.NoError(t, err)

	firstId := getFirstId(result)

	randomName := randomString(12)
	QueueParamsJson := CreateQueueJson{randomName, 60, "called_number", 10}
	fileName := "../test/jsonin/replaceQueue" + randomName + ".json"
	b, err := json.Marshal(QueueParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(replaceQueue, fileName, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestCreateDeleteRoute(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomNameQ := randomString(12)
	QueueParamsJson := CreateQueueJson{randomNameQ, 60, "called_number", 10}
	fileNameQ := "../test/jsonin/createQueue" + randomNameQ + ".json"
	bQ, errQ := json.Marshal(QueueParamsJson)
	err = ioutil.WriteFile(fileNameQ, bQ, 0644)

	errQ, result = createCliWithJsonIn(createQueue, fileNameQ)
	assert.NoError(t, errQ)

	idQ := getId(result)
	nameQ := getName(result)
	if idQ == 0 {
		t.FailNow()
	}

	queueParam := swagger.QueueSummary{int32(idQ), nameQ}
	actionParam := ActionsJson{"queue", queueParam}
	ruleJson := RulesJson{[]ActionsJson{actionParam}}
	rulesParam := []RulesJson{ruleJson}
	randomName := randomString(12)
	RouteParamsJson := CreateRouteJson{randomName, rulesParam}
	fileName := "../test/jsonin/createRoute" + randomName + ".json"
	b, err := json.Marshal(RouteParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createRoute, fileName)
	assert.NoError(t, err)

	id := getId(result)
	if id == 0 {
		t.FailNow()
	}

	createCliWithId(deleteRoute, id)
	os.Remove(fileName)

	createCliWithId(deleteQueue, idQ)
	os.Remove(fileNameQ)
}

func TestListReplaceRoute(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listRoutes)
	assert.NoError(t, err)

	firstId := getFirstId(result)

	randomNameQ := randomString(12)
	QueueParamsJson := CreateQueueJson{randomNameQ, 60, "called_number", 10}
	fileNameQ := "../test/jsonin/createQueue" + randomNameQ + ".json"
	bQ, errQ := json.Marshal(QueueParamsJson)
	err = ioutil.WriteFile(fileNameQ, bQ, 0644)

	errQ, result = createCliWithJsonIn(createQueue, fileNameQ)
	assert.NoError(t, errQ)

	idQ := getId(result)
	nameQ := getName(result)
	if idQ == 0 {
		t.FailNow()
	}

	queueParam := swagger.QueueSummary{int32(idQ), nameQ}
	actionParam := ActionsJson{"queue", queueParam}
	ruleJson := RulesJson{[]ActionsJson{actionParam}}
	rulesParam := []RulesJson{ruleJson}
	randomName := randomString(12)
	RouteParamsJson := CreateRouteJson{randomName, rulesParam}
	fileName := "../test/jsonin/replaceRoute" + randomName + ".json"
	b, err := json.Marshal(RouteParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(replaceRoute, fileName, firstId)
	assert.NoError(t, err)

	os.Remove(fileNameQ)
	os.Remove(fileName)
}

func TestCreateSms(t *testing.T) {

	var res map[string]interface{}
	var err error

	err, res = createCli(listExtensions)
	assert.NoError(t, err)
	extensionId := getFirstId(res)

	from := "+12105988591"
	to := "+16096140970"
	text := "Another message for create"
	randomName := randomString(12)
	SmsParamsJson := swagger.CreateSmsParams{from, to, text, int32(extensionId)}
	fileName := "../test/jsonin/createSms" + randomName + ".json"
	b, err := json.Marshal(SmsParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result := createCreateSmsCliWithJsonIn(createSms, fileName, from, to, text)
	if result == nil {
		t.Log("Create message response is nil")
	}

	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestListPatchSms(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listSms)
	assert.NoError(t, err)

	firstId := getFirstIdString(result)

	randomName := randomString(12)
	SmsPatchJson := swagger.PatchSmsParams{"1"}
	fileName := "../test/jsonin/patchSms" + randomName + ".json"
	b, err := json.Marshal(SmsPatchJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createReplaceCliWithJsonInString(patchSms, fileName, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestCreateSubaccount(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomName := randomString(12)
	randomPassword := randomString(12)
	AddressObject1 := AddressJson{"100 Main St", "San Diego", "CA", "92129", "US"}
	contactObject := ContactJson{"Bobby", AddressObject1, "+18585553333", "asd@sd.co"}
	AddressObject2 := AddressJson{"100 Main St", "San Diego", "CA", "92129", "US"}
	billingContactObject := ContactJson{"Bobby", AddressObject2, "+18585553333", "asd@sd.co"}
	SubaccountParamsJson := CreateSubaccountJson{randomName, randomPassword, contactObject, billingContactObject}
	fileName := "../test/jsonin/createSubaccount" + randomName + ".json"
	b, err := json.Marshal(SubaccountParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonIn(createSubaccount, fileName)
	assert.NoError(t, err)

	id := getId(result)
	if id == 0 {
		t.FailNow()
	}
	os.Remove(fileName)
}

func TestCreateDeletePricing(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listSubaccounts)
	assert.NoError(t, err)
	subaccountId := getFirstId(result)

	randomName := randomString(12)
	PricingParamsJson := swagger.CreatePricingParams{1024, "One month free trial", 1505643813}
	fileName := "../test/jsonin/createPricing" + randomName + ".json"
	b, err := json.Marshal(PricingParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCliWithJsonInAndId(createPricing, fileName, int(subaccountId))
	assert.NoError(t, err)

	id := getId(result)
	if id == 0 {
		t.FailNow()
	}

	createCliWithId(deletePricing, id)
	os.Remove(fileName)
}

func TestCreateDeleteTrunk(t *testing.T) {

	var result map[string]interface{}
	var err error

	randomName := randomString(12)
	trunkName := randomName
	trunkUri := "SIP/1234@phone.com:5060"
	trunkConcurrentCalls := 60
	trunkMaxMinutes := 800
	TrunkParamsJson := CreateTrunkJson{trunkName, trunkUri, int32(trunkConcurrentCalls), int32(trunkMaxMinutes)}
	fileName := "../test/jsonin/createTrunk" + randomName + ".json"
	b, err := json.Marshal(TrunkParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createCreateTrunkCliWithJsonIn(createTrunk, fileName, trunkName, trunkUri, trunkConcurrentCalls, trunkMaxMinutes)
	assert.NoError(t, err)

	id := getId(result)
	if id == 0 {
		t.FailNow()
	}

	createCliWithId(deleteTrunk, id)
	os.Remove(fileName)
}

func TestListReplaceTrunk(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listTrunks)
	assert.NoError(t, err)

	firstId := getFirstId(result)

	randomName := randomString(12)
	trunkName := randomName
	trunkUri := "SIP/1234@phone.com:5060"
	trunkConcurrentCalls := 80
	trunkMaxMinutes := 800
	TrunkParamsJson := CreateTrunkJson{trunkName, trunkUri, int32(trunkConcurrentCalls), int32(trunkMaxMinutes)}
	fileName := "../test/jsonin/createTrunk" + randomName + ".json"
	b, err := json.Marshal(TrunkParamsJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createReplaceTrunkCliWithJsonIn(replaceTrunk, fileName, trunkName, trunkUri, trunkConcurrentCalls, trunkMaxMinutes, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}

func TestListPatchVoicemail(t *testing.T) {

	var result map[string]interface{}
	var err error

	err, result = createCli(listVoicemail)
	assert.NoError(t, err)

	firstId := getFirstIdString(result)

	randomName := randomString(12)
	VoicemailPatchJson := swagger.PatchVoicemailParams{"1"}
	fileName := "../test/jsonin/patchVoicemail" + randomName + ".json"
	b, err := json.Marshal(VoicemailPatchJson)
	err = ioutil.WriteFile(fileName, b, 0644)

	err, result = createReplaceCliWithJsonInString(patchVoicemail, fileName, firstId)
	assert.NoError(t, err)
	os.Remove(fileName)
}
