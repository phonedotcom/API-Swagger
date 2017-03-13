# swagger_client.ContactsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_extension_contact**](ContactsApi.md#create_account_extension_contact) | **POST** /accounts/{account_id}/extensions/{extension_id}/contacts | Add a new address book contact for an extension
[**delete_account_extension_contact**](ContactsApi.md#delete_account_extension_contact) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | 
[**get_account_extension_contact**](ContactsApi.md#get_account_extension_contact) | **GET** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | Retrieve the details of an address book contact
[**list_account_extension_contacts**](ContactsApi.md#list_account_extension_contacts) | **GET** /accounts/{account_id}/extensions/{extension_id}/contacts | Show a list of address book contacts
[**replace_account_extension_contact**](ContactsApi.md#replace_account_extension_contact) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | 


# **create_account_extension_contact**
> ContactFull create_account_extension_contact(account_id, extension_id, data=data)

Add a new address book contact for an extension

For more on the input fields, see Account Contacts.

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
api_instance = swagger_client.ContactsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
data = swagger_client.CreateContactParams() # CreateContactParams | Contact data (optional)

try: 
    # Add a new address book contact for an extension
    api_response = api_instance.create_account_extension_contact(account_id, extension_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ContactsApi->create_account_extension_contact: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **data** | [**CreateContactParams**](CreateContactParams.md)| Contact data | [optional] 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_account_extension_contact**
> DeleteContact delete_account_extension_contact(account_id, extension_id, contact_id)





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
api_instance = swagger_client.ContactsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
contact_id = 56 # int | Contact ID

try: 
    # 
    api_response = api_instance.delete_account_extension_contact(account_id, extension_id, contact_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ContactsApi->delete_account_extension_contact: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **contact_id** | **int**| Contact ID | 

### Return type

[**DeleteContact**](DeleteContact.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_extension_contact**
> ContactFull get_account_extension_contact(account_id, extension_id, contact_id)

Retrieve the details of an address book contact

For more info on the fields shown, see Account Contacts.

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
api_instance = swagger_client.ContactsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
contact_id = 56 # int | Contact ID

try: 
    # Retrieve the details of an address book contact
    api_response = api_instance.get_account_extension_contact(account_id, extension_id, contact_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ContactsApi->get_account_extension_contact: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **contact_id** | **int**| Contact ID | 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_extension_contacts**
> ListContacts list_account_extension_contacts(account_id, extension_id, filters_id=filters_id, filters_group_id=filters_group_id, filters_updated_at=filters_updated_at, sort_id=sort_id, sort_updated_at=sort_updated_at, limit=limit, offset=offset, fields=fields)

Show a list of address book contacts

See Account Contacts for more info on the fields in each item.

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
api_instance = swagger_client.ContactsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_group_id = ['filters_group_id_example'] # list[str] | Group filter (optional)
filters_updated_at = ['filters_updated_at_example'] # list[str] | Updated At filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_updated_at = 'sort_updated_at_example' # str | Updated At sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Show a list of address book contacts
    api_response = api_instance.list_account_extension_contacts(account_id, extension_id, filters_id=filters_id, filters_group_id=filters_group_id, filters_updated_at=filters_updated_at, sort_id=sort_id, sort_updated_at=sort_updated_at, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ContactsApi->list_account_extension_contacts: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **filters_id** | [**list[str]**](str.md)| ID filter | [optional] 
 **filters_group_id** | [**list[str]**](str.md)| Group filter | [optional] 
 **filters_updated_at** | [**list[str]**](str.md)| Updated At filter | [optional] 
 **sort_id** | **str**| ID sorting | [optional] 
 **sort_updated_at** | **str**| Updated At sorting | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 

### Return type

[**ListContacts**](ListContacts.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace_account_extension_contact**
> ContactFull replace_account_extension_contact(account_id, extension_id, contact_id, data=data)



For more on the input fields, see Account Contacts.

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
api_instance = swagger_client.ContactsApi()
account_id = 56 # int | Account ID
extension_id = 56 # int | Extension ID
contact_id = 56 # int | Contact ID
data = swagger_client.CreateContactParams() # CreateContactParams | Contact data (optional)

try: 
    # 
    api_response = api_instance.replace_account_extension_contact(account_id, extension_id, contact_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ContactsApi->replace_account_extension_contact: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **extension_id** | **int**| Extension ID | 
 **contact_id** | **int**| Contact ID | 
 **data** | [**CreateContactParams**](CreateContactParams.md)| Contact data | [optional] 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

