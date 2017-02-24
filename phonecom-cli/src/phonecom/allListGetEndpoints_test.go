package main

import (
  "testing"
)

func TestListAccounts(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listAccounts)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getAccount, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListApplications(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listApplications)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getApplication, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListCallLogs(t *testing.T) {

  var err error

  err, _ = createCli(listCallLogs)
  assertErrorNotNull(t, err)
}

func TestListDevices(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listDevices)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getDevice, int(firstId))

  assertErrorNotNull(t, err)
}

func TestExpressServiceCodes(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listExpressServiceCodes)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getExpressServiceCode, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListExtensions(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listExtensions)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getExtension, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListCallerIds(t *testing.T) {

  var err error

  err, _ = createCli(getCallerId)
  assertErrorNotNull(t, err)

  assertErrorNotNull(t, err)
}

func TestListContacts(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listContacts)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getContact, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListGroups(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listGroups)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getGroup, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListMedia(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listMedia)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getRecording, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListMenus(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listMenus)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getMenu, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListPhoneNumbers(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listPhoneNumbers)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getPhoneNumber, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListQueues(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listQueues)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getQueue, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListRoutes(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listRoutes)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getRoute, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListSchedules(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listSchedules)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getSchedule, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListSms(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listSms)
  assertErrorNotNull(t, err)

  firstId := getFirstIdString(json)
  err, json = createGetCliStringId(getSms, firstId)

  assertErrorNotNull(t, err)
}

func TestListSubaccounts(t *testing.T) {

  var err error

  err, _ = createCli(listSubaccounts)
  assertErrorNotNull(t, err)

  assertErrorNotNull(t, err)
}

func TestListTrunks(t *testing.T) {

  var json map[string] interface{}
  var err error

  err, json = createCli(listTrunks)
  assertErrorNotNull(t, err)

  firstId := getFirstId(json)
  err, json = createGetOrRemoveCli(getTrunk, int(firstId))

  assertErrorNotNull(t, err)
}

func TestListAvailablePhoneNumbers(t *testing.T) {

  var err error

  err, _ = createCli(listAvailablePhoneNumbers)
  assertErrorNotNull(t, err)
}

func TestListAvailablePhoneNumberRegions(t *testing.T) {

  var err error

  err, _ = createCli(listAvailablePhoneNumberRegions)
  assertErrorNotNull(t, err)
}
