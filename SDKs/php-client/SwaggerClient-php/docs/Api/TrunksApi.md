# Swagger\Client\TrunksApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountTrunk**](TrunksApi.md#createAccountTrunk) | **POST** /accounts/{account_id}/trunks | Add a trunk record with SIP information
[**deleteAccountTrunk**](TrunksApi.md#deleteAccountTrunk) | **DELETE** /accounts/{account_id}/trunks/{trunk_id} | Delete a trunk from account
[**getAccountTrunk**](TrunksApi.md#getAccountTrunk) | **GET** /accounts/{account_id}/trunks/{trunk_id} | Show details of an individual trunk
[**listAccountTrunks**](TrunksApi.md#listAccountTrunks) | **GET** /accounts/{account_id}/trunks | Get a list of trunks for an account
[**replaceAccountTrunk**](TrunksApi.md#replaceAccountTrunk) | **PUT** /accounts/{account_id}/trunks/{trunk_id} | Replace parameters in a trunk


# **createAccountTrunk**
> \Swagger\Client\Model\TrunkFull createAccountTrunk($account_id, $data)

Add a trunk record with SIP information

For more on the input fields, see Account Trunks.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\TrunksApi();
$account_id = 56; // int | Account ID
$data = new \Swagger\Client\Model\CreateTrunkParams(); // \Swagger\Client\Model\CreateTrunkParams | Trunk data

try {
    $result = $api_instance->createAccountTrunk($account_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TrunksApi->createAccountTrunk: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **data** | [**\Swagger\Client\Model\CreateTrunkParams**](../Model/\Swagger\Client\Model\CreateTrunkParams.md)| Trunk data |

### Return type

[**\Swagger\Client\Model\TrunkFull**](../Model/TrunkFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **deleteAccountTrunk**
> \Swagger\Client\Model\DeleteTrunk deleteAccountTrunk($account_id, $trunk_id)

Delete a trunk from account

This service deletes a trunk from the account. For more on the properties of trunks, see Account Trunks.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\TrunksApi();
$account_id = 56; // int | Account ID
$trunk_id = 56; // int | Trunk ID

try {
    $result = $api_instance->deleteAccountTrunk($account_id, $trunk_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TrunksApi->deleteAccountTrunk: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **trunk_id** | **int**| Trunk ID |

### Return type

[**\Swagger\Client\Model\DeleteTrunk**](../Model/DeleteTrunk.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getAccountTrunk**
> \Swagger\Client\Model\TrunkFull getAccountTrunk($account_id, $trunk_id)

Show details of an individual trunk

This service shows the details of an individual Trunk.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\TrunksApi();
$account_id = 56; // int | Account ID
$trunk_id = 56; // int | Trunk ID

try {
    $result = $api_instance->getAccountTrunk($account_id, $trunk_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TrunksApi->getAccountTrunk: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **trunk_id** | **int**| Trunk ID |

### Return type

[**\Swagger\Client\Model\TrunkFull**](../Model/TrunkFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountTrunks**
> \Swagger\Client\Model\ListTrunks listAccountTrunks($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields)

Get a list of trunks for an account

See Account Trunks for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\TrunksApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_name = array("filters_name_example"); // string[] | Name filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_name = "sort_name_example"; // string | Name sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountTrunks($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TrunksApi->listAccountTrunks: ', $e->getMessage(), PHP_EOL;
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

[**\Swagger\Client\Model\ListTrunks**](../Model/ListTrunks.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **replaceAccountTrunk**
> \Swagger\Client\Model\TrunkFull replaceAccountTrunk($account_id, $trunk_id, $data)

Replace parameters in a trunk

For more on the input fields, see Account Trunks.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\TrunksApi();
$account_id = 56; // int | Account ID
$trunk_id = 56; // int | Trunk ID
$data = new \Swagger\Client\Model\CreateTrunkParams(); // \Swagger\Client\Model\CreateTrunkParams | Trunk data

try {
    $result = $api_instance->replaceAccountTrunk($account_id, $trunk_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling TrunksApi->replaceAccountTrunk: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **trunk_id** | **int**| Trunk ID |
 **data** | [**\Swagger\Client\Model\CreateTrunkParams**](../Model/\Swagger\Client\Model\CreateTrunkParams.md)| Trunk data |

### Return type

[**\Swagger\Client\Model\TrunkFull**](../Model/TrunkFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

