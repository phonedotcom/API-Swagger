# \GroupsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountExtensionContactGroup**](GroupsApi.md#CreateAccountExtensionContactGroup) | **Post** /accounts/{account_id}/extensions/{extension_id}/contact-groups | 
[**DeleteAccountExtensionContactGroup**](GroupsApi.md#DeleteAccountExtensionContactGroup) | **Delete** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | Delete an addressbook group
[**GetAccountExtensionContactGroup**](GroupsApi.md#GetAccountExtensionContactGroup) | **Get** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 
[**ListAccountExtensionContactGroups**](GroupsApi.md#ListAccountExtensionContactGroups) | **Get** /accounts/{account_id}/extensions/{extension_id}/contact-groups | Show a list of contact groups belonging to an extension
[**ReplaceAccountExtensionContactGroup**](GroupsApi.md#ReplaceAccountExtensionContactGroup) | **Put** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 


# **CreateAccountExtensionContactGroup**
> GroupFull CreateAccountExtensionContactGroup($accountId, $extensionId, $data)



See Account Contact Groups for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **data** | [**CreateGroupParams**](CreateGroupParams.md)| Group name | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountExtensionContactGroup**
> DeleteGroup DeleteAccountExtensionContactGroup($accountId, $extensionId, $groupId)

Delete an addressbook group




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **groupId** | **int32**| Group ID | 

### Return type

[**DeleteGroup**](DeleteGroup.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountExtensionContactGroup**
> GroupFull GetAccountExtensionContactGroup($accountId, $extensionId, $groupId)



See Account Contact Groups for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **groupId** | **int32**| Group ID | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountExtensionContactGroups**
> ListGroups ListAccountExtensionContactGroups($accountId, $extensionId, $filtersId, $filtersName, $sortId, $sortName, $limit, $offset, $fields)

Show a list of contact groups belonging to an extension

See Account Contact Groups for details on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersName** | [**[]string**](string.md)| Name filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListGroups**](ListGroups.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountExtensionContactGroup**
> GroupFull ReplaceAccountExtensionContactGroup($accountId, $extensionId, $groupId, $data)



See Account Contact Groups for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **groupId** | **int32**| Group ID | 
 **data** | [**CreateGroupParams**](CreateGroupParams.md)| Group name | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

