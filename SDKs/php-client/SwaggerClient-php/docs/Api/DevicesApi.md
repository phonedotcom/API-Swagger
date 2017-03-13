# Swagger\Client\DevicesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountDevice**](DevicesApi.md#createAccountDevice) | **POST** /accounts/{account_id}/devices | Register a generic VoIP device
[**getAccountDevice**](DevicesApi.md#getAccountDevice) | **GET** /accounts/{account_id}/devices/{device_id} | Show details of an individual VoIP device
[**listAccountDevices**](DevicesApi.md#listAccountDevices) | **GET** /accounts/{account_id}/devices | Get a list of VoIP devices associated with your account
[**replaceAccountDevice**](DevicesApi.md#replaceAccountDevice) | **PUT** /accounts/{account_id}/devices/{device_id} | Update the settings for an individual VoIP device


# **createAccountDevice**
> \Swagger\Client\Model\DeviceFull createAccountDevice($account_id, $data)

Register a generic VoIP device



### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\DevicesApi();
$account_id = 56; // int | Account ID
$data = new \Swagger\Client\Model\CreateDeviceParams(); // \Swagger\Client\Model\CreateDeviceParams | Device data

try {
    $result = $api_instance->createAccountDevice($account_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DevicesApi->createAccountDevice: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **data** | [**\Swagger\Client\Model\CreateDeviceParams**](../Model/\Swagger\Client\Model\CreateDeviceParams.md)| Device data | [optional]

### Return type

[**\Swagger\Client\Model\DeviceFull**](../Model/DeviceFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getAccountDevice**
> \Swagger\Client\Model\DeviceFull getAccountDevice($account_id, $device_id)

Show details of an individual VoIP device



### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\DevicesApi();
$account_id = 56; // int | Account ID
$device_id = 56; // int | Device ID

try {
    $result = $api_instance->getAccountDevice($account_id, $device_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DevicesApi->getAccountDevice: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **device_id** | **int**| Device ID |

### Return type

[**\Swagger\Client\Model\DeviceFull**](../Model/DeviceFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountDevices**
> \Swagger\Client\Model\ListDevices listAccountDevices($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields)

Get a list of VoIP devices associated with your account



### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\DevicesApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_name = array("filters_name_example"); // string[] | Name filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_name = "sort_name_example"; // string | Name sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountDevices($account_id, $filters_id, $filters_name, $sort_id, $sort_name, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DevicesApi->listAccountDevices: ', $e->getMessage(), PHP_EOL;
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

[**\Swagger\Client\Model\ListDevices**](../Model/ListDevices.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **replaceAccountDevice**
> \Swagger\Client\Model\DeviceFull replaceAccountDevice($account_id, $device_id, $data)

Update the settings for an individual VoIP device



### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\DevicesApi();
$account_id = 56; // int | Account ID
$device_id = 56; // int | Device ID
$data = new \Swagger\Client\Model\CreateDeviceParams(); // \Swagger\Client\Model\CreateDeviceParams | Device data

try {
    $result = $api_instance->replaceAccountDevice($account_id, $device_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DevicesApi->replaceAccountDevice: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **device_id** | **int**| Device ID |
 **data** | [**\Swagger\Client\Model\CreateDeviceParams**](../Model/\Swagger\Client\Model\CreateDeviceParams.md)| Device data | [optional]

### Return type

[**\Swagger\Client\Model\DeviceFull**](../Model/DeviceFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

