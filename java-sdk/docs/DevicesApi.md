# DevicesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountDevice**](DevicesApi.md#createAccountDevice) | **POST** /accounts/{account_id}/devices | Register a generic VoIP device
[**getAccountDevice**](DevicesApi.md#getAccountDevice) | **GET** /accounts/{account_id}/device/{device_id} | Show details of an individual VoIP device
[**listAccountDevices**](DevicesApi.md#listAccountDevices) | **GET** /accounts/{account_id}/devices | Get a list of VoIP devices associated with your account
[**replaceAccountDevice**](DevicesApi.md#replaceAccountDevice) | **PUT** /accounts/{account_id}/device/{device_id} | Update the settings for an individual VoIP device


<a name="createAccountDevice"></a>
# **createAccountDevice**
> DeviceFull createAccountDevice(accountId, data)

Register a generic VoIP device



### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.DevicesApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

DevicesApi apiInstance = new DevicesApi();
Integer accountId = 56; // Integer | Account ID
CreateDeviceParams data = new CreateDeviceParams(); // CreateDeviceParams | Device data
try {
    DeviceFull result = apiInstance.createAccountDevice(accountId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling DevicesApi#createAccountDevice");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **data** | [**CreateDeviceParams**](CreateDeviceParams.md)| Device data | [optional]

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getAccountDevice"></a>
# **getAccountDevice**
> DeviceFull getAccountDevice(accountId, deviceId)

Show details of an individual VoIP device



### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.DevicesApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

DevicesApi apiInstance = new DevicesApi();
Integer accountId = 56; // Integer | Account ID
Integer deviceId = 56; // Integer | Device ID
try {
    DeviceFull result = apiInstance.getAccountDevice(accountId, deviceId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling DevicesApi#getAccountDevice");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **deviceId** | **Integer**| Device ID |

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listAccountDevices"></a>
# **listAccountDevices**
> ListDevicesFull listAccountDevices(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields)

Get a list of VoIP devices associated with your account



### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.DevicesApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

DevicesApi apiInstance = new DevicesApi();
Integer accountId = 56; // Integer | Account ID
List<String> filtersId = Arrays.asList("filtersId_example"); // List<String> | ID filter
List<String> filtersName = Arrays.asList("filtersName_example"); // List<String> | Name filter
String sortId = "sortId_example"; // String | ID sorting
String sortName = "sortName_example"; // String | Name sorting
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
try {
    ListDevicesFull result = apiInstance.listAccountDevices(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling DevicesApi#listAccountDevices");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **filtersId** | [**List&lt;String&gt;**](String.md)| ID filter | [optional]
 **filtersName** | [**List&lt;String&gt;**](String.md)| Name filter | [optional]
 **sortId** | **String**| ID sorting | [optional]
 **sortName** | **String**| Name sorting | [optional]
 **limit** | **Integer**| Max results | [optional]
 **offset** | **Integer**| Results to skip | [optional]
 **fields** | **String**| Field set | [optional]

### Return type

[**ListDevicesFull**](ListDevicesFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="replaceAccountDevice"></a>
# **replaceAccountDevice**
> DeviceFull replaceAccountDevice(accountId, deviceId, data)

Update the settings for an individual VoIP device



### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.DevicesApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

DevicesApi apiInstance = new DevicesApi();
Integer accountId = 56; // Integer | Account ID
Integer deviceId = 56; // Integer | Device ID
CreateDeviceParams data = new CreateDeviceParams(); // CreateDeviceParams | Device data
try {
    DeviceFull result = apiInstance.replaceAccountDevice(accountId, deviceId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling DevicesApi#replaceAccountDevice");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **deviceId** | **Integer**| Device ID |
 **data** | [**CreateDeviceParams**](CreateDeviceParams.md)| Device data | [optional]

### Return type

[**DeviceFull**](DeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

