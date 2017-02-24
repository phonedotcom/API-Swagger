
# ExtensionFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Integer** | ID of the extension. This is the internal Phone.com ID, not the extension number callers may dial. | 
**name** | **String** | User-supplied name for the extension. On POST, leaving this empty will result in an auto-generated value. On PUT, this field is required. | 
**extension** | **Integer** | Extension number that callers may dial. On POST, leaving this empty will result in an auto-generated value. On PUT, this field is required. | 
**fullName** | **String** | Full name of the individual or department to which this extension is assigned | 
**usageType** | **String** | Can be \&quot;limited\&quot; or \&quot;unlimited\&quot;. In most cases, changing this will affect your monthly bill. Please see our Control Panel or contact Customer Service for pricing. | 
**deviceMembership** | [**DeviceMembership**](DeviceMembership.md) |  |  [optional]
**timezone** | **String** | Time zone. Can be in any commonly recognized format, such as \&quot;America/Los_Angeles\&quot;. | 
**nameGreeting** | [**MediaSummary**](MediaSummary.md) | Greeting that communicates the extension&#39;s name. Output is a Greeting Summary Object. Input must be a Greeting Lookup Object. | 
**includeInDirectory** | **Boolean** | Whether this extension should be included in the dial-by-name directory for this account. Boolean. | 
**callerId** | **String** | Phone number to use as Caller ID for outgoing calls. Must be a phone number belonging to this account, or one of any additional authorized phone numbers. You can use our List Caller Ids service to see a current list. To unassign, you may set this to \&quot;private\&quot;, NULL, or an empty string. | 
**localAreaCode** | **String** | For outbound calls, this is the North American area code that this extension is calling from. | 
**enableCallWaiting** | **Boolean** | Whether Call Waiting is enabled. Boolean. Default is TRUE. | 
**enableOutboundCalls** | **Boolean** | Whether outgoing calls are enabled. Boolean. Default is TRUE. | 
**voicemail** | [**Voicemail**](Voicemail.md) |  | 
**callNotifications** | [**Notification**](Notification.md) | Call Notifications Object. See below for details. | 
**route** | [**RouteSummary**](RouteSummary.md) | Route which will handle incoming voice and fax calls. Only valid on PUT requests, not POST. Output is a Route Summary Object if the route is named, otherwise the Full Route Object will be shown. Input must be a Route Lookup Object pointing to a named route. Route must belong to this extension already. |  [optional]



