# \MediaApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountMedia**](MediaApi.md#GetAccountMedia) | **Get** /accounts/{account_id}/media/{recording_id} | Show details of an individual media recording (Greeting or Hold Music)
[**ListAccountMedia**](MediaApi.md#ListAccountMedia) | **Get** /accounts/{account_id}/media | Get a list of media recordings for an account


# **GetAccountMedia**
> MediaFull GetAccountMedia($accountId, $recordingId)

Show details of an individual media recording (Greeting or Hold Music)

Get individual media recording


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **recordingId** | **int32**| Recording ID | 

### Return type

[**MediaFull**](MediaFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountMedia**
> ListMedia ListAccountMedia($accountId, $filtersId, $filtersName, $sortId, $sortName, $limit, $offset, $fields)

Get a list of media recordings for an account

See Account Menus for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersName** | [**[]string**](string.md)| Name filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListMedia**](ListMedia.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

