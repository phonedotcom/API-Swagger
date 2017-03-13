# Swagger\Client\GroupsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountExtensionContactGroup**](GroupsApi.md#createAccountExtensionContactGroup) | **POST** /accounts/{account_id}/extensions/{extension_id}/contact-groups | 
[**deleteAccountExtensionContactGroup**](GroupsApi.md#deleteAccountExtensionContactGroup) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | Delete an addressbook group
[**getAccountExtensionContactGroup**](GroupsApi.md#getAccountExtensionContactGroup) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 
[**listAccountExtensionContactGroups**](GroupsApi.md#listAccountExtensionContactGroups) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups | Show a list of contact groups belonging to an extension
[**replaceAccountExtensionContactGroup**](GroupsApi.md#replaceAccountExtensionContactGroup) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 


# **createAccountExtensionContactGroup**
> \Swagger\Client\Model\GroupFull createAccountExtensionContactGroup($account_id, $extension_id, $data)



See Account Contact Groups for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\GroupsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$data = new \Swagger\Client\Model\CreateGroupParams(); // \Swagger\Client\Model\CreateGroupParams | Group name

try {
    $result = $api_instance->createAccountExtensionContactGroup($account_id, $extension_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GroupsApi->createAccountExtensionContactGroup: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **data** | [**\Swagger\Client\Model\CreateGroupParams**](../Model/\Swagger\Client\Model\CreateGroupParams.md)| Group name |

### Return type

[**\Swagger\Client\Model\GroupFull**](../Model/GroupFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **deleteAccountExtensionContactGroup**
> \Swagger\Client\Model\DeleteGroup deleteAccountExtensionContactGroup($account_id, $extension_id, $group_id)

Delete an addressbook group



### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\GroupsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$group_id = 56; // int | Group ID

try {
    $result = $api_instance->deleteAccountExtensionContactGroup($account_id, $extension_id, $group_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GroupsApi->deleteAccountExtensionContactGroup: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **group_id** | **int**| Group ID |

### Return type

[**\Swagger\Client\Model\DeleteGroup**](../Model/DeleteGroup.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getAccountExtensionContactGroup**
> \Swagger\Client\Model\GroupFull getAccountExtensionContactGroup($account_id, $extension_id, $group_id)



See Account Contact Groups for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\GroupsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$group_id = 56; // int | Group ID

try {
    $result = $api_instance->getAccountExtensionContactGroup($account_id, $extension_id, $group_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GroupsApi->getAccountExtensionContactGroup: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **group_id** | **int**| Group ID |

### Return type

[**\Swagger\Client\Model\GroupFull**](../Model/GroupFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountExtensionContactGroups**
> \Swagger\Client\Model\ListGroups listAccountExtensionContactGroups($account_id, $extension_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields)

Show a list of contact groups belonging to an extension

See Account Contact Groups for details on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\GroupsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_name = array("filters_name_example"); // string[] | Name filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_name = "sort_name_example"; // string | Name sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountExtensionContactGroups($account_id, $extension_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GroupsApi->listAccountExtensionContactGroups: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **filters_id** | [**string[]**](../Model/string.md)| ID filter | [optional]
 **filters_name** | [**string[]**](../Model/string.md)| Name filter | [optional]
 **sort_id** | **string**| ID sorting | [optional]
 **sort_name** | **string**| Name sorting | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]

### Return type

[**\Swagger\Client\Model\ListGroups**](../Model/ListGroups.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **replaceAccountExtensionContactGroup**
> \Swagger\Client\Model\GroupFull replaceAccountExtensionContactGroup($account_id, $extension_id, $group_id, $data)



See Account Contact Groups for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\GroupsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$group_id = 56; // int | Group ID
$data = new \Swagger\Client\Model\CreateGroupParams(); // \Swagger\Client\Model\CreateGroupParams | Group name

try {
    $result = $api_instance->replaceAccountExtensionContactGroup($account_id, $extension_id, $group_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling GroupsApi->replaceAccountExtensionContactGroup: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **group_id** | **int**| Group ID |
 **data** | [**\Swagger\Client\Model\CreateGroupParams**](../Model/\Swagger\Client\Model\CreateGroupParams.md)| Group name |

### Return type

[**\Swagger\Client\Model\GroupFull**](../Model/GroupFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

