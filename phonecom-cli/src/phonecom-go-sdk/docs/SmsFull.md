# SmsFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique SMS ID. Read-only. | [default to null]
**From** | **string** | Caller ID number to display on the incoming/outgoing SMS message. For an outgoing message, it must be a phone number associated with your Phone.com account. | [default to null]
**To** | [**[]Recipient**](Recipient.md) | An array of SMS recipients. | [default to null]
**Direction** | **string** | Direction of SMS. &#39;in&#39; for Incoming messages, &#39;out&#39; for Outgoing messages. | [default to null]
**CreatedEpoch** | **int32** | Unix time stamp representing the UTC time that the object was created in the Phone.com API system. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Date string representing the UTC time that the object was created in the Phone.com API system. | [default to null]
**Text** | **string** | Body of the SMS text | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


