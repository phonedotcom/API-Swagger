# RouteFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Integer ID. Read-only. | [optional] [default to null]
**Name** | **string** | Name | [optional] [default to null]
**Extension** | [**ExtensionSummary**](ExtensionSummary.md) | Extension to which this route belongs. Output is an Extension Summary Object. Input must be an Extension Lookup Object. Optional. Cannot be changed after a route is created. | [optional] [default to null]
**Rules** | [**[]RuleSet**](RuleSet.md) | Array of Rule Set Objects. Required. See below for details. When processing incoming calls, the first matching rule set will be used, and all others will be ignored. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


