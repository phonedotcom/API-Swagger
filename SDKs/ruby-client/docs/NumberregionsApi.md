# SwaggerClient::NumberregionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list_available_phone_number_regions**](NumberregionsApi.md#list_available_phone_number_regions) | **GET** /phone-numbers/available/regions | 


# **list_available_phone_number_regions**
> ListPhoneNumbersRegions list_available_phone_number_regions(opts)



This service lists the quantities of available phone numbers by region.

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

api_instance = SwaggerClient::NumberregionsApi.new

opts = { 
  filters_country_code: ["filters_country_code_example"], # Array<String> | Country Code filter
  filters_npa: ["filters_npa_example"], # Array<String> | Area Code filter (North America only)
  filters_nxx: ["filters_nxx_example"], # Array<String> | 2nd set of 3 digits filter (North America only)
  filters_is_toll_free: ["filters_is_toll_free_example"], # Array<String> | Toll-free status filter
  filters_city: ["filters_city_example"], # Array<String> | City filter
  filters_province_postal_code: ["filters_province_postal_code_example"], # Array<String> | State or Province filter
  filters_country_postal_code: ["filters_country_postal_code_example"], # Array<String> | Country filter
  sort_country_code: "sort_country_code_example", # String | International calling code sorting
  sort_npa: "sort_npa_example", # String | Area Code sorting (North America only)
  sort_nxx: "sort_nxx_example", # String | 2nd set of 3 digits sorting (North America)
  sort_is_toll_free: "sort_is_toll_free_example", # String | Toll Free status sorting
  sort_city: "sort_city_example", # String | City sorting
  sort_province_postal_code: "sort_province_postal_code_example", # String | State or Province sorting
  sort_country_postal_code: "sort_country_postal_code_example", # String | Country sorting
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
  group_by: ["group_by_example"] # Array<String> | Fields to group by (supports the same set of fields as the filters and sorting)
}

begin
  #
  result = api_instance.list_available_phone_number_regions(opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling NumberregionsApi->list_available_phone_number_regions: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters_country_code** | [**Array&lt;String&gt;**](String.md)| Country Code filter | [optional] 
 **filters_npa** | [**Array&lt;String&gt;**](String.md)| Area Code filter (North America only) | [optional] 
 **filters_nxx** | [**Array&lt;String&gt;**](String.md)| 2nd set of 3 digits filter (North America only) | [optional] 
 **filters_is_toll_free** | [**Array&lt;String&gt;**](String.md)| Toll-free status filter | [optional] 
 **filters_city** | [**Array&lt;String&gt;**](String.md)| City filter | [optional] 
 **filters_province_postal_code** | [**Array&lt;String&gt;**](String.md)| State or Province filter | [optional] 
 **filters_country_postal_code** | [**Array&lt;String&gt;**](String.md)| Country filter | [optional] 
 **sort_country_code** | **String**| International calling code sorting | [optional] 
 **sort_npa** | **String**| Area Code sorting (North America only) | [optional] 
 **sort_nxx** | **String**| 2nd set of 3 digits sorting (North America) | [optional] 
 **sort_is_toll_free** | **String**| Toll Free status sorting | [optional] 
 **sort_city** | **String**| City sorting | [optional] 
 **sort_province_postal_code** | **String**| State or Province sorting | [optional] 
 **sort_country_postal_code** | **String**| Country sorting | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 
 **group_by** | [**Array&lt;String&gt;**](String.md)| Fields to group by (supports the same set of fields as the filters and sorting) | [optional] 

### Return type

[**ListPhoneNumbersRegions**](ListPhoneNumbersRegions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



