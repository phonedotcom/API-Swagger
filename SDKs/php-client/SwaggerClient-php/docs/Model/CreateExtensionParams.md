# CreateExtensionParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**caller_id** | **string** | Caller ID | [optional] 
**usage_type** | **string** | Extension type | [optional] 
**allows_call_waiting** | **bool** | Allows call waiting | [optional] 
**extension** | **int** | Extension number (auto-generated if empty) | [optional] 
**include_in_directory** | **bool** | Include in dial-by-name directory | [optional] 
**name** | **string** | Name (auto-generated if empty) | [optional] 
**full_name** | **string** | Contact name | [optional] 
**timezone** | **string** | Timezone | [optional] 
**name_greeting** | **object** | Recording lookup object | [optional] 
**voicemail_greeting_alternate** | **object** | Recording lookup object | [optional] 
**local_area_code** | **int** | Local area code | [optional] 
**voicemail_greeting_enable_leave_message_prompt** | **bool** | Enable the \&quot;leave a message\&quot; prompt for voicemail | [optional] 
**voicemail_enabled** | **bool** | Voicemail enabled | [optional] 
**enable_outbound_calls** | **bool** | Enable outgoing calls | [optional] 
**enable_call_waiting** | **bool** | Enable Call Waiting | [optional] 
**voicemail_password** | **int** | Voicemail password | [optional] 
**voicemail_greeting_type** | **string** | Voicemail greeting type | [optional] 
**voicemail_greeting_standard** | **object** | Recording lookup object | [optional] 
**voicemail_transcription** | **string** | Voicemail transcription type | [optional] 
**voicemail_notifications_emails** | **string[]** | Email notifications for voicemails. Can be a single email or an array of emails | [optional] 
**voicemail_notifications_sms** | **string** | SMS notifications for voicemails | [optional] 
**call_notifications_emails** | **string[]** | Email notifications for calls. Can be a single email or an array of emails | [optional] 
**call_notifications_sms** | **string** | SMS notifications for calls | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


