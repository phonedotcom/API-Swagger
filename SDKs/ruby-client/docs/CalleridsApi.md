# SwaggerClient::CalleridsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_caller_ids**](CalleridsApi.md#get_caller_ids) | **GET** /accounts/{account_id}/extensions/{extension_id}/caller-ids | Show the Caller ID options a given extension can use


# **get_caller_ids**
> ListCallerIds get_caller_ids(account_id, extension_id, , opts)

Show the Caller ID options a given extension can use

Get Caller ID

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

api_instance = SwaggerClient::CalleridsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

opts = { 
  filters_number: ["filters_number_example"], # Array<String> | Number filter
  filters_name: ["filters_name_example"], # Array<String> | Name filter
  sort_number: "sort_number_example", # String | Number sorting
  sort_name: "sort_name_example", # String | Name sorting
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
}

begin
  #Show the Caller ID options a given extension can use
  result = api_instance.get_caller_ids(account_id, extension_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling CalleridsApi->get_caller_ids: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **filters_number** | [**Array&lt;String&gt;**](String.md)| Number filter | [optional] 
 **filters_name** | [**Array&lt;String&gt;**](String.md)| Name filter | [optional] 
 **sort_number** | **String**| Number sorting | [optional] 
 **sort_name** | **String**| Name sorting | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 

### Return type

[**ListCallerIds**](ListCallerIds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



