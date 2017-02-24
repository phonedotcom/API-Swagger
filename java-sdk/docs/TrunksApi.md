# TrunksApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountTrunk**](TrunksApi.md#createAccountTrunk) | **POST** /accounts/{account_id}/trunks | Add a trunk record with SIP information
[**deleteAccountTrunk**](TrunksApi.md#deleteAccountTrunk) | **DELETE** /accounts/{account_id}/trunks/{trunk_id} | Delete a trunk from account
[**getAccountTrunk**](TrunksApi.md#getAccountTrunk) | **GET** /accounts/{account_id}/trunks/{trunk_id} | Show details of an individual trunk
[**listAccountTrunks**](TrunksApi.md#listAccountTrunks) | **GET** /accounts/{account_id}/trunks | Get a list of trunks for an account
[**replaceAccountTrunk**](TrunksApi.md#replaceAccountTrunk) | **PUT** /accounts/{account_id}/trunks/{trunk_id} | Replace parameters in a trunk


<a name="createAccountTrunk"></a>
# **createAccountTrunk**
> TrunkFull createAccountTrunk(accountId, data)

Add a trunk record with SIP information

For more on the input fields, see Account Trunks.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.TrunksApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

TrunksApi apiInstance = new TrunksApi();
Integer accountId = 56; // Integer | Account ID
CreateTrunkParams data = new CreateTrunkParams(); // CreateTrunkParams | Trunk data
try {
    TrunkFull result = apiInstance.createAccountTrunk(accountId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling TrunksApi#createAccountTrunk");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **data** | [**CreateTrunkParams**](CreateTrunkParams.md)| Trunk data |

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteAccountTrunk"></a>
# **deleteAccountTrunk**
> DeleteTrunk deleteAccountTrunk(accountId, trunkId)

Delete a trunk from account

This service deletes a trunk from the account. For more on the properties of trunks, see Account Trunks.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.TrunksApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

TrunksApi apiInstance = new TrunksApi();
Integer accountId = 56; // Integer | Account ID
Integer trunkId = 56; // Integer | Trunk ID
try {
    DeleteTrunk result = apiInstance.deleteAccountTrunk(accountId, trunkId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling TrunksApi#deleteAccountTrunk");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **trunkId** | **Integer**| Trunk ID |

### Return type

[**DeleteTrunk**](DeleteTrunk.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getAccountTrunk"></a>
# **getAccountTrunk**
> TrunkFull getAccountTrunk(accountId, trunkId)

Show details of an individual trunk

This service shows the details of an individual Trunk.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.TrunksApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

TrunksApi apiInstance = new TrunksApi();
Integer accountId = 56; // Integer | Account ID
Integer trunkId = 56; // Integer | Trunk ID
try {
    TrunkFull result = apiInstance.getAccountTrunk(accountId, trunkId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling TrunksApi#getAccountTrunk");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **trunkId** | **Integer**| Trunk ID |

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listAccountTrunks"></a>
# **listAccountTrunks**
> ListTrunksFull listAccountTrunks(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields)

Get a list of trunks for an account

See Account Trunks for more info on the properties.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.TrunksApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

TrunksApi apiInstance = new TrunksApi();
Integer accountId = 56; // Integer | Account ID
List<String> filtersId = Arrays.asList("filtersId_example"); // List<String> | ID filter
List<String> filtersName = Arrays.asList("filtersName_example"); // List<String> | Name filter
String sortId = "sortId_example"; // String | ID sorting
String sortName = "sortName_example"; // String | Name sorting
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
try {
    ListTrunksFull result = apiInstance.listAccountTrunks(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling TrunksApi#listAccountTrunks");
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

[**ListTrunksFull**](ListTrunksFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="replaceAccountTrunk"></a>
# **replaceAccountTrunk**
> TrunkFull replaceAccountTrunk(accountId, trunkId, data)

Replace parameters in a trunk

For more on the input fields, see Account Trunks.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.TrunksApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

TrunksApi apiInstance = new TrunksApi();
Integer accountId = 56; // Integer | Account ID
Integer trunkId = 56; // Integer | Trunk ID
CreateTrunkParams data = new CreateTrunkParams(); // CreateTrunkParams | Trunk data
try {
    TrunkFull result = apiInstance.replaceAccountTrunk(accountId, trunkId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling TrunksApi#replaceAccountTrunk");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **trunkId** | **Integer**| Trunk ID |
 **data** | [**CreateTrunkParams**](CreateTrunkParams.md)| Trunk data |

### Return type

[**TrunkFull**](TrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

