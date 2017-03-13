# Greeting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **string** | The greeting to play. Can be \&quot;name\&quot; for the name_greeting as described above, \&quot;standard\&quot; for the standard greeting, or \&quot;alternate\&quot; for an alternate greeting. See below for details. | [optional] 
**alternate** | [**\Swagger\Client\Model\MediaSummary**](MediaSummary.md) | Greeting to be played when type&#x3D;\&quot;alternate\&quot;. Output is a Greeting Summary Object. Input must be a Greeting Lookup Object. | [optional] 
**standard** | [**\Swagger\Client\Model\MediaSummary**](MediaSummary.md) | Greeting to be played when type&#x3D;\&quot;standard\&quot;. Output is a Greeting Summary Object. Input must be a Greeting Lookup Object. | [optional] 
**enable_leave_message_prompt** | **bool** | Whether to prompt the caller with the following words after the voicemail greeting has been played: \&quot;Please leave your message after the tone. When finished, hang up or press the pound key.\&quot; Boolean. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


