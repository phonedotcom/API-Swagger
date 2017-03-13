# SwaggerClient::ExpressservicecodesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_account_express_srv_code**](ExpressservicecodesApi.md#get_account_express_srv_code) | **GET** /accounts/{account_id}/express-service-codes/{code_id} | Show details of an account Express Service Code
[**list_account_express_srv_codes**](ExpressservicecodesApi.md#list_account_express_srv_codes) | **GET** /accounts/{account_id}/express-service-codes | Get the Express Service Code associated with your account in list format


# **get_account_express_srv_code**
> ExpressServiceCodeFull get_account_express_srv_code(account_id, code_id)

Show details of an account Express Service Code

This service shows the details of an Account Express Service Code.

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

api_instance = SwaggerClient::ExpressservicecodesApi.new

account_id = 56 # Integer | Account ID

code_id = 56 # Integer | Device ID


begin
  #Show details of an account Express Service Code
  result = api_instance.get_account_express_srv_code(account_id, code_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ExpressservicecodesApi->get_account_express_srv_code: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **code_id** | **Integer**| Device ID | 

### Return type

[**ExpressServiceCodeFull**](ExpressServiceCodeFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_express_srv_codes**
> ListExpressServiceCodes list_account_express_srv_codes(account_id, , opts)

Get the Express Service Code associated with your account in list format

See Express Service Codes for more detail.

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

api_instance = SwaggerClient::ExpressservicecodesApi.new

account_id = 56 # Integer | Account ID

opts = { 
  filters_id: ["filters_id_example"], # Array<String> | ID filter
}

begin
  #Get the Express Service Code associated with your account in list format
  result = api_instance.list_account_express_srv_codes(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ExpressservicecodesApi->list_account_express_srv_codes: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **filters_id** | [**Array&lt;String&gt;**](String.md)| ID filter | [optional] 

### Return type

[**ListExpressServiceCodes**](ListExpressServiceCodes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



