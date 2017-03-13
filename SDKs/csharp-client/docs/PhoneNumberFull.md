# IO.Swagger.Model.PhoneNumberFull
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int?** | Integer Phone number ID. This is the internal Phone.com ID for this number, not the phone number itself. Read-only. | [optional] 
**Name** | **string** | Name | [optional] 
**PhoneNumber** | **string** | Phone number, in E.164 format | [optional] 
**BlockIncoming** | **bool?** | Whether to block incoming calls. Boolean. | [optional] 
**BlockAnonymous** | **bool?** | Whether to block anonymous calls. Boolean. | [optional] 
**Route** | [**RouteSummary**](RouteSummary.md) | The Route assigned to handle incoming calls for this number, if any. Output is a Route Summary Object, or NULL if not set. Input can be a Route Lookup Object or NULL to unset. | [optional] 
**CallerId** | [**CallerIdPhoneNumber**](CallerIdPhoneNumber.md) | Caller ID Object, or NULL | [optional] 
**SmsForwarding** | [**SmsForwarding**](SmsForwarding.md) | SMS Forwarding Object, or NULL | [optional] 
**CallNotifications** | [**CallNotifications**](CallNotifications.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

