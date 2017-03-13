# DefaultApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ping**](DefaultApi.md#ping) | **GET** /ping | The default API command


<a name="ping"></a>
# **ping**
> PingResponse ping()

The default API command

The default API command

### Example
```java
// Import classes:
//import io.swagger.client.api.DefaultApi;

DefaultApi apiInstance = new DefaultApi();
try {
    PingResponse result = apiInstance.ping();
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling DefaultApi#ping");
    e.printStackTrace();
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

