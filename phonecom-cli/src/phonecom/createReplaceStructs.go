package main

import "github.com/phonedotcom/API-SDK-go"

type CreateDeviceJson struct {
	Name string `json:"name"`
}

type ReplaceDeviceJson struct {
	DeviceId int32  `json:"device_id"`
	Name     string `json:"name"`
}

type CreateMenuJson struct {
	Name               string `json:"name"`
	AllowExtensionDial string `json:"allow_extension_dial"`
	KeypressWaitTime   int32  `json:"keypress_wait_time"`
}

type ReplaceMenuJson struct {
	MenuId             int32  `json:"menu_id"`
	Name               string `json:"name"`
	AllowExtensionDial string `json:"allow_extension_dial"`
	KeypressWaitTime   int32  `json:"keypress_wait_time"`
}

type CreateQueueJson struct {
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
	Queue  swagger.QueueSummary `json:"queue"`
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
