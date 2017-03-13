# IO.Swagger.Api.TrunksApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountTrunk**](TrunksApi.md#createaccounttrunk) | **POST** /accounts/{account_id}/trunks | Add a trunk record with SIP information
[**DeleteAccountTrunk**](TrunksApi.md#deleteaccounttrunk) | **DELETE** /accounts/{account_id}/trunks/{trunk_id} | Delete a trunk from account
[**GetAccountTrunk**](TrunksApi.md#getaccounttrunk) | **GET** /accounts/{account_id}/trunks/{trunk_id} | Show details of an individual trunk
[**ListAccountTrunks**](TrunksApi.md#listaccounttrunks) | **GET** /accounts/{account_id}/trunks | Get a list of trunks for an account
[**ReplaceAccountTrunk**](TrunksApi.md#replaceaccounttrunk) | **PUT** /accounts/{account_id}/trunks/{trunk_id} | Replace parameters in a trunk


<a name="createaccounttrunk"></a>
# **CreateAccountTrunk**
> TrunkFull CreateAccountTrunk (int? accountId, CreateTrunkParams data)

Add a trunk record with SIP information

For more on the input fields, see Account Trunks.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAccountTrunkExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new TrunksApi();
            var accountId = 56;  // int? | Account ID
            var data = new CreateTrunkParams(); // CreateTrunkParams | Trunk data

            try
            {
                // Add a trunk record with SIP information
                TrunkFull result = apiInstance.CreateAccountTrunk(accountId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling TrunksApi.CreateAccountTrunk: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **data** | [**CreateTrunkParams**](CreateTrunkParams.md)| Trunk data | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="deleteaccounttrunk"></a>
# **DeleteAccountTrunk**
> DeleteTrunk DeleteAccountTrunk (int? accountId, int? trunkId)

Delete a trunk from account

This service deletes a trunk from the account. For more on the properties of trunks, see Account Trunks.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class DeleteAccountTrunkExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new TrunksApi();
            var accountId = 56;  // int? | Account ID
            var trunkId = 56;  // int? | Trunk ID

            try
            {
                // Delete a trunk from account
                DeleteTrunk result = apiInstance.DeleteAccountTrunk(accountId, trunkId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling TrunksApi.DeleteAccountTrunk: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **trunkId** | **int?**| Trunk ID | 

### Return type

[**DeleteTrunk**](DeleteTrunk.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getaccounttrunk"></a>
# **GetAccountTrunk**
> TrunkFull GetAccountTrunk (int? accountId, int? trunkId)

Show details of an individual trunk

This service shows the details of an individual Trunk.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountTrunkExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new TrunksApi();
            var accountId = 56;  // int? | Account ID
            var trunkId = 56;  // int? | Trunk ID

            try
            {
                // Show details of an individual trunk
                TrunkFull result = apiInstance.GetAccountTrunk(accountId, trunkId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling TrunksApi.GetAccountTrunk: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **trunkId** | **int?**| Trunk ID | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccounttrunks"></a>
# **ListAccountTrunks**
> ListTrunks ListAccountTrunks (int? accountId, List<string> filtersId = null, List<string> filtersName = null, string sortId = null, string sortName = null, int? limit = null, int? offset = null, string fields = null)

Get a list of trunks for an account

See Account Trunks for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountTrunksExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new TrunksApi();
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
                // Get a list of trunks for an account
                ListTrunks result = apiInstance.ListAccountTrunks(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling TrunksApi.ListAccountTrunks: " + e.Message );
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

[**ListTrunks**](ListTrunks.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="replaceaccounttrunk"></a>
# **ReplaceAccountTrunk**
> TrunkFull ReplaceAccountTrunk (int? accountId, int? trunkId, CreateTrunkParams data)

Replace parameters in a trunk

For more on the input fields, see Account Trunks.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ReplaceAccountTrunkExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new TrunksApi();
            var accountId = 56;  // int? | Account ID
            var trunkId = 56;  // int? | Trunk ID
            var data = new CreateTrunkParams(); // CreateTrunkParams | Trunk data

            try
            {
                // Replace parameters in a trunk
                TrunkFull result = apiInstance.ReplaceAccountTrunk(accountId, trunkId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling TrunksApi.ReplaceAccountTrunk: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **trunkId** | **int?**| Trunk ID | 
 **data** | [**CreateTrunkParams**](CreateTrunkParams.md)| Trunk data | 

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

