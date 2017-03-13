# swagger_client.QueuesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_queue**](QueuesApi.md#create_account_queue) | **POST** /accounts/{account_id}/queues | Create a queue
[**delete_account_queue**](QueuesApi.md#delete_account_queue) | **DELETE** /accounts/{account_id}/queues/{queue_id} | Delete a queue
[**get_account_queue**](QueuesApi.md#get_account_queue) | **GET** /accounts/{account_id}/queues/{queue_id} | Show details of an individual queue
[**list_account_queues**](QueuesApi.md#list_account_queues) | **GET** /accounts/{account_id}/queues | Get a list of queues for an account
[**replace_account_queue**](QueuesApi.md#replace_account_queue) | **PUT** /accounts/{account_id}/queues/{queue_id} | Replace a queue


# **create_account_queue**
> QueueFull create_account_queue(account_id, data=data)

Create a queue

For more on the input fields, see Account Queues.

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
api_instance = swagger_client.QueuesApi()
account_id = 56 # int | Account ID
data = swagger_client.CreateQueueParams() # CreateQueueParams | Queue data (optional)

try: 
    # Create a queue
    api_response = api_instance.create_account_queue(account_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling QueuesApi->create_account_queue: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional] 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_account_queue**
> DeleteQueue delete_account_queue(account_id, queue_id)

Delete a queue

This service a queue from the account. For more information on queue properties, see Account Queues.

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
api_instance = swagger_client.QueuesApi()
account_id = 56 # int | Account ID
queue_id = 56 # int | Queue ID

try: 
    # Delete a queue
    api_response = api_instance.delete_account_queue(account_id, queue_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling QueuesApi->delete_account_queue: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **queue_id** | **int**| Queue ID | 

### Return type

[**DeleteQueue**](DeleteQueue.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_queue**
> QueueFull get_account_queue(account_id, queue_id)

Show details of an individual queue

This service shows the details of an individual queue. For more on the input fields, see Account Queues.

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
api_instance = swagger_client.QueuesApi()
account_id = 56 # int | Account ID
queue_id = 56 # int | Queue ID

try: 
    # Show details of an individual queue
    api_response = api_instance.get_account_queue(account_id, queue_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling QueuesApi->get_account_queue: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **queue_id** | **int**| Queue ID | 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_queues**
> ListQueues list_account_queues(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)

Get a list of queues for an account

The List Queues service lists all the queues belong to the account. See Account Queues for more info on the properties.

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
api_instance = swagger_client.QueuesApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_name = ['filters_name_example'] # list[str] | Name filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_name = 'sort_name_example' # str | Name sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of queues for an account
    api_response = api_instance.list_account_queues(account_id, filters_id=filters_id, filters_name=filters_name, sort_id=sort_id, sort_name=sort_name, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling QueuesApi->list_account_queues: %s\n" % e)
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

[**ListQueues**](ListQueues.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace_account_queue**
> QueueFull replace_account_queue(account_id, queue_id, data=data)

Replace a queue

The Replace Queue service replaces the parameters of a queue. For more on the input fields, see Account Queues.

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
api_instance = swagger_client.QueuesApi()
account_id = 56 # int | Account ID
queue_id = 56 # int | Queue ID
data = swagger_client.CreateQueueParams() # CreateQueueParams | Queue data (optional)

try: 
    # Replace a queue
    api_response = api_instance.replace_account_queue(account_id, queue_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling QueuesApi->replace_account_queue: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **queue_id** | **int**| Queue ID | 
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional] 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

