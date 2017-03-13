# SWGContactFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**_id** | **NSNumber*** | Integer ID. Read-only. | [optional] 
**prefix** | **NSString*** | Salutation, such as Mr, Mrs, or Dr | [optional] 
**firstName** | **NSString*** | First name or given name | [optional] 
**middleName** | **NSString*** | Middle or second name | [optional] 
**lastName** | **NSString*** | Last name or surname | [optional] 
**suffix** | **NSString*** | Suffix, such as \&quot;Jr.\&quot;, \&quot;Sr.\&quot;, \&quot;II\&quot;, or \&quot;III\&quot; | [optional] 
**nickname** | **NSString*** | Nickname, or a shortened informal version of the contact&#39;s name | [optional] 
**company** | **NSString*** | Name of the contact&#39;s company | [optional] 
**phoneticFirstName** | **NSString*** | Phonetic first name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**phoneticMiddleName** | **NSString*** | Phonetic middle name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**phoneticLastName** | **NSString*** | Phonetic last name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**department** | **NSString*** | Name of the contact&#39;s department | [optional] 
**jobTitle** | **NSString*** | Contact&#39;s job title | [optional] 
**emails** | [**NSArray&lt;SWGEmail&gt;***](SWGEmail.md) | Array of Contact Email Objects. See below for details. | [optional] 
**phoneNumbers** | [**NSArray&lt;SWGPhoneNumberContact&gt;***](SWGPhoneNumberContact.md) | Array of Contact Phone Number Objects. See below for details. | [optional] 
**addresses** | [**NSArray&lt;SWGAddressListContacts&gt;***](SWGAddressListContacts.md) | Array of Contact Address Objects. See below for details. | [optional] 
**group** | [**SWGGroupListContacts***](SWGGroupListContacts.md) |  | [optional] 
**createdAt** | **NSNumber*** | Integer UNIX timestamp when the contact was created. Read-only. | [optional] 
**updatedAt** | **NSNumber*** | Integer UNIX timestamp when the contact was created. Read-only. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


