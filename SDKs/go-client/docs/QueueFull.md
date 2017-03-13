# QueueFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Integer ID. Read-only. | [default to null]
**Name** | **string** | Name. Required. | [default to null]
**Greeting** | [**MediaSummary**](MediaSummary.md) | Greeting to be played when caller first connects. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. Can be set to NULL to disable the greeting. | [optional] [default to null]
**HoldMusic** | [**HoldMusic**](HoldMusic.md) |  | [optional] [default to null]
**MaxHoldTime** | **int32** | Maximum hold time in seconds. If provided, must equal one of: 60, 120, 180, 240, 300, 600, 900, 1200, 1800, 2700, 3600. Default is 300. | [optional] [default to null]
**CallerIdType** | **string** | Caller id type to show members. If provided, must equal one of: &#39;called_number&#39;, &#39;calling_number&#39;. Default is &#39;calling_number&#39;. | [optional] [default to null]
**RingTime** | **int32** | Number of seconds to ring a member before cycling to the next member. If provided, must equal one of: 5, 10, 15, 20, 25, 30. Default is 5. | [optional] [default to null]
**Members** | [**[]Member**](Member.md) | Array of Member Objects. Non-virtual account extensions or phone numbers. See below for details. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


