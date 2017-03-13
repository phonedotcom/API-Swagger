# \SubaccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountSubaccount**](SubaccountsApi.md#CreateAccountSubaccount) | **Post** /accounts/{account_id}/subaccounts | Add a subaccount for the authenticated user or client
[**ListAccountSubaccounts**](SubaccountsApi.md#ListAccountSubaccounts) | **Get** /accounts/{account_id}/subaccounts | Get a list of subaccounts for the authenticated user or client


# **CreateAccountSubaccount**
> AccountFull CreateAccountSubaccount($accountId, $data)

Add a subaccount for the authenticated user or client

This service shows the details of an individual Subaccount.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateSubaccountParams**](CreateSubaccountParams.md)| SMS data | 

### Return type

[**AccountFull**](AccountFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountSubaccounts**
> ListAccounts ListAccountSubaccounts($accountId, $filtersId, $sortId, $limit, $offset, $fields)

Get a list of subaccounts for the authenticated user or client

This service lists the Subaccount of the authenticated client. In most cases, there will not be any.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
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

