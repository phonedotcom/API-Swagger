# IO.Swagger.Api.RoutesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoute**](RoutesApi.md#createroute) | **POST** /accounts/{account_id}/routes | Add a new address book contact for an extension
[**DeleteAccountRoute**](RoutesApi.md#deleteaccountroute) | **DELETE** /accounts/{account_id}/routes/{route_id} | 
[**GetAccountRoute**](RoutesApi.md#getaccountroute) | **GET** /accounts/{account_id}/routes/{route_id} | Show details of an individual route
[**ListAccountRoutes**](RoutesApi.md#listaccountroutes) | **GET** /accounts/{account_id}/routes | Get a list of routes for an account
[**ReplaceAccountRoute**](RoutesApi.md#replaceaccountroute) | **PUT** /accounts/{account_id}/routes/{route_id} | 


<a name="createroute"></a>
# **CreateRoute**
> RouteFull CreateRoute (int? accountId, CreateRouteParams data = null)

Add a new address book contact for an extension

For more on the input fields, see Intro to Routes.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateRouteExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new RoutesApi();
            var accountId = 56;  // int? | Account ID
            var data = new CreateRouteParams(); // CreateRouteParams | Route data (optional) 

            try
            {
                // Add a new address book contact for an extension
                RouteFull result = apiInstance.CreateRoute(accountId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling RoutesApi.CreateRoute: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **data** | [**CreateRouteParams**](CreateRouteParams.md)| Route data | [optional] 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="deleteaccountroute"></a>
# **DeleteAccountRoute**
> DeleteRoute DeleteAccountRoute (int? accountId, int? routeId)





### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class DeleteAccountRouteExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new RoutesApi();
            var accountId = 56;  // int? | Account ID
            var routeId = 56;  // int? | Route ID

            try
            {
                // 
                DeleteRoute result = apiInstance.DeleteAccountRoute(accountId, routeId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling RoutesApi.DeleteAccountRoute: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **routeId** | **int?**| Route ID | 

### Return type

[**DeleteRoute**](DeleteRoute.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getaccountroute"></a>
# **GetAccountRoute**
> RouteFull GetAccountRoute (int? accountId, int? routeId)

Show details of an individual route

This service shows the details of an individual route.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountRouteExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new RoutesApi();
            var accountId = 56;  // int? | Account ID
            var routeId = 56;  // int? | Route ID

            try
            {
                // Show details of an individual route
                RouteFull result = apiInstance.GetAccountRoute(accountId, routeId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling RoutesApi.GetAccountRoute: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **routeId** | **int?**| Route ID | 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountroutes"></a>
# **ListAccountRoutes**
> ListRoutes ListAccountRoutes (int? accountId, List<string> filtersId = null, List<string> filtersName = null, string sortId = null, string sortName = null, int? limit = null, int? offset = null, string fields = null)

Get a list of routes for an account

See Intro to Routes for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountRoutesExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new RoutesApi();
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
                // Get a list of routes for an account
                ListRoutes result = apiInstance.ListAccountRoutes(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling RoutesApi.ListAccountRoutes: " + e.Message );
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

[**ListRoutes**](ListRoutes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="replaceaccountroute"></a>
# **ReplaceAccountRoute**
> RouteFull ReplaceAccountRoute (int? accountId, int? routeId, CreateRouteParams data = null)



For more on the input fields, see Intro to Routes.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ReplaceAccountRouteExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new RoutesApi();
            var accountId = 56;  // int? | Account ID
            var routeId = 56;  // int? | Route ID
            var data = new CreateRouteParams(); // CreateRouteParams | Route data (optional) 

            try
            {
                // 
                RouteFull result = apiInstance.ReplaceAccountRoute(accountId, routeId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling RoutesApi.ReplaceAccountRoute: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **routeId** | **int?**| Route ID | 
 **data** | [**CreateRouteParams**](CreateRouteParams.md)| Route data | [optional] 

### Return type

[**RouteFull**](RouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

