# IO.Swagger.Api.AccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](AccountsApi.md#getaccount) | **GET** /accounts/{account_id} | Retrieve details of an individual account
[**ListAccounts**](AccountsApi.md#listaccounts) | **GET** /accounts | Get a list of accounts visible to the authenticated user or client


<a name="getaccount"></a>
# **GetAccount**
> AccountFull GetAccount (int? accountId)

Retrieve details of an individual account

This service shows the details of an individual account. See Accounts for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new AccountsApi();
            var accountId = 56;  // int? | Account ID

            try
            {
                // Retrieve details of an individual account
                AccountFull result = apiInstance.GetAccount(accountId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling AccountsApi.GetAccount: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 

### Return type

[**AccountFull**](AccountFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccounts"></a>
# **ListAccounts**
> ListAccounts ListAccounts (List<string> filtersId = null, string sortId = null, int? limit = null, int? offset = null, string fields = null)

Get a list of accounts visible to the authenticated user or client

This service lists the accounts accessible to the authenticated client. In most cases, there will only be one such account. See Accounts for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new AccountsApi();
            var filtersId = new List<string>(); // List<string> | ID filter (optional) 
            var sortId = sortId_example;  // string | ID sorting (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // Get a list of accounts visible to the authenticated user or client
                ListAccounts result = apiInstance.ListAccounts(filtersId, sortId, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling AccountsApi.ListAccounts: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

