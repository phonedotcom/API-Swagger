# swagger_client.MenusApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_menu**](MenusApi.md#create_account_menu) | **POST** /accounts/{account_id}/menus | Create an individual menu
[**delete_account_menu**](MenusApi.md#delete_account_menu) | **DELETE** /accounts/{account_id}/menus/{menu_id} | Delete an individual menu
[**get_account_menu**](MenusApi.md#get_account_menu) | **GET** /accounts/{account_id}/menus/{menu_id} | Show details of an individual menu
[**list_account_menus**](MenusApi.md#list_account_menus) | **GET** /accounts/{account_id}/menus | Get a list of menus for an account
[**replace_account_menu**](MenusApi.md#replace_account_menu) | **PUT** /accounts/{account_id}/menus/{menu_id} | Replace an individual menu


# **create_account_menu**
> MenuFull create_account_menu(account_id, data=data)

Create an individual menu

This service creates an individual menu. See Account Menus for more info on the properties.

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
api_instance = swagger_client.MenusApi()
account_id = 56 # int | Account ID
data = swagger_client.CreateMenuParams() # CreateMenuParams | Menu data (optional)

try: 
    # Create an individual menu
    api_response = api_instance.create_account_menu(account_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MenusApi->create_account_menu: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **data** | [**CreateMenuParams**](CreateMenuParams.md)| Menu data | [optional] 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_account_menu**
> DeleteMenu delete_account_menu(account_id, menu_id)

Delete an individual menu

This service shows the details of an individual menu.

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
api_instance = swagger_client.MenusApi()
account_id = 56 # int | Account ID
menu_id = 56 # int | Menu ID

try: 
    # Delete an individual menu
    api_response = api_instance.delete_account_menu(account_id, menu_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MenusApi->delete_account_menu: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **menu_id** | **int**| Menu ID | 

### Return type

[**DeleteMenu**](DeleteMenu.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_menu**
> MenuFull get_account_menu(account_id, menu_id)

Show details of an individual menu

This service shows the details of an individual Menu.

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
api_instance = swagger_client.MenusApi()
account_id = 56 # int | Account ID
menu_id = 56 # int | Menu ID

try: 
    # Show details of an individual menu
    api_response = api_instance.get_account_menu(account_id, menu_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MenusApi->get_account_menu: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **menu_id** | **int**| Menu ID | 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_menus**
> ListMenus list_account_menus(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)

Get a list of menus for an account

See Account Menus for more info on the properties.

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
api_instance = swagger_client.MenusApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_name = ['filters_name_example'] # list[str] | Name filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_name = 'sort_name_example' # str | Name sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of menus for an account
    api_response = api_instance.list_account_menus(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MenusApi->list_account_menus: %s\n" % e)
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

[**ListMenus**](ListMenus.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace_account_menu**
> MenuFull replace_account_menu(account_id, menu_id, data=data)

Replace an individual menu

This service replaces the details of an individual Menu.

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
api_instance = swagger_client.MenusApi()
account_id = 56 # int | Account ID
menu_id = 56 # int | Menu ID
data = swagger_client.ReplaceMenuParams() # ReplaceMenuParams | Menu data (optional)

try: 
    # Replace an individual menu
    api_response = api_instance.replace_account_menu(account_id, menu_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling MenusApi->replace_account_menu: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **menu_id** | **int**| Menu ID | 
 **data** | [**ReplaceMenuParams**](ReplaceMenuParams.md)| Menu data | [optional] 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

