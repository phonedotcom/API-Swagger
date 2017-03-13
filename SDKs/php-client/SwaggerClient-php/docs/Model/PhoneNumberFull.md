# PhoneNumberFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Integer Phone number ID. This is the internal Phone.com ID for this number, not the phone number itself. Read-only. | [optional] 
**name** | **string** | Name | [optional] 
**phone_number** | **string** | Phone number, in E.164 format | [optional] 
**block_incoming** | **bool** | Whether to block incoming calls. Boolean. | [optional] 
**block_anonymous** | **bool** | Whether to block anonymous calls. Boolean. | [optional] 
**route** | [**\Swagger\Client\Model\RouteSummary**](RouteSummary.md) | The Route assigned to handle incoming calls for this number, if any. Output is a Route Summary Object, or NULL if not set. Input can be a Route Lookup Object or NULL to unset. | [optional] 
**caller_id** | [**\Swagger\Client\Model\CallerIdPhoneNumber**](CallerIdPhoneNumber.md) | Caller ID Object, or NULL | [optional] 
**sms_forwarding** | [**\Swagger\Client\Model\SmsForwarding**](SmsForwarding.md) | SMS Forwarding Object, or NULL | [optional] 
**call_notifications** | [**\Swagger\Client\Model\CallNotifications**](CallNotifications.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


