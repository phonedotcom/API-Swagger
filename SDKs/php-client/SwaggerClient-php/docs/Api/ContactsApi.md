# Swagger\Client\ContactsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountExtensionContact**](ContactsApi.md#createAccountExtensionContact) | **POST** /accounts/{account_id}/extensions/{extension_id}/contacts | Add a new address book contact for an extension
[**deleteAccountExtensionContact**](ContactsApi.md#deleteAccountExtensionContact) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | 
[**getAccountExtensionContact**](ContactsApi.md#getAccountExtensionContact) | **GET** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | Retrieve the details of an address book contact
[**listAccountExtensionContacts**](ContactsApi.md#listAccountExtensionContacts) | **GET** /accounts/{account_id}/extensions/{extension_id}/contacts | Show a list of address book contacts
[**replaceAccountExtensionContact**](ContactsApi.md#replaceAccountExtensionContact) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | 


# **createAccountExtensionContact**
> \Swagger\Client\Model\ContactFull createAccountExtensionContact($account_id, $extension_id, $data)

Add a new address book contact for an extension

For more on the input fields, see Account Contacts.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ContactsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$data = new \Swagger\Client\Model\CreateContactParams(); // \Swagger\Client\Model\CreateContactParams | Contact data

try {
    $result = $api_instance->createAccountExtensionContact($account_id, $extension_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContactsApi->createAccountExtensionContact: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **data** | [**\Swagger\Client\Model\CreateContactParams**](../Model/\Swagger\Client\Model\CreateContactParams.md)| Contact data | [optional]

### Return type

[**\Swagger\Client\Model\ContactFull**](../Model/ContactFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **deleteAccountExtensionContact**
> \Swagger\Client\Model\DeleteContact deleteAccountExtensionContact($account_id, $extension_id, $contact_id)





### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ContactsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$contact_id = 56; // int | Contact ID

try {
    $result = $api_instance->deleteAccountExtensionContact($account_id, $extension_id, $contact_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContactsApi->deleteAccountExtensionContact: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **contact_id** | **int**| Contact ID |

### Return type

[**\Swagger\Client\Model\DeleteContact**](../Model/DeleteContact.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getAccountExtensionContact**
> \Swagger\Client\Model\ContactFull getAccountExtensionContact($account_id, $extension_id, $contact_id)

Retrieve the details of an address book contact

For more info on the fields shown, see Account Contacts.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ContactsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$contact_id = 56; // int | Contact ID

try {
    $result = $api_instance->getAccountExtensionContact($account_id, $extension_id, $contact_id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContactsApi->getAccountExtensionContact: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **contact_id** | **int**| Contact ID |

### Return type

[**\Swagger\Client\Model\ContactFull**](../Model/ContactFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listAccountExtensionContacts**
> \Swagger\Client\Model\ListContacts listAccountExtensionContacts($account_id, $extension_id, $filters_id, $filters_group_id, $filters_updated_at, $sort_id, $sort_updated_at, $limit, $offset, $fields)

Show a list of address book contacts

See Account Contacts for more info on the fields in each item.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ContactsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$filters_id = array("filters_id_example"); // string[] | ID filter
$filters_group_id = array("filters_group_id_example"); // string[] | Group filter
$filters_updated_at = array("filters_updated_at_example"); // string[] | Updated At filter
$sort_id = "sort_id_example"; // string | ID sorting
$sort_updated_at = "sort_updated_at_example"; // string | Updated At sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAccountExtensionContacts($account_id, $extension_id, $filters_id, $filters_group_id, $filters_updated_at, $sort_id, $sort_updated_at, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContactsApi->listAccountExtensionContacts: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **filters_id** | [**string[]**](../Model/string.md)| ID filter | [optional]
 **filters_group_id** | [**string[]**](../Model/string.md)| Group filter | [optional]
 **filters_updated_at** | [**string[]**](../Model/string.md)| Updated At filter | [optional]
 **sort_id** | **string**| ID sorting | [optional]
 **sort_updated_at** | **string**| Updated At sorting | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]

### Return type

[**\Swagger\Client\Model\ListContacts**](../Model/ListContacts.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **replaceAccountExtensionContact**
> \Swagger\Client\Model\ContactFull replaceAccountExtensionContact($account_id, $extension_id, $contact_id, $data)



For more on the input fields, see Account Contacts.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\ContactsApi();
$account_id = 56; // int | Account ID
$extension_id = 56; // int | Extension ID
$contact_id = 56; // int | Contact ID
$data = new \Swagger\Client\Model\CreateContactParams(); // \Swagger\Client\Model\CreateContactParams | Contact data

try {
    $result = $api_instance->replaceAccountExtensionContact($account_id, $extension_id, $contact_id, $data);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling ContactsApi->replaceAccountExtensionContact: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID |
 **extension_id** | **int**| Extension ID |
 **contact_id** | **int**| Contact ID |
 **data** | [**\Swagger\Client\Model\CreateContactParams**](../Model/\Swagger\Client\Model\CreateContactParams.md)| Contact data | [optional]

### Return type

[**\Swagger\Client\Model\ContactFull**](../Model/ContactFull.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

