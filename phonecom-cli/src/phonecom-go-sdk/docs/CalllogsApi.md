# \CalllogsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountCallLogs**](CalllogsApi.md#ListAccountCallLogs) | **Get** /accounts/{account_id}/call-logs | Get a list of call details associated with your account


# **ListAccountCallLogs**
> ListCallLogs ListAccountCallLogs($accountId, $filtersId, $filtersStartTime, $filtersCreatedAt, $filtersDirection, $filtersCalledNumber, $filtersType, $sortId, $sortStartTime, $sortCreatedAt, $limit, $offset, $fields)

Get a list of call details associated with your account

See Call Logs for more detail.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersStartTime** | [**[]string**](string.md)| Call start time filter | [optional] 
 **filtersCreatedAt** | **string**| Call log creation time filter | [optional] 
 **filtersDirection** | **string**| Call direction filter: in or out | [optional] 
 **filtersCalledNumber** | **string**| Called number | [optional] 
 **filtersType** | **string**| Call type, such as &#39;call&#39;, &#39;fax&#39;, &#39;audiogram&#39; | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortStartTime** | **string**| Sorting by call start time: asc or desc | [optional] 
 **sortCreatedAt** | **string**| Sorting by call log creation time: asc or desc | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListCallLogs**](ListCallLogs.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

