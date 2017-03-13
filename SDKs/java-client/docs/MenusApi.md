# MenusApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountMenu**](MenusApi.md#createAccountMenu) | **POST** /accounts/{account_id}/menus | Create an individual menu
[**deleteAccountMenu**](MenusApi.md#deleteAccountMenu) | **DELETE** /accounts/{account_id}/menus/{menu_id} | Delete an individual menu
[**getAccountMenu**](MenusApi.md#getAccountMenu) | **GET** /accounts/{account_id}/menus/{menu_id} | Show details of an individual menu
[**listAccountMenus**](MenusApi.md#listAccountMenus) | **GET** /accounts/{account_id}/menus | Get a list of menus for an account
[**replaceAccountMenu**](MenusApi.md#replaceAccountMenu) | **PUT** /accounts/{account_id}/menus/{menu_id} | Replace an individual menu


<a name="createAccountMenu"></a>
# **createAccountMenu**
> MenuFull createAccountMenu(accountId, data)

Create an individual menu

This service creates an individual menu. See Account Menus for more info on the properties.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.MenusApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

MenusApi apiInstance = new MenusApi();
Integer accountId = 56; // Integer | Account ID
CreateMenuParams data = new CreateMenuParams(); // CreateMenuParams | Menu data
try {
    MenuFull result = apiInstance.createAccountMenu(accountId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling MenusApi#createAccountMenu");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **data** | [**CreateMenuParams**](CreateMenuParams.md)| Menu data | [optional]

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteAccountMenu"></a>
# **deleteAccountMenu**
> DeleteMenu deleteAccountMenu(accountId, menuId)

Delete an individual menu

This service shows the details of an individual menu.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.MenusApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

MenusApi apiInstance = new MenusApi();
Integer accountId = 56; // Integer | Account ID
Integer menuId = 56; // Integer | Menu ID
try {
    DeleteMenu result = apiInstance.deleteAccountMenu(accountId, menuId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling MenusApi#deleteAccountMenu");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **menuId** | **Integer**| Menu ID |

### Return type

[**DeleteMenu**](DeleteMenu.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getAccountMenu"></a>
# **getAccountMenu**
> MenuFull getAccountMenu(accountId, menuId)

Show details of an individual menu

This service shows the details of an individual Menu.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.MenusApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

MenusApi apiInstance = new MenusApi();
Integer accountId = 56; // Integer | Account ID
Integer menuId = 56; // Integer | Menu ID
try {
    MenuFull result = apiInstance.getAccountMenu(accountId, menuId);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling MenusApi#getAccountMenu");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **menuId** | **Integer**| Menu ID |

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listAccountMenus"></a>
# **listAccountMenus**
> ListMenus listAccountMenus(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields)

Get a list of menus for an account

See Account Menus for more info on the properties.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.MenusApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

MenusApi apiInstance = new MenusApi();
Integer accountId = 56; // Integer | Account ID
List<String> filtersId = Arrays.asList("filtersId_example"); // List<String> | ID filter
List<String> filtersName = Arrays.asList("filtersName_example"); // List<String> | Name filter
String sortId = "sortId_example"; // String | ID sorting
String sortName = "sortName_example"; // String | Name sorting
Integer limit = 56; // Integer | Max results
Integer offset = 56; // Integer | Results to skip
String fields = "fields_example"; // String | Field set
try {
    ListMenus result = apiInstance.listAccountMenus(accountId, filtersId, filtersName, sortId, sortName, limit, offset, fields);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling MenusApi#listAccountMenus");
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

[**ListMenus**](ListMenus.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="replaceAccountMenu"></a>
# **replaceAccountMenu**
> MenuFull replaceAccountMenu(accountId, menuId, data)

Replace an individual menu

This service replaces the details of an individual Menu.

### Example
```java
// Import classes:
//import io.swagger.client.ApiClient;
//import io.swagger.client.ApiException;
//import io.swagger.client.Configuration;
//import io.swagger.client.auth.*;
//import io.swagger.client.api.MenusApi;

ApiClient defaultClient = Configuration.getDefaultApiClient();

// Configure API key authorization: apiKey
ApiKeyAuth apiKey = (ApiKeyAuth) defaultClient.getAuthentication("apiKey");
apiKey.setApiKey("YOUR API KEY");
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//apiKey.setApiKeyPrefix("Token");

MenusApi apiInstance = new MenusApi();
Integer accountId = 56; // Integer | Account ID
Integer menuId = 56; // Integer | Menu ID
ReplaceMenuParams data = new ReplaceMenuParams(); // ReplaceMenuParams | Menu data
try {
    MenuFull result = apiInstance.replaceAccountMenu(accountId, menuId, data);
    System.out.println(result);
} catch (ApiException e) {
    System.err.println("Exception when calling MenusApi#replaceAccountMenu");
    e.printStackTrace();
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **Integer**| Account ID |
 **menuId** | **Integer**| Menu ID |
 **data** | [**ReplaceMenuParams**](ReplaceMenuParams.md)| Menu data | [optional]

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

