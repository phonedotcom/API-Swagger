# TrunkFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Integer Trunk ID. Read-only. | 
**name** | **string** | Name. Required. | 
**uri** | **string** | Fully-qualified SIP URI. Required. | 
**max_concurrent_calls** | **int** | Max concurrent calls. Default is 10. | 
**max_minutes_per_month** | **int** | Max minutes per month. Default is 750. | 
**greeting** | [**\Swagger\Client\Model\MediaSummary**](MediaSummary.md) | Greeting. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. | 
**error_message** | [**\Swagger\Client\Model\MediaSummary**](MediaSummary.md) | Error Message. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. | 
**codecs** | **string[]** | Custom audio codec configuration, if any is needed. If provided, must be a simple array containing the prioritized list of desired codecs. Supported codecs are: g711u 64k, g711u 56k, g711a 64k, g711a 56k, g7231, g728, g729, g729A, g729B, g729AB, gms full, rfc2833, t38, ilbc, h263, g722, g722_1, g729D, g729E, amr, amr_wb, efr, evrc, h264, mpeg4, red, cng, SIP Info to 2833 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


