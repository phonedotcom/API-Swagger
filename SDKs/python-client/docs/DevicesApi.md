# swagger_client.DevicesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_device**](DevicesApi.md#create_account_device) | **POST** /accounts/{account_id}/devices | Register a generic VoIP device
[**get_account_device**](DevicesApi.md#get_account_device) | **GET** /accounts/{account_id}/devices/{device_id} | Show details of an individual VoIP device
[**list_account_devices**](DevicesApi.md#list_account_devices) | **GET** /accounts/{account_id}/devices | Get a list of VoIP devices associated with your account
[**replace_account_device**](DevicesApi.md#replace_account_device) | **PUT** /accounts/{account_id}/devices/{device_id} | Update the settings for an individual VoIP device


# **create_account_device**
> DeviceFull create_account_device(account_id, data=data)

Register a generic VoIP device



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
api_instance = swagger_client.DevicesApi()
account_id = 56 # int | Account ID
data = swagger_client.CreateDeviceParams() # CreateDeviceParams | Device data (optional)

try: 
    # Register a generic VoIP device
    api_response = api_instance.create_account_device(account_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DevicesApi->create_account_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **data** | [**CreateDeviceParams**](CreateDeviceParams.md)| Device data | [optional] 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_device**
> DeviceFull get_account_device(account_id, device_id)

Show details of an individual VoIP device



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
api_instance = swagger_client.DevicesApi()
account_id = 56 # int | Account ID
device_id = 56 # int | Device ID

try: 
    # Show details of an individual VoIP device
    api_response = api_instance.get_account_device(account_id, device_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DevicesApi->get_account_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **device_id** | **int**| Device ID | 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_devices**
> ListDevices list_account_devices(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)

Get a list of VoIP devices associated with your account



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
api_instance = swagger_client.DevicesApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_name = ['filters_name_example'] # list[str] | Name filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_name = 'sort_name_example' # str | Name sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of VoIP devices associated with your account
    api_response = api_instance.list_account_devices(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DevicesApi->list_account_devices: %s\n" % e)
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

[**ListDevices**](ListDevices.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace_account_device**
> DeviceFull replace_account_device(account_id, device_id, data=data)

Update the settings for an individual VoIP device



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
api_instance = swagger_client.DevicesApi()
account_id = 56 # int | Account ID
device_id = 56 # int | Device ID
data = swagger_client.CreateDeviceParams() # CreateDeviceParams | Device data (optional)

try: 
    # Update the settings for an individual VoIP device
    api_response = api_instance.replace_account_device(account_id, device_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DevicesApi->replace_account_device: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **device_id** | **int**| Device ID | 
 **data** | [**CreateDeviceParams**](CreateDeviceParams.md)| Device data | [optional] 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

