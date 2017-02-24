# \ExtensionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountExtension**](ExtensionsApi.md#CreateAccountExtension) | **Post** /accounts/{account_id}/extensions | Create an individual extension
[**GetAccountExtension**](ExtensionsApi.md#GetAccountExtension) | **Get** /accounts/{account_id}/extensions/{extension_id} | Show details of an individual extension
[**ListAccountExtensions**](ExtensionsApi.md#ListAccountExtensions) | **Get** /accounts/{account_id}/extensions | Get a list of extensions visible to the authenticated user or client
[**ReplaceAccountExtension**](ExtensionsApi.md#ReplaceAccountExtension) | **Put** /accounts/{account_id}/extensions/{extension_id} | Replace an individual extension


# **CreateAccountExtension**
> ExtensionFull CreateAccountExtension($accountId, $data)

Create an individual extension

This service shows how to create a virtual extension.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateExtensionParams**](CreateExtensionParams.md)| Account Extensions Data | [optional] 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountExtension**
> ExtensionFull GetAccountExtension($accountId, $extensionId)

Show details of an individual extension

This service shows the details of an individual Extension.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountExtensions**
> ListExtensions ListAccountExtensions($accountId, $filtersId, $filtersExtension, $filtersName, $sortId, $sortExtension, $sortName, $limit, $offset, $fields)

Get a list of extensions visible to the authenticated user or client

This service lists the visible extensions on a given account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersExtension** | [**[]string**](string.md)| Extension filter | [optional] 
 **filtersName** | [**[]string**](string.md)| Name filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortExtension** | **string**| Extension sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListExtensions**](ListExtensions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountExtension**
> ExtensionFull ReplaceAccountExtension($accountId, $extensionId, $data)

Replace an individual extension

This service shows how to update an individual extension.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **data** | [**ReplaceExtensionParams**](ReplaceExtensionParams.md)| Account Extensions Data | [optional] 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

