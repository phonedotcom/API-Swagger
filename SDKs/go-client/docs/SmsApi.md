# \SmsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountSms**](SmsApi.md#CreateAccountSms) | **Post** /accounts/{account_id}/sms | Send a SMS to one or a group of recipients
[**GetAccountSms**](SmsApi.md#GetAccountSms) | **Get** /accounts/{account_id}/sms/{sms_id} | Show details of an individual SMS
[**ListAccountSms**](SmsApi.md#ListAccountSms) | **Get** /accounts/{account_id}/sms | Get a list of SMS messages for an account


# **CreateAccountSms**
> SmsFull CreateAccountSms($accountId, $data)

Send a SMS to one or a group of recipients

For more on the input fields, see Intro to SMS.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateSmsParams**](CreateSmsParams.md)| SMS data | 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountSms**
> SmsFull GetAccountSms($accountId, $smsId)

Show details of an individual SMS

This service shows the details of an individual sms. See Intro to SMS for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **smsId** | **string**| SMS ID | 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountSms**
> ListSms ListAccountSms($accountId, $filtersId, $filtersDirection, $filtersFrom, $sortId, $sortCreatedAt, $limit, $offset, $fields)

Get a list of SMS messages for an account

See Intro to SMS for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersDirection** | **string**| Direction filter | [optional] 
 **filtersFrom** | **string**| Caller ID filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortCreatedAt** | **string**| Sort by created time of message | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListSms**](ListSms.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

