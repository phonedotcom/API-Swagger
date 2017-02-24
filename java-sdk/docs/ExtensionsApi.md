# ExtensionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountExtension**](ExtensionsApi.md#createAccountExtension) | **POST** /accounts/{account_id}/extensions | Create an individual extension
[**getAccountExtension**](ExtensionsApi.md#getAccountExtension) | **GET** /accounts/{account_id}/extensions/{extension_id} | Show details of an individual extension
[**listAccountExtensions**](ExtensionsApi.md#listAccountExtensions) | **GET** /accounts/{account_id}/extensions | Get a list of extensions visible to the authenticated user or client
[**replaceAccountExtension**](ExtensionsApi.md#replaceAccountExtension) | **PUT** /accounts/{account_id}/extensions/{extension_id} | Replace an individual extension


<a name="createAccountExtension"></a>
# **createAccountExtension**
> ExtensionFull createAccountExtension(accountId, data)

Create an individual extension

This service shows how to create a virtual extension.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.ExtensionsApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

ExtensionsApi apiInstance = new ExtensionsApi();
Integer accountId = 56; // Integer | Account ID
CreateExtensionParams data = new CreateExtensionParams(); // CreateExtensionParams | Account Extensions Data
try {
    ExtensionFull result = apiInstance.createAccountExtension(accountId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling ExtensionsApi#createAccountExtension");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **data** | [**CreateExtensionParams**](CreateExtensionParams.md)| Account Extensions Data | [optional]

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getAccountExtension"></a>
# **getAccountExtension**
> ExtensionFull getAccountExtension(accountId, extensionId)

Show details of an individual extension

This service shows the details of an individual Extension.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.ExtensionsApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

ExtensionsApi apiInstance = new ExtensionsApi();
Integer accountId = 56; // Integer | Account ID
Integer extensionId = 56; // Integer | Extension ID
try {
    ExtensionFull result = apiInstance.getAccountExtension(accountId, extensionId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling ExtensionsApi#getAccountExtension");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **extensionId** | **Integer**| Extension ID |

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listAccountExtensions"></a>
# **listAccountExtensions**
> ListExtensionsFull listAccountExtensions(accountId, filtersId, filtersExtension, filtersName, sortId, sortExtension, sortName, limit, offset, fields)

Get a list of extensions visible to the authenticated user or client

This service lists the visible extensions on a given account.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.ExtensionsApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

ExtensionsApi apiInstance = new ExtensionsApi();
Integer accountId = 56; // Integer | Account ID
List<String> filtersId = Arrays.asList("filtersId_example"); // List<String> | ID filter
List<String> filtersExtension = Arrays.asList("filtersExtension_example"); // List<String> | Extension filter
List<String> filtersName = Arrays.asList("filtersName_example"); // List<String> | Name filter
String sortId = "sortId_example"; // String | ID sorting
String sortExtension = "sortExtension_example"; // String | Extension sorting
String sortName = "sortName_example"; // String | Name sorting
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
try {
    ListExtensionsFull result = apiInstance.listAccountExtensions(accountId, filtersId, filtersExtension, filtersName, sortId, sortExtension, sortName, limit, offset, fields);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling ExtensionsApi#listAccountExtensions");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **filtersId** | [**List&lt;String&gt;**](String.md)| ID filter | [optional]
 **filtersExtension** | [**List&lt;String&gt;**](String.md)| Extension filter | [optional]
 **filtersName** | [**List&lt;String&gt;**](String.md)| Name filter | [optional]
 **sortId** | **String**| ID sorting | [optional]
 **sortExtension** | **String**| Extension sorting | [optional]
 **sortName** | **String**| Name sorting | [optional]
 **limit** | **Integer**| Max results | [optional]
 **offset** | **Integer**| Results to skip | [optional]
 **fields** | **String**| Field set | [optional]

### Return type

[**ListExtensionsFull**](ListExtensionsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="replaceAccountExtension"></a>
# **replaceAccountExtension**
> ExtensionFull replaceAccountExtension(accountId, extensionId, data)

Replace an individual extension

This service shows how to update an individual extension.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.ExtensionsApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

ExtensionsApi apiInstance = new ExtensionsApi();
Integer accountId = 56; // Integer | Account ID
Integer extensionId = 56; // Integer | Extension ID
ReplaceExtensionParams data = new ReplaceExtensionParams(); // ReplaceExtensionParams | Account Extensions Data
try {
    ExtensionFull result = apiInstance.replaceAccountExtension(accountId, extensionId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling ExtensionsApi#replaceAccountExtension");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **extensionId** | **Integer**| Extension ID |
 **data** | [**ReplaceExtensionParams**](ReplaceExtensionParams.md)| Account Extensions Data | [optional]

### Return type

[**ExtensionFull**](ExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

