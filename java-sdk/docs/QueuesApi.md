# QueuesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountQueue**](QueuesApi.md#createAccountQueue) | **POST** /accounts/{account_id}/queues | Create a queue
[**deleteAccountQueue**](QueuesApi.md#deleteAccountQueue) | **DELETE** /accounts/{account_id}/queues/{queue_id} | Delete a queue
[**getAccountQueue**](QueuesApi.md#getAccountQueue) | **GET** /accounts/{account_id}/queues/{queue_id} | Show details of an individual queue
[**listAccountQueues**](QueuesApi.md#listAccountQueues) | **GET** /accounts/{account_id}/queues | Get a list of queues for an account
[**replaceAccountQueue**](QueuesApi.md#replaceAccountQueue) | **PUT** /accounts/{account_id}/queues/{queue_id} | Replace a queue


<a name="createAccountQueue"></a>
# **createAccountQueue**
> QueueFull createAccountQueue(accountId, data)

Create a queue

For more on the input fields, see Account Queues.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.QueuesApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

QueuesApi apiInstance = new QueuesApi();
Integer accountId = 56; // Integer | Account ID
CreateQueueParams data = new CreateQueueParams(); // CreateQueueParams | Queue data
try {
    QueueFull result = apiInstance.createAccountQueue(accountId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling QueuesApi#createAccountQueue");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional]

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteAccountQueue"></a>
# **deleteAccountQueue**
> DeleteQueue deleteAccountQueue(accountId, queueId)

Delete a queue

This service a queue from the account. For more information on queue properties, see Account Queues.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.QueuesApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

QueuesApi apiInstance = new QueuesApi();
Integer accountId = 56; // Integer | Account ID
Integer queueId = 56; // Integer | Queue ID
try {
    DeleteQueue result = apiInstance.deleteAccountQueue(accountId, queueId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling QueuesApi#deleteAccountQueue");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **queueId** | **Integer**| Queue ID |

### Return type

[**DeleteQueue**](DeleteQueue.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getAccountQueue"></a>
# **getAccountQueue**
> QueueFull getAccountQueue(accountId, queueId)

Show details of an individual queue

This service shows the details of an individual queue. For more on the input fields, see Account Queues.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.QueuesApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

QueuesApi apiInstance = new QueuesApi();
Integer accountId = 56; // Integer | Account ID
Integer queueId = 56; // Integer | Queue ID
try {
    QueueFull result = apiInstance.getAccountQueue(accountId, queueId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling QueuesApi#getAccountQueue");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **queueId** | **Integer**| Queue ID |

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listAccountQueues"></a>
# **listAccountQueues**
> ListQueuesFull listAccountQueues(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields)

Get a list of queues for an account

The List Queues service lists all the queues belong to the account. See Account Queues for more info on the properties.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.QueuesApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

QueuesApi apiInstance = new QueuesApi();
Integer accountId = 56; // Integer | Account ID
List<String> filtersId = Arrays.asList("filtersId_example"); // List<String> | ID filter
List<String> filtersName = Arrays.asList("filtersName_example"); // List<String> | Name filter
String sortId = "sortId_example"; // String | ID sorting
String sortName = "sortName_example"; // String | Name sorting
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
try {
    ListQueuesFull result = apiInstance.listAccountQueues(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling QueuesApi#listAccountQueues");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **filtersId** | [**List&lt;String&gt;**](String.md)| ID filter | [optional]
 **filtersName** | [**List&lt;String&gt;**](String.md)| Name filter | [optional]
 **sortId** | **String**| ID sorting | [optional]
 **sortName** | **String**| Name sorting | [optional]
 **limit** | **Integer**| Max results | [optional]
 **offset** | **Integer**| Results to skip | [optional]
 **fields** | **String**| Field set | [optional]

### Return type

[**ListQueuesFull**](ListQueuesFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="replaceAccountQueue"></a>
# **replaceAccountQueue**
> QueueFull replaceAccountQueue(accountId, queueId, data)

Replace a queue

The Replace Queue service replaces the parameters of a queue. For more on the input fields, see Account Queues.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.QueuesApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

QueuesApi apiInstance = new QueuesApi();
Integer accountId = 56; // Integer | Account ID
Integer queueId = 56; // Integer | Queue ID
CreateQueueParams data = new CreateQueueParams(); // CreateQueueParams | Queue data
try {
    QueueFull result = apiInstance.replaceAccountQueue(accountId, queueId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling QueuesApi#replaceAccountQueue");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **queueId** | **Integer**| Queue ID |
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional]

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

