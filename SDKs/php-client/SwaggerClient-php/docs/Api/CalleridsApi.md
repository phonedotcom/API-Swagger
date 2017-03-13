# Swagger\Client\CalleridsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getCallerIds**](CalleridsApi.md#getCallerIds) | **GET** /accounts/{account_id}/extensions/{extension_id}/caller-ids | Show the Caller ID options a given extension can use


# **getCallerIds**
> \Swagger\Client\Model\ListCallerIds getCallerIds($account_id, $extension_id, $filters_number, $filters_name, $sort_number, $sort_name, $limit, $offset, $fields)

Show the Caller ID options a given extension can use

Get Caller ID

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\CalleridsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$filters_number = array("filters_number_example"); // string[] | Number filter
$filters_name = array("filters_name_example"); // string[] | Name filter
$sort_number = "sort_number_example"; // string | Number sorting
$sort_name = "sort_name_example"; // string | Name sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->getCallerIds($account_id, $extension_id, $filters_number, $filters_name, $sort_number, $sort_name, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CalleridsApi->getCallerIds: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **filters_number** | [**string[]**](../Model/string.md)| Number filter | [optional]
 **filters_name** | [**string[]**](../Model/string.md)| Name filter | [optional]
 **sort_number** | **string**| Number sorting | [optional]
 **sort_name** | **string**| Name sorting | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]

### Return type

[**\Swagger\Client\Model\ListCallerIds**](../Model/ListCallerIds.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

