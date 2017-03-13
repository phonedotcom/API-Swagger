# IO.Swagger.Api.ExpressservicecodesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountExpressSrvCode**](ExpressservicecodesApi.md#getaccountexpresssrvcode) | **GET** /accounts/{account_id}/express-service-codes/{code_id} | Show details of an account Express Service Code
[**ListAccountExpressSrvCodes**](ExpressservicecodesApi.md#listaccountexpresssrvcodes) | **GET** /accounts/{account_id}/express-service-codes | Get the Express Service Code associated with your account in list format


<a name="getaccountexpresssrvcode"></a>
# **GetAccountExpressSrvCode**
> ExpressServiceCodeFull GetAccountExpressSrvCode (int? accountId, int? codeId)

Show details of an account Express Service Code

This service shows the details of an Account Express Service Code.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountExpressSrvCodeExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new ExpressservicecodesApi();
            var accountId = 56;  // int? | Account ID
            var codeId = 56;  // int? | Device ID

            try
            {
                // Show details of an account Express Service Code
                ExpressServiceCodeFull result = apiInstance.GetAccountExpressSrvCode(accountId, codeId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling ExpressservicecodesApi.GetAccountExpressSrvCode: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **codeId** | **int?**| Device ID | 

### Return type

[**ExpressServiceCodeFull**](ExpressServiceCodeFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountexpresssrvcodes"></a>
# **ListAccountExpressSrvCodes**
> ListExpressServiceCodes ListAccountExpressSrvCodes (int? accountId, List<string> filtersId = null)

Get the Express Service Code associated with your account in list format

See Express Service Codes for more detail.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountExpressSrvCodesExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new ExpressservicecodesApi();
            var accountId = 56;  // int? | Account ID
            var filtersId = new List<string>(); // List<string> | ID filter (optional) 

            try
            {
                // Get the Express Service Code associated with your account in list format
                ListExpressServiceCodes result = apiInstance.ListAccountExpressSrvCodes(accountId, filtersId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling ExpressservicecodesApi.ListAccountExpressSrvCodes: " + e.Message );
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

### Return type

[**ListExpressServiceCodes**](ListExpressServiceCodes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

