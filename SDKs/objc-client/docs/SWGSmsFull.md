# SWGSmsFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**_id** | **NSString*** | Unique SMS ID. Read-only. | 
**from** | **NSString*** | Caller ID number to display on the incoming/outgoing SMS message. For an outgoing message, it must be a phone number associated with your Phone.com account. | 
**to** | [**NSArray&lt;SWGRecipient&gt;***](SWGRecipient.md) | An array of SMS recipients. | 
**direction** | **NSString*** | Direction of SMS. &#39;in&#39; for Incoming messages, &#39;out&#39; for Outgoing messages. | 
**createdEpoch** | **NSNumber*** | Unix time stamp representing the UTC time that the object was created in the Phone.com API system. | 
**createdAt** | **NSDate*** | Date string representing the UTC time that the object was created in the Phone.com API system. | 
**text** | **NSString*** | Body of the SMS text | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


