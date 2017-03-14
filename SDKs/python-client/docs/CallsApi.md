# swagger_client.CallsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_call**](CallsApi.md#create_account_call) | **POST** /accounts/{account_id}/calls | Make a phone call


# **create_account_call**
> CallFull create_account_call(account_id, data=data)

Make a phone call



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
api_instance = swagger_client.CallsApi()
account_id = 56 # int | Account ID
data = swagger_client.CreateCallParams() # CreateCallParams | Call data (optional)

try: 
    # Make a phone call
    api_response = api_instance.create_account_call(account_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling CallsApi->create_account_call: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **data** | [**CreateCallParams**](CreateCallParams.md)| Call data | [optional] 

### Return type

[**CallFull**](CallFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

