package main

import (
  "testing"
  "encoding/json"
  "time"
  "math/rand"
  "io/ioutil"
  "os"
)

func TestCreateDevice(t *testing.T) {

  var res map[string] interface{}
  var err error

  randomName := randomString(12)
  fileName := "../test/jsonin/createDevice" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\":1315091,\n\t\"name\":\""+ randomName +"\"\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)

  err, res = createCliWithJsonIn(createDevice, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(res)

  if (id == 0) {
    t.FailNow()
  }

}

func randomString(strlen int) string {

  rand.Seed(time.Now().UTC().UnixNano())
  const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
  result := make([]byte, strlen)
  for i := 0; i < strlen; i++ {
    result[i] = chars[rand.Intn(len(chars))]
  }
  return string(result)
}

type Device struct {
  accountId int `json:"account_id"`
  name string `json:"name"`
}

func (d *Device) marshalDevice() ([]byte, error) {
  return json.Marshal(&struct {
    accountId     int `json:"account_id"`
    name          string `json:"name"`}{

    accountId: d.accountId,
    name:      d.name,
  })
}

func TestListReplaceDevice(t *testing.T) {

  var res map[string] interface{}
  var err error

  randomName := randomString(12)
  fileName := "../test/jsonin/replaceDevice" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\":1315091,\n\t\"device_id\":111,\n\t\"name\":\""+ randomName +"\"\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)

  err, res = createCliWithJsonIn(createDevice, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(res)
  if (id == 0) {
    t.FailNow()
  }

}

func TestCreateExtension(t *testing.T) {

  var json map[string] interface{}
  var err error

  randomName := randomString(12)
  fileName := "../test/jsonin/createExtension" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"caller_id\": \"+12019570328\",\n\t\"usage_type\": \"unlimited\",\n\t\"allows_call_waiting\": true,\n\t\"extension\": 12340,\n\t\"include_in_directory\": true,\n\t\"name\": \"The name\",\n\t\"full_name\": \"The full name\",\n\t\"timezone\": \"America/Los_Angeles\",\n\t\"local_area_code\": 619,\n\t\"voicemail[greeting][enable_leave_message_prompt]\": true,\n\t\"voicemail[enabled]\": false,\n\t\"enable_outbound_calls\": true,\n\t\"enable_call_waiting\": false,\n\t\"voicemail[password]\": 12345,\n\t\"voicemail[greeting][type]\": \"standard\",\n\t\"voicemail[transcription]\": \"automated\",\n\t\"voicemail[notifications][sms]\": \"+18587741111\",\n\t\"call_notifications[sms]\": \"+18587748888\"\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)

  err, json = createCliWithJsonIn(createExtension, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}

func TestListReplaceExtension(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/replaceExtension" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"extension_id\": 1767963,\n\t\"name\": \"new bobby\",\n\t\"timezone\": \"America/Los_Angeles\",\n\t\"include_in_directory\": true,\n\t\"extension\": 111,\n\t\"enable_outbound_calls\": true,\n\t\"usage_type\": \"unlimited\",\n\t\"voicemail[password]\": \"000000\",\n\t\"full_name\": \"bobby McFerrin\",\n\t\"enable_call_waiting\": true,\n\t\"voicemail[greeting][type]\": \"standard\",\n\t\"caller_id\": \"private\",\n\t\"local_area_code\": 619,\n\t\"voicemail[enabled]\": true,\n\t\"voicemail[greeting][enable_leave_message_prompt]\": true,\n\t\"voicemail[transcription]\": \"automated\",\n\t\"voicemail[notifications][sms]\": \"+18587741111\",\n\t\"call_notifications[sms]\": \"+18587748888\"\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createExtension, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}

func TestCreateDeleteContact(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/createContact" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"extension_id\": 1764590,\n\t\"first_name\": \"Geordi\",\n\t\"middle_name\": \"middle name\",\n\t\"last_name\": \"last name\",\n\t\"prefix\": \"prefix\",\n\t\"phonetic_first_name\": \"phoneticFirstName\",\n\t\"phonetic_middle_name\": \"phoneticMiddleName\",\n\t\"phonetic_last_name\": \"phoneticLastName\",\n\t\"suffix\": \"suffix\",\n\t\"nickname\": \"nickname\",\n\t\"company\": \"company\",\n\t\"department\": \"department\",\n\t\"job_title\": \"jobTitle\"\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createContact, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteContact, id);
}

func TestCreateDeleteGroup(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/createGroup" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"extension_id\": 1764590,\n\t\"name\": \"Ferengi Traders\"\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createGroup, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteGroup, id);

}

func TestCreateDeleteMenu(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/createMenu" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"name\": \"the menu Name\",\n\t\"allow_extension_dial\": true,\n\t\"keypress_wait_time\": 3\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createMenu, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteMenu, id);

}

func TestListReplaceMenu(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/replaceMenu" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"menu_id\":123123,\n\t\"name\": \"the menu Name\",\n\t\"allow_extension_dial\": true,\n\t\"keypress_wait_time\": 3\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createMenu, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}

func TestCreatePhoneNumber(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/createPhoneNumber" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"phone_number\": \"+12019570329\",\n\t\"name\": \"Phone Name Now\",\n\t\"block_incoming\": true,\n\t\"block_anonymous\": true,\n\t\"caller_id[name]\": \"Phone N\",\n\t\"caller_id[type]\": \"business\",\n\t\"sms_forwarding[type]\": \"extension\",\n\t\"call_notifications[sms]\": \"+18587740222\"\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createPhoneNumber, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}

func TestListReplacePhoneNumber(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/replacePhoneNumber" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"number_id\": \"+12019570329\",\n\t\"name\": \"Phone Name Now\",\n\t\"block_incoming\": true,\n\t\"block_anonymous\": true,\n\t\"caller_id[name]\": \"Phone N\",\n\t\"caller_id[type]\": \"business\",\n\t\"sms_forwarding[type]\": \"extension\",\n\t\"call_notifications[sms]\": \"+18587740222\"\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createPhoneNumber, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}

func TestCreateDeleteQueue(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/createQueue" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"name\": \"sample queue\",\n\t\"max_hold_time\": 60,\n\t\"caller_id_type\": \"called_number\",\n\t\"ring_time\": 10\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createQueue, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteQueue, id);

}

func TestListReplaceQueue(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/replaceQueue" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"queue_id\":123213,\n\t\"name\": \"sample queue\",\n\t\"max_hold_time\": 60,\n\t\"caller_id_type\": \"called_number\",\n\t\"ring_time\": 10\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createQueue, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}

func TestCreateDeleteRoute(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/createRoute" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"name\": \"API V3 Test\",\n\t\"rules\": [\n\t\t{\n\t\t\t\"actions\": [\n\t\t\t\t{\n\t\t\t\t\t\"action\": \"queue\",\n\t\t\t\t\t\"queue\": {\n\t\t\t\t\t\t\"id\": 21971,\n\t\t\t\t\t\t\"name\": \"sample queue\"\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t]\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createRoute, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteRoute, id);

}

func TestListReplaceRoute(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/replaceRoute" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"route_id\":12313,\n\t\"name\": \"API V3 Test\",\n\t\"rules\": [\n\t\t{\n\t\t\t\"actions\": [\n\t\t\t\t{\n\t\t\t\t\t\"action\": \"queue\",\n\t\t\t\t\t\"queue\": {\n\t\t\t\t\t\t\"id\": 21971,\n\t\t\t\t\t\t\"name\": \"sample queue\"\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t]\n\t\t}\n\t]\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createRoute, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}

func TestCreateSms(t *testing.T) {

  var json map[string] interface{}
  var err error

  randomName := randomString(12)
  fileName := "../test/jsonin/createSms" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"from\": \"+16309624775\",\n\t\"to\": \"+12019570328\",\n\t\"text\": \"Sms work too\",\n\t\"extension_id\": 1764595\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)

  err, json = createCliWithJsonIn(createSms, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}

func TestCreateSubaccount(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/createSubaccount" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"username\": \"\",\n\t\"password\": \"\"\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)



  err, json = createCliWithJsonIn(createSubaccount, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}

func TestCreateDeleteTrunk(t *testing.T) {

  var json map[string] interface{}
  var err error



  randomName := randomString(12)
  fileName := "../test/jsonin/createTrunk" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"name\": \"The trunk name\",\n\t\"uri\": \"SIP/1234@phone.com:5060\",\n\t\"max_concurrent_calls\": 60,\n\t\"max_minutes_per_month\": 800\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)

  err, json = createCliWithJsonIn(createTrunk, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

  createGetOrRemoveCli(deleteTrunk, id);

}

func TestListReplaceTrunk(t *testing.T) {

  var json map[string] interface{}
  var err error

  randomName := randomString(12)
  fileName := "../test/jsonin/createTrunk" + randomName + ".json"
  d1 := []byte("{\n\t\"account_id\": 1315091,\n\t\"trunk_id\":123213,\n\t\"name\": \"The trunk name\",\n\t\"uri\": \"SIP/1234@phone.com:5060\",\n\t\"max_concurrent_calls\": 60,\n\t\"max_minutes_per_month\": 800\n}")
  err = ioutil.WriteFile(fileName, d1, 0644)

  err, json = createCliWithJsonIn(createTrunk, fileName)
  os.Remove(fileName)
  assertErrorNotNull(t, err)

  id := getId(json)
  if (id == 0) {
    t.FailNow()
  }

}
