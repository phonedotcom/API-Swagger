# SwaggerClient::DevicesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_device**](DevicesApi.md#create_account_device) | **POST** /accounts/{account_id}/devices | Register a generic VoIP device
[**get_account_device**](DevicesApi.md#get_account_device) | **GET** /accounts/{account_id}/devices/{device_id} | Show details of an individual VoIP device
[**list_account_devices**](DevicesApi.md#list_account_devices) | **GET** /accounts/{account_id}/devices | Get a list of VoIP devices associated with your account
[**replace_account_device**](DevicesApi.md#replace_account_device) | **PUT** /accounts/{account_id}/devices/{device_id} | Update the settings for an individual VoIP device


# **create_account_device**
> DeviceFull create_account_device(account_id, , opts)

Register a generic VoIP device



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

api_instance = SwaggerClient::DevicesApi.new

account_id = 56 # Integer | Account ID

opts = { 
  data: SwaggerClient::CreateDeviceParams.new # CreateDeviceParams | Device data
}

begin
  #Register a generic VoIP device
  result = api_instance.create_account_device(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling DevicesApi->create_account_device: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **data** | [**CreateDeviceParams**](CreateDeviceParams.md)| Device data | [optional] 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_account_device**
> DeviceFull get_account_device(account_id, device_id)

Show details of an individual VoIP device



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

api_instance = SwaggerClient::DevicesApi.new

account_id = 56 # Integer | Account ID

device_id = 56 # Integer | Device ID


begin
  #Show details of an individual VoIP device
  result = api_instance.get_account_device(account_id, device_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling DevicesApi->get_account_device: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **device_id** | **Integer**| Device ID | 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_devices**
> ListDevices list_account_devices(account_id, , opts)

Get a list of VoIP devices associated with your account



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

api_instance = SwaggerClient::DevicesApi.new

account_id = 56 # Integer | Account ID

opts = { 
  filters_id: ["filters_id_example"], # Array<String> | ID filter
  filters_name: ["filters_name_example"], # Array<String> | Name filter
  sort_id: "sort_id_example", # String | ID sorting
  sort_name: "sort_name_example", # String | Name sorting
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
}

begin
  #Get a list of VoIP devices associated with your account
  result = api_instance.list_account_devices(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling DevicesApi->list_account_devices: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **filters_id** | [**Array&lt;String&gt;**](String.md)| ID filter | [optional] 
 **filters_name** | [**Array&lt;String&gt;**](String.md)| Name filter | [optional] 
 **sort_id** | **String**| ID sorting | [optional] 
 **sort_name** | **String**| Name sorting | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 

### Return type

[**ListDevices**](ListDevices.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **replace_account_device**
> DeviceFull replace_account_device(account_id, device_id, opts)

Update the settings for an individual VoIP device



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

api_instance = SwaggerClient::DevicesApi.new

account_id = 56 # Integer | Account ID

device_id = 56 # Integer | Device ID

opts = { 
  data: SwaggerClient::CreateDeviceParams.new # CreateDeviceParams | Device data
}

begin
  #Update the settings for an individual VoIP device
  result = api_instance.replace_account_device(account_id, device_id, opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling DevicesApi->replace_account_device: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **device_id** | **Integer**| Device ID | 
 **data** | [**CreateDeviceParams**](CreateDeviceParams.md)| Device data | [optional] 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



