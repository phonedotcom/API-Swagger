# \NumberregionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailablePhoneNumberRegions**](NumberregionsApi.md#ListAvailablePhoneNumberRegions) | **Get** /phone-numbers/available/regions | 


# **ListAvailablePhoneNumberRegions**
> ListPhoneNumbersRegions ListAvailablePhoneNumberRegions($filtersCountryCode, $filtersNpa, $filtersNxx, $filtersIsTollFree, $filtersCity, $filtersProvincePostalCode, $filtersCountryPostalCode, $sortCountryCode, $sortNpa, $sortNxx, $sortIsTollFree, $sortCity, $sortProvincePostalCode, $sortCountryPostalCode, $limit, $offset, $fields, $groupBy)



This service lists the quantities of available phone numbers by region.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersCountryCode** | [**[]string**](string.md)| Country Code filter | [optional] 
 **filtersNpa** | [**[]string**](string.md)| Area Code filter (North America only) | [optional] 
 **filtersNxx** | [**[]string**](string.md)| 2nd set of 3 digits filter (North America only) | [optional] 
 **filtersIsTollFree** | [**[]string**](string.md)| Toll-free status filter | [optional] 
 **filtersCity** | [**[]string**](string.md)| City filter | [optional] 
 **filtersProvincePostalCode** | [**[]string**](string.md)| State or Province filter | [optional] 
 **filtersCountryPostalCode** | [**[]string**](string.md)| Country filter | [optional] 
 **sortCountryCode** | **string**| International calling code sorting | [optional] 
 **sortNpa** | **string**| Area Code sorting (North America only) | [optional] 
 **sortNxx** | **string**| 2nd set of 3 digits sorting (North America) | [optional] 
 **sortIsTollFree** | **string**| Toll Free status sorting | [optional] 
 **sortCity** | **string**| City sorting | [optional] 
 **sortProvincePostalCode** | **string**| State or Province sorting | [optional] 
 **sortCountryPostalCode** | **string**| Country sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 
 **groupBy** | [**[]string**](string.md)| Fields to group by (supports the same set of fields as the filters and sorting) | [optional] 

### Return type

[**ListPhoneNumbersRegions**](ListPhoneNumbersRegions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

