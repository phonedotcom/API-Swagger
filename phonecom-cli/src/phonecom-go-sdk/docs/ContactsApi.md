# \ContactsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountExtensionContact**](ContactsApi.md#CreateAccountExtensionContact) | **Post** /accounts/{account_id}/extensions/{extension_id}/contacts | Add a new address book contact for an extension
[**DeleteAccountExtensionContact**](ContactsApi.md#DeleteAccountExtensionContact) | **Delete** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | 
[**GetAccountExtensionContact**](ContactsApi.md#GetAccountExtensionContact) | **Get** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | Retrieve the details of an address book contact
[**ListAccountExtensionContacts**](ContactsApi.md#ListAccountExtensionContacts) | **Get** /accounts/{account_id}/extensions/{extension_id}/contacts | Show a list of address book contacts
[**ReplaceAccountExtensionContact**](ContactsApi.md#ReplaceAccountExtensionContact) | **Put** /accounts/{account_id}/extensions/{extension_id}/contacts | 


# **CreateAccountExtensionContact**
> ContactFull CreateAccountExtensionContact($accountId, $extensionId, $data)

Add a new address book contact for an extension

For more on the input fields, see Account Contacts.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **data** | [**CreateContactParams**](CreateContactParams.md)| Contact data | [optional] 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAccountExtensionContact**
> DeleteContact DeleteAccountExtensionContact($accountId, $extensionId, $contactId)






### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **contactId** | **int32**| Contact ID | 

### Return type

[**DeleteContact**](DeleteContact.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountExtensionContact**
> ContactFull GetAccountExtensionContact($accountId, $extensionId, $contactId)

Retrieve the details of an address book contact

For more info on the fields shown, see Account Contacts.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **contactId** | **int32**| Contact ID | 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAccountExtensionContacts**
> ListContacts ListAccountExtensionContacts($accountId, $extensionId, $filtersId, $filtersGroupId, $filtersUpdatedAt, $sortId, $sortUpdatedAt, $limit, $offset, $fields)

Show a list of address book contacts

See Account Contacts for more info on the fields in each item.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **filtersId** | [**[]string**](string.md)| ID filter | [optional] 
 **filtersGroupId** | [**[]string**](string.md)| Group filter | [optional] 
 **filtersUpdatedAt** | [**[]string**](string.md)| Updated At filter | [optional] 
 **sortId** | **string**| ID sorting | [optional] 
 **sortUpdatedAt** | **string**| Updated At sorting | [optional] 
 **limit** | **int32**| Max results | [optional] 
 **offset** | **int32**| Results to skip | [optional] 
 **fields** | **string**| Field set | [optional] 

### Return type

[**ListContacts**](ListContacts.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceAccountExtensionContact**
> ContactFull ReplaceAccountExtensionContact($accountId, $extensionId, $data)



For more on the input fields, see Account Contacts.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **int32**| Account ID | 
 **extensionId** | **int32**| Extension ID | 
 **data** | [**CreateContactParams**](CreateContactParams.md)| Contact data | [optional] 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

