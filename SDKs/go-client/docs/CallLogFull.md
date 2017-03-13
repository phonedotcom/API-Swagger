# CallLogFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID | [optional] [default to null]
**Uuid** | **string** | Internal system id, may be null | [optional] [default to null]
**Extension** | [**ExtensionSummary**](ExtensionSummary.md) | Account extension | [optional] [default to null]
**CallerId** | **string** | Call made from this phone number | [optional] [default to null]
**CalledNumber** | **string** | Call made to this phone number | [optional] [default to null]
**StartTime** | **string** | Call start time | [optional] [default to null]
**CreatedAt** | **string** | Call log creation time. Same as call end time + time needed to create call log | [optional] [default to null]
**Direction** | **string** | Call direction: in or out | [optional] [default to null]
**Type_** | **string** | Call type: call, fax, audiogram ... | [optional] [default to null]
**CallDuration** | **int32** | Call duration in seconds | [optional] [default to null]
**IsMonitored** | **string** | Was call being monitored? | [optional] [default to null]
**CallNumber** | **string** | Internal system call reference number | [optional] [default to null]
**FinalAction** | **string** | Last action of call flow | [optional] [default to null]
**CallRecording** | **string** | URL of call recording if available. Empty string if call recording does not exist | [optional] [default to null]
**Details** | [**[]CallDetails**](CallDetails.md) | A list of call flows from beginning of call to end of call. | [optional] [default to null]
**CallerCnam** | **string** | Internal system caller id / name | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


