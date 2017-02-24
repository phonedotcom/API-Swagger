
# TrunkFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Integer** | Integer Trunk ID. Read-only. | 
**name** | **String** | Name. Required. | 
**uri** | **String** | Fully-qualified SIP URI. Required. | 
**maxConcurrentCalls** | **Integer** | Max concurrent calls. Default is 10. | 
**maxMinutesPerMonth** | **Integer** | Max minutes per month. Default is 750. | 
**greeting** | [**MediaSummary**](MediaSummary.md) | Greeting. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. | 
**errorMessage** | [**MediaSummary**](MediaSummary.md) | Error Message. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. | 
**codecs** | **List&lt;String&gt;** | Custom audio codec configuration, if any is needed. If provided, must be a simple array containing the prioritized list of desired codecs. Supported codecs are: g711u 64k, g711u 56k, g711a 64k, g711a 56k, g7231, g728, g729, g729A, g729B, g729AB, gms full, rfc2833, t38, ilbc, h263, g722, g722_1, g729D, g729E, amr, amr_wb, efr, evrc, h264, mpeg4, red, cng, SIP Info to 2833 | 



