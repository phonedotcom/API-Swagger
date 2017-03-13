# Voicemail

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**enabled** | **bool** | Whether voicemail is enabled. Boolean. | [optional] 
**password** | **str** | Password for accessing voicemail box. Must be digits only. | [optional] 
**greeting** | [**Greeting**](Greeting.md) |  | [optional] 
**attachments** | **str** | If notification emails are being used, this defines the format of the audio attachments. Can be \&quot;wav\&quot; for WAV format, \&quot;mp3\&quot; for MP3 format, or NULL to disable attachments. | [optional] 
**notifications** | [**Notification**](Notification.md) | Voicemail Notifications Object. See below for details. Can be set to NULL to disable notifications. | [optional] 
**transcription** | **str** | Type of voicemail transcription to use. Can be \&quot;human\&quot; for high-quality manual transcriptions by human operators, \&quot;automated\&quot; for machine-generated transcriptions, or NULL to omit trancriptions. Changing this option will affect your monthly bill. Please see our Control Panel or contact Customer Service for details. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


