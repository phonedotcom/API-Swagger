# \CalleridsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCallerIds**](CalleridsApi.md#GetCallerIds) | **Get** /accounts/{account_id}/extensions/{extension_id}/caller-ids | Show the Caller ID options a given extension can use


# **GetCallerIds**
> ListCallerIds GetCallerIds($accountId, $extensionId, $filtersNumber, $filtersName, $sortNumber, $sortName, $limit, $offset, $fields)

Show the Caller ID options a given extension can use

Get Caller ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **filtersNumber** | [**[]string**](string.md)| Number filter | [optional] 
 **filtersName** | [**[]string**](string.md)| Name filter | [optional] 
 **sortNumber** | **string**| Number sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListCallerIds**](ListCallerIds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

