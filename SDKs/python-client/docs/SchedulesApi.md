# swagger_client.SchedulesApi

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
api_instance = swagger_client.SchedulesApi()
account_id = 56 # int | Account ID
schedule_id = 56 # int | Schedule ID

try: 
    # Show details of an individual schedule
    api_response = api_instance.get_account_schedule(account_id, schedule_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchedulesApi->get_account_schedule: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **schedule_id** | **int**| Schedule ID | 

### Return type

[**ScheduleFull**](ScheduleFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_schedules**
> ListSchedules list_account_schedules(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)

Get a list of schedules for an account

See Intro to Schedules for more info on the properties.

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
api_instance = swagger_client.SchedulesApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_name = ['filters_name_example'] # list[str] | Name filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_name = 'sort_name_example' # str | Name sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of schedules for an account
    api_response = api_instance.list_account_schedules(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SchedulesApi->list_account_schedules: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **filters_id** | [**list[str]**](str.md)| ID filter | [optional] 
 **filters_name** | [**list[str]**](str.md)| Name filter | [optional] 
 **sort_id** | **str**| ID sorting | [optional] 
 **sort_name** | **str**| Name sorting | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 

### Return type

[**ListSchedules**](ListSchedules.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

