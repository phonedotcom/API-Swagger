# IO.Swagger.Api.SubaccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountSubaccount**](SubaccountsApi.md#createaccountsubaccount) | **POST** /accounts/{account_id}/subaccounts | Add a subaccount for the authenticated user or client
[**ListAccountSubaccounts**](SubaccountsApi.md#listaccountsubaccounts) | **GET** /accounts/{account_id}/subaccounts | Get a list of subaccounts for the authenticated user or client


<a name="createaccountsubaccount"></a>
# **CreateAccountSubaccount**
> AccountFull CreateAccountSubaccount (int? accountId, CreateSubaccountParams data)

Add a subaccount for the authenticated user or client

This service shows the details of an individual Subaccount.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAccountSubaccountExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new SubaccountsApi();
            var accountId = 56;  // int? | Account ID
            var data = new CreateSubaccountParams(); // CreateSubaccountParams | SMS data

            try
            {
                // Add a subaccount for the authenticated user or client
                AccountFull result = apiInstance.CreateAccountSubaccount(accountId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SubaccountsApi.CreateAccountSubaccount: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **data** | [**CreateSubaccountParams**](CreateSubaccountParams.md)| SMS data | 

### Return type

[**AccountFull**](AccountFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountsubaccounts"></a>
# **ListAccountSubaccounts**
> ListAccounts ListAccountSubaccounts (int? accountId, List<string> filtersId = null, string sortId = null, int? limit = null, int? offset = null, string fields = null)

Get a list of subaccounts for the authenticated user or client

This service lists the Subaccount of the authenticated client. In most cases, there will not be any.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountSubaccountsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new SubaccountsApi();
            var accountId = 56;  // int? | Account ID
            var filtersId = new List<string>(); // List<string> | ID filter (optional) 
            var sortId = sortId_example;  // string | ID sorting (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // Get a list of subaccounts for the authenticated user or client
                ListAccounts result = apiInstance.ListAccountSubaccounts(accountId, filtersId, sortId, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SubaccountsApi.ListAccountSubaccounts: " + e.Message );
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
 **sortId** | **string**| ID sorting | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListAccounts**](ListAccounts.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

