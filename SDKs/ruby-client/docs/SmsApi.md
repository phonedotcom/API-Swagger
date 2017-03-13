# SwaggerClient::SmsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_sms**](SmsApi.md#create_account_sms) | **POST** /accounts/{account_id}/sms | Send a SMS to one or a group of recipients
[**get_account_sms**](SmsApi.md#get_account_sms) | **GET** /accounts/{account_id}/sms/{sms_id} | Show details of an individual SMS
[**list_account_sms**](SmsApi.md#list_account_sms) | **GET** /accounts/{account_id}/sms | Get a list of SMS messages for an account


# **create_account_sms**
> SmsFull create_account_sms(account_id, data)

Send a SMS to one or a group of recipients

For more on the input fields, see Intro to SMS.

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

api_instance = SwaggerClient::SmsApi.new

account_id = 56 # Integer | Account ID

data = SwaggerClient::CreateSmsParams.new # CreateSmsParams | SMS data


begin
  #Send a SMS to one or a group of recipients
  result = api_instance.create_account_sms(account_id, data)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SmsApi->create_account_sms: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **data** | [**CreateSmsParams**](CreateSmsParams.md)| SMS data | 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_account_sms**
> SmsFull get_account_sms(account_id, sms_id)

Show details of an individual SMS

This service shows the details of an individual sms. See Intro to SMS for more info on the properties.

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

api_instance = SwaggerClient::SmsApi.new

account_id = 56 # Integer | Account ID

sms_id = "sms_id_example" # String | SMS ID


begin
  #Show details of an individual SMS
  result = api_instance.get_account_sms(account_id, sms_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SmsApi->get_account_sms: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **sms_id** | **String**| SMS ID | 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_sms**
> ListSms list_account_sms(account_id, , opts)

Get a list of SMS messages for an account

See Intro to SMS for more info on the properties.

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

api_instance = SwaggerClient::SmsApi.new

account_id = 56 # Integer | Account ID

opts = { 
  filters_id: ["filters_id_example"], # Array<String> | ID filter
  filters_direction: "filters_direction_example", # String | Direction filter
  filters_from: "filters_from_example", # String | Caller ID filter
  sort_id: "sort_id_example", # String | ID sorting
  sort_created_at: "sort_created_at_example", # String | Sort by created time of message
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
}

begin
  #Get a list of SMS messages for an account
  result = api_instance.list_account_sms(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling SmsApi->list_account_sms: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **filters_id** | [**Array&lt;String&gt;**](String.md)| ID filter | [optional] 
 **filters_direction** | **String**| Direction filter | [optional] 
 **filters_from** | **String**| Caller ID filter | [optional] 
 **sort_id** | **String**| ID sorting | [optional] 
 **sort_created_at** | **String**| Sort by created time of message | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 

### Return type

[**ListSms**](ListSms.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



