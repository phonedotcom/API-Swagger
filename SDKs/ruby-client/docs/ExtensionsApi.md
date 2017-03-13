# SwaggerClient::ExtensionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_extension**](ExtensionsApi.md#create_account_extension) | **POST** /accounts/{account_id}/extensions | Create an individual extension
[**get_account_extension**](ExtensionsApi.md#get_account_extension) | **GET** /accounts/{account_id}/extensions/{extension_id} | Show details of an individual extension
[**list_account_extensions**](ExtensionsApi.md#list_account_extensions) | **GET** /accounts/{account_id}/extensions | Get a list of extensions visible to the authenticated user or client
[**replace_account_extension**](ExtensionsApi.md#replace_account_extension) | **PUT** /accounts/{account_id}/extensions/{extension_id} | Replace an individual extension


# **create_account_extension**
> ExtensionFull create_account_extension(account_id, , opts)

Create an individual extension

This service shows how to create a virtual extension.

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

api_instance = SwaggerClient::ExtensionsApi.new

account_id = 56 # Integer | Account ID

opts = { 
  data: SwaggerClient::CreateExtensionParams.new # CreateExtensionParams | Account Extensions Data
}

begin
  #Create an individual extension
  result = api_instance.create_account_extension(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ExtensionsApi->create_account_extension: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **data** | [**CreateExtensionParams**](CreateExtensionParams.md)| Account Extensions Data | [optional] 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_account_extension**
> ExtensionFull get_account_extension(account_id, extension_id, )

Show details of an individual extension

This service shows the details of an individual Extension.

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

api_instance = SwaggerClient::ExtensionsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID


begin
  #Show details of an individual extension
  result = api_instance.get_account_extension(account_id, extension_id, )
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ExtensionsApi->get_account_extension: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_extensions**
> ListExtensions list_account_extensions(account_id, , opts)

Get a list of extensions visible to the authenticated user or client

This service lists the visible extensions on a given account.

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

api_instance = SwaggerClient::ExtensionsApi.new

account_id = 56 # Integer | Account ID

opts = { 
  filters_id: ["filters_id_example"], # Array<String> | ID filter
  filters_extension: ["filters_extension_example"], # Array<String> | Extension filter
  filters_name: ["filters_name_example"], # Array<String> | Name filter
  sort_id: "sort_id_example", # String | ID sorting
  sort_extension: "sort_extension_example", # String | Extension sorting
  sort_name: "sort_name_example", # String | Name sorting
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
}

begin
  #Get a list of extensions visible to the authenticated user or client
  result = api_instance.list_account_extensions(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ExtensionsApi->list_account_extensions: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **filters_id** | [**Array&lt;String&gt;**](String.md)| ID filter | [optional] 
 **filters_extension** | [**Array&lt;String&gt;**](String.md)| Extension filter | [optional] 
 **filters_name** | [**Array&lt;String&gt;**](String.md)| Name filter | [optional] 
 **sort_id** | **String**| ID sorting | [optional] 
 **sort_extension** | **String**| Extension sorting | [optional] 
 **sort_name** | **String**| Name sorting | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 

### Return type

[**ListExtensions**](ListExtensions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **replace_account_extension**
> ExtensionFull replace_account_extension(account_id, extension_id, , opts)

Replace an individual extension

This service shows how to update an individual extension.

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

api_instance = SwaggerClient::ExtensionsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

opts = { 
  data: SwaggerClient::ReplaceExtensionParams.new # ReplaceExtensionParams | Account Extensions Data
}

begin
  #Replace an individual extension
  result = api_instance.replace_account_extension(account_id, extension_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ExtensionsApi->replace_account_extension: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **data** | [**ReplaceExtensionParams**](ReplaceExtensionParams.md)| Account Extensions Data | [optional] 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



