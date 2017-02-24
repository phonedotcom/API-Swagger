
# RouteFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Integer** | Integer ID. Read-only. |  [optional]
**name** | **String** | Name |  [optional]
**extension** | [**ExtensionSummary**](ExtensionSummary.md) | Extension to which this route belongs. Output is an Extension Summary Object. Input must be an Extension Lookup Object. Optional. Cannot be changed after a route is created. |  [optional]
**rules** | [**List&lt;RuleSet&gt;**](RuleSet.md) | Array of Rule Set Objects. Required. See below for details. When processing incoming calls, the first matching rule set will be used, and all others will be ignored. |  [optional]



