# IO.Swagger.Api.ExtensionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountExtension**](ExtensionsApi.md#createaccountextension) | **POST** /accounts/{account_id}/extensions | Create an individual extension
[**GetAccountExtension**](ExtensionsApi.md#getaccountextension) | **GET** /accounts/{account_id}/extensions/{extension_id} | Show details of an individual extension
[**ListAccountExtensions**](ExtensionsApi.md#listaccountextensions) | **GET** /accounts/{account_id}/extensions | Get a list of extensions visible to the authenticated user or client
[**ReplaceAccountExtension**](ExtensionsApi.md#replaceaccountextension) | **PUT** /accounts/{account_id}/extensions/{extension_id} | Replace an individual extension


<a name="createaccountextension"></a>
# **CreateAccountExtension**
> ExtensionFull CreateAccountExtension (int? accountId, CreateExtensionParams data = null)

Create an individual extension

This service shows how to create a virtual extension.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAccountExtensionExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new ExtensionsApi();
            var accountId = 56;  // int? | Account ID
            var data = new CreateExtensionParams(); // CreateExtensionParams | Account Extensions Data (optional) 

            try
            {
                // Create an individual extension
                ExtensionFull result = apiInstance.CreateAccountExtension(accountId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling ExtensionsApi.CreateAccountExtension: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **data** | [**CreateExtensionParams**](CreateExtensionParams.md)| Account Extensions Data | [optional] 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getaccountextension"></a>
# **GetAccountExtension**
> ExtensionFull GetAccountExtension (int? accountId, int? extensionId)

Show details of an individual extension

This service shows the details of an individual Extension.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountExtensionExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new ExtensionsApi();
            var accountId = 56;  // int? | Account ID
            var extensionId = 56;  // int? | Extension ID

            try
            {
                // Show details of an individual extension
                ExtensionFull result = apiInstance.GetAccountExtension(accountId, extensionId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling ExtensionsApi.GetAccountExtension: " + e.Message );
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

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountextensions"></a>
# **ListAccountExtensions**
> ListExtensions ListAccountExtensions (int? accountId, List<string> filtersId = null, List<string> filtersExtension = null, List<string> filtersName = null, string sortId = null, string sortExtension = null, string sortName = null, int? limit = null, int? offset = null, string fields = null)

Get a list of extensions visible to the authenticated user or client

This service lists the visible extensions on a given account.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountExtensionsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new ExtensionsApi();
            var accountId = 56;  // int? | Account ID
            var filtersId = new List<string>(); // List<string> | ID filter (optional) 
            var filtersExtension = new List<string>(); // List<string> | Extension filter (optional) 
            var filtersName = new List<string>(); // List<string> | Name filter (optional) 
            var sortId = sortId_example;  // string | ID sorting (optional) 
            var sortExtension = sortExtension_example;  // string | Extension sorting (optional) 
            var sortName = sortName_example;  // string | Name sorting (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // Get a list of extensions visible to the authenticated user or client
                ListExtensions result = apiInstance.ListAccountExtensions(accountId, filtersId, filtersExtension, filtersName, sortId, sortExtension, sortName, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling ExtensionsApi.ListAccountExtensions: " + e.Message );
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
 **filtersExtension** | [**List&lt;string&gt;**](string.md)| Extension filter | [optional] 
 **filtersName** | [**List&lt;string&gt;**](string.md)| Name filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortExtension** | **string**| Extension sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListExtensions**](ListExtensions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="replaceaccountextension"></a>
# **ReplaceAccountExtension**
> ExtensionFull ReplaceAccountExtension (int? accountId, int? extensionId, ReplaceExtensionParams data = null)

Replace an individual extension

This service shows how to update an individual extension.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ReplaceAccountExtensionExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new ExtensionsApi();
            var accountId = 56;  // int? | Account ID
            var extensionId = 56;  // int? | Extension ID
            var data = new ReplaceExtensionParams(); // ReplaceExtensionParams | Account Extensions Data (optional) 

            try
            {
                // Replace an individual extension
                ExtensionFull result = apiInstance.ReplaceAccountExtension(accountId, extensionId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling ExtensionsApi.ReplaceAccountExtension: " + e.Message );
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
 **data** | [**ReplaceExtensionParams**](ReplaceExtensionParams.md)| Account Extensions Data | [optional] 

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

