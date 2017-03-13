# SwaggerClient::ContactFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Integer** | Integer ID. Read-only. | [optional] 
**prefix** | **String** | Salutation, such as Mr, Mrs, or Dr | [optional] 
**first_name** | **String** | First name or given name | [optional] 
**middle_name** | **String** | Middle or second name | [optional] 
**last_name** | **String** | Last name or surname | [optional] 
**suffix** | **String** | Suffix, such as \&quot;Jr.\&quot;, \&quot;Sr.\&quot;, \&quot;II\&quot;, or \&quot;III\&quot; | [optional] 
**nickname** | **String** | Nickname, or a shortened informal version of the contact&#39;s name | [optional] 
**company** | **String** | Name of the contact&#39;s company | [optional] 
**phonetic_first_name** | **String** | Phonetic first name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**phonetic_middle_name** | **String** | Phonetic middle name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**phonetic_last_name** | **String** | Phonetic last name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**department** | **String** | Name of the contact&#39;s department | [optional] 
**job_title** | **String** | Contact&#39;s job title | [optional] 
**emails** | [**Array&lt;Email&gt;**](Email.md) | Array of Contact Email Objects. See below for details. | [optional] 
**phone_numbers** | [**Array&lt;PhoneNumberContact&gt;**](PhoneNumberContact.md) | Array of Contact Phone Number Objects. See below for details. | [optional] 
**addresses** | [**Array&lt;AddressListContacts&gt;**](AddressListContacts.md) | Array of Contact Address Objects. See below for details. | [optional] 
**group** | [**GroupListContacts**](GroupListContacts.md) |  | [optional] 
**created_at** | **Integer** | Integer UNIX timestamp when the contact was created. Read-only. | [optional] 
**updated_at** | **Integer** | Integer UNIX timestamp when the contact was created. Read-only. | [optional] 


