# IO.Swagger.Api.DevicesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountDevice**](DevicesApi.md#createaccountdevice) | **POST** /accounts/{account_id}/devices | Register a generic VoIP device
[**GetAccountDevice**](DevicesApi.md#getaccountdevice) | **GET** /accounts/{account_id}/devices/{device_id} | Show details of an individual VoIP device
[**ListAccountDevices**](DevicesApi.md#listaccountdevices) | **GET** /accounts/{account_id}/devices | Get a list of VoIP devices associated with your account
[**ReplaceAccountDevice**](DevicesApi.md#replaceaccountdevice) | **PUT** /accounts/{account_id}/devices/{device_id} | Update the settings for an individual VoIP device


<a name="createaccountdevice"></a>
# **CreateAccountDevice**
> DeviceFull CreateAccountDevice (int? accountId, CreateDeviceParams data = null)

Register a generic VoIP device



### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAccountDeviceExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new DevicesApi();
            var accountId = 56;  // int? | Account ID
            var data = new CreateDeviceParams(); // CreateDeviceParams | Device data (optional) 

            try
            {
                // Register a generic VoIP device
                DeviceFull result = apiInstance.CreateAccountDevice(accountId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DevicesApi.CreateAccountDevice: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **data** | [**CreateDeviceParams**](CreateDeviceParams.md)| Device data | [optional] 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getaccountdevice"></a>
# **GetAccountDevice**
> DeviceFull GetAccountDevice (int? accountId, int? deviceId)

Show details of an individual VoIP device



### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountDeviceExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new DevicesApi();
            var accountId = 56;  // int? | Account ID
            var deviceId = 56;  // int? | Device ID

            try
            {
                // Show details of an individual VoIP device
                DeviceFull result = apiInstance.GetAccountDevice(accountId, deviceId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DevicesApi.GetAccountDevice: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **deviceId** | **int?**| Device ID | 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountdevices"></a>
# **ListAccountDevices**
> ListDevices ListAccountDevices (int? accountId, List<string> filtersId = null, List<string> filtersName = null, string sortId = null, string sortName = null, int? limit = null, int? offset = null, string fields = null)

Get a list of VoIP devices associated with your account



### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountDevicesExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new DevicesApi();
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
                // Get a list of VoIP devices associated with your account
                ListDevices result = apiInstance.ListAccountDevices(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DevicesApi.ListAccountDevices: " + e.Message );
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

[**ListDevices**](ListDevices.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="replaceaccountdevice"></a>
# **ReplaceAccountDevice**
> DeviceFull ReplaceAccountDevice (int? accountId, int? deviceId, CreateDeviceParams data = null)

Update the settings for an individual VoIP device



### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ReplaceAccountDeviceExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new DevicesApi();
            var accountId = 56;  // int? | Account ID
            var deviceId = 56;  // int? | Device ID
            var data = new CreateDeviceParams(); // CreateDeviceParams | Device data (optional) 

            try
            {
                // Update the settings for an individual VoIP device
                DeviceFull result = apiInstance.ReplaceAccountDevice(accountId, deviceId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling DevicesApi.ReplaceAccountDevice: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **deviceId** | **int?**| Device ID | 
 **data** | [**CreateDeviceParams**](CreateDeviceParams.md)| Device data | [optional] 

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

