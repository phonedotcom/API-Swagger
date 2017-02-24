# \AvailablenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailablePhoneNumbers**](AvailablenumbersApi.md#ListAvailablePhoneNumbers) | **Get** /phone-numbers/available | 


# **ListAvailablePhoneNumbers**
> ListAvailableNumbers ListAvailablePhoneNumbers($filtersPhoneNumber, $filtersCountryCode, $filtersNpa, $filtersNxx, $filtersXxxx, $filtersCity, $filtersProvince, $filtersCountry, $filtersPrice, $filtersCategory, $sortInternal, $sortPrice, $sortPhoneNumber, $limit, $offset, $fields)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersPhoneNumber** | [**[]string**](string.md)| Phone number filter | [optional] 
 **filtersCountryCode** | [**[]string**](string.md)| Country Code filter | [optional] 
 **filtersNpa** | [**[]string**](string.md)| Area Code filter (North America only) | [optional] 
 **filtersNxx** | [**[]string**](string.md)| 2nd set of 3 digits filter (North America only) | [optional] 
 **filtersXxxx** | [**[]string**](string.md)| NANP XXXX filter | [optional] 
 **filtersCity** | [**[]string**](string.md)| City filter | [optional] 
 **filtersProvince** | [**[]string**](string.md)| State or Province (postal code) filter | [optional] 
 **filtersCountry** | [**[]string**](string.md)| Country (postal code) filter | [optional] 
 **filtersPrice** | [**[]string**](string.md)| Price filter | [optional] 
 **filtersCategory** | [**[]string**](string.md)| Category filter | [optional] 
 **sortInternal** | **string**| Internal (quasi-random) sorting | [optional] 
 **sortPrice** | **string**| Price sorting | [optional] 
 **sortPhoneNumber** | **string**| Phone number sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListAvailableNumbers**](ListAvailableNumbers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

