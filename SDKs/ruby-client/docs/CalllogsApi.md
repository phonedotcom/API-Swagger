# SwaggerClient::CalllogsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list_account_call_logs**](CalllogsApi.md#list_account_call_logs) | **GET** /accounts/{account_id}/call-logs | Get a list of call details associated with your account


# **list_account_call_logs**
> ListCallLogs list_account_call_logs(account_id, , opts)

Get a list of call details associated with your account

See Call Logs for more detail.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure API key authorization: apiKey
  config.api_key['Authorization'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  #config.api_key_prefix['Authorization'] = 'Bearer'
end

api_instance = SwaggerClient::CalllogsApi.new

account_id = 56 # Integer | Account ID

opts = { 
  filters_id: ["filters_id_example"], # Array<String> | ID filter
  filters_start_time: ["filters_start_time_example"], # Array<String> | Call start time filter
  filters_created_at: "filters_created_at_example", # String | Call log creation time filter
  filters_direction: "filters_direction_example", # String | Call direction filter: in or out
  filters_called_number: "filters_called_number_example", # String | Called number
  filters_type: "filters_type_example", # String | Call type, such as 'call', 'fax', 'audiogram'
  filters_extension: ["filters_extension_example"], # Array<String> | Extension filter
  sort_id: "sort_id_example", # String | ID sorting
  sort_start_time: "sort_start_time_example", # String | Sorting by call start time: asc or desc
  sort_created_at: "sort_created_at_example", # String | Sorting by call log creation time: asc or desc
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
}

begin
  #Get a list of call details associated with your account
  result = api_instance.list_account_call_logs(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling CalllogsApi->list_account_call_logs: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **filters_id** | [**Array&lt;String&gt;**](String.md)| ID filter | [optional] 
 **filters_start_time** | [**Array&lt;String&gt;**](String.md)| Call start time filter | [optional] 
 **filters_created_at** | **String**| Call log creation time filter | [optional] 
 **filters_direction** | **String**| Call direction filter: in or out | [optional] 
 **filters_called_number** | **String**| Called number | [optional] 
 **filters_type** | **String**| Call type, such as &#39;call&#39;, &#39;fax&#39;, &#39;audiogram&#39; | [optional] 
 **filters_extension** | [**Array&lt;String&gt;**](String.md)| Extension filter | [optional] 
 **sort_id** | **String**| ID sorting | [optional] 
 **sort_start_time** | **String**| Sorting by call start time: asc or desc | [optional] 
 **sort_created_at** | **String**| Sorting by call log creation time: asc or desc | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 

### Return type

[**ListCallLogs**](ListCallLogs.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



