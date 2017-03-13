# IO.Swagger.Api.QueuesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountQueue**](QueuesApi.md#createaccountqueue) | **POST** /accounts/{account_id}/queues | Create a queue
[**DeleteAccountQueue**](QueuesApi.md#deleteaccountqueue) | **DELETE** /accounts/{account_id}/queues/{queue_id} | Delete a queue
[**GetAccountQueue**](QueuesApi.md#getaccountqueue) | **GET** /accounts/{account_id}/queues/{queue_id} | Show details of an individual queue
[**ListAccountQueues**](QueuesApi.md#listaccountqueues) | **GET** /accounts/{account_id}/queues | Get a list of queues for an account
[**ReplaceAccountQueue**](QueuesApi.md#replaceaccountqueue) | **PUT** /accounts/{account_id}/queues/{queue_id} | Replace a queue


<a name="createaccountqueue"></a>
# **CreateAccountQueue**
> QueueFull CreateAccountQueue (int? accountId, CreateQueueParams data = null)

Create a queue

For more on the input fields, see Account Queues.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAccountQueueExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new QueuesApi();
            var accountId = 56;  // int? | Account ID
            var data = new CreateQueueParams(); // CreateQueueParams | Queue data (optional) 

            try
            {
                // Create a queue
                QueueFull result = apiInstance.CreateAccountQueue(accountId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling QueuesApi.CreateAccountQueue: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional] 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="deleteaccountqueue"></a>
# **DeleteAccountQueue**
> DeleteQueue DeleteAccountQueue (int? accountId, int? queueId)

Delete a queue

This service a queue from the account. For more information on queue properties, see Account Queues.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class DeleteAccountQueueExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new QueuesApi();
            var accountId = 56;  // int? | Account ID
            var queueId = 56;  // int? | Queue ID

            try
            {
                // Delete a queue
                DeleteQueue result = apiInstance.DeleteAccountQueue(accountId, queueId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling QueuesApi.DeleteAccountQueue: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **queueId** | **int?**| Queue ID | 

### Return type

[**DeleteQueue**](DeleteQueue.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getaccountqueue"></a>
# **GetAccountQueue**
> QueueFull GetAccountQueue (int? accountId, int? queueId)

Show details of an individual queue

This service shows the details of an individual queue. For more on the input fields, see Account Queues.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountQueueExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new QueuesApi();
            var accountId = 56;  // int? | Account ID
            var queueId = 56;  // int? | Queue ID

            try
            {
                // Show details of an individual queue
                QueueFull result = apiInstance.GetAccountQueue(accountId, queueId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling QueuesApi.GetAccountQueue: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **queueId** | **int?**| Queue ID | 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountqueues"></a>
# **ListAccountQueues**
> ListQueues ListAccountQueues (int? accountId, List<string> filtersId = null, List<string> filtersName = null, string sortId = null, string sortName = null, int? limit = null, int? offset = null, string fields = null)

Get a list of queues for an account

The List Queues service lists all the queues belong to the account. See Account Queues for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountQueuesExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new QueuesApi();
            var accountId = 56;  // int? | Account ID
            var filtersId = new List<string>(); // List<string> | ID filter (optional) 
            var filtersName = new List<string>(); // List<string> | Name filter (optional) 
            var sortId = sortId_example;  // string | ID sorting (optional) 
            var sortName = sortName_example;  // string | Name sorting (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // Get a list of queues for an account
                ListQueues result = apiInstance.ListAccountQueues(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling QueuesApi.ListAccountQueues: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **filtersId** | [**List&lt;string&gt;**](string.md)| ID filter | [optional] 
 **filtersName** | [**List&lt;string&gt;**](string.md)| Name filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListQueues**](ListQueues.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="replaceaccountqueue"></a>
# **ReplaceAccountQueue**
> QueueFull ReplaceAccountQueue (int? accountId, int? queueId, CreateQueueParams data = null)

Replace a queue

The Replace Queue service replaces the parameters of a queue. For more on the input fields, see Account Queues.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ReplaceAccountQueueExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new QueuesApi();
            var accountId = 56;  // int? | Account ID
            var queueId = 56;  // int? | Queue ID
            var data = new CreateQueueParams(); // CreateQueueParams | Queue data (optional) 

            try
            {
                // Replace a queue
                QueueFull result = apiInstance.ReplaceAccountQueue(accountId, queueId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling QueuesApi.ReplaceAccountQueue: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **queueId** | **int?**| Queue ID | 
 **data** | [**CreateQueueParams**](CreateQueueParams.md)| Queue data | [optional] 

### Return type

[**QueueFull**](QueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

