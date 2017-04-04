package main

import "github.com/phonedotcom/API-SDK-go"

type CreateDeviceJson struct {
	Name string `json:"name"`
}

type ReplaceDeviceJson struct {
	DeviceId int32  `json:"device_id"`
	Name     string `json:"name"`
}

type CreateExtensionJson struct {
	Voicemail           swagger.Voicemail         `json:"voicemail"`
	CallNotifications   swagger.CallNotifications `json:"call_notifications"`
	CallerId            string                    `json:"caller_id"`
	UsageType           string                    `json:"usage_type"`
	AllowsCallWaiting   bool                      `json:"allows_call_waiting"`
	Extension           int32                     `json:"extension"`
	IncludeInDirectory  bool                      `json:"include_in_directory"`
	Name                string                    `json:"name"`
	FullName            string                    `json:"full_name"`
	Timezone            string                    `json:"timezone"`
	NameGreeting        interface{}               `json:"name_greeting"`
	LocalAreaCode       int32                     `json:"local_area_code"`
	EnableOutboundCalls bool                      `json:"enable_outbound_calls"`
	EnableCallWaiting   bool                      `json:"enable_call_waiting"`
}

type ReplaceExtensionJson struct {
	ExtensionId         int32                     `json:"extension_id"`
	Voicemail           swagger.Voicemail         `json:"voicemail"`
	CallNotifications   swagger.CallNotifications `json:"call_notifications"`
	NameGreeting        interface{}               `json:"name_greeting"`
	Name                string                    `json:"name"`
	Timezone            string                    `json:"timezone"`
	IncludeInDirectory  bool                      `json:"include_in_directory"`
	Extension           int32                     `json:"extension"`
	EnableOutboundCalls bool                      `json:"enable_outbound_calls"`
	UsageType           string                    `json:"usage_type"`
	FullName            string                    `json:"full_name"`
	EnableCallWaiting   bool                      `json:"enable_call_waiting"`
	CallerId            string                    `json:"caller_id"`
	LocalAreaCode       int32                     `json:"local_area_code"`
	Route               string                    `json:"route"`
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
	PhoneNumber       string                      `json:"phone_number"`
	Route             interface{}                 `json:"route"`
	Name              string                      `json:"name"`
	BlockIncoming     bool                        `json:"block_incoming"`
	BlockAnonymous    bool                        `json:"block_anonymous"`
	CallerId          swagger.CallerIdPhoneNumber `json:"caller_id"`
	SmsForwarding     swagger.SmsForwardingParams `json:"sms_forwarding"`
	CallNotifications swagger.CallNotifications   `json:"call_notifications"`
}

type ReplacePhoneNumberJson struct {
	NumberId          int32                       `json:"number_id"`
	Route             interface{}                 `json:"route"`
	Name              string                      `json:"name"`
	BlockIncoming     bool                        `json:"block_incoming"`
	BlockAnonymous    bool                        `json:"block_anonymous"`
	CallerId          swagger.CallerIdPhoneNumber `json:"caller_id"`
	SmsForwarding     swagger.SmsForwardingParams `json:"sms_forwarding"`
	PoolItem          interface{}                 `json:"pool_item"`
	CallNotifications swagger.CallNotifications   `json:"call_notifications"`
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
