# Swagger\Client\NumberregionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listAvailablePhoneNumberRegions**](NumberregionsApi.md#listAvailablePhoneNumberRegions) | **GET** /phone-numbers/available/regions | 


# **listAvailablePhoneNumberRegions**
> \Swagger\Client\Model\ListPhoneNumbersRegions listAvailablePhoneNumberRegions($filters_country_code, $filters_npa, $filters_nxx, $filters_is_toll_free, $filters_city, $filters_province_postal_code, $filters_country_postal_code, $sort_country_code, $sort_npa, $sort_nxx, $sort_is_toll_free, $sort_city, $sort_province_postal_code, $sort_country_postal_code, $limit, $offset, $fields, $group_by)



This service lists the quantities of available phone numbers by region.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

// Configure API key authorization: apiKey
Swagger\Client\Configuration::getDefaultConfiguration()->setApiKey('Authorization', 'YOUR_API_KEY');
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
// Swagger\Client\Configuration::getDefaultConfiguration()->setApiKeyPrefix('Authorization', 'Bearer');

$api_instance = new Swagger\Client\Api\NumberregionsApi();
$filters_country_code = array("filters_country_code_example"); // string[] | Country Code filter
$filters_npa = array("filters_npa_example"); // string[] | Area Code filter (North America only)
$filters_nxx = array("filters_nxx_example"); // string[] | 2nd set of 3 digits filter (North America only)
$filters_is_toll_free = array("filters_is_toll_free_example"); // string[] | Toll-free status filter
$filters_city = array("filters_city_example"); // string[] | City filter
$filters_province_postal_code = array("filters_province_postal_code_example"); // string[] | State or Province filter
$filters_country_postal_code = array("filters_country_postal_code_example"); // string[] | Country filter
$sort_country_code = "sort_country_code_example"; // string | International calling code sorting
$sort_npa = "sort_npa_example"; // string | Area Code sorting (North America only)
$sort_nxx = "sort_nxx_example"; // string | 2nd set of 3 digits sorting (North America)
$sort_is_toll_free = "sort_is_toll_free_example"; // string | Toll Free status sorting
$sort_city = "sort_city_example"; // string | City sorting
$sort_province_postal_code = "sort_province_postal_code_example"; // string | State or Province sorting
$sort_country_postal_code = "sort_country_postal_code_example"; // string | Country sorting
$limit = 56; // int | Max results
$offset = 56; // int | Results to skip
$fields = "fields_example"; // string | Field set
$group_by = array("group_by_example"); // string[] | Fields to group by (supports the same set of fields as the filters and sorting)

try {
    $result = $api_instance->listAvailablePhoneNumberRegions($filters_country_code, $filters_npa, $filters_nxx, $filters_is_toll_free, $filters_city, $filters_province_postal_code, $filters_country_postal_code, $sort_country_code, $sort_npa, $sort_nxx, $sort_is_toll_free, $sort_city, $sort_province_postal_code, $sort_country_postal_code, $limit, $offset, $fields, $group_by);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling NumberregionsApi->listAvailablePhoneNumberRegions: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters_country_code** | [**string[]**](../Model/string.md)| Country Code filter | [optional]
 **filters_npa** | [**string[]**](../Model/string.md)| Area Code filter (North America only) | [optional]
 **filters_nxx** | [**string[]**](../Model/string.md)| 2nd set of 3 digits filter (North America only) | [optional]
 **filters_is_toll_free** | [**string[]**](../Model/string.md)| Toll-free status filter | [optional]
 **filters_city** | [**string[]**](../Model/string.md)| City filter | [optional]
 **filters_province_postal_code** | [**string[]**](../Model/string.md)| State or Province filter | [optional]
 **filters_country_postal_code** | [**string[]**](../Model/string.md)| Country filter | [optional]
 **sort_country_code** | **string**| International calling code sorting | [optional]
 **sort_npa** | **string**| Area Code sorting (North America only) | [optional]
 **sort_nxx** | **string**| 2nd set of 3 digits sorting (North America) | [optional]
 **sort_is_toll_free** | **string**| Toll Free status sorting | [optional]
 **sort_city** | **string**| City sorting | [optional]
 **sort_province_postal_code** | **string**| State or Province sorting | [optional]
 **sort_country_postal_code** | **string**| Country sorting | [optional]
 **limit** | **int**| Max results | [optional]
 **offset** | **int**| Results to skip | [optional]
 **fields** | **string**| Field set | [optional]
 **group_by** | [**string[]**](../Model/string.md)| Fields to group by (supports the same set of fields as the filters and sorting) | [optional]

### Return type

[**\Swagger\Client\Model\ListPhoneNumbersRegions**](../Model/ListPhoneNumbersRegions.md)

### Authorization

[apiKey](../../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

