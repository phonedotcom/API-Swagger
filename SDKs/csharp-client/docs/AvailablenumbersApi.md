# IO.Swagger.Api.AvailablenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailablePhoneNumbers**](AvailablenumbersApi.md#listavailablephonenumbers) | **GET** /phone-numbers/available | 


<a name="listavailablephonenumbers"></a>
# **ListAvailablePhoneNumbers**
> ListAvailableNumbers ListAvailablePhoneNumbers (List<string> filtersPhoneNumber = null, List<string> filtersCountryCode = null, List<string> filtersNpa = null, List<string> filtersNxx = null, List<string> filtersXxxx = null, List<string> filtersCity = null, List<string> filtersProvince = null, List<string> filtersCountry = null, List<string> filtersPrice = null, List<string> filtersCategory = null, string sortInternal = null, string sortPrice = null, string sortPhoneNumber = null, int? limit = null, int? offset = null, string fields = null)





### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAvailablePhoneNumbersExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new AvailablenumbersApi();
            var filtersPhoneNumber = new List<string>(); // List<string> | Phone number filter (optional) 
            var filtersCountryCode = new List<string>(); // List<string> | Country Code filter (optional) 
            var filtersNpa = new List<string>(); // List<string> | Area Code filter (North America only) (optional) 
            var filtersNxx = new List<string>(); // List<string> | 2nd set of 3 digits filter (North America only) (optional) 
            var filtersXxxx = new List<string>(); // List<string> | NANP XXXX filter (optional) 
            var filtersCity = new List<string>(); // List<string> | City filter (optional) 
            var filtersProvince = new List<string>(); // List<string> | State or Province (postal code) filter (optional) 
            var filtersCountry = new List<string>(); // List<string> | Country (postal code) filter (optional) 
            var filtersPrice = new List<string>(); // List<string> | Price filter (optional) 
            var filtersCategory = new List<string>(); // List<string> | Category filter (optional) 
            var sortInternal = sortInternal_example;  // string | Internal (quasi-random) sorting (optional) 
            var sortPrice = sortPrice_example;  // string | Price sorting (optional) 
            var sortPhoneNumber = sortPhoneNumber_example;  // string | Phone number sorting (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 

            try
            {
                // 
                ListAvailableNumbers result = apiInstance.ListAvailablePhoneNumbers(filtersPhoneNumber, filtersCountryCode, filtersNpa, filtersNxx, filtersXxxx, filtersCity, filtersProvince, filtersCountry, filtersPrice, filtersCategory, sortInternal, sortPrice, sortPhoneNumber, limit, offset, fields);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling AvailablenumbersApi.ListAvailablePhoneNumbers: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersPhoneNumber** | [**List&lt;string&gt;**](string.md)| Phone number filter | [optional] 
 **filtersCountryCode** | [**List&lt;string&gt;**](string.md)| Country Code filter | [optional] 
 **filtersNpa** | [**List&lt;string&gt;**](string.md)| Area Code filter (North America only) | [optional] 
 **filtersNxx** | [**List&lt;string&gt;**](string.md)| 2nd set of 3 digits filter (North America only) | [optional] 
 **filtersXxxx** | [**List&lt;string&gt;**](string.md)| NANP XXXX filter | [optional] 
 **filtersCity** | [**List&lt;string&gt;**](string.md)| City filter | [optional] 
 **filtersProvince** | [**List&lt;string&gt;**](string.md)| State or Province (postal code) filter | [optional] 
 **filtersCountry** | [**List&lt;string&gt;**](string.md)| Country (postal code) filter | [optional] 
 **filtersPrice** | [**List&lt;string&gt;**](string.md)| Price filter | [optional] 
 **filtersCategory** | [**List&lt;string&gt;**](string.md)| Category filter | [optional] 
 **sortInternal** | **string**| Internal (quasi-random) sorting | [optional] 
 **sortPrice** | **string**| Price sorting | [optional] 
 **sortPhoneNumber** | **string**| Phone number sorting | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListAvailableNumbers**](ListAvailableNumbers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

