# Swagger\Client\ExtensionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountExtension**](ExtensionsApi.md#createAccountExtension) | **POST** /accounts/{account_id}/extensions | Create an individual extension
[**getAccountExtension**](ExtensionsApi.md#getAccountExtension) | **GET** /accounts/{account_id}/extensions/{extension_id} | Show details of an individual extension
[**listAccountExtensions**](ExtensionsApi.md#listAccountExtensions) | **GET** /accounts/{account_id}/extensions | Get a list of extensions visible to the authenticated user or client
[**replaceAccountExtension**](ExtensionsApi.md#replaceAccountExtension) | **PUT** /accounts/{account_id}/extensions/{extension_id} | Replace an individual extension


# **createAccountExtension**
> \Swagger\Client\Model\ExtensionFull createAccountExtension($account_id, $data)

Create an individual extension

This service shows how to create a virtual extension.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ExtensionsApi();
$account_id = 56; // int | Account ID
$data = new \Swagger\Client\Model\CreateExtensionParams(); // \Swagger\Client\Model\CreateExtensionParams | Account Extensions Data

try {
    $result = $api_instance->createAccountExtension($account_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ExtensionsApi->createAccountExtension: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **data** | [**\Swagger\Client\Model\CreateExtensionParams**](../Model/\Swagger\Client\Model\CreateExtensionParams.md)| Account Extensions Data | [optional]

### Return type

[**\Swagger\Client\Model\ExtensionFull**](../Model/ExtensionFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getAccountExtension**
> \Swagger\Client\Model\ExtensionFull getAccountExtension($account_id, $extension_id)

Show details of an individual extension

This service shows the details of an individual Extension.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ExtensionsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID

try {
    $result = $api_instance->getAccountExtension($account_id, $extension_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ExtensionsApi->getAccountExtension: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |

### Return type

[**\Swagger\Client\Model\ExtensionFull**](../Model/ExtensionFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountExtensions**
> \Swagger\Client\Model\ListExtensions listAccountExtensions($account_id, $filters_id, $filters_extension, $filters_name, $sort_id, $sort_extension, $sort_name, $limit, $offset, $fields)

Get a list of extensions visible to the authenticated user or client

This service lists the visible extensions on a given account.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ExtensionsApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_extension = array("filters_extension_example"); // string[] | Extension filter
$filters_name = array("filters_name_example"); // string[] | Name filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_extension = "sort_extension_example"; // string | Extension sorting
$sort_name = "sort_name_example"; // string | Name sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountExtensions($account_id, $filters_id, $filters_extension, $filters_name, $sort_id, $sort_extension, $sort_name, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ExtensionsApi->listAccountExtensions: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **filters_id** | [**string[]**](../Model/string.md)| ID filter | [optional]
 **filters_extension** | [**string[]**](../Model/string.md)| Extension filter | [optional]
 **filters_name** | [**string[]**](../Model/string.md)| Name filter | [optional]
 **sort_id** | **string**| ID sorting | [optional]
 **sort_extension** | **string**| Extension sorting | [optional]
 **sort_name** | **string**| Name sorting | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]

### Return type

[**\Swagger\Client\Model\ListExtensions**](../Model/ListExtensions.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **replaceAccountExtension**
> \Swagger\Client\Model\ExtensionFull replaceAccountExtension($account_id, $extension_id, $data)

Replace an individual extension

This service shows how to update an individual extension.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ExtensionsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$data = new \Swagger\Client\Model\ReplaceExtensionParams(); // \Swagger\Client\Model\ReplaceExtensionParams | Account Extensions Data

try {
    $result = $api_instance->replaceAccountExtension($account_id, $extension_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ExtensionsApi->replaceAccountExtension: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **data** | [**\Swagger\Client\Model\ReplaceExtensionParams**](../Model/\Swagger\Client\Model\ReplaceExtensionParams.md)| Account Extensions Data | [optional]

### Return type

[**\Swagger\Client\Model\ExtensionFull**](../Model/ExtensionFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

