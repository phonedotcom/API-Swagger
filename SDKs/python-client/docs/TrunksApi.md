# swagger_client.TrunksApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_trunk**](TrunksApi.md#create_account_trunk) | **POST** /accounts/{account_id}/trunks | Add a trunk record with SIP information
[**delete_account_trunk**](TrunksApi.md#delete_account_trunk) | **DELETE** /accounts/{account_id}/trunks/{trunk_id} | Delete a trunk from account
[**get_account_trunk**](TrunksApi.md#get_account_trunk) | **GET** /accounts/{account_id}/trunks/{trunk_id} | Show details of an individual trunk
[**list_account_trunks**](TrunksApi.md#list_account_trunks) | **GET** /accounts/{account_id}/trunks | Get a list of trunks for an account
[**replace_account_trunk**](TrunksApi.md#replace_account_trunk) | **PUT** /accounts/{account_id}/trunks/{trunk_id} | Replace parameters in a trunk


# **create_account_trunk**
> TrunkFull create_account_trunk(account_id, data)

Add a trunk record with SIP information

For more on the input fields, see Account Trunks.

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
api_instance = swagger_client.TrunksApi()
account_id = 56 # int | Account ID
data = swagger_client.CreateTrunkParams() # CreateTrunkParams | Trunk data

try: 
    # Add a trunk record with SIP information
    api_response = api_instance.create_account_trunk(account_id, data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling TrunksApi->create_account_trunk: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **data** | [**CreateTrunkParams**](CreateTrunkParams.md)| Trunk data | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_account_trunk**
> DeleteTrunk delete_account_trunk(account_id, trunk_id)

Delete a trunk from account

This service deletes a trunk from the account. For more on the properties of trunks, see Account Trunks.

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
api_instance = swagger_client.TrunksApi()
account_id = 56 # int | Account ID
trunk_id = 56 # int | Trunk ID

try: 
    # Delete a trunk from account
    api_response = api_instance.delete_account_trunk(account_id, trunk_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling TrunksApi->delete_account_trunk: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **trunk_id** | **int**| Trunk ID | 

### Return type

[**DeleteTrunk**](DeleteTrunk.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_trunk**
> TrunkFull get_account_trunk(account_id, trunk_id)

Show details of an individual trunk

This service shows the details of an individual Trunk.

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
api_instance = swagger_client.TrunksApi()
account_id = 56 # int | Account ID
trunk_id = 56 # int | Trunk ID

try: 
    # Show details of an individual trunk
    api_response = api_instance.get_account_trunk(account_id, trunk_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling TrunksApi->get_account_trunk: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **trunk_id** | **int**| Trunk ID | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_trunks**
> ListTrunks list_account_trunks(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)

Get a list of trunks for an account

See Account Trunks for more info on the properties.

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
api_instance = swagger_client.TrunksApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_name = ['filters_name_example'] # list[str] | Name filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_name = 'sort_name_example' # str | Name sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of trunks for an account
    api_response = api_instance.list_account_trunks(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling TrunksApi->list_account_trunks: %s\n" % e)
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

[**ListTrunks**](ListTrunks.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace_account_trunk**
> TrunkFull replace_account_trunk(account_id, trunk_id, data)

Replace parameters in a trunk

For more on the input fields, see Account Trunks.

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
api_instance = swagger_client.TrunksApi()
account_id = 56 # int | Account ID
trunk_id = 56 # int | Trunk ID
data = swagger_client.CreateTrunkParams() # CreateTrunkParams | Trunk data

try: 
    # Replace parameters in a trunk
    api_response = api_instance.replace_account_trunk(account_id, trunk_id, data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling TrunksApi->replace_account_trunk: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **trunk_id** | **int**| Trunk ID | 
 **data** | [**CreateTrunkParams**](CreateTrunkParams.md)| Trunk data | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

