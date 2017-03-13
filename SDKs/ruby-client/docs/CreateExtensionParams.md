# SwaggerClient::CreateExtensionParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**caller_id** | **String** | Caller ID | [optional] 
**usage_type** | **String** | Extension type | [optional] 
**allows_call_waiting** | **BOOLEAN** | Allows call waiting | [optional] 
**extension** | **Integer** | Extension number (auto-generated if empty) | [optional] 
**include_in_directory** | **BOOLEAN** | Include in dial-by-name directory | [optional] 
**name** | **String** | Name (auto-generated if empty) | [optional] 
**full_name** | **String** | Contact name | [optional] 
**timezone** | **String** | Timezone | [optional] 
**name_greeting** | **Object** | Recording lookup object | [optional] 
**voicemail_greeting_alternate** | **Object** | Recording lookup object | [optional] 
**local_area_code** | **Integer** | Local area code | [optional] 
**voicemail_greeting_enable_leave_message_prompt** | **BOOLEAN** | Enable the \&quot;leave a message\&quot; prompt for voicemail | [optional] 
**voicemail_enabled** | **BOOLEAN** | Voicemail enabled | [optional] 
**enable_outbound_calls** | **BOOLEAN** | Enable outgoing calls | [optional] 
**enable_call_waiting** | **BOOLEAN** | Enable Call Waiting | [optional] 
**voicemail_password** | **Integer** | Voicemail password | [optional] 
**voicemail_greeting_type** | **String** | Voicemail greeting type | [optional] 
**voicemail_greeting_standard** | **Object** | Recording lookup object | [optional] 
**voicemail_transcription** | **String** | Voicemail transcription type | [optional] 
**voicemail_notifications_emails** | **Array&lt;String&gt;** | Email notifications for voicemails. Can be a single email or an array of emails | [optional] 
**voicemail_notifications_sms** | **String** | SMS notifications for voicemails | [optional] 
**call_notifications_emails** | **Array&lt;String&gt;** | Email notifications for calls. Can be a single email or an array of emails | [optional] 
**call_notifications_sms** | **String** | SMS notifications for calls | [optional] 


