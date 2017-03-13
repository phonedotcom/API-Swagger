# \CallsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountCalls**](CallsApi.md#CreateAccountCalls) | **Post** /accounts/{account_id}/calls | Make a phone call


# **CreateAccountCalls**
> CallFull CreateAccountCalls($accountId, $data)

Make a phone call




### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateCallParams**](CreateCallParams.md)| Call data | [optional] 

### Return type

[**CallFull**](CallFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

