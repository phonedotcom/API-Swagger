# SwaggerClient::DefaultApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ping**](DefaultApi.md#ping) | **GET** /ping | The default API command


# **ping**
> PingResponse ping

The default API command

The default API command

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

api_instance = SwaggerClient::DefaultApi.new

begin
  #The default API command
  result = api_instance.ping
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling DefaultApi->ping: #{e}"
end
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



