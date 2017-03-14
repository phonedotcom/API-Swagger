# IO.Swagger.Api.CallsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountCall**](CallsApi.md#createaccountcall) | **POST** /accounts/{account_id}/calls | Make a phone call


<a name="createaccountcall"></a>
# **CreateAccountCall**
> CallFull CreateAccountCall (int? accountId, CreateCallParams data = null)

Make a phone call



### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAccountCallExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new CallsApi();
            var accountId = 56;  // int? | Account ID
            var data = new CreateCallParams(); // CreateCallParams | Call data (optional) 

            try
            {
                // Make a phone call
                CallFull result = apiInstance.CreateAccountCall(accountId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling CallsApi.CreateAccountCall: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **data** | [**CreateCallParams**](CreateCallParams.md)| Call data | [optional] 

### Return type

[**CallFull**](CallFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

