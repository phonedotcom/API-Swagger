# Swagger\Client\SmsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountSms**](SmsApi.md#createAccountSms) | **POST** /accounts/{account_id}/sms | Send a SMS to one or a group of recipients
[**getAccountSms**](SmsApi.md#getAccountSms) | **GET** /accounts/{account_id}/sms/{sms_id} | Show details of an individual SMS
[**listAccountSms**](SmsApi.md#listAccountSms) | **GET** /accounts/{account_id}/sms | Get a list of SMS messages for an account


# **createAccountSms**
> \Swagger\Client\Model\SmsFull createAccountSms($account_id, $data)

Send a SMS to one or a group of recipients

For more on the input fields, see Intro to SMS.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\SmsApi();
$account_id = 56; // int | Account ID
$data = new \Swagger\Client\Model\CreateSmsParams(); // \Swagger\Client\Model\CreateSmsParams | SMS data

try {
    $result = $api_instance->createAccountSms($account_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SmsApi->createAccountSms: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **data** | [**\Swagger\Client\Model\CreateSmsParams**](../Model/\Swagger\Client\Model\CreateSmsParams.md)| SMS data |

### Return type

[**\Swagger\Client\Model\SmsFull**](../Model/SmsFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getAccountSms**
> \Swagger\Client\Model\SmsFull getAccountSms($account_id, $sms_id)

Show details of an individual SMS

This service shows the details of an individual sms. See Intro to SMS for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\SmsApi();
$account_id = 56; // int | Account ID
$sms_id = "sms_id_example"; // string | SMS ID

try {
    $result = $api_instance->getAccountSms($account_id, $sms_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SmsApi->getAccountSms: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **sms_id** | **string**| SMS ID |

### Return type

[**\Swagger\Client\Model\SmsFull**](../Model/SmsFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountSms**
> \Swagger\Client\Model\ListSms listAccountSms($account_id, $filters_id, $filters_direction, $filters_from, $sort_id, $sort_created_at, $limit, $offset, $fields)

Get a list of SMS messages for an account

See Intro to SMS for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\SmsApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_direction = "filters_direction_example"; // string | Direction filter
$filters_from = "filters_from_example"; // string | Caller ID filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_created_at = "sort_created_at_example"; // string | Sort by created time of message
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountSms($account_id, $filters_id, $filters_direction, $filters_from, $sort_id, $sort_created_at, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling SmsApi->listAccountSms: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **filters_id** | [**string[]**](../Model/string.md)| ID filter | [optional]
 **filters_direction** | **string**| Direction filter | [optional]
 **filters_from** | **string**| Caller ID filter | [optional]
 **sort_id** | **string**| ID sorting | [optional]
 **sort_created_at** | **string**| Sort by created time of message | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]

### Return type

[**\Swagger\Client\Model\ListSms**](../Model/ListSms.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

