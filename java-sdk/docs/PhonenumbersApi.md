# PhonenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountPhoneNumber**](PhonenumbersApi.md#createAccountPhoneNumber) | **POST** /accounts/{account_id}/phone-numbers | Add a phone number to an account
[**getAccountPhoneNumber**](PhonenumbersApi.md#getAccountPhoneNumber) | **GET** /accounts/{account_id}/phone-numbers/{number_id} | Show details of an individual phone number
[**listAccountPhoneNumbers**](PhonenumbersApi.md#listAccountPhoneNumbers) | **GET** /accounts/{account_id}/phone-numbers | Get a list of phone numbers registered to an account
[**replaceAccountPhoneNumber**](PhonenumbersApi.md#replaceAccountPhoneNumber) | **PUT** /accounts/{account_id}/phone-numbers/{number_id} | Update the settings for an existing phone number on your account


<a name="createAccountPhoneNumber"></a>
# **createAccountPhoneNumber**
> PhoneNumberFull createAccountPhoneNumber(accountId, data)

Add a phone number to an account

See Intro to Account Phone Numbers for more info on the properties to use.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.PhonenumbersApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

PhonenumbersApi apiInstance = new PhonenumbersApi();
Integer accountId = 56; // Integer | Account ID
CreatePhoneNumberParams data = new CreatePhoneNumberParams(); // CreatePhoneNumberParams | Phone Number data
try {
    PhoneNumberFull result = apiInstance.createAccountPhoneNumber(accountId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling PhonenumbersApi#createAccountPhoneNumber");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **data** | [**CreatePhoneNumberParams**](CreatePhoneNumberParams.md)| Phone Number data | [optional]

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getAccountPhoneNumber"></a>
# **getAccountPhoneNumber**
> PhoneNumberFull getAccountPhoneNumber(accountId, numberId)

Show details of an individual phone number

See Intro to Account Phone Numbers for more info on the properties.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.PhonenumbersApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

PhonenumbersApi apiInstance = new PhonenumbersApi();
Integer accountId = 56; // Integer | Account ID
Integer numberId = 56; // Integer | Number ID
try {
    PhoneNumberFull result = apiInstance.getAccountPhoneNumber(accountId, numberId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling PhonenumbersApi#getAccountPhoneNumber");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **numberId** | **Integer**| Number ID |

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listAccountPhoneNumbers"></a>
# **listAccountPhoneNumbers**
> ListPhoneNumbersFull listAccountPhoneNumbers(accountId, filtersId, filtersName, filtersPhoneNumber, sortId, sortName, sortPhoneNumber, limit, offset, fields)

Get a list of phone numbers registered to an account

See Intro to Account Phone Numbers for more info on the properties.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.PhonenumbersApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

PhonenumbersApi apiInstance = new PhonenumbersApi();
Integer accountId = 56; // Integer | Account ID
List<String> filtersId = Arrays.asList("filtersId_example"); // List<String> | ID filter
List<String> filtersName = Arrays.asList("filtersName_example"); // List<String> | Name filter
List<String> filtersPhoneNumber = Arrays.asList("filtersPhoneNumber_example"); // List<String> | Phone number filter
String sortId = "sortId_example"; // String | ID sorting
String sortName = "sortName_example"; // String | Name sorting
String sortPhoneNumber = "sortPhoneNumber_example"; // String | Phone number sorting
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
try {
    ListPhoneNumbersFull result = apiInstance.listAccountPhoneNumbers(accountId, filtersId, filtersName, filtersPhoneNumber, sortId, sortName, sortPhoneNumber, limit, offset, fields);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling PhonenumbersApi#listAccountPhoneNumbers");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **filtersId** | [**List&lt;String&gt;**](String.md)| ID filter | [optional]
 **filtersName** | [**List&lt;String&gt;**](String.md)| Name filter | [optional]
 **filtersPhoneNumber** | [**List&lt;String&gt;**](String.md)| Phone number filter | [optional]
 **sortId** | **String**| ID sorting | [optional]
 **sortName** | **String**| Name sorting | [optional]
 **sortPhoneNumber** | **String**| Phone number sorting | [optional]
 **limit** | **Integer**| Max results | [optional]
 **offset** | **Integer**| Results to skip | [optional]
 **fields** | **String**| Field set | [optional]

### Return type

[**ListPhoneNumbersFull**](ListPhoneNumbersFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="replaceAccountPhoneNumber"></a>
# **replaceAccountPhoneNumber**
> PhoneNumberFull replaceAccountPhoneNumber(accountId, numberId, data)

Update the settings for an existing phone number on your account

See Intro to Account Phone Numbers for more info on the properties.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.PhonenumbersApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

PhonenumbersApi apiInstance = new PhonenumbersApi();
Integer accountId = 56; // Integer | Account ID
Integer numberId = 56; // Integer | Number ID
ReplacePhoneNumberParams data = new ReplacePhoneNumberParams(); // ReplacePhoneNumberParams | Phone Number data
try {
    PhoneNumberFull result = apiInstance.replaceAccountPhoneNumber(accountId, numberId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling PhonenumbersApi#replaceAccountPhoneNumber");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **numberId** | **Integer**| Number ID |
 **data** | [**ReplacePhoneNumberParams**](ReplacePhoneNumberParams.md)| Phone Number data | [optional]

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

