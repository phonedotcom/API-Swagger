# swagger_client.AvailablenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**list_available_phone_numbers**](AvailablenumbersApi.md#list_available_phone_numbers) | **GET** /phone-numbers/available | 


# **list_available_phone_numbers**
> ListAvailableNumbers list_available_phone_numbers(filters_phone_number=filters_phone_number, filters_country_code=filters_country_code, filters_npa=filters_npa, filters_nxx=filters_nxx, filters_xxxx=filters_xxxx, filters_city=filters_city, filters_province=filters_province, filters_country=filters_country, filters_price=filters_price, filters_category=filters_category, sort_internal=sort_internal, sort_price=sort_price, sort_phone_number=sort_phone_number, limit=limit, offset=offset, fields=fields)





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
api_instance = swagger_client.AvailablenumbersApi()
filters_phone_number = ['filters_phone_number_example'] # list[str] | Phone number filter (optional)
filters_country_code = ['filters_country_code_example'] # list[str] | Country Code filter (optional)
filters_npa = ['filters_npa_example'] # list[str] | Area Code filter (North America only) (optional)
filters_nxx = ['filters_nxx_example'] # list[str] | 2nd set of 3 digits filter (North America only) (optional)
filters_xxxx = ['filters_xxxx_example'] # list[str] | NANP XXXX filter (optional)
filters_city = ['filters_city_example'] # list[str] | City filter (optional)
filters_province = ['filters_province_example'] # list[str] | State or Province (postal code) filter (optional)
filters_country = ['filters_country_example'] # list[str] | Country (postal code) filter (optional)
filters_price = ['filters_price_example'] # list[str] | Price filter (optional)
filters_category = ['filters_category_example'] # list[str] | Category filter (optional)
sort_internal = 'sort_internal_example' # str | Internal (quasi-random) sorting (optional)
sort_price = 'sort_price_example' # str | Price sorting (optional)
sort_phone_number = 'sort_phone_number_example' # str | Phone number sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # 
    api_response = api_instance.list_available_phone_numbers(filters_phone_number=filters_phone_number, filters_country_code=filters_country_code, filters_npa=filters_npa, filters_nxx=filters_nxx, filters_xxxx=filters_xxxx, filters_city=filters_city, filters_province=filters_province, filters_country=filters_country, filters_price=filters_price, filters_category=filters_category, sort_internal=sort_internal, sort_price=sort_price, sort_phone_number=sort_phone_number, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling AvailablenumbersApi->list_available_phone_numbers: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters_phone_number** | [**list[str]**](str.md)| Phone number filter | [optional] 
 **filters_country_code** | [**list[str]**](str.md)| Country Code filter | [optional] 
 **filters_npa** | [**list[str]**](str.md)| Area Code filter (North America only) | [optional] 
 **filters_nxx** | [**list[str]**](str.md)| 2nd set of 3 digits filter (North America only) | [optional] 
 **filters_xxxx** | [**list[str]**](str.md)| NANP XXXX filter | [optional] 
 **filters_city** | [**list[str]**](str.md)| City filter | [optional] 
 **filters_province** | [**list[str]**](str.md)| State or Province (postal code) filter | [optional] 
 **filters_country** | [**list[str]**](str.md)| Country (postal code) filter | [optional] 
 **filters_price** | [**list[str]**](str.md)| Price filter | [optional] 
 **filters_category** | [**list[str]**](str.md)| Category filter | [optional] 
 **sort_internal** | **str**| Internal (quasi-random) sorting | [optional] 
 **sort_price** | **str**| Price sorting | [optional] 
 **sort_phone_number** | **str**| Phone number sorting | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 

### Return type

[**ListAvailableNumbers**](ListAvailableNumbers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

