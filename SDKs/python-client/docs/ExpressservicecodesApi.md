# swagger_client.ExpressservicecodesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_account_express_srv_code**](ExpressservicecodesApi.md#get_account_express_srv_code) | **GET** /accounts/{account_id}/express-service-codes/{code_id} | Show details of an account Express Service Code
[**list_account_express_srv_codes**](ExpressservicecodesApi.md#list_account_express_srv_codes) | **GET** /accounts/{account_id}/express-service-codes | Get the Express Service Code associated with your account in list format


# **get_account_express_srv_code**
> ExpressServiceCodeFull get_account_express_srv_code(account_id, code_id)

Show details of an account Express Service Code

This service shows the details of an Account Express Service Code.

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
api_instance = swagger_client.ExpressservicecodesApi()
account_id = 56 # int | Account ID
code_id = 56 # int | Device ID

try: 
    # Show details of an account Express Service Code
    api_response = api_instance.get_account_express_srv_code(account_id, code_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ExpressservicecodesApi->get_account_express_srv_code: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **code_id** | **int**| Device ID | 

### Return type

[**ExpressServiceCodeFull**](ExpressServiceCodeFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_express_srv_codes**
> ListExpressServiceCodes list_account_express_srv_codes(account_id, filters_id=filters_id)

Get the Express Service Code associated with your account in list format

See Express Service Codes for more detail.

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
api_instance = swagger_client.ExpressservicecodesApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)

try: 
    # Get the Express Service Code associated with your account in list format
    api_response = api_instance.list_account_express_srv_codes(account_id, filters_id=filters_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ExpressservicecodesApi->list_account_express_srv_codes: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **filters_id** | [**list[str]**](str.md)| ID filter | [optional] 

### Return type

[**ListExpressServiceCodes**](ListExpressServiceCodes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

