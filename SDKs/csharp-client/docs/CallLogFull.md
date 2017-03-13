# IO.Swagger.Model.CallLogFull
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID | [optional] 
**Uuid** | **string** | Internal system id, may be null | [optional] 
**Extension** | [**ExtensionSummary**](ExtensionSummary.md) | Account extension | [optional] 
**CallerId** | **string** | Call made from this phone number | [optional] 
**CalledNumber** | **string** | Call made to this phone number | [optional] 
**StartTime** | **string** | Call start time | [optional] 
**CreatedAt** | **string** | Call log creation time. Same as call end time + time needed to create call log | [optional] 
**Direction** | **string** | Call direction: in or out | [optional] 
**Type** | **string** | Call type: call, fax, audiogram ... | [optional] 
**CallDuration** | **int?** | Call duration in seconds | [optional] 
**IsMonitored** | **string** | Was call being monitored? | [optional] 
**CallNumber** | **string** | Internal system call reference number | [optional] 
**FinalAction** | **string** | Last action of call flow | [optional] 
**CallRecording** | **string** | URL of call recording if available. Empty string if call recording does not exist | [optional] 
**Details** | [**List&lt;CallDetails&gt;**](CallDetails.md) | A list of call flows from beginning of call to end of call. | [optional] 
**CallerCnam** | **string** | Internal system caller id / name | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

