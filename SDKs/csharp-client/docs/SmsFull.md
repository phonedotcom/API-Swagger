# IO.Swagger.Model.SmsFull
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique SMS ID. Read-only. | 
**From** | **string** | Caller ID number to display on the incoming/outgoing SMS message. For an outgoing message, it must be a phone number associated with your Phone.com account. | 
**To** | [**List&lt;Recipient&gt;**](Recipient.md) | An array of SMS recipients. | 
**Direction** | **string** | Direction of SMS. &#39;in&#39; for Incoming messages, &#39;out&#39; for Outgoing messages. | 
**CreatedEpoch** | **int?** | Unix time stamp representing the UTC time that the object was created in the Phone.com API system. | 
**CreatedAt** | **DateTime?** | Date string representing the UTC time that the object was created in the Phone.com API system. | 
**Text** | **string** | Body of the SMS text | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

