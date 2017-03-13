# SwaggerClient::MenusApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_menu**](MenusApi.md#create_account_menu) | **POST** /accounts/{account_id}/menus | Create an individual menu
[**delete_account_menu**](MenusApi.md#delete_account_menu) | **DELETE** /accounts/{account_id}/menus/{menu_id} | Delete an individual menu
[**get_account_menu**](MenusApi.md#get_account_menu) | **GET** /accounts/{account_id}/menus/{menu_id} | Show details of an individual menu
[**list_account_menus**](MenusApi.md#list_account_menus) | **GET** /accounts/{account_id}/menus | Get a list of menus for an account
[**replace_account_menu**](MenusApi.md#replace_account_menu) | **PUT** /accounts/{account_id}/menus/{menu_id} | Replace an individual menu


# **create_account_menu**
> MenuFull create_account_menu(account_id, , opts)

Create an individual menu

This service creates an individual menu. See Account Menus for more info on the properties.

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

api_instance = SwaggerClient::MenusApi.new

account_id = 56 # Integer | Account ID

opts = { 
  data: SwaggerClient::CreateMenuParams.new # CreateMenuParams | Menu data
}

begin
  #Create an individual menu
  result = api_instance.create_account_menu(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling MenusApi->create_account_menu: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **data** | [**CreateMenuParams**](CreateMenuParams.md)| Menu data | [optional] 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **delete_account_menu**
> DeleteMenu delete_account_menu(account_id, menu_id)

Delete an individual menu

This service shows the details of an individual menu.

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

api_instance = SwaggerClient::MenusApi.new

account_id = 56 # Integer | Account ID

menu_id = 56 # Integer | Menu ID


begin
  #Delete an individual menu
  result = api_instance.delete_account_menu(account_id, menu_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling MenusApi->delete_account_menu: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **menu_id** | **Integer**| Menu ID | 

### Return type

[**DeleteMenu**](DeleteMenu.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_account_menu**
> MenuFull get_account_menu(account_id, menu_id)

Show details of an individual menu

This service shows the details of an individual Menu.

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

api_instance = SwaggerClient::MenusApi.new

account_id = 56 # Integer | Account ID

menu_id = 56 # Integer | Menu ID


begin
  #Show details of an individual menu
  result = api_instance.get_account_menu(account_id, menu_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling MenusApi->get_account_menu: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **menu_id** | **Integer**| Menu ID | 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_menus**
> ListMenus list_account_menus(account_id, , opts)

Get a list of menus for an account

See Account Menus for more info on the properties.

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

api_instance = SwaggerClient::MenusApi.new

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
  #Get a list of menus for an account
  result = api_instance.list_account_menus(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling MenusApi->list_account_menus: #{e}"
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

[**ListMenus**](ListMenus.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **replace_account_menu**
> MenuFull replace_account_menu(account_id, menu_id, opts)

Replace an individual menu

This service replaces the details of an individual Menu.

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

api_instance = SwaggerClient::MenusApi.new

account_id = 56 # Integer | Account ID

menu_id = 56 # Integer | Menu ID

opts = { 
  data: SwaggerClient::ReplaceMenuParams.new # ReplaceMenuParams | Menu data
}

begin
  #Replace an individual menu
  result = api_instance.replace_account_menu(account_id, menu_id, opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling MenusApi->replace_account_menu: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **menu_id** | **Integer**| Menu ID | 
 **data** | [**ReplaceMenuParams**](ReplaceMenuParams.md)| Menu data | [optional] 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



