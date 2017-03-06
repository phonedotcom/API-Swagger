package main

import (
  "testing"
  "encoding/json"
  "io/ioutil"
  "os"
  "phonecom-go-sdk"
)

func TestCreateDevice(t *testing.T) {

  var result map[string] interface{}
  var err error

  randomName := randomString(12)
  DeviceParamsJson := CreateDeviceJson{randomName}
  fileName := "../test/jsonin/createDevice" + randomName + ".json"
  b, err := json.Marshal(DeviceParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createDevice, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)

  if (id == 0) {
    t.FailNow()
  }

  defer os.Remove(fileName)
}

func TestListReplaceDevice(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listDevices)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)

  assertErrorNotNull(t, err)

  randomName := randomString(12)
  DeviceParamsJson := ReplaceDeviceJson{int32(firstId), randomName}
  fileName := "../test/jsonin/replaceDevice" + randomName + ".json"
  b, err := json.Marshal(DeviceParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  t.Log(firstId)
  err, result = createReplaceCliWithJsonIn(replaceDevice, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

func TestCreateExtension(t *testing.T) {

  var result map[string] interface{}
  var err error

  randomName := randomString(12)
  randomNum := randomNumber(10, 9999999)
  ExtensionParamsJson := CreateExtensionJson{"+12019570328", "unlimited", true, int32(randomNum), true, "The name", "The full name", "America/Los_Angeles", 619, true, false, true, false, 12345, "standard", "automated", "+18587741111", "+18587748888"}
  fileName := "../test/jsonin/createExtension" + randomName + ".json"
  b, err := json.Marshal(ExtensionParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createExtension, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }
  os.Remove(fileName)
}

func TestListReplaceExtension(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listExtensions)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)

  randomName := randomString(12)
  ExtensionParamsJson := ReplaceExtensionJson{int32(firstId), randomName, "America/Los_Angeles", true, 111, true, "unlimited", 12344, "bobby McFerrin", true, "standard", "private", 619, true, true, "automated", "+18587741111", "+18587748888"}
  fileName := "../test/jsonin/replaceExtension" + randomName + ".json"
  b, err := json.Marshal(ExtensionParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceCliWithJsonIn(replaceExtension, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

func TestCreateDeleteContact(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listExtensions)
  assertErrorNotNull(t, err)

  extensionId := getFirstId(result)

  randomName := randomString(12)
  ContactParamsJson := CreateContactJson{int32(extensionId), "Geordi", "middle name", "last name", "prefix", "phoneticFirstName", "phoneticMiddleName", "phoneticLastName", "suffix", "nickname", "company", "department", "jobTitle"}
  fileName := "../test/jsonin/createContact" + randomName + ".json"
  b, err := json.Marshal(ContactParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createContact, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createCliWithIdIdSecondary(deleteContact, extensionId, id);
  os.Remove(fileName)
}

func TestListReplaceContact(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listExtensions)
  assertErrorNotNull(t, err)

  extensionId := getFirstId(result)

  err, result = createCliWithId(listContacts, extensionId)
  assertErrorNotNull(t, err)

  contactId := getFirstId(result)

  randomName := randomString(12)
  ContactParamsJson := swagger.CreateContactParams{randomName, "middle name", "last name", "prefix", "phoneticFirstName", "phoneticMiddleName", "phoneticLastName", "suffix", "nickname", "company", "department", "jobTitle", nil, nil, nil, nil}
  fileName := "../test/jsonin/replaceContact" + randomName + ".json"
  b, err := json.MarshalIndent(ContactParamsJson, "", "  ")
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceContactCliWithJsonIn(replaceContact, fileName, extensionId, contactId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

func TestCreateDeleteGroup(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listExtensions)
  assertErrorNotNull(t, err)

  extensionId := getFirstId(result)

  randomName := randomString(12)
  fileName := "../test/jsonin/createGroup" + randomName + ".json"
  GroupParamsJson := CreateGroupJson{int32(extensionId), "Ferengi Traders"}
  b, err := json.Marshal(GroupParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createGroup, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createCliWithIdIdSecondary(deleteGroup, extensionId, id);
  os.Remove(fileName)
}

func TestListReplaceGroup(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listExtensions)
  assertErrorNotNull(t, err)

  extensionId := getFirstId(result)

  err, result = createCliWithId(listGroups, extensionId)
  assertErrorNotNull(t, err)

  groupId := getFirstId(result)

  randomName := randomString(12)
  groupParamsJson := swagger.CreateGroupParams{randomName}
  fileName := "../test/jsonin/replaceGroup" + randomName + ".json"
  b, err := json.Marshal(groupParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceContactCliWithJsonIn(replaceGroup, fileName, extensionId, groupId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

func TestCreateDeleteMenu(t *testing.T) {

  var result map[string] interface{}
  var err error


  randomName := randomString(12)
  MenuParamsJson := CreateMenuJson{randomName, true, 3}
  fileName := "../test/jsonin/createMenu" + randomName + ".json"
  b, err := json.Marshal(MenuParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)


  err, result = createCliWithJsonIn(createMenu, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)

  if (id <= 0) {
    t.Fatal()
  }

  createCliWithId(deleteMenu, id);
  os.Remove(fileName)
}

func TestListReplaceMenu(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listMenus)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)

  randomName := randomString(12)
  MenuParamsJson := ReplaceMenuJson{int32(firstId), randomName, false, 5}
  fileName := "../test/jsonin/replaceMenu" + randomName + ".json"
  b, err := json.Marshal(MenuParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceCliWithJsonIn(replaceMenu, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

func TestCreatePhoneNumber(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listAvailablePhoneNumbers)
  assertErrorNotNull(t, err)

  firstId := getFirstAvailablePhoneNumber(result)

  randomName := randomString(12)
  PhoneNumberParamsJson := CreatePhoneNumberJson{firstId, "Phone Name Now", true, true, "Phone N", "business", "extension", "+18587740222"}
  fileName := "../test/jsonin/createPhoneNumber" + randomName + ".json"
  b, err := json.Marshal(PhoneNumberParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createPhoneNumber, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }
  os.Remove(fileName)
}

func TestListReplacePhoneNumber(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listPhoneNumbers)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)

  randomName := randomString(12)
  PhoneNumberParamsJson := ReplacePhoneNumberJson{int32(firstId), "Robert", true, true, "Phone N", "business", "extension", "+18587740222"}
  fileName := "../test/jsonin/replacePhoneNumber" + randomName + ".json"
  b, err := json.Marshal(PhoneNumberParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceCliWithJsonIn(replacePhoneNumber, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

func TestCreateDeleteQueue(t *testing.T) {

  var result map[string] interface{}
  var err error

  randomName := randomString(12)
  QueueParamsJson := CreateQueueJson{randomName, 60, "called_number", 10}
  fileName := "../test/jsonin/createQueue" + randomName + ".json"
  b, err := json.Marshal(QueueParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createQueue, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createCliWithId(deleteQueue, id);
  os.Remove(fileName)
}

func TestListReplaceQueue(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listQueues)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)

  randomName := randomString(12)
  QueueParamsJson := ReplaceQueueJson{int32(firstId), randomName, 60, "called_number", 10}
  fileName := "../test/jsonin/replaceQueue" + randomName + ".json"
  b, err := json.Marshal(QueueParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceCliWithJsonIn(replaceQueue, fileName, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

func TestCreateDeleteRoute(t *testing.T) {

  var result map[string] interface{}
  var err error

  randomNameQ := randomString(12)
  QueueParamsJson := CreateQueueJson{randomNameQ, 60, "called_number", 10}
  fileNameQ := "../test/jsonin/createQueue" + randomNameQ + ".json"
  bQ, errQ := json.Marshal(QueueParamsJson)
  err = ioutil.WriteFile(fileNameQ, bQ, 0644)

  errQ, result = createCliWithJsonIn(createQueue, fileNameQ)
  assertErrorNotNull(t, errQ)

  idQ := getId(result)
  nameQ := getName(result)
  if (idQ == 0) {
    t.FailNow()
  }

  queueParam := QueueJson{int32(idQ), nameQ}
  actionParam := ActionsJson{"queue", queueParam}
  ruleJson := RulesJson{[]ActionsJson{actionParam}}
  rulesParam := []RulesJson{ruleJson}
  randomName := randomString(12)
  RouteParamsJson := CreateRouteJson{randomName, rulesParam}
  fileName := "../test/jsonin/createRoute" + randomName + ".json"
  b, err := json.Marshal(RouteParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createCliWithJsonIn(createRoute, fileName)
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createCliWithId(deleteRoute, id);
  os.Remove(fileName)

  createCliWithId(deleteQueue, idQ);
  os.Remove(fileNameQ)
}

func TestListReplaceRoute(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listRoutes)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)

  randomNameQ := randomString(12)
  QueueParamsJson := CreateQueueJson{randomNameQ, 60, "called_number", 10}
  fileNameQ := "../test/jsonin/createQueue" + randomNameQ + ".json"
  bQ, errQ := json.Marshal(QueueParamsJson)
  err = ioutil.WriteFile(fileNameQ, bQ, 0644)

  errQ, result = createCliWithJsonIn(createQueue, fileNameQ)
  assertErrorNotNull(t, errQ)

  idQ := getId(result)
  nameQ := getName(result)
  if (idQ == 0) {
    t.FailNow()
  }

  queueParam := QueueJson{int32(idQ), nameQ}
  actionParam := ActionsJson{"queue", queueParam}
  ruleJson := RulesJson{[]ActionsJson{actionParam}}
  rulesParam := []RulesJson{ruleJson}
  randomName := randomString(12)
  RouteParamsJson := ReplaceRouteJson{int32(firstId), randomName, rulesParam}
  fileName := "../test/jsonin/replaceRoute" + randomName + ".json"
  b, err := json.Marshal(RouteParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceCliWithJsonIn(replaceRoute, fileName, firstId)
  assertErrorNotNull(t, err)

  os.Remove(fileNameQ)
  os.Remove(fileName)
}

func TestCreateSms(t *testing.T) {

  var err error

  from := "+16309624775"
  to := "+12019570328"
  text := "Another message for create"
  randomName := randomString(12)
  SmsParamsJson := CreateSmsJson{from, to, text, 1767963}
  fileName := "../test/jsonin/createSms" + randomName + ".json"
  b, err := json.Marshal(SmsParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result := createCreateSmsCliWithJsonIn(createSms, fileName, from, to, text)
  if (result == nil) {
    t.Log("Create message response is nil")
  }

  assertErrorNotNull(t, err)
  os.Remove(fileName)
}

func TestCreateSubaccount(t *testing.T) {

  var result map[string] interface{}
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
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }
  os.Remove(fileName)
}

func TestCreateDeleteTrunk(t *testing.T) {

  var result map[string] interface{}
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
  assertErrorNotNull(t, err)

  id := getId(result)
  if (id == 0) {
    t.FailNow()
  }

  createCliWithId(deleteTrunk, id);
  os.Remove(fileName)
}

func TestListReplaceTrunk(t *testing.T) {

  var result map[string] interface{}
  var err error

  err, result = createCli(listTrunks)
  assertErrorNotNull(t, err)

  firstId := getFirstId(result)

  randomName := randomString(12)
  trunkName := randomName
  trunkUri := "SIP/1234@phone.com:5060"
  trunkConcurrentCalls := 80
  trunkMaxMinutes := 800
  TrunkParamsJson := ReplaceTrunkJson{int32(firstId), trunkName, trunkUri, int32(trunkConcurrentCalls), int32(trunkMaxMinutes)}
  fileName := "../test/jsonin/createTrunk" + randomName + ".json"
  b, err := json.Marshal(TrunkParamsJson)
  err = ioutil.WriteFile(fileName, b, 0644)

  err, result = createReplaceTrunkCliWithJsonIn(replaceTrunk, fileName, trunkName, trunkUri, trunkConcurrentCalls, trunkMaxMinutes, firstId)
  assertErrorNotNull(t, err)
  os.Remove(fileName)
}
