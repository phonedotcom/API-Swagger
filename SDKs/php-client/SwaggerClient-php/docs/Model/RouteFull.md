# RouteFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Integer ID. Read-only. | [optional] 
**name** | **string** | Name | [optional] 
**extension** | [**\Swagger\Client\Model\ExtensionSummary**](ExtensionSummary.md) | Extension to which this route belongs. Output is an Extension Summary Object. Input must be an Extension Lookup Object. Optional. Cannot be changed after a route is created. | [optional] 
**rules** | [**\Swagger\Client\Model\RuleSet[]**](RuleSet.md) | Array of Rule Set Objects. Required. See below for details. When processing incoming calls, the first matching rule set will be used, and all others will be ignored. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


