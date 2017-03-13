# swagger_client.SmsApi

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
api_instance = swagger_client.SmsApi()
account_id = 56 # int | Account ID
data = swagger_client.CreateSmsParams() # CreateSmsParams | SMS data

try: 
    # Send a SMS to one or a group of recipients
    api_response = api_instance.create_account_sms(account_id, data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SmsApi->create_account_sms: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **data** | [**CreateSmsParams**](CreateSmsParams.md)| SMS data | 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_sms**
> SmsFull get_account_sms(account_id, sms_id)

Show details of an individual SMS

This service shows the details of an individual sms. See Intro to SMS for more info on the properties.

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
api_instance = swagger_client.SmsApi()
account_id = 56 # int | Account ID
sms_id = 'sms_id_example' # str | SMS ID

try: 
    # Show details of an individual SMS
    api_response = api_instance.get_account_sms(account_id, sms_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SmsApi->get_account_sms: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **sms_id** | **str**| SMS ID | 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_sms**
> ListSms list_account_sms(account_id, filters_id=filters_id, filters_direction=filters_direction, filters_from=filters_from, sort_id=sort_id, sort_created_at=sort_created_at, limit=limit, offset=offset, fields=fields)

Get a list of SMS messages for an account

See Intro to SMS for more info on the properties.

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
api_instance = swagger_client.SmsApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_direction = 'filters_direction_example' # str | Direction filter (optional)
filters_from = 'filters_from_example' # str | Caller ID filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_created_at = 'sort_created_at_example' # str | Sort by created time of message (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of SMS messages for an account
    api_response = api_instance.list_account_sms(account_id, filters_id=filters_id, filters_direction=filters_direction, filters_from=filters_from, sort_id=sort_id, sort_created_at=sort_created_at, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SmsApi->list_account_sms: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **filters_id** | [**list[str]**](str.md)| ID filter | [optional] 
 **filters_direction** | **str**| Direction filter | [optional] 
 **filters_from** | **str**| Caller ID filter | [optional] 
 **sort_id** | **str**| ID sorting | [optional] 
 **sort_created_at** | **str**| Sort by created time of message | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 

### Return type

[**ListSms**](ListSms.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

