# IO.Swagger.Api.PhonenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountPhoneNumber**](PhonenumbersApi.md#createaccountphonenumber) | **POST** /accounts/{account_id}/phone-numbers | Add a phone number to an account
[**GetAccountPhoneNumber**](PhonenumbersApi.md#getaccountphonenumber) | **GET** /accounts/{account_id}/phone-numbers/{number_id} | Show details of an individual phone number
[**ListAccountPhoneNumbers**](PhonenumbersApi.md#listaccountphonenumbers) | **GET** /accounts/{account_id}/phone-numbers | Get a list of phone numbers registered to an account
[**ReplaceAccountPhoneNumber**](PhonenumbersApi.md#replaceaccountphonenumber) | **PUT** /accounts/{account_id}/phone-numbers/{number_id} | Update the settings for an existing phone number on your account


<a name="createaccountphonenumber"></a>
# **CreateAccountPhoneNumber**
> PhoneNumberFull CreateAccountPhoneNumber (int? accountId, CreatePhoneNumberParams data = null)

Add a phone number to an account

See Intro to Account Phone Numbers for more info on the properties to use.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class CreateAccountPhoneNumberExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new PhonenumbersApi();
            var accountId = 56;  // int? | Account ID
            var data = new CreatePhoneNumberParams(); // CreatePhoneNumberParams | Phone Number data (optional) 

            try
            {
                // Add a phone number to an account
                PhoneNumberFull result = apiInstance.CreateAccountPhoneNumber(accountId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling PhonenumbersApi.CreateAccountPhoneNumber: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **data** | [**CreatePhoneNumberParams**](CreatePhoneNumberParams.md)| Phone Number data | [optional] 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getaccountphonenumber"></a>
# **GetAccountPhoneNumber**
> PhoneNumberFull GetAccountPhoneNumber (int? accountId, int? numberId)

Show details of an individual phone number

See Intro to Account Phone Numbers for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class GetAccountPhoneNumberExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new PhonenumbersApi();
            var accountId = 56;  // int? | Account ID
            var numberId = 56;  // int? | Number ID

            try
            {
                // Show details of an individual phone number
                PhoneNumberFull result = apiInstance.GetAccountPhoneNumber(accountId, numberId);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling PhonenumbersApi.GetAccountPhoneNumber: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **numberId** | **int?**| Number ID | 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="listaccountphonenumbers"></a>
# **ListAccountPhoneNumbers**
> ListPhoneNumbers ListAccountPhoneNumbers (int? accountId, List<string> filtersId = null, List<string> filtersName = null, List<string> filtersPhoneNumber = null, string sortId = null, string sortName = null, string sortPhoneNumber = null, int? limit = null, int? offset = null, string fields = null)

Get a list of phone numbers registered to an account

See Intro to Account Phone Numbers for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAccountPhoneNumbersExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new PhonenumbersApi();
            var accountId = 56;  // int? | Account ID
            var filtersId = new List<string>(); // List<string> | ID filter (optional) 
            var filtersName = new List<string>(); // List<string> | Name filter (optional) 
            var filtersPhoneNumber = new List<string>(); // List<string> | Phone number filter (optional) 
            var sortId = sortId_example;  // string | ID sorting (optional) 
            var sortName = sortName_example;  // string | Name sorting (optional) 
            var sortPhoneNumber = sortPhoneNumber_example;  // string | Phone number sorting (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // Get a list of phone numbers registered to an account
                ListPhoneNumbers result = apiInstance.ListAccountPhoneNumbers(accountId, filtersId, filtersName, filtersPhoneNumber, sortId, sortName, sortPhoneNumber, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling PhonenumbersApi.ListAccountPhoneNumbers: " + e.Message );
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
 **filtersPhoneNumber** | [**List&lt;string&gt;**](string.md)| Phone number filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **sortPhoneNumber** | **string**| Phone number sorting | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListPhoneNumbers**](ListPhoneNumbers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="replaceaccountphonenumber"></a>
# **ReplaceAccountPhoneNumber**
> PhoneNumberFull ReplaceAccountPhoneNumber (int? accountId, int? numberId, ReplacePhoneNumberParams data = null)

Update the settings for an existing phone number on your account

See Intro to Account Phone Numbers for more info on the properties.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ReplaceAccountPhoneNumberExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new PhonenumbersApi();
            var accountId = 56;  // int? | Account ID
            var numberId = 56;  // int? | Number ID
            var data = new ReplacePhoneNumberParams(); // ReplacePhoneNumberParams | Phone Number data (optional) 

            try
            {
                // Update the settings for an existing phone number on your account
                PhoneNumberFull result = apiInstance.ReplaceAccountPhoneNumber(accountId, numberId, data);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling PhonenumbersApi.ReplaceAccountPhoneNumber: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int?**| Account ID | 
 **numberId** | **int?**| Number ID | 
 **data** | [**ReplacePhoneNumberParams**](ReplacePhoneNumberParams.md)| Phone Number data | [optional] 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

