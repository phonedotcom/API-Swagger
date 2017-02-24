# MenuFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Integer Menu ID. Read-only. | [optional] [default to null]
**Name** | **string** | Name. Required. Unique. | [optional] [default to null]
**AllowExtensionDial** | **bool** | Boolean. Determines whether a caller can enter an extension number to bypass the menu. | [optional] [default to null]
**KeypressWaitTime** | **int32** | Boolean. Determines whether a caller can enter an extension number to bypass the menu. | [optional] [default to null]
**Greeting** | [**MediaSummary**](MediaSummary.md) | Greeting that is played when a caller enters a menu. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. | [optional] [default to null]
**KeypressError** | [**MediaSummary**](MediaSummary.md) | Message that is played when the caller makes a keypress error. Output is a Media Summary Object. Input must be a Media Lookup Object. Must refer to a media recording that has is_hold_music set to FALSE. | [optional] [default to null]
**TimeoutHandler** | [**RouteSummary**](RouteSummary.md) | Route that will be entered when the caller fails to choose a menu option within the allotted time. Output is a Route Summary Object if the route is named, otherwise the Full Route Object will be shown. Input must be a Route Lookup Object pointing to a named route. | [optional] [default to null]
**Options** | [**[]Option**](Option.md) | Array of menu option objects. See below for details. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


