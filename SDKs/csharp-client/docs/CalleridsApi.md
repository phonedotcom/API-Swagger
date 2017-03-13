# IO.Swagger.Api.CalleridsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCallerIds**](CalleridsApi.md#getcallerids) | **GET** /accounts/{account_id}/extensions/{extension_id}/caller-ids | Show the Caller ID options a given extension can use


<a name="getcallerids"></a>
# **GetCallerIds**
> ListCallerIds GetCallerIds (int? accountId, int? extensionId, List<string> filtersNumber = null, List<string> filtersName = null, string sortNumber = null, string sortName = null, int? limit = null, int? offset = null, string fields = null)

Show the Caller ID options a given extension can use

Get Caller ID

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetCallerIdsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new CalleridsApi();
            var accountId = 56;  // int? | Account ID
            var extensionId = 56;  // int? | Extension ID
            var filtersNumber = new List<string>(); // List<string> | Number filter (optional) 
            var filtersName = new List<string>(); // List<string> | Name filter (optional) 
            var sortNumber = sortNumber_example;  // string | Number sorting (optional) 
            var sortName = sortName_example;  // string | Name sorting (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // Show the Caller ID options a given extension can use
                ListCallerIds result = apiInstance.GetCallerIds(accountId, extensionId, filtersNumber, filtersName, sortNumber, sortName, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling CalleridsApi.GetCallerIds: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **extensionId** | **int?**| Extension ID | 
 **filtersNumber** | [**List&lt;string&gt;**](string.md)| Number filter | [optional] 
 **filtersName** | [**List&lt;string&gt;**](string.md)| Name filter | [optional] 
 **sortNumber** | **string**| Number sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListCallerIds**](ListCallerIds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

