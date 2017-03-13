# SmsFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | Unique SMS ID. Read-only. | 
**from** | **string** | Caller ID number to display on the incoming/outgoing SMS message. For an outgoing message, it must be a phone number associated with your Phone.com account. | 
**to** | [**\Swagger\Client\Model\Recipient[]**](Recipient.md) | An array of SMS recipients. | 
**direction** | **string** | Direction of SMS. &#39;in&#39; for Incoming messages, &#39;out&#39; for Outgoing messages. | 
**created_epoch** | **int** | Unix time stamp representing the UTC time that the object was created in the Phone.com API system. | 
**created_at** | [**\DateTime**](\DateTime.md) | Date string representing the UTC time that the object was created in the Phone.com API system. | 
**text** | **string** | Body of the SMS text | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


