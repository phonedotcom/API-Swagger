
# PhoneNumberFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Integer** | Integer Phone number ID. This is the internal Phone.com ID for this number, not the phone number itself. Read-only. |  [optional]
**name** | **String** | Name |  [optional]
**phoneNumber** | **String** | Phone number, in E.164 format |  [optional]
**blockIncoming** | **Boolean** | Whether to block incoming calls. Boolean. |  [optional]
**blockAnonymous** | **Boolean** | Whether to block anonymous calls. Boolean. |  [optional]
**route** | [**RouteSummary**](RouteSummary.md) | The Route assigned to handle incoming calls for this number, if any. Output is a Route Summary Object, or NULL if not set. Input can be a Route Lookup Object or NULL to unset. |  [optional]
**callerId** | [**CallerIdPhoneNumber**](CallerIdPhoneNumber.md) | Caller ID Object, or NULL |  [optional]
**smsForwarding** | [**SmsForwarding**](SmsForwarding.md) | SMS Forwarding Object, or NULL |  [optional]
**callNotifications** | [**CallNotifications**](CallNotifications.md) |  |  [optional]



