# SwaggerClient::RoutesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_route**](RoutesApi.md#create_route) | **POST** /accounts/{account_id}/routes | Add a new address book contact for an extension
[**delete_account_route**](RoutesApi.md#delete_account_route) | **DELETE** /accounts/{account_id}/routes/{route_id} | 
[**get_account_route**](RoutesApi.md#get_account_route) | **GET** /accounts/{account_id}/routes/{route_id} | Show details of an individual route
[**list_account_routes**](RoutesApi.md#list_account_routes) | **GET** /accounts/{account_id}/routes | Get a list of routes for an account
[**replace_account_route**](RoutesApi.md#replace_account_route) | **PUT** /accounts/{account_id}/routes/{route_id} | 


# **create_route**
> RouteFull create_route(account_id, , opts)

Add a new address book contact for an extension

For more on the input fields, see Intro to Routes.

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

api_instance = SwaggerClient::RoutesApi.new

account_id = 56 # Integer | Account ID

opts = { 
  data: SwaggerClient::CreateRouteParams.new # CreateRouteParams | Route data
}

begin
  #Add a new address book contact for an extension
  result = api_instance.create_route(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling RoutesApi->create_route: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **data** | [**CreateRouteParams**](CreateRouteParams.md)| Route data | [optional] 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **delete_account_route**
> DeleteRoute delete_account_route(account_id, route_id)





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

api_instance = SwaggerClient::RoutesApi.new

account_id = 56 # Integer | Account ID

route_id = 56 # Integer | Route ID


begin
  #
  result = api_instance.delete_account_route(account_id, route_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling RoutesApi->delete_account_route: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **route_id** | **Integer**| Route ID | 

### Return type

[**DeleteRoute**](DeleteRoute.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_account_route**
> RouteFull get_account_route(account_id, route_id)

Show details of an individual route

This service shows the details of an individual route.

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

api_instance = SwaggerClient::RoutesApi.new

account_id = 56 # Integer | Account ID

route_id = 56 # Integer | Route ID


begin
  #Show details of an individual route
  result = api_instance.get_account_route(account_id, route_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling RoutesApi->get_account_route: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **route_id** | **Integer**| Route ID | 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_routes**
> ListRoutes list_account_routes(account_id, , opts)

Get a list of routes for an account

See Intro to Routes for more info on the properties.

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

api_instance = SwaggerClient::RoutesApi.new

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
  #Get a list of routes for an account
  result = api_instance.list_account_routes(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling RoutesApi->list_account_routes: #{e}"
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

[**ListRoutes**](ListRoutes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **replace_account_route**
> RouteFull replace_account_route(account_id, route_id, opts)



For more on the input fields, see Intro to Routes.

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

api_instance = SwaggerClient::RoutesApi.new

account_id = 56 # Integer | Account ID

route_id = 56 # Integer | Route ID

opts = { 
  data: SwaggerClient::CreateRouteParams.new # CreateRouteParams | Route data
}

begin
  #
  result = api_instance.replace_account_route(account_id, route_id, opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling RoutesApi->replace_account_route: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **route_id** | **Integer**| Route ID | 
 **data** | [**CreateRouteParams**](CreateRouteParams.md)| Route data | [optional] 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



