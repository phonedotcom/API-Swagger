# SwaggerClient::AvailablenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list_available_phone_numbers**](AvailablenumbersApi.md#list_available_phone_numbers) | **GET** /phone-numbers/available | 


# **list_available_phone_numbers**
> ListAvailableNumbers list_available_phone_numbers(opts)





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

api_instance = SwaggerClient::AvailablenumbersApi.new

opts = { 
  filters_phone_number: ["filters_phone_number_example"], # Array<String> | Phone number filter
  filters_country_code: ["filters_country_code_example"], # Array<String> | Country Code filter
  filters_npa: ["filters_npa_example"], # Array<String> | Area Code filter (North America only)
  filters_nxx: ["filters_nxx_example"], # Array<String> | 2nd set of 3 digits filter (North America only)
  filters_xxxx: ["filters_xxxx_example"], # Array<String> | NANP XXXX filter
  filters_city: ["filters_city_example"], # Array<String> | City filter
  filters_province: ["filters_province_example"], # Array<String> | State or Province (postal code) filter
  filters_country: ["filters_country_example"], # Array<String> | Country (postal code) filter
  filters_price: ["filters_price_example"], # Array<String> | Price filter
  filters_category: ["filters_category_example"], # Array<String> | Category filter
  sort_internal: "sort_internal_example", # String | Internal (quasi-random) sorting
  sort_price: "sort_price_example", # String | Price sorting
  sort_phone_number: "sort_phone_number_example", # String | Phone number sorting
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
}

begin
  #
  result = api_instance.list_available_phone_numbers(opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling AvailablenumbersApi->list_available_phone_numbers: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters_phone_number** | [**Array&lt;String&gt;**](String.md)| Phone number filter | [optional] 
 **filters_country_code** | [**Array&lt;String&gt;**](String.md)| Country Code filter | [optional] 
 **filters_npa** | [**Array&lt;String&gt;**](String.md)| Area Code filter (North America only) | [optional] 
 **filters_nxx** | [**Array&lt;String&gt;**](String.md)| 2nd set of 3 digits filter (North America only) | [optional] 
 **filters_xxxx** | [**Array&lt;String&gt;**](String.md)| NANP XXXX filter | [optional] 
 **filters_city** | [**Array&lt;String&gt;**](String.md)| City filter | [optional] 
 **filters_province** | [**Array&lt;String&gt;**](String.md)| State or Province (postal code) filter | [optional] 
 **filters_country** | [**Array&lt;String&gt;**](String.md)| Country (postal code) filter | [optional] 
 **filters_price** | [**Array&lt;String&gt;**](String.md)| Price filter | [optional] 
 **filters_category** | [**Array&lt;String&gt;**](String.md)| Category filter | [optional] 
 **sort_internal** | **String**| Internal (quasi-random) sorting | [optional] 
 **sort_price** | **String**| Price sorting | [optional] 
 **sort_phone_number** | **String**| Phone number sorting | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 

### Return type

[**ListAvailableNumbers**](ListAvailableNumbers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



