# Swagger\Client\QueuesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountQueue**](QueuesApi.md#createAccountQueue) | **POST** /accounts/{account_id}/queues | Create a queue
[**deleteAccountQueue**](QueuesApi.md#deleteAccountQueue) | **DELETE** /accounts/{account_id}/queues/{queue_id} | Delete a queue
[**getAccountQueue**](QueuesApi.md#getAccountQueue) | **GET** /accounts/{account_id}/queues/{queue_id} | Show details of an individual queue
[**listAccountQueues**](QueuesApi.md#listAccountQueues) | **GET** /accounts/{account_id}/queues | Get a list of queues for an account
[**replaceAccountQueue**](QueuesApi.md#replaceAccountQueue) | **PUT** /accounts/{account_id}/queues/{queue_id} | Replace a queue


# **createAccountQueue**
> \Swagger\Client\Model\QueueFull createAccountQueue($account_id, $data)

Create a queue

For more on the input fields, see Account Queues.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\QueuesApi();
$account_id = 56; // int | Account ID
$data = new \Swagger\Client\Model\CreateQueueParams(); // \Swagger\Client\Model\CreateQueueParams | Queue data

try {
    $result = $api_instance->createAccountQueue($account_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling QueuesApi->createAccountQueue: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **data** | [**\Swagger\Client\Model\CreateQueueParams**](../Model/\Swagger\Client\Model\CreateQueueParams.md)| Queue data | [optional]

### Return type

[**\Swagger\Client\Model\QueueFull**](../Model/QueueFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **deleteAccountQueue**
> \Swagger\Client\Model\DeleteQueue deleteAccountQueue($account_id, $queue_id)

Delete a queue

This service a queue from the account. For more information on queue properties, see Account Queues.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\QueuesApi();
$account_id = 56; // int | Account ID
$queue_id = 56; // int | Queue ID

try {
    $result = $api_instance->deleteAccountQueue($account_id, $queue_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling QueuesApi->deleteAccountQueue: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **queue_id** | **int**| Queue ID |

### Return type

[**\Swagger\Client\Model\DeleteQueue**](../Model/DeleteQueue.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getAccountQueue**
> \Swagger\Client\Model\QueueFull getAccountQueue($account_id, $queue_id)

Show details of an individual queue

This service shows the details of an individual queue. For more on the input fields, see Account Queues.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\QueuesApi();
$account_id = 56; // int | Account ID
$queue_id = 56; // int | Queue ID

try {
    $result = $api_instance->getAccountQueue($account_id, $queue_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling QueuesApi->getAccountQueue: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **queue_id** | **int**| Queue ID |

### Return type

[**\Swagger\Client\Model\QueueFull**](../Model/QueueFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountQueues**
> \Swagger\Client\Model\ListQueues listAccountQueues($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields)

Get a list of queues for an account

The List Queues service lists all the queues belong to the account. See Account Queues for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\QueuesApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_name = array("filters_name_example"); // string[] | Name filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_name = "sort_name_example"; // string | Name sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountQueues($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling QueuesApi->listAccountQueues: ', $e->getMessage(), PHP_EOL;
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

[**\Swagger\Client\Model\ListQueues**](../Model/ListQueues.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **replaceAccountQueue**
> \Swagger\Client\Model\QueueFull replaceAccountQueue($account_id, $queue_id, $data)

Replace a queue

The Replace Queue service replaces the parameters of a queue. For more on the input fields, see Account Queues.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\QueuesApi();
$account_id = 56; // int | Account ID
$queue_id = 56; // int | Queue ID
$data = new \Swagger\Client\Model\CreateQueueParams(); // \Swagger\Client\Model\CreateQueueParams | Queue data

try {
    $result = $api_instance->replaceAccountQueue($account_id, $queue_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling QueuesApi->replaceAccountQueue: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **queue_id** | **int**| Queue ID |
 **data** | [**\Swagger\Client\Model\CreateQueueParams**](../Model/\Swagger\Client\Model\CreateQueueParams.md)| Queue data | [optional]

### Return type

[**\Swagger\Client\Model\QueueFull**](../Model/QueueFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

