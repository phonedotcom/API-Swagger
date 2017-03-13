# SwaggerClient::ReplaceExtensionParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**voicemail_greeting_alternate** | **Object** | Recording lookup object | [optional] 
**name_greeting** | **Object** | Recording lookup object | [optional] 
**name** | **String** | Name (required) | [optional] 
**timezone** | **String** | Timezone | [optional] 
**include_in_directory** | **BOOLEAN** | Include in dial-by-name directory | [optional] 
**extension** | **Integer** | Extension number (required) | [optional] 
**enable_outbound_calls** | **BOOLEAN** | Enable outgoing calls | [optional] 
**usage_type** | **String** | Extension type | [optional] 
**voicemail_password** | **Integer** | Voicemail password | [optional] 
**full_name** | **String** | Contact name | [optional] 
**enable_call_waiting** | **BOOLEAN** | Enable Call Waiting | [optional] 
**voicemail_greeting_standard** | **Object** | Recording lookup object | [optional] 
**voicemail_greeting_type** | **String** | Voicemail greeting type | [optional] 
**caller_id** | **String** | Caller ID | [optional] 
**local_area_code** | **Integer** | Local area code | [optional] 
**voicemail_enabled** | **BOOLEAN** | Voicemail enabled | [optional] 
**voicemail_greeting_enable_leave_message_prompt** | **BOOLEAN** | Use leave message prompt after voicemail | [optional] 
**voicemail_transcription** | **String** | Voicemail transcription type | [optional] 
**voicemail_notifications_emails** | **Array&lt;String&gt;** | Email notifications for voicemails. Can be a single email or an array of emails | [optional] 
**voicemail_notifications_sms** | **String** | SMS notifications for voicemails | [optional] 
**call_notifications_emails** | **Array&lt;String&gt;** | Email notifications for calls. Can be a single email or an array of emails | [optional] 
**call_notifications_sms** | **String** | SMS notifications for calls | [optional] 
**route** | **Array&lt;String&gt;** | Route object lookup (must belong to this extension) | [optional] 


