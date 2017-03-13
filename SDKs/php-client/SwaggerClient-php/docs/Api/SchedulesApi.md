# Swagger\Client\SchedulesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAccountSchedule**](SchedulesApi.md#getAccountSchedule) | **GET** /accounts/{account_id}/schedules/{schedule_id} | Show details of an individual schedule
[**listAccountSchedules**](SchedulesApi.md#listAccountSchedules) | **GET** /accounts/{account_id}/schedules | Get a list of schedules for an account


# **getAccountSchedule**
> \Swagger\Client\Model\ScheduleFull getAccountSchedule($account_id, $schedule_id)

Show details of an individual schedule

This service shows the details of an individual schedule.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\SchedulesApi();
$account_id = 56; // int | Account ID
$schedule_id = 56; // int | Schedule ID

try {
    $result = $api_instance->getAccountSchedule($account_id, $schedule_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SchedulesApi->getAccountSchedule: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **schedule_id** | **int**| Schedule ID |

### Return type

[**\Swagger\Client\Model\ScheduleFull**](../Model/ScheduleFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountSchedules**
> \Swagger\Client\Model\ListSchedules listAccountSchedules($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields)

Get a list of schedules for an account

See Intro to Schedules for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\SchedulesApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_name = array("filters_name_example"); // string[] | Name filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_name = "sort_name_example"; // string | Name sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountSchedules($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SchedulesApi->listAccountSchedules: ', $e->getMessage(), PHP_EOL;
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

[**\Swagger\Client\Model\ListSchedules**](../Model/ListSchedules.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

