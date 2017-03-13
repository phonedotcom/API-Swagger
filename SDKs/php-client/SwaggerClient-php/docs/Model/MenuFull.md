# MenuFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Integer Menu ID. Read-only. | [optional] 
**name** | **string** | Name. Required. Unique. | [optional] 
**allow_extension_dial** | **bool** | Boolean. Determines whether a caller can enter an extension number to bypass the menu. | [optional] 
**keypress_wait_time** | **int** | Boolean. Determines whether a caller can enter an extension number to bypass the menu. | [optional] 
**greeting** | [**\Swagger\Client\Model\MediaSummary**](MediaSummary.md) | Greeting that is played when a caller enters a menu. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. | [optional] 
**keypress_error** | [**\Swagger\Client\Model\MediaSummary**](MediaSummary.md) | Message that is played when the caller makes a keypress error. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. | [optional] 
**timeout_handler** | [**\Swagger\Client\Model\RouteSummary**](RouteSummary.md) | Route that will be entered when the caller fails to choose a menu option within the allotted time. Output is a Route Summary Object if the route is named, otherwise the Full Route Object will be shown. Input must be a Route Lookup Object pointing to a named route. | [optional] 
**options** | [**\Swagger\Client\Model\Option[]**](Option.md) | Array of menu option objects. See below for details. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


