# SwaggerClient::CreatePhoneNumberParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**phone_number** | **Object** | Phone number | [optional] 
**route** | **Object** | Route lookup object | [optional] 
**name** | **String** | Phone Name | [optional] 
**block_incoming** | **BOOLEAN** | Block incoming calls | [optional] 
**block_anonymous** | **BOOLEAN** | Block anonymous calls | [optional] 
**caller_id_name** | **String** | Caller ID name | [optional] 
**caller_id_type** | **String** | Caller ID type | [optional] 
**sms_forwarding_type** | **String** | &#39;application&#39; or &#39;extension&#39; | [optional] 
**sms_forwarding_application** | **Object** | Application lookup object | [optional] 
**sms_forwarding_extension** | **Object** | Extension lookup object | [optional] 
**call_notifications_emails** | **Array&lt;String&gt;** | Call notifications for emails. Can be a single email or an array of emails | [optional] 
**call_notifications_sms** | **String** | Call notification for SMS | [optional] 


