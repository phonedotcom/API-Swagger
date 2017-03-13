# swagger_client.GroupsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_extension_contact_group**](GroupsApi.md#create_account_extension_contact_group) | **POST** /accounts/{account_id}/extensions/{extension_id}/contact-groups | 
[**delete_account_extension_contact_group**](GroupsApi.md#delete_account_extension_contact_group) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | Delete an addressbook group
[**get_account_extension_contact_group**](GroupsApi.md#get_account_extension_contact_group) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 
[**list_account_extension_contact_groups**](GroupsApi.md#list_account_extension_contact_groups) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups | Show a list of contact groups belonging to an extension
[**replace_account_extension_contact_group**](GroupsApi.md#replace_account_extension_contact_group) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 


# **create_account_extension_contact_group**
> GroupFull create_account_extension_contact_group(account_id, extension_id, data)



See Account Contact Groups for more info on the properties.

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
api_instance = swagger_client.GroupsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
data = swagger_client.CreateGroupParams() # CreateGroupParams | Group name

try: 
    # 
    api_response = api_instance.create_account_extension_contact_group(account_id, extension_id, data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling GroupsApi->create_account_extension_contact_group: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **data** | [**CreateGroupParams**](CreateGroupParams.md)| Group name | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_account_extension_contact_group**
> DeleteGroup delete_account_extension_contact_group(account_id, extension_id, group_id)

Delete an addressbook group



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
api_instance = swagger_client.GroupsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
group_id = 56 # int | Group ID

try: 
    # Delete an addressbook group
    api_response = api_instance.delete_account_extension_contact_group(account_id, extension_id, group_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling GroupsApi->delete_account_extension_contact_group: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **group_id** | **int**| Group ID | 

### Return type

[**DeleteGroup**](DeleteGroup.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_extension_contact_group**
> GroupFull get_account_extension_contact_group(account_id, extension_id, group_id)



See Account Contact Groups for more info on the properties.

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
api_instance = swagger_client.GroupsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
group_id = 56 # int | Group ID

try: 
    # 
    api_response = api_instance.get_account_extension_contact_group(account_id, extension_id, group_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling GroupsApi->get_account_extension_contact_group: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **group_id** | **int**| Group ID | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_extension_contact_groups**
> ListGroups list_account_extension_contact_groups(account_id, extension_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)

Show a list of contact groups belonging to an extension

See Account Contact Groups for details on the properties.

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
api_instance = swagger_client.GroupsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_name = ['filters_name_example'] # list[str] | Name filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_name = 'sort_name_example' # str | Name sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Show a list of contact groups belonging to an extension
    api_response = api_instance.list_account_extension_contact_groups(account_id, extension_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling GroupsApi->list_account_extension_contact_groups: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **filters_id** | [**list[str]**](str.md)| ID filter | [optional] 
 **filters_name** | [**list[str]**](str.md)| Name filter | [optional] 
 **sort_id** | **str**| ID sorting | [optional] 
 **sort_name** | **str**| Name sorting | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 

### Return type

[**ListGroups**](ListGroups.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace_account_extension_contact_group**
> GroupFull replace_account_extension_contact_group(account_id, extension_id, group_id, data)



See Account Contact Groups for more info on the properties.

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
api_instance = swagger_client.GroupsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
group_id = 56 # int | Group ID
data = swagger_client.CreateGroupParams() # CreateGroupParams | Group name

try: 
    # 
    api_response = api_instance.replace_account_extension_contact_group(account_id, extension_id, group_id, data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling GroupsApi->replace_account_extension_contact_group: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **group_id** | **int**| Group ID | 
 **data** | [**CreateGroupParams**](CreateGroupParams.md)| Group name | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

