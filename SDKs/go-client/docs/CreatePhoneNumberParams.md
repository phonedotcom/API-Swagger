# CreatePhoneNumberParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | [**interface{}**](interface{}.md) | Phone number | [optional] [default to null]
**Route** | [**interface{}**](interface{}.md) | Route lookup object | [optional] [default to null]
**Name** | **string** | Phone Name | [optional] [default to null]
**BlockIncoming** | **bool** | Block incoming calls | [optional] [default to null]
**BlockAnonymous** | **bool** | Block anonymous calls | [optional] [default to null]
**CallerIdName** | **string** | Caller ID name | [optional] [default to null]
**CallerIdType** | **string** | Caller ID type | [optional] [default to null]
**SmsForwardingType** | **string** | &#39;application&#39; or &#39;extension&#39; | [optional] [default to null]
**SmsForwardingApplication** | [**interface{}**](interface{}.md) | Application lookup object | [optional] [default to null]
**SmsForwardingExtension** | [**interface{}**](interface{}.md) | Extension lookup object | [optional] [default to null]
**CallNotificationsEmails** | **[]string** | Call notifications for emails. Can be a single email or an array of emails | [optional] [default to null]
**CallNotificationsSms** | **string** | Call notification for SMS | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


