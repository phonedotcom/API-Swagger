# ExtensionFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the extension. This is the internal Phone.com ID, not the extension number callers may dial. | [optional] [default to null]
**Name** | **string** | User-supplied name for the extension. On POST, leaving this empty will result in an auto-generated value. On PUT, this field is required. | [optional] [default to null]
**Extension** | **int32** | Extension number that callers may dial. On POST, leaving this empty will result in an auto-generated value. On PUT, this field is required. | [optional] [default to null]
**FullName** | **string** | Full name of the individual or department to which this extension is assigned | [optional] [default to null]
**UsageType** | **string** | Can be \&quot;limited\&quot; or \&quot;unlimited\&quot;. In most cases, changing this will affect your monthly bill. Please see our Control Panel or contact Customer Service for pricing. | [optional] [default to null]
**DeviceMembership** | [**DeviceMembership**](DeviceMembership.md) |  | [optional] [default to null]
**Timezone** | **string** | Time zone. Can be in any commonly recognized format, such as \&quot;America/Los_Angeles\&quot;. | [optional] [default to null]
**NameGreeting** | [**MediaSummary**](MediaSummary.md) | Greeting that communicates the extension&#39;s name. Output is a Greeting Summary Object. Input must be a Greeting Lookup Object. | [optional] [default to null]
**IncludeInDirectory** | **bool** | Whether this extension should be included in the dial-by-name directory for this account. Boolean. | [optional] [default to null]
**CallerId** | **string** | Phone number to use as Caller ID for outgoing calls. Must be a phone number belonging to this account, or one of any additional authorized phone numbers. You can use our List Caller Ids service to see a current list. To unassign, you may set this to \&quot;private\&quot;, NULL, or an empty string. | [optional] [default to null]
**LocalAreaCode** | **string** | For outbound calls, this is the North American area code that this extension is calling from. | [optional] [default to null]
**EnableCallWaiting** | **bool** | Whether Call Waiting is enabled. Boolean. Default is TRUE. | [optional] [default to null]
**EnableOutboundCalls** | **bool** | Whether outgoing calls are enabled. Boolean. Default is TRUE. | [optional] [default to null]
**Voicemail** | [**Voicemail**](Voicemail.md) |  | [optional] [default to null]
**CallNotifications** | [**Notification**](Notification.md) | Call Notifications Object. See below for details. | [optional] [default to null]
**Route** | [**RouteSummary**](RouteSummary.md) | Route which will handle incoming voice and fax calls. Only valid on PUT requests, not POST. Output is a Route Summary Object if the route is named, otherwise the Full Route Object will be shown. Input must be a Route Lookup Object pointing to a named route. Route must belong to this extension already. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


