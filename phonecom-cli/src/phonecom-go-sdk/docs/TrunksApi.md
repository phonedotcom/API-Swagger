# \TrunksApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountTrunk**](TrunksApi.md#CreateAccountTrunk) | **Post** /accounts/{account_id}/trunks | Add a trunk record with SIP information
[**DeleteAccountTrunk**](TrunksApi.md#DeleteAccountTrunk) | **Delete** /accounts/{account_id}/trunks/{trunk_id} | Delete a trunk from account
[**GetAccountTrunk**](TrunksApi.md#GetAccountTrunk) | **Get** /accounts/{account_id}/trunks/{trunk_id} | Show details of an individual trunk
[**ListAccountTrunks**](TrunksApi.md#ListAccountTrunks) | **Get** /accounts/{account_id}/trunks | Get a list of trunks for an account
[**ReplaceAccountTrunk**](TrunksApi.md#ReplaceAccountTrunk) | **Put** /accounts/{account_id}/trunks/{trunk_id} | Replace parameters in a trunk


# **CreateAccountTrunk**
> TrunkFull CreateAccountTrunk($accountId, $data)

Add a trunk record with SIP information

For more on the input fields, see Account Trunks.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateTrunkParams**](CreateTrunkParams.md)| Trunk data | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountTrunk**
> DeleteTrunk DeleteAccountTrunk($accountId, $trunkId)

Delete a trunk from account

This service deletes a trunk from the account. For more on the properties of trunks, see Account Trunks.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **trunkId** | **int32**| Trunk ID | 

### Return type

[**DeleteTrunk**](DeleteTrunk.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountTrunk**
> TrunkFull GetAccountTrunk($accountId, $trunkId)

Show details of an individual trunk

This service shows the details of an individual Trunk.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **trunkId** | **int32**| Trunk ID | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountTrunks**
> ListTrunks ListAccountTrunks($accountId, $filtersId, $filtersName, $sortId, $sortName, $limit, $offset, $fields)

Get a list of trunks for an account

See Account Trunks for more info on the properties.


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

[**ListTrunks**](ListTrunks.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountTrunk**
> TrunkFull ReplaceAccountTrunk($accountId, $trunkId, $data)

Replace parameters in a trunk

For more on the input fields, see Account Trunks.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **trunkId** | **int32**| Trunk ID | 
 **data** | [**CreateTrunkParams**](CreateTrunkParams.md)| Trunk data | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

