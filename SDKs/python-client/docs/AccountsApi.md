# swagger_client.AccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_account**](AccountsApi.md#get_account) | **GET** /accounts/{account_id} | Retrieve details of an individual account
[**list_accounts**](AccountsApi.md#list_accounts) | **GET** /accounts | Get a list of accounts visible to the authenticated user or client


# **get_account**
> AccountFull get_account(account_id)

Retrieve details of an individual account

This service shows the details of an individual account. See Accounts for more info on the properties.

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
api_instance = swagger_client.AccountsApi()
account_id = 56 # int | Account ID

try: 
    # Retrieve details of an individual account
    api_response = api_instance.get_account(account_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling AccountsApi->get_account: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 

### Return type

[**AccountFull**](AccountFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_accounts**
> ListAccounts list_accounts(filters_id=filters_id, sort_id=sort_id, limit=limit, offset=offset, fields=fields)

Get a list of accounts visible to the authenticated user or client

This service lists the accounts accessible to the authenticated client. In most cases, there will only be one such account. See Accounts for more info on the properties.

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
api_instance = swagger_client.AccountsApi()
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of accounts visible to the authenticated user or client
    api_response = api_instance.list_accounts(filters_id=filters_id, sort_id=sort_id, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling AccountsApi->list_accounts: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters_id** | [**list[str]**](str.md)| ID filter | [optional] 
 **sort_id** | **str**| ID sorting | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 

### Return type

[**ListAccounts**](ListAccounts.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

