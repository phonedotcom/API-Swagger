# CallsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountCalls**](CallsApi.md#createAccountCalls) | **POST** /accounts/{account_id}/calls | Make a phone call


<a name="createAccountCalls"></a>
# **createAccountCalls**
> CallFull createAccountCalls(accountId, data)

Make a phone call



### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.CallsApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

CallsApi apiInstance = new CallsApi();
Integer accountId = 56; // Integer | Account ID
CreateCallParams data = new CreateCallParams(); // CreateCallParams | Call data
try {
    CallFull result = apiInstance.createAccountCalls(accountId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling CallsApi#createAccountCalls");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **data** | [**CreateCallParams**](CreateCallParams.md)| Call data | [optional]

### Return type

[**CallFull**](CallFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

