# SwaggerClient::ApplicationsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_account_application**](ApplicationsApi.md#get_account_application) | **GET** /accounts/{account_id}/applications/{application_id} | Show details of an individual application
[**list_account_applications**](ApplicationsApi.md#list_account_applications) | **GET** /accounts/{account_id}/applications | Get a list of applications you have defined


# **get_account_application**
> ApplicationFull get_account_application(account_id, application_id)

Show details of an individual application



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

api_instance = SwaggerClient::ApplicationsApi.new

account_id = 56 # Integer | Account ID

application_id = 56 # Integer | Application ID


begin
  #Show details of an individual application
  result = api_instance.get_account_application(account_id, application_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ApplicationsApi->get_account_application: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **application_id** | **Integer**| Application ID | 

### Return type

[**ApplicationFull**](ApplicationFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_applications**
> ListApplications list_account_applications(account_id, , opts)

Get a list of applications you have defined

Get a list of an account available applications

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

api_instance = SwaggerClient::ApplicationsApi.new

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
  #Get a list of applications you have defined
  result = api_instance.list_account_applications(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ApplicationsApi->list_account_applications: #{e}"
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

[**ListApplications**](ListApplications.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



