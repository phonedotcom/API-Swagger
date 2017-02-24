# \PhonenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountPhoneNumber**](PhonenumbersApi.md#CreateAccountPhoneNumber) | **Post** /accounts/{account_id}/phone-numbers | Add a phone number to an account
[**GetAccountPhoneNumber**](PhonenumbersApi.md#GetAccountPhoneNumber) | **Get** /accounts/{account_id}/phone-numbers/{number_id} | Show details of an individual phone number
[**ListAccountPhoneNumbers**](PhonenumbersApi.md#ListAccountPhoneNumbers) | **Get** /accounts/{account_id}/phone-numbers | Get a list of phone numbers registered to an account
[**ReplaceAccountPhoneNumber**](PhonenumbersApi.md#ReplaceAccountPhoneNumber) | **Put** /accounts/{account_id}/phone-numbers/{number_id} | Update the settings for an existing phone number on your account


# **CreateAccountPhoneNumber**
> PhoneNumberFull CreateAccountPhoneNumber($accountId, $data)

Add a phone number to an account

See Intro to Account Phone Numbers for more info on the properties to use.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreatePhoneNumberParams**](CreatePhoneNumberParams.md)| Phone Number data | [optional] 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountPhoneNumber**
> PhoneNumberFull GetAccountPhoneNumber($accountId, $numberId)

Show details of an individual phone number

See Intro to Account Phone Numbers for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **numberId** | **int32**| Number ID | 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountPhoneNumbers**
> ListPhoneNumbers ListAccountPhoneNumbers($accountId, $filtersId, $filtersName, $filtersPhoneNumber, $sortId, $sortName, $sortPhoneNumber, $limit, $offset, $fields)

Get a list of phone numbers registered to an account

See Intro to Account Phone Numbers for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersName** | [**[]string**](string.md)| Name filter | [optional] 
 **filtersPhoneNumber** | [**[]string**](string.md)| Phone number filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **sortPhoneNumber** | **string**| Phone number sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListPhoneNumbers**](ListPhoneNumbers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountPhoneNumber**
> PhoneNumberFull ReplaceAccountPhoneNumber($accountId, $numberId, $data)

Update the settings for an existing phone number on your account

See Intro to Account Phone Numbers for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **numberId** | **int32**| Number ID | 
 **data** | [**ReplacePhoneNumberParams**](ReplacePhoneNumberParams.md)| Phone Number data | [optional] 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

