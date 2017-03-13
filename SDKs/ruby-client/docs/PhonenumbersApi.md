# SwaggerClient::PhonenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_phone_number**](PhonenumbersApi.md#create_account_phone_number) | **POST** /accounts/{account_id}/phone-numbers | Add a phone number to an account
[**get_account_phone_number**](PhonenumbersApi.md#get_account_phone_number) | **GET** /accounts/{account_id}/phone-numbers/{number_id} | Show details of an individual phone number
[**list_account_phone_numbers**](PhonenumbersApi.md#list_account_phone_numbers) | **GET** /accounts/{account_id}/phone-numbers | Get a list of phone numbers registered to an account
[**replace_account_phone_number**](PhonenumbersApi.md#replace_account_phone_number) | **PUT** /accounts/{account_id}/phone-numbers/{number_id} | Update the settings for an existing phone number on your account


# **create_account_phone_number**
> PhoneNumberFull create_account_phone_number(account_id, , opts)

Add a phone number to an account

See Intro to Account Phone Numbers for more info on the properties to use.

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

api_instance = SwaggerClient::PhonenumbersApi.new

account_id = 56 # Integer | Account ID

opts = { 
  data: SwaggerClient::CreatePhoneNumberParams.new # CreatePhoneNumberParams | Phone Number data
}

begin
  #Add a phone number to an account
  result = api_instance.create_account_phone_number(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling PhonenumbersApi->create_account_phone_number: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **data** | [**CreatePhoneNumberParams**](CreatePhoneNumberParams.md)| Phone Number data | [optional] 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_account_phone_number**
> PhoneNumberFull get_account_phone_number(account_id, number_id)

Show details of an individual phone number

See Intro to Account Phone Numbers for more info on the properties.

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

api_instance = SwaggerClient::PhonenumbersApi.new

account_id = 56 # Integer | Account ID

number_id = 56 # Integer | Number ID


begin
  #Show details of an individual phone number
  result = api_instance.get_account_phone_number(account_id, number_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling PhonenumbersApi->get_account_phone_number: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **number_id** | **Integer**| Number ID | 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_phone_numbers**
> ListPhoneNumbers list_account_phone_numbers(account_id, , opts)

Get a list of phone numbers registered to an account

See Intro to Account Phone Numbers for more info on the properties.

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

api_instance = SwaggerClient::PhonenumbersApi.new

account_id = 56 # Integer | Account ID

opts = { 
  filters_id: ["filters_id_example"], # Array<String> | ID filter
  filters_name: ["filters_name_example"], # Array<String> | Name filter
  filters_phone_number: ["filters_phone_number_example"], # Array<String> | Phone number filter
  sort_id: "sort_id_example", # String | ID sorting
  sort_name: "sort_name_example", # String | Name sorting
  sort_phone_number: "sort_phone_number_example", # String | Phone number sorting
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
}

begin
  #Get a list of phone numbers registered to an account
  result = api_instance.list_account_phone_numbers(account_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling PhonenumbersApi->list_account_phone_numbers: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **filters_id** | [**Array&lt;String&gt;**](String.md)| ID filter | [optional] 
 **filters_name** | [**Array&lt;String&gt;**](String.md)| Name filter | [optional] 
 **filters_phone_number** | [**Array&lt;String&gt;**](String.md)| Phone number filter | [optional] 
 **sort_id** | **String**| ID sorting | [optional] 
 **sort_name** | **String**| Name sorting | [optional] 
 **sort_phone_number** | **String**| Phone number sorting | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 

### Return type

[**ListPhoneNumbers**](ListPhoneNumbers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **replace_account_phone_number**
> PhoneNumberFull replace_account_phone_number(account_id, number_id, opts)

Update the settings for an existing phone number on your account

See Intro to Account Phone Numbers for more info on the properties.

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

api_instance = SwaggerClient::PhonenumbersApi.new

account_id = 56 # Integer | Account ID

number_id = 56 # Integer | Number ID

opts = { 
  data: SwaggerClient::ReplacePhoneNumberParams.new # ReplacePhoneNumberParams | Phone Number data
}

begin
  #Update the settings for an existing phone number on your account
  result = api_instance.replace_account_phone_number(account_id, number_id, opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling PhonenumbersApi->replace_account_phone_number: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **number_id** | **Integer**| Number ID | 
 **data** | [**ReplacePhoneNumberParams**](ReplacePhoneNumberParams.md)| Phone Number data | [optional] 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



