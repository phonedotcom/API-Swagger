# IO.Swagger.Api.NumberregionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailablePhoneNumberRegions**](NumberregionsApi.md#listavailablephonenumberregions) | **GET** /phone-numbers/available/regions | 


<a name="listavailablephonenumberregions"></a>
# **ListAvailablePhoneNumberRegions**
> ListPhoneNumbersRegions ListAvailablePhoneNumberRegions (List<string> filtersCountryCode = null, List<string> filtersNpa = null, List<string> filtersNxx = null, List<string> filtersIsTollFree = null, List<string> filtersCity = null, List<string> filtersProvincePostalCode = null, List<string> filtersCountryPostalCode = null, string sortCountryCode = null, string sortNpa = null, string sortNxx = null, string sortIsTollFree = null, string sortCity = null, string sortProvincePostalCode = null, string sortCountryPostalCode = null, int? limit = null, int? offset = null, string fields = null, List<string> groupBy = null)



This service lists the quantities of available phone numbers by region.

### Example
```csharp
using System;
using System.Diagnostics;
using IO.Swagger.Api;
using IO.Swagger.Client;
using IO.Swagger.Model;

namespace Example
{
    public class ListAvailablePhoneNumberRegionsExample
    {
        public void main()
        {
            
            // Configure API key authorization: apiKey
            Configuration.Default.ApiKey.Add("Authorization", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // Configuration.Default.ApiKeyPrefix.Add("Authorization", "Bearer");

            var apiInstance = new NumberregionsApi();
            var filtersCountryCode = new List<string>(); // List<string> | Country Code filter (optional) 
            var filtersNpa = new List<string>(); // List<string> | Area Code filter (North America only) (optional) 
            var filtersNxx = new List<string>(); // List<string> | 2nd set of 3 digits filter (North America only) (optional) 
            var filtersIsTollFree = new List<string>(); // List<string> | Toll-free status filter (optional) 
            var filtersCity = new List<string>(); // List<string> | City filter (optional) 
            var filtersProvincePostalCode = new List<string>(); // List<string> | State or Province filter (optional) 
            var filtersCountryPostalCode = new List<string>(); // List<string> | Country filter (optional) 
            var sortCountryCode = sortCountryCode_example;  // string | International calling code sorting (optional) 
            var sortNpa = sortNpa_example;  // string | Area Code sorting (North America only) (optional) 
            var sortNxx = sortNxx_example;  // string | 2nd set of 3 digits sorting (North America) (optional) 
            var sortIsTollFree = sortIsTollFree_example;  // string | Toll Free status sorting (optional) 
            var sortCity = sortCity_example;  // string | City sorting (optional) 
            var sortProvincePostalCode = sortProvincePostalCode_example;  // string | State or Province sorting (optional) 
            var sortCountryPostalCode = sortCountryPostalCode_example;  // string | Country sorting (optional) 
            var limit = 56;  // int? | Max results (optional) 
            var offset = 56;  // int? | Results to skip (optional) 
            var fields = fields_example;  // string | Field set (optional) 
            var groupBy = new List<string>(); // List<string> | Fields to group by (supports the same set of fields as the filters and sorting) (optional) 

            try
            {
                // 
                ListPhoneNumbersRegions result = apiInstance.ListAvailablePhoneNumberRegions(filtersCountryCode, filtersNpa, filtersNxx, filtersIsTollFree, filtersCity, filtersProvincePostalCode, filtersCountryPostalCode, sortCountryCode, sortNpa, sortNxx, sortIsTollFree, sortCity, sortProvincePostalCode, sortCountryPostalCode, limit, offset, fields, groupBy);
                Debug.WriteLine(result);
            }
            catch (Exception e)
            {
                Debug.Print("Exception when calling NumberregionsApi.ListAvailablePhoneNumberRegions: " + e.Message );
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersCountryCode** | [**List&lt;string&gt;**](string.md)| Country Code filter | [optional] 
 **filtersNpa** | [**List&lt;string&gt;**](string.md)| Area Code filter (North America only) | [optional] 
 **filtersNxx** | [**List&lt;string&gt;**](string.md)| 2nd set of 3 digits filter (North America only) | [optional] 
 **filtersIsTollFree** | [**List&lt;string&gt;**](string.md)| Toll-free status filter | [optional] 
 **filtersCity** | [**List&lt;string&gt;**](string.md)| City filter | [optional] 
 **filtersProvincePostalCode** | [**List&lt;string&gt;**](string.md)| State or Province filter | [optional] 
 **filtersCountryPostalCode** | [**List&lt;string&gt;**](string.md)| Country filter | [optional] 
 **sortCountryCode** | **string**| International calling code sorting | [optional] 
 **sortNpa** | **string**| Area Code sorting (North America only) | [optional] 
 **sortNxx** | **string**| 2nd set of 3 digits sorting (North America) | [optional] 
 **sortIsTollFree** | **string**| Toll Free status sorting | [optional] 
 **sortCity** | **string**| City sorting | [optional] 
 **sortProvincePostalCode** | **string**| State or Province sorting | [optional] 
 **sortCountryPostalCode** | **string**| Country sorting | [optional] 
 **limit** | **int?**| Max results | [optional] 
 **offset** | **int?**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 
 **groupBy** | [**List&lt;string&gt;**](string.md)| Fields to group by (supports the same set of fields as the filters and sorting) | [optional] 

### Return type

[**ListPhoneNumbersRegions**](ListPhoneNumbersRegions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

