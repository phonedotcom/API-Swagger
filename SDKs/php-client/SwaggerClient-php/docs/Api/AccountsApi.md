# Swagger\Client\AccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAccount**](AccountsApi.md#getAccount) | **GET** /accounts/{account_id} | Retrieve details of an individual account
[**listAccounts**](AccountsApi.md#listAccounts) | **GET** /accounts | Get a list of accounts visible to the authenticated user or client


# **getAccount**
> \Swagger\Client\Model\AccountFull getAccount($account_id)

Retrieve details of an individual account

This service shows the details of an individual account. See Accounts for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\AccountsApi();
$account_id = 56; // int | Account ID

try {
    $result = $api_instance->getAccount($account_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AccountsApi->getAccount: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |

### Return type

[**\Swagger\Client\Model\AccountFull**](../Model/AccountFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccounts**
> \Swagger\Client\Model\ListAccounts listAccounts($filters_id, $sort_id, $limit, $offset, $fields)

Get a list of accounts visible to the authenticated user or client

This service lists the accounts accessible to the authenticated client. In most cases, there will only be one such account. See Accounts for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\AccountsApi();
$filters_id = array("filters_id_example"); // string[] | ID filter
$sort_id = "sort_id_example"; // string | ID sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccounts($filters_id, $sort_id, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AccountsApi->listAccounts: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

