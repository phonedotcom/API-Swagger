# swagger_client.ExtensionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_extension**](ExtensionsApi.md#create_account_extension) | **POST** /accounts/{account_id}/extensions | Create an individual extension
[**get_account_extension**](ExtensionsApi.md#get_account_extension) | **GET** /accounts/{account_id}/extensions/{extension_id} | Show details of an individual extension
[**list_account_extensions**](ExtensionsApi.md#list_account_extensions) | **GET** /accounts/{account_id}/extensions | Get a list of extensions visible to the authenticated user or client
[**replace_account_extension**](ExtensionsApi.md#replace_account_extension) | **PUT** /accounts/{account_id}/extensions/{extension_id} | Replace an individual extension


# **create_account_extension**
> ExtensionFull create_account_extension(account_id, data=data)

Create an individual extension

This service shows how to create a virtual extension.

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
api_instance = swagger_client.ExtensionsApi()
account_id = 56 # int | Account ID
data = swagger_client.CreateExtensionParams() # CreateExtensionParams | Account Extensions Data (optional)

try: 
    # Create an individual extension
    api_response = api_instance.create_account_extension(account_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ExtensionsApi->create_account_extension: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **data** | [**CreateExtensionParams**](CreateExtensionParams.md)| Account Extensions Data | [optional] 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_extension**
> ExtensionFull get_account_extension(account_id, extension_id)

Show details of an individual extension

This service shows the details of an individual Extension.

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
api_instance = swagger_client.ExtensionsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID

try: 
    # Show details of an individual extension
    api_response = api_instance.get_account_extension(account_id, extension_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ExtensionsApi->get_account_extension: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_extensions**
> ListExtensions list_account_extensions(account_id, filters_id=filters_id, filters_extension=filters_extension, filters_name=filters_name, sort_id=sort_id, sort_extension=sort_extension, sort_name=sort_name, limit=limit, offset=offset, fields=fields)

Get a list of extensions visible to the authenticated user or client

This service lists the visible extensions on a given account.

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
api_instance = swagger_client.ExtensionsApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_extension = ['filters_extension_example'] # list[str] | Extension filter (optional)
filters_name = ['filters_name_example'] # list[str] | Name filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_extension = 'sort_extension_example' # str | Extension sorting (optional)
sort_name = 'sort_name_example' # str | Name sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of extensions visible to the authenticated user or client
    api_response = api_instance.list_account_extensions(account_id, filters_id=filters_id, filters_extension=filters_extension, filters_name=filters_name, sort_id=sort_id, sort_extension=sort_extension, sort_name=sort_name, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ExtensionsApi->list_account_extensions: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **filters_id** | [**list[str]**](str.md)| ID filter | [optional] 
 **filters_extension** | [**list[str]**](str.md)| Extension filter | [optional] 
 **filters_name** | [**list[str]**](str.md)| Name filter | [optional] 
 **sort_id** | **str**| ID sorting | [optional] 
 **sort_extension** | **str**| Extension sorting | [optional] 
 **sort_name** | **str**| Name sorting | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 

### Return type

[**ListExtensions**](ListExtensions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace_account_extension**
> ExtensionFull replace_account_extension(account_id, extension_id, data=data)

Replace an individual extension

This service shows how to update an individual extension.

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
api_instance = swagger_client.ExtensionsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
data = swagger_client.ReplaceExtensionParams() # ReplaceExtensionParams | Account Extensions Data (optional)

try: 
    # Replace an individual extension
    api_response = api_instance.replace_account_extension(account_id, extension_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ExtensionsApi->replace_account_extension: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **data** | [**ReplaceExtensionParams**](ReplaceExtensionParams.md)| Account Extensions Data | [optional] 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

