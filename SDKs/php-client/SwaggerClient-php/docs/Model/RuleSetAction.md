# RuleSetAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**action** | **string** | Required. | [optional] 
**extension** | [**\Swagger\Client\Model\ExtensionSummary**](ExtensionSummary.md) | Extension that this action refers to. Output is an Extension Summary Object. Input must be an Extension Lookup Object. Required. | [optional] 
**items** | [**\Swagger\Client\Model\RuleSetForwardItem[]**](RuleSetForwardItem.md) | This action is for forwarding calls to any number of extensions or phone numbers. The forwarding is handled in parallel, meaning that all phone numbers and/or extensions will ring simultaneously. When the call is answered by any single phone number or extension, ringing will stop for all of them. Subsequent actions in this rule set will be performed if the call is not answered before the timeout period is reached, or if it is forwarded to an extension that has its own route and that route does not result in any actions that disconnect the call or take over call handling. | [optional] 
**timeout** | **int** | Seconds that our routing engine should wait until moving on. Optional. Must be an integer between 5 and 90. Default is 5 seconds. | [optional] 
**hold_music** | [**\Swagger\Client\Model\MediaSummary**](MediaSummary.md) | Hold Music to be played while callers are waiting. Output is a Media Summary Object. Input must be a Media Lookup Object. Optional. Must refer to a media recording that has is_hold_music set to TRUE. Default is to play a standard ring tone. | [optional] 
**greeting** | [**\Swagger\Client\Model\MediaSummary**](MediaSummary.md) | Greeting that this action refers to. Output is a Media Summary Object. Input must be a Media Lookup Object. Required. Must refer to a media recording that has is_hold_music set to FALSE. | [optional] 
**duration** | **int** | Required. Seconds that the caller should be placed on hold before being moved onto the next action. Must be an integer between 1 and 60 seconds. | [optional] 
**menu** | [**\Swagger\Client\Model\MenuSummary**](MenuSummary.md) | Menu that this action refers to. Required. Output is a Menu Summary Object. Input must be a Menu Lookup Object. | [optional] 
**queue** | [**\Swagger\Client\Model\QueueSummary**](QueueSummary.md) | Queue that this action refers to. Required. Output is a Queue Summary Object. Input must be a Queue Lookup Object. | [optional] 
**trunk** | [**\Swagger\Client\Model\TrunkSummary**](TrunkSummary.md) | Trunk that this action refers to. Required. Output is a Trunk Summary Object. Input must be a Trunk Lookup Object. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


