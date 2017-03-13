# Swagger\Client\RoutesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createRoute**](RoutesApi.md#createRoute) | **POST** /accounts/{account_id}/routes | Add a new address book contact for an extension
[**deleteAccountRoute**](RoutesApi.md#deleteAccountRoute) | **DELETE** /accounts/{account_id}/routes/{route_id} | 
[**getAccountRoute**](RoutesApi.md#getAccountRoute) | **GET** /accounts/{account_id}/routes/{route_id} | Show details of an individual route
[**listAccountRoutes**](RoutesApi.md#listAccountRoutes) | **GET** /accounts/{account_id}/routes | Get a list of routes for an account
[**replaceAccountRoute**](RoutesApi.md#replaceAccountRoute) | **PUT** /accounts/{account_id}/routes/{route_id} | 


# **createRoute**
> \Swagger\Client\Model\RouteFull createRoute($account_id, $data)

Add a new address book contact for an extension

For more on the input fields, see Intro to Routes.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\RoutesApi();
$account_id = 56; // int | Account ID
$data = new \Swagger\Client\Model\CreateRouteParams(); // \Swagger\Client\Model\CreateRouteParams | Route data

try {
    $result = $api_instance->createRoute($account_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RoutesApi->createRoute: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **data** | [**\Swagger\Client\Model\CreateRouteParams**](../Model/\Swagger\Client\Model\CreateRouteParams.md)| Route data | [optional]

### Return type

[**\Swagger\Client\Model\RouteFull**](../Model/RouteFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **deleteAccountRoute**
> \Swagger\Client\Model\DeleteRoute deleteAccountRoute($account_id, $route_id)





### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\RoutesApi();
$account_id = 56; // int | Account ID
$route_id = 56; // int | Route ID

try {
    $result = $api_instance->deleteAccountRoute($account_id, $route_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RoutesApi->deleteAccountRoute: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **route_id** | **int**| Route ID |

### Return type

[**\Swagger\Client\Model\DeleteRoute**](../Model/DeleteRoute.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getAccountRoute**
> \Swagger\Client\Model\RouteFull getAccountRoute($account_id, $route_id)

Show details of an individual route

This service shows the details of an individual route.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\RoutesApi();
$account_id = 56; // int | Account ID
$route_id = 56; // int | Route ID

try {
    $result = $api_instance->getAccountRoute($account_id, $route_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RoutesApi->getAccountRoute: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **route_id** | **int**| Route ID |

### Return type

[**\Swagger\Client\Model\RouteFull**](../Model/RouteFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountRoutes**
> \Swagger\Client\Model\ListRoutes listAccountRoutes($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields)

Get a list of routes for an account

See Intro to Routes for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\RoutesApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_name = array("filters_name_example"); // string[] | Name filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_name = "sort_name_example"; // string | Name sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountRoutes($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RoutesApi->listAccountRoutes: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **filters_id** | [**string[]**](../Model/string.md)| ID filter | [optional]
 **filters_name** | [**string[]**](../Model/string.md)| Name filter | [optional]
 **sort_id** | **string**| ID sorting | [optional]
 **sort_name** | **string**| Name sorting | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]

### Return type

[**\Swagger\Client\Model\ListRoutes**](../Model/ListRoutes.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **replaceAccountRoute**
> \Swagger\Client\Model\RouteFull replaceAccountRoute($account_id, $route_id, $data)



For more on the input fields, see Intro to Routes.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\RoutesApi();
$account_id = 56; // int | Account ID
$route_id = 56; // int | Route ID
$data = new \Swagger\Client\Model\CreateRouteParams(); // \Swagger\Client\Model\CreateRouteParams | Route data

try {
    $result = $api_instance->replaceAccountRoute($account_id, $route_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RoutesApi->replaceAccountRoute: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **route_id** | **int**| Route ID |
 **data** | [**\Swagger\Client\Model\CreateRouteParams**](../Model/\Swagger\Client\Model\CreateRouteParams.md)| Route data | [optional]

### Return type

[**\Swagger\Client\Model\RouteFull**](../Model/RouteFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

