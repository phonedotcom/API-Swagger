# swagger_client.SubaccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_subaccount**](SubaccountsApi.md#create_account_subaccount) | **POST** /accounts/{account_id}/subaccounts | Add a subaccount for the authenticated user or client
[**list_account_subaccounts**](SubaccountsApi.md#list_account_subaccounts) | **GET** /accounts/{account_id}/subaccounts | Get a list of subaccounts for the authenticated user or client


# **create_account_subaccount**
> AccountFull create_account_subaccount(account_id, data)

Add a subaccount for the authenticated user or client

This service shows the details of an individual Subaccount.

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
api_instance = swagger_client.SubaccountsApi()
account_id = 56 # int | Account ID
data = swagger_client.CreateSubaccountParams() # CreateSubaccountParams | SMS data

try: 
    # Add a subaccount for the authenticated user or client
    api_response = api_instance.create_account_subaccount(account_id, data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SubaccountsApi->create_account_subaccount: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **data** | [**CreateSubaccountParams**](CreateSubaccountParams.md)| SMS data | 

### Return type

[**AccountFull**](AccountFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_subaccounts**
> ListAccounts list_account_subaccounts(account_id, filters_id=filters_id, sort_id=sort_id, limit=limit, offset=offset, fields=fields)

Get a list of subaccounts for the authenticated user or client

This service lists the Subaccount of the authenticated client. In most cases, there will not be any.

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
api_instance = swagger_client.SubaccountsApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of subaccounts for the authenticated user or client
    api_response = api_instance.list_account_subaccounts(account_id, filters_id=filters_id, sort_id=sort_id, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SubaccountsApi->list_account_subaccounts: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
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

