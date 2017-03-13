# Swagger\Client\ExpressservicecodesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAccountExpressSrvCode**](ExpressservicecodesApi.md#getAccountExpressSrvCode) | **GET** /accounts/{account_id}/express-service-codes/{code_id} | Show details of an account Express Service Code
[**listAccountExpressSrvCodes**](ExpressservicecodesApi.md#listAccountExpressSrvCodes) | **GET** /accounts/{account_id}/express-service-codes | Get the Express Service Code associated with your account in list format


# **getAccountExpressSrvCode**
> \Swagger\Client\Model\ExpressServiceCodeFull getAccountExpressSrvCode($account_id, $code_id)

Show details of an account Express Service Code

This service shows the details of an Account Express Service Code.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ExpressservicecodesApi();
$account_id = 56; // int | Account ID
$code_id = 56; // int | Device ID

try {
    $result = $api_instance->getAccountExpressSrvCode($account_id, $code_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ExpressservicecodesApi->getAccountExpressSrvCode: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **code_id** | **int**| Device ID |

### Return type

[**\Swagger\Client\Model\ExpressServiceCodeFull**](../Model/ExpressServiceCodeFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountExpressSrvCodes**
> \Swagger\Client\Model\ListExpressServiceCodes listAccountExpressSrvCodes($account_id, $filters_id)

Get the Express Service Code associated with your account in list format

See Express Service Codes for more detail.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ExpressservicecodesApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter

try {
    $result = $api_instance->listAccountExpressSrvCodes($account_id, $filters_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ExpressservicecodesApi->listAccountExpressSrvCodes: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **filters_id** | [**string[]**](../Model/string.md)| ID filter | [optional]

### Return type

[**\Swagger\Client\Model\ListExpressServiceCodes**](../Model/ListExpressServiceCodes.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

