# SWGPhoneNumberFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**_id** | **NSNumber*** | Integer Phone number ID. This is the internal Phone.com ID for this number, not the phone number itself. Read-only. | [optional] 
**name** | **NSString*** | Name | [optional] 
**phoneNumber** | **NSString*** | Phone number, in E.164 format | [optional] 
**blockIncoming** | **NSNumber*** | Whether to block incoming calls. Boolean. | [optional] 
**blockAnonymous** | **NSNumber*** | Whether to block anonymous calls. Boolean. | [optional] 
**route** | [**SWGRouteSummary***](SWGRouteSummary.md) | The Route assigned to handle incoming calls for this number, if any. Output is a Route Summary Object, or NULL if not set. Input can be a Route Lookup Object or NULL to unset. | [optional] 
**callerId** | [**SWGCallerIdPhoneNumber***](SWGCallerIdPhoneNumber.md) | Caller ID Object, or NULL | [optional] 
**smsForwarding** | [**SWGSmsForwarding***](SWGSmsForwarding.md) | SMS Forwarding Object, or NULL | [optional] 
**callNotifications** | [**SWGCallNotifications***](SWGCallNotifications.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


