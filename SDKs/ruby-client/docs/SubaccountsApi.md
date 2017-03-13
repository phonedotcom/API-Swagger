# SwaggerClient::SubaccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_subaccount**](SubaccountsApi.md#create_account_subaccount) | **POST** /accounts/{account_id}/subaccounts | Add a subaccount for the authenticated user or client
[**list_account_subaccounts**](SubaccountsApi.md#list_account_subaccounts) | **GET** /accounts/{account_id}/subaccounts | Get a list of subaccounts for the authenticated user or client


# **create_account_subaccount**
> AccountFull create_account_subaccount(account_id, data)

Add a subaccount for the authenticated user or client

This service shows the details of an individual Subaccount.

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

api_instance = SwaggerClient::SubaccountsApi.new

account_id = 56 # Integer | Account ID

data = SwaggerClient::CreateSubaccountParams.new # CreateSubaccountParams | SMS data


begin
  #Add a subaccount for the authenticated user or client
  result = api_instance.create_account_subaccount(account_id, data)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SubaccountsApi->create_account_subaccount: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **data** | [**CreateSubaccountParams**](CreateSubaccountParams.md)| SMS data | 

### Return type

[**AccountFull**](AccountFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_subaccounts**
> ListAccounts list_account_subaccounts(account_id, , opts)

Get a list of subaccounts for the authenticated user or client

This service lists the Subaccount of the authenticated client. In most cases, there will not be any.

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

api_instance = SwaggerClient::SubaccountsApi.new

account_id = 56 # Integer | Account ID

opts = { 
  filters_id: ["filters_id_example"], # Array<String> | ID filter
  sort_id: "sort_id_example", # String | ID sorting
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
}

begin
  #Get a list of subaccounts for the authenticated user or client
  result = api_instance.list_account_subaccounts(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SubaccountsApi->list_account_subaccounts: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
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



