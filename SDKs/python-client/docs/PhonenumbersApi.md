# swagger_client.PhonenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_account_phone_number**](PhonenumbersApi.md#create_account_phone_number) | **POST** /accounts/{account_id}/phone-numbers | Add a phone number to an account
[**get_account_phone_number**](PhonenumbersApi.md#get_account_phone_number) | **GET** /accounts/{account_id}/phone-numbers/{number_id} | Show details of an individual phone number
[**list_account_phone_numbers**](PhonenumbersApi.md#list_account_phone_numbers) | **GET** /accounts/{account_id}/phone-numbers | Get a list of phone numbers registered to an account
[**replace_account_phone_number**](PhonenumbersApi.md#replace_account_phone_number) | **PUT** /accounts/{account_id}/phone-numbers/{number_id} | Update the settings for an existing phone number on your account


# **create_account_phone_number**
> PhoneNumberFull create_account_phone_number(account_id, data=data)

Add a phone number to an account

See Intro to Account Phone Numbers for more info on the properties to use.

### Example 
```python
from __future__ import print_statement
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# Configure API key authorization: apiKey
swagger_client.configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# swagger_client.configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = swagger_client.PhonenumbersApi()
account_id = 56 # int | Account ID
data = swagger_client.CreatePhoneNumberParams() # CreatePhoneNumberParams | Phone Number data (optional)

try: 
    # Add a phone number to an account
    api_response = api_instance.create_account_phone_number(account_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling PhonenumbersApi->create_account_phone_number: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **data** | [**CreatePhoneNumberParams**](CreatePhoneNumberParams.md)| Phone Number data | [optional] 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_account_phone_number**
> PhoneNumberFull get_account_phone_number(account_id, number_id)

Show details of an individual phone number

See Intro to Account Phone Numbers for more info on the properties.

### Example 
```python
from __future__ import print_statement
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# Configure API key authorization: apiKey
swagger_client.configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# swagger_client.configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = swagger_client.PhonenumbersApi()
account_id = 56 # int | Account ID
number_id = 56 # int | Number ID

try: 
    # Show details of an individual phone number
    api_response = api_instance.get_account_phone_number(account_id, number_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling PhonenumbersApi->get_account_phone_number: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **number_id** | **int**| Number ID | 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_account_phone_numbers**
> ListPhoneNumbers list_account_phone_numbers(account_id, filters_id=filters_id, filters_name=filters_name, filters_phone_number=filters_phone_number, sort_id=sort_id, sort_name=sort_name, sort_phone_number=sort_phone_number, limit=limit, offset=offset, fields=fields)

Get a list of phone numbers registered to an account

See Intro to Account Phone Numbers for more info on the properties.

### Example 
```python
from __future__ import print_statement
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# Configure API key authorization: apiKey
swagger_client.configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# swagger_client.configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = swagger_client.PhonenumbersApi()
account_id = 56 # int | Account ID
filters_id = ['filters_id_example'] # list[str] | ID filter (optional)
filters_name = ['filters_name_example'] # list[str] | Name filter (optional)
filters_phone_number = ['filters_phone_number_example'] # list[str] | Phone number filter (optional)
sort_id = 'sort_id_example' # str | ID sorting (optional)
sort_name = 'sort_name_example' # str | Name sorting (optional)
sort_phone_number = 'sort_phone_number_example' # str | Phone number sorting (optional)
limit = 56 # int | Max results (optional)
offset = 56 # int | Results to skip (optional)
fields = 'fields_example' # str | Field set (optional)

try: 
    # Get a list of phone numbers registered to an account
    api_response = api_instance.list_account_phone_numbers(account_id, filters_id=filters_id, filters_name=filters_name, filters_phone_number=filters_phone_number, sort_id=sort_id, sort_name=sort_name, sort_phone_number=sort_phone_number, limit=limit, offset=offset, fields=fields)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling PhonenumbersApi->list_account_phone_numbers: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **filters_id** | [**list[str]**](str.md)| ID filter | [optional] 
 **filters_name** | [**list[str]**](str.md)| Name filter | [optional] 
 **filters_phone_number** | [**list[str]**](str.md)| Phone number filter | [optional] 
 **sort_id** | **str**| ID sorting | [optional] 
 **sort_name** | **str**| Name sorting | [optional] 
 **sort_phone_number** | **str**| Phone number sorting | [optional] 
 **limit** | **int**| Max results | [optional] 
 **offset** | **int**| Results to skip | [optional] 
 **fields** | **str**| Field set | [optional] 

### Return type

[**ListPhoneNumbers**](ListPhoneNumbers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replace_account_phone_number**
> PhoneNumberFull replace_account_phone_number(account_id, number_id, data=data)

Update the settings for an existing phone number on your account

See Intro to Account Phone Numbers for more info on the properties.

### Example 
```python
from __future__ import print_statement
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# Configure API key authorization: apiKey
swagger_client.configuration.api_key['Authorization'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# swagger_client.configuration.api_key_prefix['Authorization'] = 'Bearer'

# create an instance of the API class
api_instance = swagger_client.PhonenumbersApi()
account_id = 56 # int | Account ID
number_id = 56 # int | Number ID
data = swagger_client.ReplacePhoneNumberParams() # ReplacePhoneNumberParams | Phone Number data (optional)

try: 
    # Update the settings for an existing phone number on your account
    api_response = api_instance.replace_account_phone_number(account_id, number_id, data=data)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling PhonenumbersApi->replace_account_phone_number: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account_id** | **int**| Account ID | 
 **number_id** | **int**| Number ID | 
 **data** | [**ReplacePhoneNumberParams**](ReplacePhoneNumberParams.md)| Phone Number data | [optional] 

### Return type

[**PhoneNumberFull**](PhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

