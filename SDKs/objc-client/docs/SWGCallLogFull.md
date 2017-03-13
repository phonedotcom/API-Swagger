# SWGCallLogFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**_id** | **NSString*** | ID | [optional] 
**uuid** | **NSString*** | Internal system id, may be null | [optional] 
**extension** | [**SWGExtensionSummary***](SWGExtensionSummary.md) | Account extension | [optional] 
**callerId** | **NSString*** | Call made from this phone number | [optional] 
**calledNumber** | **NSString*** | Call made to this phone number | [optional] 
**startTime** | **NSString*** | Call start time | [optional] 
**createdAt** | **NSString*** | Call log creation time. Same as call end time + time needed to create call log | [optional] 
**direction** | **NSString*** | Call direction: in or out | [optional] 
**type** | **NSString*** | Call type: call, fax, audiogram ... | [optional] 
**callDuration** | **NSNumber*** | Call duration in seconds | [optional] 
**isMonitored** | **NSString*** | Was call being monitored? | [optional] 
**callNumber** | **NSString*** | Internal system call reference number | [optional] 
**finalAction** | **NSString*** | Last action of call flow | [optional] 
**callRecording** | **NSString*** | URL of call recording if available. Empty string if call recording does not exist | [optional] 
**details** | [**NSArray&lt;SWGCallDetails&gt;***](SWGCallDetails.md) | A list of call flows from beginning of call to end of call. | [optional] 
**callerCnam** | **NSString*** | Internal system caller id / name | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


