# ReplaceExtensionParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VoicemailGreetingAlternate** | [**interface{}**](interface{}.md) | Recording lookup object | [optional] [default to null]
**NameGreeting** | [**interface{}**](interface{}.md) | Recording lookup object | [optional] [default to null]
**Name** | **string** | Name (required) | [optional] [default to null]
**Timezone** | **string** | Timezone | [optional] [default to null]
**IncludeInDirectory** | **bool** | Include in dial-by-name directory | [optional] [default to null]
**Extension** | **int32** | Extension number (required) | [optional] [default to null]
**EnableOutboundCalls** | **bool** | Enable outgoing calls | [optional] [default to null]
**UsageType** | **string** | Extension type | [optional] [default to null]
**VoicemailPassword** | **int32** | Voicemail password | [optional] [default to null]
**FullName** | **string** | Contact name | [optional] [default to null]
**EnableCallWaiting** | **bool** | Enable Call Waiting | [optional] [default to null]
**VoicemailGreetingStandard** | [**interface{}**](interface{}.md) | Recording lookup object | [optional] [default to null]
**VoicemailGreetingType** | **string** | Voicemail greeting type | [optional] [default to null]
**CallerId** | **string** | Caller ID | [optional] [default to null]
**LocalAreaCode** | **int32** | Local area code | [optional] [default to null]
**VoicemailEnabled** | **bool** | Voicemail enabled | [optional] [default to null]
**VoicemailGreetingEnableLeaveMessagePrompt** | **bool** | Use leave message prompt after voicemail | [optional] [default to null]
**VoicemailTranscription** | **string** | Voicemail transcription type | [optional] [default to null]
**VoicemailNotificationsEmails** | **[]string** | Email notifications for voicemails. Can be a single email or an array of emails | [optional] [default to null]
**VoicemailNotificationsSms** | **string** | SMS notifications for voicemails | [optional] [default to null]
**CallNotificationsEmails** | **[]string** | Email notifications for calls. Can be a single email or an array of emails | [optional] [default to null]
**CallNotificationsSms** | **string** | SMS notifications for calls | [optional] [default to null]
**Route** | **[]string** | Route object lookup (must belong to this extension) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


