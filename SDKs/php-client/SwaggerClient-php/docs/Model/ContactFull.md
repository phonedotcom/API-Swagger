# ContactFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Integer ID. Read-only. | [optional] 
**prefix** | **string** | Salutation, such as Mr, Mrs, or Dr | [optional] 
**first_name** | **string** | First name or given name | [optional] 
**middle_name** | **string** | Middle or second name | [optional] 
**last_name** | **string** | Last name or surname | [optional] 
**suffix** | **string** | Suffix, such as \&quot;Jr.\&quot;, \&quot;Sr.\&quot;, \&quot;II\&quot;, or \&quot;III\&quot; | [optional] 
**nickname** | **string** | Nickname, or a shortened informal version of the contact&#39;s name | [optional] 
**company** | **string** | Name of the contact&#39;s company | [optional] 
**phonetic_first_name** | **string** | Phonetic first name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**phonetic_middle_name** | **string** | Phonetic middle name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**phonetic_last_name** | **string** | Phonetic last name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**department** | **string** | Name of the contact&#39;s department | [optional] 
**job_title** | **string** | Contact&#39;s job title | [optional] 
**emails** | [**\Swagger\Client\Model\Email[]**](Email.md) | Array of Contact Email Objects. See below for details. | [optional] 
**phone_numbers** | [**\Swagger\Client\Model\PhoneNumberContact[]**](PhoneNumberContact.md) | Array of Contact Phone Number Objects. See below for details. | [optional] 
**addresses** | [**\Swagger\Client\Model\AddressListContacts[]**](AddressListContacts.md) | Array of Contact Address Objects. See below for details. | [optional] 
**group** | [**\Swagger\Client\Model\GroupListContacts**](GroupListContacts.md) |  | [optional] 
**created_at** | **int** | Integer UNIX timestamp when the contact was created. Read-only. | [optional] 
**updated_at** | **int** | Integer UNIX timestamp when the contact was created. Read-only. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


