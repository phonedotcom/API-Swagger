# CallsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountCall**](CallsApi.md#createAccountCall) | **POST** /accounts/{account_id}/calls | Make a phone call


<a name="createAccountCall"></a>
# **createAccountCall**
> CallFull createAccountCall(accountId, data)

Make a phone call



### Example
```java
// Import classes:
//import io.swagger.client.api.CallsApi;

CallsApi apiInstance = new CallsApi();
Integer accountId = 56; // Integer | Account ID
CreateCallParams data = new CreateCallParams(); // CreateCallParams | Call data
try {
    CallFull result = apiInstance.createAccountCall(accountId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling CallsApi#createAccountCall");
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

