# PhoneNumberFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Integer Phone number ID. This is the internal Phone.com ID for this number, not the phone number itself. Read-only. | [optional] [default to null]
**Name** | **string** | Name | [optional] [default to null]
**PhoneNumber** | **string** | Phone number, in E.164 format | [optional] [default to null]
**BlockIncoming** | **bool** | Whether to block incoming calls. Boolean. | [optional] [default to null]
**BlockAnonymous** | **bool** | Whether to block anonymous calls. Boolean. | [optional] [default to null]
**Route** | [**RouteSummary**](RouteSummary.md) | The Route assigned to handle incoming calls for this number, if any. Output is a Route Summary Object, or NULL if not set. Input can be a Route Lookup Object or NULL to unset. | [optional] [default to null]
**CallerId** | [**CallerIdPhoneNumber**](CallerIdPhoneNumber.md) | Caller ID Object, or NULL | [optional] [default to null]
**SmsForwarding** | [**SmsForwarding**](SmsForwarding.md) | SMS Forwarding Object, or NULL | [optional] [default to null]
**CallNotifications** | [**CallNotifications**](CallNotifications.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


