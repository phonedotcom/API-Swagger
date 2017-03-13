# SwaggerClient::QueuesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_queue**](QueuesApi.md#create_account_queue) | **POST** /accounts/{account_id}/queues | Create a queue
[**delete_account_queue**](QueuesApi.md#delete_account_queue) | **DELETE** /accounts/{account_id}/queues/{queue_id} | Delete a queue
[**get_account_queue**](QueuesApi.md#get_account_queue) | **GET** /accounts/{account_id}/queues/{queue_id} | Show details of an individual queue
[**list_account_queues**](QueuesApi.md#list_account_queues) | **GET** /accounts/{account_id}/queues | Get a list of queues for an account
[**replace_account_queue**](QueuesApi.md#replace_account_queue) | **PUT** /accounts/{account_id}/queues/{queue_id} | Replace a queue


# **create_account_queue**
> QueueFull create_account_queue(account_id, , opts)

Create a queue

For more on the input fields, see Account Queues.

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

api_instance = SwaggerClient::QueuesApi.new

account_id = 56 # Integer | Account ID

opts = { 
  data: SwaggerClient::CreateQueueParams.new # CreateQueueParams | Queue data
}

begin
  #Create a queue
  result = api_instance.create_account_queue(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling QueuesApi->create_account_queue: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional] 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **delete_account_queue**
> DeleteQueue delete_account_queue(account_id, queue_id)

Delete a queue

This service a queue from the account. For more information on queue properties, see Account Queues.

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

api_instance = SwaggerClient::QueuesApi.new

account_id = 56 # Integer | Account ID

queue_id = 56 # Integer | Queue ID


begin
  #Delete a queue
  result = api_instance.delete_account_queue(account_id, queue_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling QueuesApi->delete_account_queue: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **queue_id** | **Integer**| Queue ID | 

### Return type

[**DeleteQueue**](DeleteQueue.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_account_queue**
> QueueFull get_account_queue(account_id, queue_id)

Show details of an individual queue

This service shows the details of an individual queue. For more on the input fields, see Account Queues.

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

api_instance = SwaggerClient::QueuesApi.new

account_id = 56 # Integer | Account ID

queue_id = 56 # Integer | Queue ID


begin
  #Show details of an individual queue
  result = api_instance.get_account_queue(account_id, queue_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling QueuesApi->get_account_queue: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **queue_id** | **Integer**| Queue ID | 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_queues**
> ListQueues list_account_queues(account_id, , opts)

Get a list of queues for an account

The List Queues service lists all the queues belong to the account. See Account Queues for more info on the properties.

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

api_instance = SwaggerClient::QueuesApi.new

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
  #Get a list of queues for an account
  result = api_instance.list_account_queues(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling QueuesApi->list_account_queues: #{e}"
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

[**ListQueues**](ListQueues.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **replace_account_queue**
> QueueFull replace_account_queue(account_id, queue_id, opts)

Replace a queue

The Replace Queue service replaces the parameters of a queue. For more on the input fields, see Account Queues.

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

api_instance = SwaggerClient::QueuesApi.new

account_id = 56 # Integer | Account ID

queue_id = 56 # Integer | Queue ID

opts = { 
  data: SwaggerClient::CreateQueueParams.new # CreateQueueParams | Queue data
}

begin
  #Replace a queue
  result = api_instance.replace_account_queue(account_id, queue_id, opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling QueuesApi->replace_account_queue: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **queue_id** | **Integer**| Queue ID | 
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional] 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



