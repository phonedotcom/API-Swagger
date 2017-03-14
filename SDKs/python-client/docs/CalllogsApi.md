# swagger_client.CalllogsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_account_call_logs**](CalllogsApi.md#get_account_call_logs) | **GET** /accounts/{account_id}/call-logs/{call_id} | Show details of an individual Call Log entry
[**list_account_call_logs**](CalllogsApi.md#list_account_call_logs) | **GET** /accounts/{account_id}/call-logs | Get a list of call details associated with your account


# **get_account_call_logs**
> CallLogFull get_account_call_logs(account_id, call_id)

Show details of an individual Call Log entry

See Call Logs for more detail.

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
api_instance = swagger_client.CalllogsApi()
account_id = 56 # int | Account ID
call_id = 'call_id_example' # str | Call ID

try: 
    # Show details of an individual Call Log entry
    api_response = api_instance.get_account_call_logs(account_id, call_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling CalllogsApi->get_account_call_logs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **call_id** | **str**| Call ID | 

### Return type

[**CallLogFull**](CallLogFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_call_logs**
> ListCallLogs list_account_call_logs(account_id, filters_id=filters_id, filters_start_time=filters_start_time, filters_created_at=filters_created_at, filters_direction=filters_direction, filters_called_number=filters_called_number, filters_type=filters_type, filters_extension=filters_extension, sort_id=sort_id, sort_start_time=sort_start_time, sort_created_at=sort_created_at, limit=limit, offset=offset, fields=fields)

Get a list of call details associated with your account

See Call Logs for more detail.

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
api_instance = swagger_client.CalllogsApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_start_time = ['filters_start_time_example'] # list[str] | Call start time filter (optional)
filters_created_at = 'filters_created_at_example' # str | Call log creation time filter (optional)
filters_direction = 'filters_direction_example' # str | Call direction filter: in or out (optional)
filters_called_number = 'filters_called_number_example' # str | Called number (optional)
filters_type = 'filters_type_example' # str | Call type, such as 'call', 'fax', 'audiogram' (optional)
filters_extension = ['filters_extension_example'] # list[str] | Extension filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_start_time = 'sort_start_time_example' # str | Sorting by call start time: asc or desc (optional)
sort_created_at = 'sort_created_at_example' # str | Sorting by call log creation time: asc or desc (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of call details associated with your account
    api_response = api_instance.list_account_call_logs(account_id, filters_id=filters_id, filters_start_time=filters_start_time, filters_created_at=filters_created_at, filters_direction=filters_direction, filters_called_number=filters_called_number, filters_type=filters_type, filters_extension=filters_extension, sort_id=sort_id, sort_start_time=sort_start_time, sort_created_at=sort_created_at, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling CalllogsApi->list_account_call_logs: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **filters_id** | [**list[str]**](str.md)| ID filter | [optional] 
 **filters_start_time** | [**list[str]**](str.md)| Call start time filter | [optional] 
 **filters_created_at** | **str**| Call log creation time filter | [optional] 
 **filters_direction** | **str**| Call direction filter: in or out | [optional] 
 **filters_called_number** | **str**| Called number | [optional] 
 **filters_type** | **str**| Call type, such as &#39;call&#39;, &#39;fax&#39;, &#39;audiogram&#39; | [optional] 
 **filters_extension** | [**list[str]**](str.md)| Extension filter | [optional] 
 **sort_id** | **str**| ID sorting | [optional] 
 **sort_start_time** | **str**| Sorting by call start time: asc or desc | [optional] 
 **sort_created_at** | **str**| Sorting by call log creation time: asc or desc | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 

### Return type

[**ListCallLogs**](ListCallLogs.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

