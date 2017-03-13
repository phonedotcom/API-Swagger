# swagger_client.CalleridsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_caller_ids**](CalleridsApi.md#get_caller_ids) | **GET** /accounts/{account_id}/extensions/{extension_id}/caller-ids | Show the Caller ID options a given extension can use


# **get_caller_ids**
> ListCallerIds get_caller_ids(account_id, extension_id, filters_number=filters_number, filters_name=filters_name, sort_number=sort_number, sort_name=sort_name, limit=limit, offset=offset, fields=fields)

Show the Caller ID options a given extension can use

Get Caller ID

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
api_instance = swagger_client.CalleridsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
filters_number = ['filters_number_example'] # list[str] | Number filter (optional)
filters_name = ['filters_name_example'] # list[str] | Name filter (optional)
sort_number = 'sort_number_example' # str | Number sorting (optional)
sort_name = 'sort_name_example' # str | Name sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Show the Caller ID options a given extension can use
    api_response = api_instance.get_caller_ids(account_id, extension_id, filters_number=filters_number, filters_name=filters_name, sort_number=sort_number, sort_name=sort_name, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling CalleridsApi->get_caller_ids: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **filters_number** | [**list[str]**](str.md)| Number filter | [optional] 
 **filters_name** | [**list[str]**](str.md)| Name filter | [optional] 
 **sort_number** | **str**| Number sorting | [optional] 
 **sort_name** | **str**| Name sorting | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 

### Return type

[**ListCallerIds**](ListCallerIds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

