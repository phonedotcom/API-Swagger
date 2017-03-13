# SwaggerClient::AccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_account**](AccountsApi.md#get_account) | **GET** /accounts/{account_id} | Retrieve details of an individual account
[**list_accounts**](AccountsApi.md#list_accounts) | **GET** /accounts | Get a list of accounts visible to the authenticated user or client


# **get_account**
> AccountFull get_account(account_id, )

Retrieve details of an individual account

This service shows the details of an individual account. See Accounts for more info on the properties.

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

api_instance = SwaggerClient::AccountsApi.new

account_id = 56 # Integer | Account ID


begin
  #Retrieve details of an individual account
  result = api_instance.get_account(account_id, )
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling AccountsApi->get_account: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 

### Return type

[**AccountFull**](AccountFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_accounts**
> ListAccounts list_accounts(opts)

Get a list of accounts visible to the authenticated user or client

This service lists the accounts accessible to the authenticated client. In most cases, there will only be one such account. See Accounts for more info on the properties.

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

api_instance = SwaggerClient::AccountsApi.new

opts = { 
  filters_id: ["filters_id_example"], # Array<String> | ID filter
  sort_id: "sort_id_example", # String | ID sorting
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
}

begin
  #Get a list of accounts visible to the authenticated user or client
  result = api_instance.list_accounts(opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling AccountsApi->list_accounts: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters_id** | [**Array&lt;String&gt;**](String.md)| ID filter | [optional] 
 **sort_id** | **String**| ID sorting | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 

### Return type

[**ListAccounts**](ListAccounts.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



