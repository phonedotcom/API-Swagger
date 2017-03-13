# ContactFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** | Integer ID. Read-only. | [optional] 
**prefix** | **str** | Salutation, such as Mr, Mrs, or Dr | [optional] 
**first_name** | **str** | First name or given name | [optional] 
**middle_name** | **str** | Middle or second name | [optional] 
**last_name** | **str** | Last name or surname | [optional] 
**suffix** | **str** | Suffix, such as \&quot;Jr.\&quot;, \&quot;Sr.\&quot;, \&quot;II\&quot;, or \&quot;III\&quot; | [optional] 
**nickname** | **str** | Nickname, or a shortened informal version of the contact&#39;s name | [optional] 
**company** | **str** | Name of the contact&#39;s company | [optional] 
**phonetic_first_name** | **str** | Phonetic first name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**phonetic_middle_name** | **str** | Phonetic middle name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**phonetic_last_name** | **str** | Phonetic last name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**department** | **str** | Name of the contact&#39;s department | [optional] 
**job_title** | **str** | Contact&#39;s job title | [optional] 
**emails** | [**list[Email]**](Email.md) | Array of Contact Email Objects. See below for details. | [optional] 
**phone_numbers** | [**list[PhoneNumberContact]**](PhoneNumberContact.md) | Array of Contact Phone Number Objects. See below for details. | [optional] 
**addresses** | [**list[AddressListContacts]**](AddressListContacts.md) | Array of Contact Address Objects. See below for details. | [optional] 
**group** | [**GroupListContacts**](GroupListContacts.md) |  | [optional] 
**created_at** | **int** | Integer UNIX timestamp when the contact was created. Read-only. | [optional] 
**updated_at** | **int** | Integer UNIX timestamp when the contact was created. Read-only. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


