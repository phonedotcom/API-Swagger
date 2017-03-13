# swagger_client.RoutesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_route**](RoutesApi.md#create_route) | **POST** /accounts/{account_id}/routes | Add a new address book contact for an extension
[**delete_account_route**](RoutesApi.md#delete_account_route) | **DELETE** /accounts/{account_id}/routes/{route_id} | 
[**get_account_route**](RoutesApi.md#get_account_route) | **GET** /accounts/{account_id}/routes/{route_id} | Show details of an individual route
[**list_account_routes**](RoutesApi.md#list_account_routes) | **GET** /accounts/{account_id}/routes | Get a list of routes for an account
[**replace_account_route**](RoutesApi.md#replace_account_route) | **PUT** /accounts/{account_id}/routes/{route_id} | 


# **create_route**
> RouteFull create_route(account_id, data=data)

Add a new address book contact for an extension

For more on the input fields, see Intro to Routes.

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
api_instance = swagger_client.RoutesApi()
account_id = 56 # int | Account ID
data = swagger_client.CreateRouteParams() # CreateRouteParams | Route data (optional)

try: 
    # Add a new address book contact for an extension
    api_response = api_instance.create_route(account_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RoutesApi->create_route: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **data** | [**CreateRouteParams**](CreateRouteParams.md)| Route data | [optional] 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_account_route**
> DeleteRoute delete_account_route(account_id, route_id)





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
api_instance = swagger_client.RoutesApi()
account_id = 56 # int | Account ID
route_id = 56 # int | Route ID

try: 
    # 
    api_response = api_instance.delete_account_route(account_id, route_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RoutesApi->delete_account_route: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **route_id** | **int**| Route ID | 

### Return type

[**DeleteRoute**](DeleteRoute.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_route**
> RouteFull get_account_route(account_id, route_id)

Show details of an individual route

This service shows the details of an individual route.

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
api_instance = swagger_client.RoutesApi()
account_id = 56 # int | Account ID
route_id = 56 # int | Route ID

try: 
    # Show details of an individual route
    api_response = api_instance.get_account_route(account_id, route_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RoutesApi->get_account_route: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **route_id** | **int**| Route ID | 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_routes**
> ListRoutes list_account_routes(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)

Get a list of routes for an account

See Intro to Routes for more info on the properties.

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
api_instance = swagger_client.RoutesApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_name = ['filters_name_example'] # list[str] | Name filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_name = 'sort_name_example' # str | Name sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of routes for an account
    api_response = api_instance.list_account_routes(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RoutesApi->list_account_routes: %s\n" % e)
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

[**ListRoutes**](ListRoutes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace_account_route**
> RouteFull replace_account_route(account_id, route_id, data=data)



For more on the input fields, see Intro to Routes.

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
api_instance = swagger_client.RoutesApi()
account_id = 56 # int | Account ID
route_id = 56 # int | Route ID
data = swagger_client.CreateRouteParams() # CreateRouteParams | Route data (optional)

try: 
    # 
    api_response = api_instance.replace_account_route(account_id, route_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling RoutesApi->replace_account_route: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **route_id** | **int**| Route ID | 
 **data** | [**CreateRouteParams**](CreateRouteParams.md)| Route data | [optional] 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

