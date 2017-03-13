# IO.Swagger.Api.SmsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountSms**](SmsApi.md#createaccountsms) | **POST** /accounts/{account_id}/sms | Send a SMS to one or a group of recipients
[**GetAccountSms**](SmsApi.md#getaccountsms) | **GET** /accounts/{account_id}/sms/{sms_id} | Show details of an individual SMS
[**ListAccountSms**](SmsApi.md#listaccountsms) | **GET** /accounts/{account_id}/sms | Get a list of SMS messages for an account


<a name="createaccountsms"></a>
# **CreateAccountSms**
> SmsFull CreateAccountSms (int? accountId, CreateSmsParams data)

Send a SMS to one or a group of recipients

For more on the input fields, see Intro to SMS.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAccountSmsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new SmsApi();
            var accountId = 56;  // int? | Account ID
            var data = new CreateSmsParams(); // CreateSmsParams | SMS data

            try
            {
                // Send a SMS to one or a group of recipients
                SmsFull result = apiInstance.CreateAccountSms(accountId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SmsApi.CreateAccountSms: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **data** | [**CreateSmsParams**](CreateSmsParams.md)| SMS data | 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getaccountsms"></a>
# **GetAccountSms**
> SmsFull GetAccountSms (int? accountId, string smsId)

Show details of an individual SMS

This service shows the details of an individual sms. See Intro to SMS for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountSmsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new SmsApi();
            var accountId = 56;  // int? | Account ID
            var smsId = smsId_example;  // string | SMS ID

            try
            {
                // Show details of an individual SMS
                SmsFull result = apiInstance.GetAccountSms(accountId, smsId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SmsApi.GetAccountSms: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **smsId** | **string**| SMS ID | 

### Return type

[**SmsFull**](SmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountsms"></a>
# **ListAccountSms**
> ListSms ListAccountSms (int? accountId, List<string> filtersId = null, string filtersDirection = null, string filtersFrom = null, string sortId = null, string sortCreatedAt = null, int? limit = null, int? offset = null, string fields = null)

Get a list of SMS messages for an account

See Intro to SMS for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountSmsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new SmsApi();
            var accountId = 56;  // int? | Account ID
            var filtersId = new List<string>(); // List<string> | ID filter (optional) 
            var filtersDirection = filtersDirection_example;  // string | Direction filter (optional) 
            var filtersFrom = filtersFrom_example;  // string | Caller ID filter (optional) 
            var sortId = sortId_example;  // string | ID sorting (optional) 
            var sortCreatedAt = sortCreatedAt_example;  // string | Sort by created time of message (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // Get a list of SMS messages for an account
                ListSms result = apiInstance.ListAccountSms(accountId, filtersId, filtersDirection, filtersFrom, sortId, sortCreatedAt, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling SmsApi.ListAccountSms: " + e.Message );
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
 **filtersDirection** | **string**| Direction filter | [optional] 
 **filtersFrom** | **string**| Caller ID filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortCreatedAt** | **string**| Sort by created time of message | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListSms**](ListSms.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

