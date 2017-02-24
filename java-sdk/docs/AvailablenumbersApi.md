# AvailablenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listAvailablePhoneNumbers**](AvailablenumbersApi.md#listAvailablePhoneNumbers) | **GET** /phone-numbers/available | 


<a name="listAvailablePhoneNumbers"></a>
# **listAvailablePhoneNumbers**
> ListAvailableNumbersFull listAvailablePhoneNumbers(filtersPhoneNumber, filtersCountryCode, filtersNpa, filtersNxx, filtersXxxx, filtersCity, filtersProvince, filtersCountry, filtersPrice, filtersCategory, filtersIsTollFree, sortInternal, sortPrice, sortPhoneNumber, limit, offset, fields)





### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.AvailablenumbersApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

AvailablenumbersApi apiInstance = new AvailablenumbersApi();
List<String> filtersPhoneNumber = Arrays.asList("filtersPhoneNumber_example"); // List<String> | Phone number filter
List<String> filtersCountryCode = Arrays.asList("filtersCountryCode_example"); // List<String> | Country Code filter
List<String> filtersNpa = Arrays.asList("filtersNpa_example"); // List<String> | Area Code filter (North America only)
List<String> filtersNxx = Arrays.asList("filtersNxx_example"); // List<String> | 2nd set of 3 digits filter (North America only)
List<String> filtersXxxx = Arrays.asList("filtersXxxx_example"); // List<String> | NANP XXXX filter
List<String> filtersCity = Arrays.asList("filtersCity_example"); // List<String> | City filter
List<String> filtersProvince = Arrays.asList("filtersProvince_example"); // List<String> | State or Province (postal code) filter
List<String> filtersCountry = Arrays.asList("filtersCountry_example"); // List<String> | Country (postal code) filter
List<String> filtersPrice = Arrays.asList("filtersPrice_example"); // List<String> | Price filter
List<String> filtersCategory = Arrays.asList("filtersCategory_example"); // List<String> | Category filter
List<String> filtersIsTollFree = Arrays.asList("filtersIsTollFree_example"); // List<String> | Toll-free status filter
String sortInternal = "sortInternal_example"; // String | Internal (quasi-random) sorting
String sortPrice = "sortPrice_example"; // String | Price sorting
String sortPhoneNumber = "sortPhoneNumber_example"; // String | Phone number sorting
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
try {
    ListAvailableNumbersFull result = apiInstance.listAvailablePhoneNumbers(filtersPhoneNumber, filtersCountryCode, filtersNpa, filtersNxx, filtersXxxx, filtersCity, filtersProvince, filtersCountry, filtersPrice, filtersCategory, filtersIsTollFree, sortInternal, sortPrice, sortPhoneNumber, limit, offset, fields);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling AvailablenumbersApi#listAvailablePhoneNumbers");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersPhoneNumber** | [**List&lt;String&gt;**](String.md)| Phone number filter | [optional]
 **filtersCountryCode** | [**List&lt;String&gt;**](String.md)| Country Code filter | [optional]
 **filtersNpa** | [**List&lt;String&gt;**](String.md)| Area Code filter (North America only) | [optional]
 **filtersNxx** | [**List&lt;String&gt;**](String.md)| 2nd set of 3 digits filter (North America only) | [optional]
 **filtersXxxx** | [**List&lt;String&gt;**](String.md)| NANP XXXX filter | [optional]
 **filtersCity** | [**List&lt;String&gt;**](String.md)| City filter | [optional]
 **filtersProvince** | [**List&lt;String&gt;**](String.md)| State or Province (postal code) filter | [optional]
 **filtersCountry** | [**List&lt;String&gt;**](String.md)| Country (postal code) filter | [optional]
 **filtersPrice** | [**List&lt;String&gt;**](String.md)| Price filter | [optional]
 **filtersCategory** | [**List&lt;String&gt;**](String.md)| Category filter | [optional]
 **filtersIsTollFree** | [**List&lt;String&gt;**](String.md)| Toll-free status filter | [optional]
 **sortInternal** | **String**| Internal (quasi-random) sorting | [optional]
 **sortPrice** | **String**| Price sorting | [optional]
 **sortPhoneNumber** | **String**| Phone number sorting | [optional]
 **limit** | **Integer**| Max results | [optional]
 **offset** | **Integer**| Results to skip | [optional]
 **fields** | **String**| Field set | [optional]

### Return type

[**ListAvailableNumbersFull**](ListAvailableNumbersFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

