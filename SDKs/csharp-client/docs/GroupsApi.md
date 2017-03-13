# IO.Swagger.Api.GroupsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountExtensionContactGroup**](GroupsApi.md#createaccountextensioncontactgroup) | **POST** /accounts/{account_id}/extensions/{extension_id}/contact-groups | 
[**DeleteAccountExtensionContactGroup**](GroupsApi.md#deleteaccountextensioncontactgroup) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | Delete an addressbook group
[**GetAccountExtensionContactGroup**](GroupsApi.md#getaccountextensioncontactgroup) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 
[**ListAccountExtensionContactGroups**](GroupsApi.md#listaccountextensioncontactgroups) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups | Show a list of contact groups belonging to an extension
[**ReplaceAccountExtensionContactGroup**](GroupsApi.md#replaceaccountextensioncontactgroup) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 


<a name="createaccountextensioncontactgroup"></a>
# **CreateAccountExtensionContactGroup**
> GroupFull CreateAccountExtensionContactGroup (int? accountId, int? extensionId, CreateGroupParams data)



See Account Contact Groups for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAccountExtensionContactGroupExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new GroupsApi();
            var accountId = 56;  // int? | Account ID
            var extensionId = 56;  // int? | Extension ID
            var data = new CreateGroupParams(); // CreateGroupParams | Group name

            try
            {
                // 
                GroupFull result = apiInstance.CreateAccountExtensionContactGroup(accountId, extensionId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling GroupsApi.CreateAccountExtensionContactGroup: " + e.Message );
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
 **data** | [**CreateGroupParams**](CreateGroupParams.md)| Group name | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="deleteaccountextensioncontactgroup"></a>
# **DeleteAccountExtensionContactGroup**
> DeleteGroup DeleteAccountExtensionContactGroup (int? accountId, int? extensionId, int? groupId)

Delete an addressbook group



### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class DeleteAccountExtensionContactGroupExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new GroupsApi();
            var accountId = 56;  // int? | Account ID
            var extensionId = 56;  // int? | Extension ID
            var groupId = 56;  // int? | Group ID

            try
            {
                // Delete an addressbook group
                DeleteGroup result = apiInstance.DeleteAccountExtensionContactGroup(accountId, extensionId, groupId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling GroupsApi.DeleteAccountExtensionContactGroup: " + e.Message );
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
 **groupId** | **int?**| Group ID | 

### Return type

[**DeleteGroup**](DeleteGroup.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getaccountextensioncontactgroup"></a>
# **GetAccountExtensionContactGroup**
> GroupFull GetAccountExtensionContactGroup (int? accountId, int? extensionId, int? groupId)



See Account Contact Groups for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountExtensionContactGroupExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new GroupsApi();
            var accountId = 56;  // int? | Account ID
            var extensionId = 56;  // int? | Extension ID
            var groupId = 56;  // int? | Group ID

            try
            {
                // 
                GroupFull result = apiInstance.GetAccountExtensionContactGroup(accountId, extensionId, groupId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling GroupsApi.GetAccountExtensionContactGroup: " + e.Message );
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
 **groupId** | **int?**| Group ID | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountextensioncontactgroups"></a>
# **ListAccountExtensionContactGroups**
> ListGroups ListAccountExtensionContactGroups (int? accountId, int? extensionId, List<string> filtersId = null, List<string> filtersName = null, string sortId = null, string sortName = null, int? limit = null, int? offset = null, string fields = null)

Show a list of contact groups belonging to an extension

See Account Contact Groups for details on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountExtensionContactGroupsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new GroupsApi();
            var accountId = 56;  // int? | Account ID
            var extensionId = 56;  // int? | Extension ID
            var filtersId = new List<string>(); // List<string> | ID filter (optional) 
            var filtersName = new List<string>(); // List<string> | Name filter (optional) 
            var sortId = sortId_example;  // string | ID sorting (optional) 
            var sortName = sortName_example;  // string | Name sorting (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // Show a list of contact groups belonging to an extension
                ListGroups result = apiInstance.ListAccountExtensionContactGroups(accountId, extensionId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling GroupsApi.ListAccountExtensionContactGroups: " + e.Message );
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
 **filtersId** | [**List&lt;string&gt;**](string.md)| ID filter | [optional] 
 **filtersName** | [**List&lt;string&gt;**](string.md)| Name filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListGroups**](ListGroups.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="replaceaccountextensioncontactgroup"></a>
# **ReplaceAccountExtensionContactGroup**
> GroupFull ReplaceAccountExtensionContactGroup (int? accountId, int? extensionId, int? groupId, CreateGroupParams data)



See Account Contact Groups for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ReplaceAccountExtensionContactGroupExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new GroupsApi();
            var accountId = 56;  // int? | Account ID
            var extensionId = 56;  // int? | Extension ID
            var groupId = 56;  // int? | Group ID
            var data = new CreateGroupParams(); // CreateGroupParams | Group name

            try
            {
                // 
                GroupFull result = apiInstance.ReplaceAccountExtensionContactGroup(accountId, extensionId, groupId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling GroupsApi.ReplaceAccountExtensionContactGroup: " + e.Message );
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
 **groupId** | **int?**| Group ID | 
 **data** | [**CreateGroupParams**](CreateGroupParams.md)| Group name | 

### Return type

[**GroupFull**](GroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

