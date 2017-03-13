# CallLogFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | ID | [optional] 
**uuid** | **string** | Internal system id, may be null | [optional] 
**extension** | [**\Swagger\Client\Model\ExtensionSummary**](ExtensionSummary.md) | Account extension | [optional] 
**caller_id** | **string** | Call made from this phone number | [optional] 
**called_number** | **string** | Call made to this phone number | [optional] 
**start_time** | **string** | Call start time | [optional] 
**created_at** | **string** | Call log creation time. Same as call end time + time needed to create call log | [optional] 
**direction** | **string** | Call direction: in or out | [optional] 
**type** | **string** | Call type: call, fax, audiogram ... | [optional] 
**call_duration** | **int** | Call duration in seconds | [optional] 
**is_monitored** | **string** | Was call being monitored? | [optional] 
**call_number** | **string** | Internal system call reference number | [optional] 
**final_action** | **string** | Last action of call flow | [optional] 
**call_recording** | **string** | URL of call recording if available. Empty string if call recording does not exist | [optional] 
**details** | [**\Swagger\Client\Model\CallDetails[]**](CallDetails.md) | A list of call flows from beginning of call to end of call. | [optional] 
**caller_cnam** | **string** | Internal system caller id / name | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


