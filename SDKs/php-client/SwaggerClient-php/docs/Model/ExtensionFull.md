# ExtensionFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | ID of the extension. This is the internal Phone.com ID, not the extension number callers may dial. | [optional] 
**name** | **string** | User-supplied name for the extension. On POST, leaving this empty will result in an auto-generated value. On PUT, this field is required. | [optional] 
**extension** | **int** | Extension number that callers may dial. On POST, leaving this empty will result in an auto-generated value. On PUT, this field is required. | [optional] 
**full_name** | **string** | Full name of the individual or department to which this extension is assigned | [optional] 
**usage_type** | **string** | Can be \&quot;limited\&quot; or \&quot;unlimited\&quot;. In most cases, changing this will affect your monthly bill. Please see our Control Panel or contact Customer Service for pricing. | [optional] 
**device_membership** | [**\Swagger\Client\Model\DeviceMembership**](DeviceMembership.md) |  | [optional] 
**timezone** | **string** | Time zone. Can be in any commonly recognized format, such as \&quot;America/Los_Angeles\&quot;. | [optional] 
**name_greeting** | [**\Swagger\Client\Model\MediaSummary**](MediaSummary.md) | Greeting that communicates the extension&#39;s name. Output is a Greeting Summary Object. Input must be a Greeting Lookup Object. | [optional] 
**include_in_directory** | **bool** | Whether this extension should be included in the dial-by-name directory for this account. Boolean. | [optional] 
**caller_id** | **string** | Phone number to use as Caller ID for outgoing calls. Must be a phone number belonging to this account, or one of any additional authorized phone numbers. You can use our List Caller Ids service to see a current list. To unassign, you may set this to \&quot;private\&quot;, NULL, or an empty string. | [optional] 
**local_area_code** | **string** | For outbound calls, this is the North American area code that this extension is calling from. | [optional] 
**enable_call_waiting** | **bool** | Whether Call Waiting is enabled. Boolean. Default is TRUE. | [optional] 
**enable_outbound_calls** | **bool** | Whether outgoing calls are enabled. Boolean. Default is TRUE. | [optional] 
**voicemail** | [**\Swagger\Client\Model\Voicemail**](Voicemail.md) |  | [optional] 
**call_notifications** | [**\Swagger\Client\Model\Notification**](Notification.md) | Call Notifications Object. See below for details. | [optional] 
**route** | [**\Swagger\Client\Model\RouteSummary**](RouteSummary.md) | Route which will handle incoming voice and fax calls. Only valid on PUT requests, not POST. Output is a Route Summary Object if the route is named, otherwise the Full Route Object will be shown. Input must be a Route Lookup Object pointing to a named route. Route must belong to this extension already. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


