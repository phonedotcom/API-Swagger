# SwaggerClient::CallsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_calls**](CallsApi.md#create_account_calls) | **POST** /accounts/{account_id}/calls | Make a phone call


# **create_account_calls**
> CallFull create_account_calls(account_id, , opts)

Make a phone call



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

api_instance = SwaggerClient::CallsApi.new

account_id = 56 # Integer | Account ID

opts = { 
  data: SwaggerClient::CreateCallParams.new # CreateCallParams | Call data
}

begin
  #Make a phone call
  result = api_instance.create_account_calls(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling CallsApi->create_account_calls: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **data** | [**CreateCallParams**](CreateCallParams.md)| Call data | [optional] 

### Return type

[**CallFull**](CallFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



