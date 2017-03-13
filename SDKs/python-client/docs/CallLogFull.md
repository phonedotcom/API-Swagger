# CallLogFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | ID | [optional] 
**uuid** | **str** | Internal system id, may be null | [optional] 
**extension** | [**ExtensionSummary**](ExtensionSummary.md) | Account extension | [optional] 
**caller_id** | **str** | Call made from this phone number | [optional] 
**called_number** | **str** | Call made to this phone number | [optional] 
**start_time** | **str** | Call start time | [optional] 
**created_at** | **str** | Call log creation time. Same as call end time + time needed to create call log | [optional] 
**direction** | **str** | Call direction: in or out | [optional] 
**type** | **str** | Call type: call, fax, audiogram ... | [optional] 
**call_duration** | **int** | Call duration in seconds | [optional] 
**is_monitored** | **str** | Was call being monitored? | [optional] 
**call_number** | **str** | Internal system call reference number | [optional] 
**final_action** | **str** | Last action of call flow | [optional] 
**call_recording** | **str** | URL of call recording if available. Empty string if call recording does not exist | [optional] 
**details** | [**list[CallDetails]**](CallDetails.md) | A list of call flows from beginning of call to end of call. | [optional] 
**caller_cnam** | **str** | Internal system caller id / name | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


