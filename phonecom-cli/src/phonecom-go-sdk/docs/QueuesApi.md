# \QueuesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountQueue**](QueuesApi.md#CreateAccountQueue) | **Post** /accounts/{account_id}/queues | Create a queue
[**DeleteAccountQueue**](QueuesApi.md#DeleteAccountQueue) | **Delete** /accounts/{account_id}/queues/{queue_id} | Delete a queue
[**GetAccountQueue**](QueuesApi.md#GetAccountQueue) | **Get** /accounts/{account_id}/queues/{queue_id} | Show details of an individual queue
[**ListAccountQueues**](QueuesApi.md#ListAccountQueues) | **Get** /accounts/{account_id}/queues | Get a list of queues for an account
[**ReplaceAccountQueue**](QueuesApi.md#ReplaceAccountQueue) | **Put** /accounts/{account_id}/queues/{queue_id} | Replace a queue


# **CreateAccountQueue**
> QueueFull CreateAccountQueue($accountId, $data)

Create a queue

For more on the input fields, see Account Queues.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional] 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountQueue**
> DeleteQueue DeleteAccountQueue($accountId, $queueId)

Delete a queue

This service a queue from the account. For more information on queue properties, see Account Queues.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **queueId** | **int32**| Queue ID | 

### Return type

[**DeleteQueue**](DeleteQueue.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountQueue**
> QueueFull GetAccountQueue($accountId, $queueId)

Show details of an individual queue

This service shows the details of an individual queue. For more on the input fields, see Account Queues.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **queueId** | **int32**| Queue ID | 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountQueues**
> ListQueues ListAccountQueues($accountId, $filtersId, $filtersName, $sortId, $sortName, $limit, $offset, $fields)

Get a list of queues for an account

The List Queues service lists all the queues belong to the account. See Account Queues for more info on the properties.


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

[**ListQueues**](ListQueues.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountQueue**
> QueueFull ReplaceAccountQueue($accountId, $queueId, $data)

Replace a queue

The Replace Queue service replaces the parameters of a queue. For more on the input fields, see Account Queues.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **queueId** | **int32**| Queue ID | 
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional] 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

