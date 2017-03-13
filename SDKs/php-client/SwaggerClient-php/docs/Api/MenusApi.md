# Swagger\Client\MenusApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountMenu**](MenusApi.md#createAccountMenu) | **POST** /accounts/{account_id}/menus | Create an individual menu
[**deleteAccountMenu**](MenusApi.md#deleteAccountMenu) | **DELETE** /accounts/{account_id}/menus/{menu_id} | Delete an individual menu
[**getAccountMenu**](MenusApi.md#getAccountMenu) | **GET** /accounts/{account_id}/menus/{menu_id} | Show details of an individual menu
[**listAccountMenus**](MenusApi.md#listAccountMenus) | **GET** /accounts/{account_id}/menus | Get a list of menus for an account
[**replaceAccountMenu**](MenusApi.md#replaceAccountMenu) | **PUT** /accounts/{account_id}/menus/{menu_id} | Replace an individual menu


# **createAccountMenu**
> \Swagger\Client\Model\MenuFull createAccountMenu($account_id, $data)

Create an individual menu

This service creates an individual menu. See Account Menus for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\MenusApi();
$account_id = 56; // int | Account ID
$data = new \Swagger\Client\Model\CreateMenuParams(); // \Swagger\Client\Model\CreateMenuParams | Menu data

try {
    $result = $api_instance->createAccountMenu($account_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling MenusApi->createAccountMenu: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **data** | [**\Swagger\Client\Model\CreateMenuParams**](../Model/\Swagger\Client\Model\CreateMenuParams.md)| Menu data | [optional]

### Return type

[**\Swagger\Client\Model\MenuFull**](../Model/MenuFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **deleteAccountMenu**
> \Swagger\Client\Model\DeleteMenu deleteAccountMenu($account_id, $menu_id)

Delete an individual menu

This service shows the details of an individual menu.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\MenusApi();
$account_id = 56; // int | Account ID
$menu_id = 56; // int | Menu ID

try {
    $result = $api_instance->deleteAccountMenu($account_id, $menu_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling MenusApi->deleteAccountMenu: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **menu_id** | **int**| Menu ID |

### Return type

[**\Swagger\Client\Model\DeleteMenu**](../Model/DeleteMenu.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getAccountMenu**
> \Swagger\Client\Model\MenuFull getAccountMenu($account_id, $menu_id)

Show details of an individual menu

This service shows the details of an individual Menu.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\MenusApi();
$account_id = 56; // int | Account ID
$menu_id = 56; // int | Menu ID

try {
    $result = $api_instance->getAccountMenu($account_id, $menu_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling MenusApi->getAccountMenu: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **menu_id** | **int**| Menu ID |

### Return type

[**\Swagger\Client\Model\MenuFull**](../Model/MenuFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountMenus**
> \Swagger\Client\Model\ListMenus listAccountMenus($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields)

Get a list of menus for an account

See Account Menus for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\MenusApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_name = array("filters_name_example"); // string[] | Name filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_name = "sort_name_example"; // string | Name sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountMenus($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling MenusApi->listAccountMenus: ', $e->getMessage(), PHP_EOL;
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

[**\Swagger\Client\Model\ListMenus**](../Model/ListMenus.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **replaceAccountMenu**
> \Swagger\Client\Model\MenuFull replaceAccountMenu($account_id, $menu_id, $data)

Replace an individual menu

This service replaces the details of an individual Menu.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\MenusApi();
$account_id = 56; // int | Account ID
$menu_id = 56; // int | Menu ID
$data = new \Swagger\Client\Model\ReplaceMenuParams(); // \Swagger\Client\Model\ReplaceMenuParams | Menu data

try {
    $result = $api_instance->replaceAccountMenu($account_id, $menu_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling MenusApi->replaceAccountMenu: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **menu_id** | **int**| Menu ID |
 **data** | [**\Swagger\Client\Model\ReplaceMenuParams**](../Model/\Swagger\Client\Model\ReplaceMenuParams.md)| Menu data | [optional]

### Return type

[**\Swagger\Client\Model\MenuFull**](../Model/MenuFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

