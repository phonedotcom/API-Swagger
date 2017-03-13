# SwaggerClient::TrunksApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_trunk**](TrunksApi.md#create_account_trunk) | **POST** /accounts/{account_id}/trunks | Add a trunk record with SIP information
[**delete_account_trunk**](TrunksApi.md#delete_account_trunk) | **DELETE** /accounts/{account_id}/trunks/{trunk_id} | Delete a trunk from account
[**get_account_trunk**](TrunksApi.md#get_account_trunk) | **GET** /accounts/{account_id}/trunks/{trunk_id} | Show details of an individual trunk
[**list_account_trunks**](TrunksApi.md#list_account_trunks) | **GET** /accounts/{account_id}/trunks | Get a list of trunks for an account
[**replace_account_trunk**](TrunksApi.md#replace_account_trunk) | **PUT** /accounts/{account_id}/trunks/{trunk_id} | Replace parameters in a trunk


# **create_account_trunk**
> TrunkFull create_account_trunk(account_id, data)

Add a trunk record with SIP information

For more on the input fields, see Account Trunks.

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

api_instance = SwaggerClient::TrunksApi.new

account_id = 56 # Integer | Account ID

data = SwaggerClient::CreateTrunkParams.new # CreateTrunkParams | Trunk data


begin
  #Add a trunk record with SIP information
  result = api_instance.create_account_trunk(account_id, data)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling TrunksApi->create_account_trunk: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **data** | [**CreateTrunkParams**](CreateTrunkParams.md)| Trunk data | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **delete_account_trunk**
> DeleteTrunk delete_account_trunk(account_id, trunk_id)

Delete a trunk from account

This service deletes a trunk from the account. For more on the properties of trunks, see Account Trunks.

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

api_instance = SwaggerClient::TrunksApi.new

account_id = 56 # Integer | Account ID

trunk_id = 56 # Integer | Trunk ID


begin
  #Delete a trunk from account
  result = api_instance.delete_account_trunk(account_id, trunk_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling TrunksApi->delete_account_trunk: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **trunk_id** | **Integer**| Trunk ID | 

### Return type

[**DeleteTrunk**](DeleteTrunk.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_account_trunk**
> TrunkFull get_account_trunk(account_id, trunk_id)

Show details of an individual trunk

This service shows the details of an individual Trunk.

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

api_instance = SwaggerClient::TrunksApi.new

account_id = 56 # Integer | Account ID

trunk_id = 56 # Integer | Trunk ID


begin
  #Show details of an individual trunk
  result = api_instance.get_account_trunk(account_id, trunk_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling TrunksApi->get_account_trunk: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **trunk_id** | **Integer**| Trunk ID | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_trunks**
> ListTrunks list_account_trunks(account_id, , opts)

Get a list of trunks for an account

See Account Trunks for more info on the properties.

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

api_instance = SwaggerClient::TrunksApi.new

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
  #Get a list of trunks for an account
  result = api_instance.list_account_trunks(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling TrunksApi->list_account_trunks: #{e}"
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

[**ListTrunks**](ListTrunks.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **replace_account_trunk**
> TrunkFull replace_account_trunk(account_id, trunk_iddata)

Replace parameters in a trunk

For more on the input fields, see Account Trunks.

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

api_instance = SwaggerClient::TrunksApi.new

account_id = 56 # Integer | Account ID

trunk_id = 56 # Integer | Trunk ID

data = SwaggerClient::CreateTrunkParams.new # CreateTrunkParams | Trunk data


begin
  #Replace parameters in a trunk
  result = api_instance.replace_account_trunk(account_id, trunk_iddata)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling TrunksApi->replace_account_trunk: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **trunk_id** | **Integer**| Trunk ID | 
 **data** | [**CreateTrunkParams**](CreateTrunkParams.md)| Trunk data | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



