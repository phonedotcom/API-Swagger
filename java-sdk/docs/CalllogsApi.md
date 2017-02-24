# CalllogsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listAccountCallLogs**](CalllogsApi.md#listAccountCallLogs) | **GET** /accounts/{account_id}/call-logs | Get a list of call details associated with your account


<a name="listAccountCallLogs"></a>
# **listAccountCallLogs**
> ListCallLogsFull listAccountCallLogs(accountId, filtersId, filtersStartTime, filtersCreatedAt, filtersDirection, filtersCalledNumber, filtersType, filtersExtension, sortId, sortStartTime, sortCreatedAt, limit, offset, fields)

Get a list of call details associated with your account

Fetch a list of call logs associated with your account.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.CalllogsApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

CalllogsApi apiInstance = new CalllogsApi();
Integer accountId = 56; // Integer | Account ID
List<String> filtersId = Arrays.asList("filtersId_example"); // List<String> | ID filter
List<String> filtersStartTime = Arrays.asList("filtersStartTime_example"); // List<String> | Call start time filter
String filtersCreatedAt = "filtersCreatedAt_example"; // String | Call log creation time filter
String filtersDirection = "filtersDirection_example"; // String | Call direction filter: in or out
String filtersCalledNumber = "filtersCalledNumber_example"; // String | Called number
String filtersType = "filtersType_example"; // String | Call type, such as 'call', 'fax', 'audiogram'
List<String> filtersExtension = Arrays.asList("filtersExtension_example"); // List<String> | Extension filter
String sortId = "sortId_example"; // String | ID sorting
String sortStartTime = "sortStartTime_example"; // String | Sorting by call start time: asc or desc
String sortCreatedAt = "sortCreatedAt_example"; // String | Sorting by call log creation time: asc or desc
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
try {
    ListCallLogsFull result = apiInstance.listAccountCallLogs(accountId, filtersId, filtersStartTime, filtersCreatedAt, filtersDirection, filtersCalledNumber, filtersType, filtersExtension, sortId, sortStartTime, sortCreatedAt, limit, offset, fields);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling CalllogsApi#listAccountCallLogs");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **filtersId** | [**List&lt;String&gt;**](String.md)| ID filter | [optional]
 **filtersStartTime** | [**List&lt;String&gt;**](String.md)| Call start time filter | [optional]
 **filtersCreatedAt** | **String**| Call log creation time filter | [optional]
 **filtersDirection** | **String**| Call direction filter: in or out | [optional]
 **filtersCalledNumber** | **String**| Called number | [optional]
 **filtersType** | **String**| Call type, such as &#39;call&#39;, &#39;fax&#39;, &#39;audiogram&#39; | [optional]
 **filtersExtension** | [**List&lt;String&gt;**](String.md)| Extension filter | [optional]
 **sortId** | **String**| ID sorting | [optional]
 **sortStartTime** | **String**| Sorting by call start time: asc or desc | [optional]
 **sortCreatedAt** | **String**| Sorting by call log creation time: asc or desc | [optional]
 **limit** | **Integer**| Max results | [optional]
 **offset** | **Integer**| Results to skip | [optional]
 **fields** | **String**| Field set | [optional]

### Return type

[**ListCallLogsFull**](ListCallLogsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

