# swagger_client.NumberregionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list_available_phone_number_regions**](NumberregionsApi.md#list_available_phone_number_regions) | **GET** /phone-numbers/available/regions | 


# **list_available_phone_number_regions**
> ListPhoneNumbersRegions list_available_phone_number_regions(filters_country_code=filters_country_code, filters_npa=filters_npa, filters_nxx=filters_nxx, filters_is_toll_free=filters_is_toll_free, filters_city=filters_city, filters_province_postal_code=filters_province_postal_code, filters_country_postal_code=filters_country_postal_code, sort_country_code=sort_country_code, sort_npa=sort_npa, sort_nxx=sort_nxx, sort_is_toll_free=sort_is_toll_free, sort_city=sort_city, sort_province_postal_code=sort_province_postal_code, sort_country_postal_code=sort_country_postal_code, limit=limit, offset=offset, fields=fields, group_by=group_by)



This service lists the quantities of available phone numbers by region.

### Example 
```python
from __future__ import print_statement
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# Configure API key authorization: apiKey
swagger_client.configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# swagger_client.configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = swagger_client.NumberregionsApi()
filters_country_code = ['filters_country_code_example'] # list[str] | Country Code filter (optional)
filters_npa = ['filters_npa_example'] # list[str] | Area Code filter (North America only) (optional)
filters_nxx = ['filters_nxx_example'] # list[str] | 2nd set of 3 digits filter (North America only) (optional)
filters_is_toll_free = ['filters_is_toll_free_example'] # list[str] | Toll-free status filter (optional)
filters_city = ['filters_city_example'] # list[str] | City filter (optional)
filters_province_postal_code = ['filters_province_postal_code_example'] # list[str] | State or Province filter (optional)
filters_country_postal_code = ['filters_country_postal_code_example'] # list[str] | Country filter (optional)
sort_country_code = 'sort_country_code_example' # str | International calling code sorting (optional)
sort_npa = 'sort_npa_example' # str | Area Code sorting (North America only) (optional)
sort_nxx = 'sort_nxx_example' # str | 2nd set of 3 digits sorting (North America) (optional)
sort_is_toll_free = 'sort_is_toll_free_example' # str | Toll Free status sorting (optional)
sort_city = 'sort_city_example' # str | City sorting (optional)
sort_province_postal_code = 'sort_province_postal_code_example' # str | State or Province sorting (optional)
sort_country_postal_code = 'sort_country_postal_code_example' # str | Country sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)
group_by = ['group_by_example'] # list[str] | Fields to group by (supports the same set of fields as the filters and sorting) (optional)

try: 
    # 
    api_response = api_instance.list_available_phone_number_regions(filters_country_code=filters_country_code, filters_npa=filters_npa, filters_nxx=filters_nxx, filters_is_toll_free=filters_is_toll_free, filters_city=filters_city, filters_province_postal_code=filters_province_postal_code, filters_country_postal_code=filters_country_postal_code, sort_country_code=sort_country_code, sort_npa=sort_npa, sort_nxx=sort_nxx, sort_is_toll_free=sort_is_toll_free, sort_city=sort_city, sort_province_postal_code=sort_province_postal_code, sort_country_postal_code=sort_country_postal_code, limit=limit, offset=offset, fields=fields, group_by=group_by)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling NumberregionsApi->list_available_phone_number_regions: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters_country_code** | [**list[str]**](str.md)| Country Code filter | [optional] 
 **filters_npa** | [**list[str]**](str.md)| Area Code filter (North America only) | [optional] 
 **filters_nxx** | [**list[str]**](str.md)| 2nd set of 3 digits filter (North America only) | [optional] 
 **filters_is_toll_free** | [**list[str]**](str.md)| Toll-free status filter | [optional] 
 **filters_city** | [**list[str]**](str.md)| City filter | [optional] 
 **filters_province_postal_code** | [**list[str]**](str.md)| State or Province filter | [optional] 
 **filters_country_postal_code** | [**list[str]**](str.md)| Country filter | [optional] 
 **sort_country_code** | **str**| International calling code sorting | [optional] 
 **sort_npa** | **str**| Area Code sorting (North America only) | [optional] 
 **sort_nxx** | **str**| 2nd set of 3 digits sorting (North America) | [optional] 
 **sort_is_toll_free** | **str**| Toll Free status sorting | [optional] 
 **sort_city** | **str**| City sorting | [optional] 
 **sort_province_postal_code** | **str**| State or Province sorting | [optional] 
 **sort_country_postal_code** | **str**| Country sorting | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 
 **group_by** | [**list[str]**](str.md)| Fields to group by (supports the same set of fields as the filters and sorting) | [optional] 

### Return type

[**ListPhoneNumbersRegions**](ListPhoneNumbersRegions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

