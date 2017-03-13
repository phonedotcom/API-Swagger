# Swagger\Client\SubaccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountSubaccount**](SubaccountsApi.md#createAccountSubaccount) | **POST** /accounts/{account_id}/subaccounts | Add a subaccount for the authenticated user or client
[**listAccountSubaccounts**](SubaccountsApi.md#listAccountSubaccounts) | **GET** /accounts/{account_id}/subaccounts | Get a list of subaccounts for the authenticated user or client


# **createAccountSubaccount**
> \Swagger\Client\Model\AccountFull createAccountSubaccount($account_id, $data)

Add a subaccount for the authenticated user or client

This service shows the details of an individual Subaccount.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\SubaccountsApi();
$account_id = 56; // int | Account ID
$data = new \Swagger\Client\Model\CreateSubaccountParams(); // \Swagger\Client\Model\CreateSubaccountParams | SMS data

try {
    $result = $api_instance->createAccountSubaccount($account_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubaccountsApi->createAccountSubaccount: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **data** | [**\Swagger\Client\Model\CreateSubaccountParams**](../Model/\Swagger\Client\Model\CreateSubaccountParams.md)| SMS data |

### Return type

[**\Swagger\Client\Model\AccountFull**](../Model/AccountFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountSubaccounts**
> \Swagger\Client\Model\ListAccounts listAccountSubaccounts($account_id, $filters_id, $sort_id, $limit, $offset, $fields)

Get a list of subaccounts for the authenticated user or client

This service lists the Subaccount of the authenticated client. In most cases, there will not be any.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\SubaccountsApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$sort_id = "sort_id_example"; // string | ID sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountSubaccounts($account_id, $filters_id, $sort_id, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SubaccountsApi->listAccountSubaccounts: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **filters_id** | [**string[]**](../Model/string.md)| ID filter | [optional]
 **sort_id** | **string**| ID sorting | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]

### Return type

[**\Swagger\Client\Model\ListAccounts**](../Model/ListAccounts.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

