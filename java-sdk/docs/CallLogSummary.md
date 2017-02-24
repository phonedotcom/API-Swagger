
# CallLogSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **String** | ID |  [optional]
**uuid** | **String** | Internal system id, may be null |  [optional]
**extension** | [**ExtensionSummary**](ExtensionSummary.md) | Account extension |  [optional]
**callerId** | **String** | Call made from this phone number |  [optional]
**calledNumber** | **String** | Call made to this phone number |  [optional]
**startTime** | **String** | Call start time |  [optional]
**createdAt** | **String** | Call log creation time. Same as call end time + time needed to create call log |  [optional]
**direction** | **String** | Call direction: in or out |  [optional]
**type** | **String** | Call type: call, fax, audiogram ... |  [optional]
**callDuration** | **Integer** | Call duration in seconds |  [optional]
**isMonitored** | **String** | Was call being monitored? |  [optional]
**callNumber** | **String** | Internal system call reference number |  [optional]
**finalAction** | **String** | Last action of call flow |  [optional]
**callRecording** | **String** | URL of call recording if available. Empty string if call recording does not exist |  [optional]



