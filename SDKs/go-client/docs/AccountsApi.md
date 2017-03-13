# \AccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /accounts/{account_id} | Retrieve details of an individual account
[**ListAccounts**](AccountsApi.md#ListAccounts) | **Get** /accounts | Get a list of accounts visible to the authenticated user or client


# **GetAccount**
> AccountFull GetAccount($accountId)

Retrieve details of an individual account

This service shows the details of an individual account. See Accounts for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 

### Return type

[**AccountFull**](AccountFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccounts**
> ListAccounts ListAccounts($filtersId, $sortId, $limit, $offset, $fields)

Get a list of accounts visible to the authenticated user or client

This service lists the accounts accessible to the authenticated client. In most cases, there will only be one such account. See Accounts for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListAccounts**](ListAccounts.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

