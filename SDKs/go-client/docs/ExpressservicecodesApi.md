# \ExpressservicecodesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountExpressSrvCode**](ExpressservicecodesApi.md#GetAccountExpressSrvCode) | **Get** /accounts/{account_id}/express-service-codes/{code_id} | Show details of an account Express Service Code
[**ListAccountExpressSrvCodes**](ExpressservicecodesApi.md#ListAccountExpressSrvCodes) | **Get** /accounts/{account_id}/express-service-codes | Get the Express Service Code associated with your account in list format


# **GetAccountExpressSrvCode**
> ExpressServiceCodeFull GetAccountExpressSrvCode($accountId, $codeId)

Show details of an account Express Service Code

This service shows the details of an Account Express Service Code.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **codeId** | **int32**| Device ID | 

### Return type

[**ExpressServiceCodeFull**](ExpressServiceCodeFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountExpressSrvCodes**
> ListExpressServiceCodes ListAccountExpressSrvCodes($accountId, $filtersId)

Get the Express Service Code associated with your account in list format

See Express Service Codes for more detail.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 

### Return type

[**ListExpressServiceCodes**](ListExpressServiceCodes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

