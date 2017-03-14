# Swagger\Client\CalllogsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAccountCallLogs**](CalllogsApi.md#getAccountCallLogs) | **GET** /accounts/{account_id}/call-logs/{call_id} | Show details of an individual Call Log entry
[**listAccountCallLogs**](CalllogsApi.md#listAccountCallLogs) | **GET** /accounts/{account_id}/call-logs | Get a list of call details associated with your account


# **getAccountCallLogs**
> \Swagger\Client\Model\CallLogFull getAccountCallLogs($account_id, $call_id)

Show details of an individual Call Log entry

See Call Logs for more detail.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\CalllogsApi();
$account_id = 56; // int | Account ID
$call_id = "call_id_example"; // string | Call ID

try {
    $result = $api_instance->getAccountCallLogs($account_id, $call_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CalllogsApi->getAccountCallLogs: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **call_id** | **string**| Call ID |

### Return type

[**\Swagger\Client\Model\CallLogFull**](../Model/CallLogFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountCallLogs**
> \Swagger\Client\Model\ListCallLogs listAccountCallLogs($account_id, $filters_id, $filters_start_time, $filters_created_at, $filters_direction, $filters_called_number, $filters_type, $filters_extension, $sort_id, $sort_start_time, $sort_created_at, $limit, $offset, $fields)

Get a list of call details associated with your account

See Call Logs for more detail.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\CalllogsApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_start_time = array("filters_start_time_example"); // string[] | Call start time filter
$filters_created_at = "filters_created_at_example"; // string | Call log creation time filter
$filters_direction = "filters_direction_example"; // string | Call direction filter: in or out
$filters_called_number = "filters_called_number_example"; // string | Called number
$filters_type = "filters_type_example"; // string | Call type, such as 'call', 'fax', 'audiogram'
$filters_extension = array("filters_extension_example"); // string[] | Extension filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_start_time = "sort_start_time_example"; // string | Sorting by call start time: asc or desc
$sort_created_at = "sort_created_at_example"; // string | Sorting by call log creation time: asc or desc
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountCallLogs($account_id, $filters_id, $filters_start_time, $filters_created_at, $filters_direction, $filters_called_number, $filters_type, $filters_extension, $sort_id, $sort_start_time, $sort_created_at, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CalllogsApi->listAccountCallLogs: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **filters_id** | [**string[]**](../Model/string.md)| ID filter | [optional]
 **filters_start_time** | [**string[]**](../Model/string.md)| Call start time filter | [optional]
 **filters_created_at** | **string**| Call log creation time filter | [optional]
 **filters_direction** | **string**| Call direction filter: in or out | [optional]
 **filters_called_number** | **string**| Called number | [optional]
 **filters_type** | **string**| Call type, such as &#39;call&#39;, &#39;fax&#39;, &#39;audiogram&#39; | [optional]
 **filters_extension** | [**string[]**](../Model/string.md)| Extension filter | [optional]
 **sort_id** | **string**| ID sorting | [optional]
 **sort_start_time** | **string**| Sorting by call start time: asc or desc | [optional]
 **sort_created_at** | **string**| Sorting by call log creation time: asc or desc | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]

### Return type

[**\Swagger\Client\Model\ListCallLogs**](../Model/ListCallLogs.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

