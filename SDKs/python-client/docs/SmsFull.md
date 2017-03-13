# SmsFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | Unique SMS ID. Read-only. | 
**_from** | **str** | Caller ID number to display on the incoming/outgoing SMS message. For an outgoing message, it must be a phone number associated with your Phone.com account. | 
**to** | [**list[Recipient]**](Recipient.md) | An array of SMS recipients. | 
**direction** | **str** | Direction of SMS. &#39;in&#39; for Incoming messages, &#39;out&#39; for Outgoing messages. | 
**created_epoch** | **int** | Unix time stamp representing the UTC time that the object was created in the Phone.com API system. | 
**created_at** | **datetime** | Date string representing the UTC time that the object was created in the Phone.com API system. | 
**text** | **str** | Body of the SMS text | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


