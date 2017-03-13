# Swagger\Client\PhonenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountPhoneNumber**](PhonenumbersApi.md#createAccountPhoneNumber) | **POST** /accounts/{account_id}/phone-numbers | Add a phone number to an account
[**getAccountPhoneNumber**](PhonenumbersApi.md#getAccountPhoneNumber) | **GET** /accounts/{account_id}/phone-numbers/{number_id} | Show details of an individual phone number
[**listAccountPhoneNumbers**](PhonenumbersApi.md#listAccountPhoneNumbers) | **GET** /accounts/{account_id}/phone-numbers | Get a list of phone numbers registered to an account
[**replaceAccountPhoneNumber**](PhonenumbersApi.md#replaceAccountPhoneNumber) | **PUT** /accounts/{account_id}/phone-numbers/{number_id} | Update the settings for an existing phone number on your account


# **createAccountPhoneNumber**
> \Swagger\Client\Model\PhoneNumberFull createAccountPhoneNumber($account_id, $data)

Add a phone number to an account

See Intro to Account Phone Numbers for more info on the properties to use.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\PhonenumbersApi();
$account_id = 56; // int | Account ID
$data = new \Swagger\Client\Model\CreatePhoneNumberParams(); // \Swagger\Client\Model\CreatePhoneNumberParams | Phone Number data

try {
    $result = $api_instance->createAccountPhoneNumber($account_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PhonenumbersApi->createAccountPhoneNumber: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **data** | [**\Swagger\Client\Model\CreatePhoneNumberParams**](../Model/\Swagger\Client\Model\CreatePhoneNumberParams.md)| Phone Number data | [optional]

### Return type

[**\Swagger\Client\Model\PhoneNumberFull**](../Model/PhoneNumberFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getAccountPhoneNumber**
> \Swagger\Client\Model\PhoneNumberFull getAccountPhoneNumber($account_id, $number_id)

Show details of an individual phone number

See Intro to Account Phone Numbers for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\PhonenumbersApi();
$account_id = 56; // int | Account ID
$number_id = 56; // int | Number ID

try {
    $result = $api_instance->getAccountPhoneNumber($account_id, $number_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PhonenumbersApi->getAccountPhoneNumber: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **number_id** | **int**| Number ID |

### Return type

[**\Swagger\Client\Model\PhoneNumberFull**](../Model/PhoneNumberFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountPhoneNumbers**
> \Swagger\Client\Model\ListPhoneNumbers listAccountPhoneNumbers($account_id, $filters_id, $filters_name, $filters_phone_number, $sort_id, $sort_name, $sort_phone_number, $limit, $offset, $fields)

Get a list of phone numbers registered to an account

See Intro to Account Phone Numbers for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\PhonenumbersApi();
$account_id = 56; // int | Account ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_name = array("filters_name_example"); // string[] | Name filter
$filters_phone_number = array("filters_phone_number_example"); // string[] | Phone number filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_name = "sort_name_example"; // string | Name sorting
$sort_phone_number = "sort_phone_number_example"; // string | Phone number sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountPhoneNumbers($account_id, $filters_id, $filters_name, $filters_phone_number, $sort_id, $sort_name, $sort_phone_number, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PhonenumbersApi->listAccountPhoneNumbers: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **filters_id** | [**string[]**](../Model/string.md)| ID filter | [optional]
 **filters_name** | [**string[]**](../Model/string.md)| Name filter | [optional]
 **filters_phone_number** | [**string[]**](../Model/string.md)| Phone number filter | [optional]
 **sort_id** | **string**| ID sorting | [optional]
 **sort_name** | **string**| Name sorting | [optional]
 **sort_phone_number** | **string**| Phone number sorting | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]

### Return type

[**\Swagger\Client\Model\ListPhoneNumbers**](../Model/ListPhoneNumbers.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **replaceAccountPhoneNumber**
> \Swagger\Client\Model\PhoneNumberFull replaceAccountPhoneNumber($account_id, $number_id, $data)

Update the settings for an existing phone number on your account

See Intro to Account Phone Numbers for more info on the properties.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\PhonenumbersApi();
$account_id = 56; // int | Account ID
$number_id = 56; // int | Number ID
$data = new \Swagger\Client\Model\ReplacePhoneNumberParams(); // \Swagger\Client\Model\ReplacePhoneNumberParams | Phone Number data

try {
    $result = $api_instance->replaceAccountPhoneNumber($account_id, $number_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling PhonenumbersApi->replaceAccountPhoneNumber: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **number_id** | **int**| Number ID |
 **data** | [**\Swagger\Client\Model\ReplacePhoneNumberParams**](../Model/\Swagger\Client\Model\ReplacePhoneNumberParams.md)| Phone Number data | [optional]

### Return type

[**\Swagger\Client\Model\PhoneNumberFull**](../Model/PhoneNumberFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

