
# SmsFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **String** | Unique SMS ID. Read-only. | 
**from** | **String** | Caller ID number to display on the incoming/outgoing SMS message. For an outgoing message, it must be a phone number associated with your Phone.com account. | 
**to** | [**List&lt;Recipient&gt;**](Recipient.md) | An array of SMS recipients. | 
**direction** | **String** | Direction of SMS. &#39;in&#39; for Incoming messages, &#39;out&#39; for Outgoing messages. | 
**createdEpoch** | **Integer** | Unix time stamp representing the UTC time that the object was created in the Phone.com API system. | 
**createdAt** | [**DateTime**](DateTime.md) | Date string representing the UTC time that the object was created in the Phone.com API system. | 
**text** | **String** | Body of the SMS text | 



