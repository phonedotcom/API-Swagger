# SwaggerClient::SchedulesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_account_schedule**](SchedulesApi.md#get_account_schedule) | **GET** /accounts/{account_id}/schedules/{schedule_id} | Show details of an individual schedule
[**list_account_schedules**](SchedulesApi.md#list_account_schedules) | **GET** /accounts/{account_id}/schedules | Get a list of schedules for an account


# **get_account_schedule**
> ScheduleFull get_account_schedule(account_id, schedule_id)

Show details of an individual schedule

This service shows the details of an individual schedule.

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

api_instance = SwaggerClient::SchedulesApi.new

account_id = 56 # Integer | Account ID

schedule_id = 56 # Integer | Schedule ID


begin
  #Show details of an individual schedule
  result = api_instance.get_account_schedule(account_id, schedule_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SchedulesApi->get_account_schedule: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **schedule_id** | **Integer**| Schedule ID | 

### Return type

[**ScheduleFull**](ScheduleFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_schedules**
> ListSchedules list_account_schedules(account_id, , opts)

Get a list of schedules for an account

See Intro to Schedules for more info on the properties.

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

api_instance = SwaggerClient::SchedulesApi.new

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
  #Get a list of schedules for an account
  result = api_instance.list_account_schedules(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SchedulesApi->list_account_schedules: #{e}"
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

[**ListSchedules**](ListSchedules.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



