# SwaggerClient::GroupsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_extension_contact_group**](GroupsApi.md#create_account_extension_contact_group) | **POST** /accounts/{account_id}/extensions/{extension_id}/contact-groups | 
[**delete_account_extension_contact_group**](GroupsApi.md#delete_account_extension_contact_group) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | Delete an addressbook group
[**get_account_extension_contact_group**](GroupsApi.md#get_account_extension_contact_group) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 
[**list_account_extension_contact_groups**](GroupsApi.md#list_account_extension_contact_groups) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups | Show a list of contact groups belonging to an extension
[**replace_account_extension_contact_group**](GroupsApi.md#replace_account_extension_contact_group) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 


# **create_account_extension_contact_group**
> GroupFull create_account_extension_contact_group(account_id, extension_id, data)



See Account Contact Groups for more info on the properties.

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

api_instance = SwaggerClient::GroupsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

data = SwaggerClient::CreateGroupParams.new # CreateGroupParams | Group name


begin
  #
  result = api_instance.create_account_extension_contact_group(account_id, extension_id, data)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling GroupsApi->create_account_extension_contact_group: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **data** | [**CreateGroupParams**](CreateGroupParams.md)| Group name | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **delete_account_extension_contact_group**
> DeleteGroup delete_account_extension_contact_group(account_id, extension_id, group_id)

Delete an addressbook group



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

api_instance = SwaggerClient::GroupsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

group_id = 56 # Integer | Group ID


begin
  #Delete an addressbook group
  result = api_instance.delete_account_extension_contact_group(account_id, extension_id, group_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling GroupsApi->delete_account_extension_contact_group: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **group_id** | **Integer**| Group ID | 

### Return type

[**DeleteGroup**](DeleteGroup.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_account_extension_contact_group**
> GroupFull get_account_extension_contact_group(account_id, extension_id, group_id)



See Account Contact Groups for more info on the properties.

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

api_instance = SwaggerClient::GroupsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

group_id = 56 # Integer | Group ID


begin
  #
  result = api_instance.get_account_extension_contact_group(account_id, extension_id, group_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling GroupsApi->get_account_extension_contact_group: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **group_id** | **Integer**| Group ID | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_extension_contact_groups**
> ListGroups list_account_extension_contact_groups(account_id, extension_id, , opts)

Show a list of contact groups belonging to an extension

See Account Contact Groups for details on the properties.

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

api_instance = SwaggerClient::GroupsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

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
  #Show a list of contact groups belonging to an extension
  result = api_instance.list_account_extension_contact_groups(account_id, extension_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling GroupsApi->list_account_extension_contact_groups: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **filters_id** | [**Array&lt;String&gt;**](String.md)| ID filter | [optional] 
 **filters_name** | [**Array&lt;String&gt;**](String.md)| Name filter | [optional] 
 **sort_id** | **String**| ID sorting | [optional] 
 **sort_name** | **String**| Name sorting | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 

### Return type

[**ListGroups**](ListGroups.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **replace_account_extension_contact_group**
> GroupFull replace_account_extension_contact_group(account_id, extension_id, group_iddata)



See Account Contact Groups for more info on the properties.

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

api_instance = SwaggerClient::GroupsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

group_id = 56 # Integer | Group ID

data = SwaggerClient::CreateGroupParams.new # CreateGroupParams | Group name


begin
  #
  result = api_instance.replace_account_extension_contact_group(account_id, extension_id, group_iddata)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling GroupsApi->replace_account_extension_contact_group: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **group_id** | **Integer**| Group ID | 
 **data** | [**CreateGroupParams**](CreateGroupParams.md)| Group name | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



