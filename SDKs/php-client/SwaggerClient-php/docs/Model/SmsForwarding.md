# SmsForwarding

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **string** | Can be \&quot;extension\&quot; or \&quot;application\&quot; | [optional] 
**extension** | [**\Swagger\Client\Model\ExtensionSummary**](ExtensionSummary.md) | Required if type &#x3D; \&quot;extension\&quot;. Extension that messages should be directed to. Output is an Extension Summary Object. Input must be an Extension Lookup Object. | [optional] 
**application** | [**\Swagger\Client\Model\ApplicationSummary**](ApplicationSummary.md) | Required if type &#x3D; \&quot;application\&quot;. Application that messages should be directed to. Output is an Application Summary Object. Input must be an Application Lookup Object. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


