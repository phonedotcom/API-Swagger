# IO.Swagger.Api.CalllogsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountCallLogs**](CalllogsApi.md#getaccountcalllogs) | **GET** /accounts/{account_id}/call-logs/{call_id} | Show details of an individual Call Log entry
[**ListAccountCallLogs**](CalllogsApi.md#listaccountcalllogs) | **GET** /accounts/{account_id}/call-logs | Get a list of call details associated with your account


<a name="getaccountcalllogs"></a>
# **GetAccountCallLogs**
> CallLogFull GetAccountCallLogs (int? accountId, string callId)

Show details of an individual Call Log entry

See Call Logs for more detail.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountCallLogsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new CalllogsApi();
            var accountId = 56;  // int? | Account ID
            var callId = callId_example;  // string | Call ID

            try
            {
                // Show details of an individual Call Log entry
                CallLogFull result = apiInstance.GetAccountCallLogs(accountId, callId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling CalllogsApi.GetAccountCallLogs: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **callId** | **string**| Call ID | 

### Return type

[**CallLogFull**](CallLogFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountcalllogs"></a>
# **ListAccountCallLogs**
> ListCallLogs ListAccountCallLogs (int? accountId, List<string> filtersId = null, List<string> filtersStartTime = null, string filtersCreatedAt = null, string filtersDirection = null, string filtersCalledNumber = null, string filtersType = null, List<string> filtersExtension = null, string sortId = null, string sortStartTime = null, string sortCreatedAt = null, int? limit = null, int? offset = null, string fields = null)

Get a list of call details associated with your account

See Call Logs for more detail.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountCallLogsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new CalllogsApi();
            var accountId = 56;  // int? | Account ID
            var filtersId = new List<string>(); // List<string> | ID filter (optional) 
            var filtersStartTime = new List<string>(); // List<string> | Call start time filter (optional) 
            var filtersCreatedAt = filtersCreatedAt_example;  // string | Call log creation time filter (optional) 
            var filtersDirection = filtersDirection_example;  // string | Call direction filter: in or out (optional) 
            var filtersCalledNumber = filtersCalledNumber_example;  // string | Called number (optional) 
            var filtersType = filtersType_example;  // string | Call type, such as 'call', 'fax', 'audiogram' (optional) 
            var filtersExtension = new List<string>(); // List<string> | Extension filter (optional) 
            var sortId = sortId_example;  // string | ID sorting (optional) 
            var sortStartTime = sortStartTime_example;  // string | Sorting by call start time: asc or desc (optional) 
            var sortCreatedAt = sortCreatedAt_example;  // string | Sorting by call log creation time: asc or desc (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // Get a list of call details associated with your account
                ListCallLogs result = apiInstance.ListAccountCallLogs(accountId, filtersId, filtersStartTime, filtersCreatedAt, filtersDirection, filtersCalledNumber, filtersType, filtersExtension, sortId, sortStartTime, sortCreatedAt, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling CalllogsApi.ListAccountCallLogs: " + e.Message );
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
 **filtersStartTime** | [**List&lt;string&gt;**](string.md)| Call start time filter | [optional] 
 **filtersCreatedAt** | **string**| Call log creation time filter | [optional] 
 **filtersDirection** | **string**| Call direction filter: in or out | [optional] 
 **filtersCalledNumber** | **string**| Called number | [optional] 
 **filtersType** | **string**| Call type, such as &#39;call&#39;, &#39;fax&#39;, &#39;audiogram&#39; | [optional] 
 **filtersExtension** | [**List&lt;string&gt;**](string.md)| Extension filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortStartTime** | **string**| Sorting by call start time: asc or desc | [optional] 
 **sortCreatedAt** | **string**| Sorting by call log creation time: asc or desc | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListCallLogs**](ListCallLogs.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

