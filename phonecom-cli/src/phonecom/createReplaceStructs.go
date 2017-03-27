package main

type CreateDeviceJson struct {
	Name string `json:"name"`
}

type ReplaceDeviceJson struct {
	DeviceId int32  `json:"device_id"`
	Name     string `json:"name"`
}

type CreateExtensionJson struct {
	CallerId                                  string `json:"caller_id"`
	UsageType                                 string `json:"usage_type"`
	AllowsCallWaiting                         bool   `json:"allows_call_waiting"`
	Extension                                 int32  `json:"extension"`
	IncludeInDirectory                        bool   `json:"include_in_directory"`
	Name                                      string `json:"name"`
	FullName                                  string `json:"full_name"`
	Timezone                                  string `json:"timezone"`
	LocalAreaCode                             int32  `json:"local_area_code"`
	VoicemailGreetingEnableLeaveMessagePrompt bool   `json:"voicemail[greeting][enable_leave_message_prompt]"`
	VoicemailEnabled                          bool   `json:"voicemail[enabled]"`
	EnableOutboundCalls                       bool   `json:"enable_outbound_calls"`
	EnableCallWaiting                         bool   `json:"enable_call_waiting"`
	VoicemailPassword                         int32  `json:"voicemail[password]"`
	VoicemailGreetingType                     string `json:"voicemail[greeting][type]"`
	VoicemailTranscription                    string `json:"voicemail[transcription]"`
	VoicemailNotificationsSms                 string `json:"voicemail[notifications][sms]"`
	CallNotificationsSms                      string `json:"call_notifications[sms]"`
}

type ReplaceExtensionJson struct {
	ExtensionId                               int32  `json:"extension_id"`
	Name                                      string `json:"name"`
	Timezone                                  string `json:"timezone"`
	IncludeInDirectory                        bool   `json:"include_in_directory"`
	Extension                                 int32  `json:"extension"`
	EnableOutboundCalls                       bool   `json:"enable_outbound_calls"`
	UsageType                                 string `json:"usage_type"`
	VoicemailPassword                         int32  `json:"voicemail[password]"`
	FullName                                  string `json:"full_name"`
	EnableCallWaiting                         bool   `json:"enable_call_waiting"`
	VoicemailGreetingType                     string `json:"voicemail[greeting][type]"`
	CallerId                                  string `json:"caller_id"`
	LocalAreaCode                             int32  `json:"local_area_code"`
	VoicemailEnabled                          bool   `json:"voicemail[enabled]"`
	VoicemailGreetingEnableLeaveMessagePrompt bool   `json:"voicemail[greeting][enable_leave_message_prompt]"`
	VoicemailTranscription                    string `json:"voicemail[transcription]"`
	VoicemailNotificationsSms                 string `json:"voicemail[notifications][sms]"`
	CallNotificationsSms                      string `json:"call_notifications[sms]"`
}

type CreateContactJson struct {
	ExtensionId        int32  `json:"extension_id"`
	FirstName          string `json:"first_name"`
	MiddleName         string `json:"middle_name"`
	LastName           string `json:"last_name"`
	Prefix             string `json:"prefix"`
	PhoneticFirstName  string `json:"phonetic_first_name"`
	PhoneticMiddleName string `json:"phonetic_middle_name"`
	PhoneticLastName   string `json:"phonetic_last_name"`
	Suffix             string `json:"suffix"`
	Nickname           string `json:"nickname"`
	Company            string `json:"company"`
	Department         string `json:"department"`
	JobTitle           string `json:"job_title"`
}

type CreateGroupJson struct {
	ExtensionId int32  `json:"extension_id"`
	Name        string `json:"name"`
}

type CreateMenuJson struct {
	Name               string `json:"name"`
	AllowExtensionDial bool   `json:"allow_extension_dial"`
	KeypressWaitTime   int32  `json:"keypress_wait_time"`
}

type ReplaceMenuJson struct {
	MenuId             int32  `json:"menu_id"`
	Name               string `json:"name"`
	AllowExtensionDial bool   `json:"allow_extension_dial"`
	KeypressWaitTime   int32  `json:"keypress_wait_time"`
}

type CreatePhoneNumberJson struct {
	PhoneNumber          string `json:"phone_number"`
	Name                 string `json:"name"`
	BlockIncoming        bool   `json:"block_incoming"`
	BlockAnonymous       bool   `json:"block_anonymous"`
	CallerIdName         string `json:"caller_id[name]"`
	CallerIdType         string `json:"caller_id[type]"`
	SmsForwardingType    string `json:"sms_forwarding[type]"`
	CallNotificationsSms string `json:"call_notifications[sms]"`
}

type ReplacePhoneNumberJson struct {
	NumberId             int32  `json:"number_id"`
	Name                 string `json:"name"`
	BlockIncoming        bool   `json:"block_incoming"`
	BlockAnonymous       bool   `json:"block_anonymous"`
	CallerIdName         string `json:"caller_id[name]"`
	CallerIdType         string `json:"caller_id[type]"`
	SmsForwardingType    string `json:"sms_forwarding[type]"`
	CallNotificationsSms string `json:"call_notifications[sms]"`
}

type CreateQueueJson struct {
	Name         string `json:"name"`
	MaxHoldTime  int32  `json:"max_hold_time"`
	CallerIdType string `json:"caller_id_type"`
	RingTime     int32  `json:"ring_time"`
}

type ReplaceQueueJson struct {
	QueueId      int32  `json:"queue_id"`
	Name         string `json:"name"`
	MaxHoldTime  int32  `json:"max_hold_time"`
	CallerIdType string `json:"caller_id_type"`
	RingTime     int32  `json:"ring_time"`
}

type CreateRouteJson struct {
	Name  string      `json:"name"`
	Rules []RulesJson `json:"rules"`
}

type RulesJson struct {
	Actions []ActionsJson `json:"actions"`
}

type ActionsJson struct {
	Action string    `json:"action"`
	Queue  QueueJson `json:"queue"`
}

type QueueJson struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

type ReplaceRouteJson struct {
	RouteId int32       `json:"route_id"`
	Name    string      `json:"name"`
	Rules   []RulesJson `json:"rules"`
}

type CreateSmsJson struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Text        string `json:"text"`
	ExtensionId int32  `json:"extension_id"`
}

type CreateSubaccountJson struct {
	Username       string      `json:"username"`
	Password       string      `json:"password"`
	Contact        ContactJson `json:"contact"`
	BillingContact ContactJson `json:"billing_contact"`
}

type ContactJson struct {
	Name         string      `json:"name"`
	Address      AddressJson `json:"address"`
	Phone        string      `json:"phone"`
	PrimaryEmail string      `json:"primary_email"`
}

type AddressJson struct {
	Line1      string `json:"line_1"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type CreateTrunkJson struct {
	Name               string `json:"name"`
	Uri                string `json:"uri"`
	MaxConcurrentCalls int32  `json:"max_concurrent_calls"`
	MaxMinutesPerMonth int32  `json:"max_minutes_per_month"`
}

type ReplaceTrunkJson struct {
	TrunkId            int32  `json:"trunk_id"`
	Name               string `json:"name"`
	Uri                string `json:"uri"`
	MaxConcurrentCalls int32  `json:"max_concurrent_calls"`
	MaxMinutesPerMonth int32  `json:"max_minutes_per_month"`
}

type GetAccountJson struct {
	AccountId int32 `json:"account_id"`
}

type GetApplicationJson struct {
	ApplicationId int32 `json:"application_id"`
}

type GetCallLogJson struct {
	CallLogId string `json:"call_log_id"`
}

type GetDeviceJson struct {
	DeviceId int32 `json:"device_id"`
}

type GetExpressServiceCodeJson struct {
	ExpressServiceCodeId int32 `json:"code_id"`
}

type GetExtensionJson struct {
	ExtensionId int32 `json:"extension_id"`
}

type GetContactJson struct {
	ExtensionId int32 `json:"extension_id"`
	ContactId   int32 `json:"contact_id"`
}

type GetGroupJson struct {
	ExtensionId int32 `json:"extension_id"`
	GroupId     int32 `json:"group_id"`
}

type GetMediaJson struct {
	MediaId int32 `json:"media_id"`
}

type GetMenuJson struct {
	MenuId int32 `json:"menu_id"`
}

type GetPhoneNumberJson struct {
	PhoneNumberId int32 `json:"number_id"`
}

type GetQueueJson struct {
	QueueId int32 `json:"queue_id"`
}

type GetRouteJson struct {
	RouteId int32 `json:"route_id"`
}

type GetScheduleJson struct {
	ScheduleId int32 `json:"schedule_id"`
}

type GetSmsJson struct {
	SmsId string `json:"sms_id"`
}

type GetTrunkJson struct {
	TrunkId int32 `json:"trunk_id"`
}
