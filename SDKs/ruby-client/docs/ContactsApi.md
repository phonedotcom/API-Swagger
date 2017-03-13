# SwaggerClient::ContactsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_extension_contact**](ContactsApi.md#create_account_extension_contact) | **POST** /accounts/{account_id}/extensions/{extension_id}/contacts | Add a new address book contact for an extension
[**delete_account_extension_contact**](ContactsApi.md#delete_account_extension_contact) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | 
[**get_account_extension_contact**](ContactsApi.md#get_account_extension_contact) | **GET** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | Retrieve the details of an address book contact
[**list_account_extension_contacts**](ContactsApi.md#list_account_extension_contacts) | **GET** /accounts/{account_id}/extensions/{extension_id}/contacts | Show a list of address book contacts
[**replace_account_extension_contact**](ContactsApi.md#replace_account_extension_contact) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id} | 


# **create_account_extension_contact**
> ContactFull create_account_extension_contact(account_id, extension_id, , opts)

Add a new address book contact for an extension

For more on the input fields, see Account Contacts.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure API key authorization: apiKey
  config.api_key['Authorization'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  #config.api_key_prefix['Authorization'] = 'Bearer'
end

api_instance = SwaggerClient::ContactsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

opts = { 
  data: SwaggerClient::CreateContactParams.new # CreateContactParams | Contact data
}

begin
  #Add a new address book contact for an extension
  result = api_instance.create_account_extension_contact(account_id, extension_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ContactsApi->create_account_extension_contact: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **data** | [**CreateContactParams**](CreateContactParams.md)| Contact data | [optional] 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **delete_account_extension_contact**
> DeleteContact delete_account_extension_contact(account_id, extension_id, contact_id)





### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure API key authorization: apiKey
  config.api_key['Authorization'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  #config.api_key_prefix['Authorization'] = 'Bearer'
end

api_instance = SwaggerClient::ContactsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

contact_id = 56 # Integer | Contact ID


begin
  #
  result = api_instance.delete_account_extension_contact(account_id, extension_id, contact_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ContactsApi->delete_account_extension_contact: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **contact_id** | **Integer**| Contact ID | 

### Return type

[**DeleteContact**](DeleteContact.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **get_account_extension_contact**
> ContactFull get_account_extension_contact(account_id, extension_id, contact_id)

Retrieve the details of an address book contact

For more info on the fields shown, see Account Contacts.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure API key authorization: apiKey
  config.api_key['Authorization'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  #config.api_key_prefix['Authorization'] = 'Bearer'
end

api_instance = SwaggerClient::ContactsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

contact_id = 56 # Integer | Contact ID


begin
  #Retrieve the details of an address book contact
  result = api_instance.get_account_extension_contact(account_id, extension_id, contact_id)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ContactsApi->get_account_extension_contact: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **contact_id** | **Integer**| Contact ID | 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **list_account_extension_contacts**
> ListContacts list_account_extension_contacts(account_id, extension_id, , opts)

Show a list of address book contacts

See Account Contacts for more info on the fields in each item.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure API key authorization: apiKey
  config.api_key['Authorization'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  #config.api_key_prefix['Authorization'] = 'Bearer'
end

api_instance = SwaggerClient::ContactsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

opts = { 
  filters_id: ["filters_id_example"], # Array<String> | ID filter
  filters_group_id: ["filters_group_id_example"], # Array<String> | Group filter
  filters_updated_at: ["filters_updated_at_example"], # Array<String> | Updated At filter
  sort_id: "sort_id_example", # String | ID sorting
  sort_updated_at: "sort_updated_at_example", # String | Updated At sorting
  limit: 56, # Integer | Max results
  offset: 56, # Integer | Results to skip
  fields: "fields_example", # String | Field set
}

begin
  #Show a list of address book contacts
  result = api_instance.list_account_extension_contacts(account_id, extension_id, , opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ContactsApi->list_account_extension_contacts: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **filters_id** | [**Array&lt;String&gt;**](String.md)| ID filter | [optional] 
 **filters_group_id** | [**Array&lt;String&gt;**](String.md)| Group filter | [optional] 
 **filters_updated_at** | [**Array&lt;String&gt;**](String.md)| Updated At filter | [optional] 
 **sort_id** | **String**| ID sorting | [optional] 
 **sort_updated_at** | **String**| Updated At sorting | [optional] 
 **limit** | **Integer**| Max results | [optional] 
 **offset** | **Integer**| Results to skip | [optional] 
 **fields** | **String**| Field set | [optional] 

### Return type

[**ListContacts**](ListContacts.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



# **replace_account_extension_contact**
> ContactFull replace_account_extension_contact(account_id, extension_id, contact_id, opts)



For more on the input fields, see Account Contacts.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure API key authorization: apiKey
  config.api_key['Authorization'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  #config.api_key_prefix['Authorization'] = 'Bearer'
end

api_instance = SwaggerClient::ContactsApi.new

account_id = 56 # Integer | Account ID

extension_id = 56 # Integer | Extension ID

contact_id = 56 # Integer | Contact ID

opts = { 
  data: SwaggerClient::CreateContactParams.new # CreateContactParams | Contact data
}

begin
  #
  result = api_instance.replace_account_extension_contact(account_id, extension_id, contact_id, opts)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling ContactsApi->replace_account_extension_contact: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **Integer**| Account ID | 
 **extension_id** | **Integer**| Extension ID | 
 **contact_id** | **Integer**| Contact ID | 
 **data** | [**CreateContactParams**](CreateContactParams.md)| Contact data | [optional] 

### Return type

[**ContactFull**](ContactFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json



