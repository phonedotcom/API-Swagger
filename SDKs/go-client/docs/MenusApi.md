# \MenusApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountMenu**](MenusApi.md#CreateAccountMenu) | **Post** /accounts/{account_id}/menus | Create an individual menu
[**DeleteAccountMenu**](MenusApi.md#DeleteAccountMenu) | **Delete** /accounts/{account_id}/menus/{menu_id} | Delete an individual menu
[**GetAccountMenu**](MenusApi.md#GetAccountMenu) | **Get** /accounts/{account_id}/menus/{menu_id} | Show details of an individual menu
[**ListAccountMenus**](MenusApi.md#ListAccountMenus) | **Get** /accounts/{account_id}/menus | Get a list of menus for an account
[**ReplaceAccountMenu**](MenusApi.md#ReplaceAccountMenu) | **Put** /accounts/{account_id}/menus/{menu_id} | Replace an individual menu


# **CreateAccountMenu**
> MenuFull CreateAccountMenu($accountId, $data)

Create an individual menu

This service creates an individual menu. See Account Menus for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **data** | [**CreateMenuParams**](CreateMenuParams.md)| Menu data | [optional] 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountMenu**
> DeleteMenu DeleteAccountMenu($accountId, $menuId)

Delete an individual menu

This service shows the details of an individual menu.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **menuId** | **int32**| Menu ID | 

### Return type

[**DeleteMenu**](DeleteMenu.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountMenu**
> MenuFull GetAccountMenu($accountId, $menuId)

Show details of an individual menu

This service shows the details of an individual Menu.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **menuId** | **int32**| Menu ID | 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountMenus**
> ListMenus ListAccountMenus($accountId, $filtersId, $filtersName, $sortId, $sortName, $limit, $offset, $fields)

Get a list of menus for an account

See Account Menus for more info on the properties.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersName** | [**[]string**](string.md)| Name filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortName** | **string**| Name sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListMenus**](ListMenus.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountMenu**
> MenuFull ReplaceAccountMenu($accountId, $menuId, $data)

Replace an individual menu

This service replaces the details of an individual Menu.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **menuId** | **int32**| Menu ID | 
 **data** | [**ReplaceMenuParams**](ReplaceMenuParams.md)| Menu data | [optional] 

### Return type

[**MenuFull**](MenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

