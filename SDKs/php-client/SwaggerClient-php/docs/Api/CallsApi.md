# Swagger\Client\CallsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountCall**](CallsApi.md#createAccountCall) | **POST** /accounts/{account_id}/calls | Make a phone call


# **createAccountCall**
> \Swagger\Client\Model\CallFull createAccountCall($account_id, $data)

Make a phone call



### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\CallsApi();
$account_id = 56; // int | Account ID
$data = new \Swagger\Client\Model\CreateCallParams(); // \Swagger\Client\Model\CreateCallParams | Call data

try {
    $result = $api_instance->createAccountCall($account_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling CallsApi->createAccountCall: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **data** | [**\Swagger\Client\Model\CreateCallParams**](../Model/\Swagger\Client\Model\CreateCallParams.md)| Call data | [optional]

### Return type

[**\Swagger\Client\Model\CallFull**](../Model/CallFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

