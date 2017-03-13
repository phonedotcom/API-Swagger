# Swagger\Client\AvailablenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listAvailablePhoneNumbers**](AvailablenumbersApi.md#listAvailablePhoneNumbers) | **GET** /phone-numbers/available | 


# **listAvailablePhoneNumbers**
> \Swagger\Client\Model\ListAvailableNumbers listAvailablePhoneNumbers($filters_phone_number, $filters_country_code, $filters_npa, $filters_nxx, $filters_xxxx, $filters_city, $filters_province, $filters_country, $filters_price, $filters_category, $sort_internal, $sort_price, $sort_phone_number, $limit, $offset, $fields)





### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\AvailablenumbersApi();
$filters_phone_number = array("filters_phone_number_example"); // string[] | Phone number filter
$filters_country_code = array("filters_country_code_example"); // string[] | Country Code filter
$filters_npa = array("filters_npa_example"); // string[] | Area Code filter (North America only)
$filters_nxx = array("filters_nxx_example"); // string[] | 2nd set of 3 digits filter (North America only)
$filters_xxxx = array("filters_xxxx_example"); // string[] | NANP XXXX filter
$filters_city = array("filters_city_example"); // string[] | City filter
$filters_province = array("filters_province_example"); // string[] | State or Province (postal code) filter
$filters_country = array("filters_country_example"); // string[] | Country (postal code) filter
$filters_price = array("filters_price_example"); // string[] | Price filter
$filters_category = array("filters_category_example"); // string[] | Category filter
$sort_internal = "sort_internal_example"; // string | Internal (quasi-random) sorting
$sort_price = "sort_price_example"; // string | Price sorting
$sort_phone_number = "sort_phone_number_example"; // string | Phone number sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set

try {
    $result = $api_instance->listAvailablePhoneNumbers($filters_phone_number, $filters_country_code, $filters_npa, $filters_nxx, $filters_xxxx, $filters_city, $filters_province, $filters_country, $filters_price, $filters_category, $sort_internal, $sort_price, $sort_phone_number, $limit, $offset, $fields);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling AvailablenumbersApi->listAvailablePhoneNumbers: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters_phone_number** | [**string[]**](../Model/string.md)| Phone number filter | [optional]
 **filters_country_code** | [**string[]**](../Model/string.md)| Country Code filter | [optional]
 **filters_npa** | [**string[]**](../Model/string.md)| Area Code filter (North America only) | [optional]
 **filters_nxx** | [**string[]**](../Model/string.md)| 2nd set of 3 digits filter (North America only) | [optional]
 **filters_xxxx** | [**string[]**](../Model/string.md)| NANP XXXX filter | [optional]
 **filters_city** | [**string[]**](../Model/string.md)| City filter | [optional]
 **filters_province** | [**string[]**](../Model/string.md)| State or Province (postal code) filter | [optional]
 **filters_country** | [**string[]**](../Model/string.md)| Country (postal code) filter | [optional]
 **filters_price** | [**string[]**](../Model/string.md)| Price filter | [optional]
 **filters_category** | [**string[]**](../Model/string.md)| Category filter | [optional]
 **sort_internal** | **string**| Internal (quasi-random) sorting | [optional]
 **sort_price** | **string**| Price sorting | [optional]
 **sort_phone_number** | **string**| Phone number sorting | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]

### Return type

[**\Swagger\Client\Model\ListAvailableNumbers**](../Model/ListAvailableNumbers.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

