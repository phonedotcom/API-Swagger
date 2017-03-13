# IO.Swagger.Api.MenusApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountMenu**](MenusApi.md#createaccountmenu) | **POST** /accounts/{account_id}/menus | Create an individual menu
[**DeleteAccountMenu**](MenusApi.md#deleteaccountmenu) | **DELETE** /accounts/{account_id}/menus/{menu_id} | Delete an individual menu
[**GetAccountMenu**](MenusApi.md#getaccountmenu) | **GET** /accounts/{account_id}/menus/{menu_id} | Show details of an individual menu
[**ListAccountMenus**](MenusApi.md#listaccountmenus) | **GET** /accounts/{account_id}/menus | Get a list of menus for an account
[**ReplaceAccountMenu**](MenusApi.md#replaceaccountmenu) | **PUT** /accounts/{account_id}/menus/{menu_id} | Replace an individual menu


<a name="createaccountmenu"></a>
# **CreateAccountMenu**
> MenuFull CreateAccountMenu (int? accountId, CreateMenuParams data = null)

Create an individual menu

This service creates an individual menu. See Account Menus for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAccountMenuExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new MenusApi();
            var accountId = 56;  // int? | Account ID
            var data = new CreateMenuParams(); // CreateMenuParams | Menu data (optional) 

            try
            {
                // Create an individual menu
                MenuFull result = apiInstance.CreateAccountMenu(accountId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling MenusApi.CreateAccountMenu: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **data** | [**CreateMenuParams**](CreateMenuParams.md)| Menu data | [optional] 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="deleteaccountmenu"></a>
# **DeleteAccountMenu**
> DeleteMenu DeleteAccountMenu (int? accountId, int? menuId)

Delete an individual menu

This service shows the details of an individual menu.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class DeleteAccountMenuExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new MenusApi();
            var accountId = 56;  // int? | Account ID
            var menuId = 56;  // int? | Menu ID

            try
            {
                // Delete an individual menu
                DeleteMenu result = apiInstance.DeleteAccountMenu(accountId, menuId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling MenusApi.DeleteAccountMenu: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **menuId** | **int?**| Menu ID | 

### Return type

[**DeleteMenu**](DeleteMenu.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getaccountmenu"></a>
# **GetAccountMenu**
> MenuFull GetAccountMenu (int? accountId, int? menuId)

Show details of an individual menu

This service shows the details of an individual Menu.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountMenuExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new MenusApi();
            var accountId = 56;  // int? | Account ID
            var menuId = 56;  // int? | Menu ID

            try
            {
                // Show details of an individual menu
                MenuFull result = apiInstance.GetAccountMenu(accountId, menuId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling MenusApi.GetAccountMenu: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **menuId** | **int?**| Menu ID | 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountmenus"></a>
# **ListAccountMenus**
> ListMenus ListAccountMenus (int? accountId, List<string> filtersId = null, List<string> filtersName = null, string sortId = null, string sortName = null, int? limit = null, int? offset = null, string fields = null)

Get a list of menus for an account

See Account Menus for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountMenusExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new MenusApi();
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
                // Get a list of menus for an account
                ListMenus result = apiInstance.ListAccountMenus(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling MenusApi.ListAccountMenus: " + e.Message );
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

[**ListMenus**](ListMenus.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="replaceaccountmenu"></a>
# **ReplaceAccountMenu**
> MenuFull ReplaceAccountMenu (int? accountId, int? menuId, ReplaceMenuParams data = null)

Replace an individual menu

This service replaces the details of an individual Menu.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ReplaceAccountMenuExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new MenusApi();
            var accountId = 56;  // int? | Account ID
            var menuId = 56;  // int? | Menu ID
            var data = new ReplaceMenuParams(); // ReplaceMenuParams | Menu data (optional) 

            try
            {
                // Replace an individual menu
                MenuFull result = apiInstance.ReplaceAccountMenu(accountId, menuId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling MenusApi.ReplaceAccountMenu: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **menuId** | **int?**| Menu ID | 
 **data** | [**ReplaceMenuParams**](ReplaceMenuParams.md)| Menu data | [optional] 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

