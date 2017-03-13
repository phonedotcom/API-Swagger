# ReplaceExtensionParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**voicemail_greeting_alternate** | **object** | Recording lookup object | [optional] 
**name_greeting** | **object** | Recording lookup object | [optional] 
**name** | **str** | Name (required) | [optional] 
**timezone** | **str** | Timezone | [optional] 
**include_in_directory** | **bool** | Include in dial-by-name directory | [optional] 
**extension** | **int** | Extension number (required) | [optional] 
**enable_outbound_calls** | **bool** | Enable outgoing calls | [optional] 
**usage_type** | **str** | Extension type | [optional] 
**voicemail_password** | **int** | Voicemail password | [optional] 
**full_name** | **str** | Contact name | [optional] 
**enable_call_waiting** | **bool** | Enable Call Waiting | [optional] 
**voicemail_greeting_standard** | **object** | Recording lookup object | [optional] 
**voicemail_greeting_type** | **str** | Voicemail greeting type | [optional] 
**caller_id** | **str** | Caller ID | [optional] 
**local_area_code** | **int** | Local area code | [optional] 
**voicemail_enabled** | **bool** | Voicemail enabled | [optional] 
**voicemail_greeting_enable_leave_message_prompt** | **bool** | Use leave message prompt after voicemail | [optional] 
**voicemail_transcription** | **str** | Voicemail transcription type | [optional] 
**voicemail_notifications_emails** | **list[str]** | Email notifications for voicemails. Can be a single email or an array of emails | [optional] 
**voicemail_notifications_sms** | **str** | SMS notifications for voicemails | [optional] 
**call_notifications_emails** | **list[str]** | Email notifications for calls. Can be a single email or an array of emails | [optional] 
**call_notifications_sms** | **str** | SMS notifications for calls | [optional] 
**route** | **list[str]** | Route object lookup (must belong to this extension) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


