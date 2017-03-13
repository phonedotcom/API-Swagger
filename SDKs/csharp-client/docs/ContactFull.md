# IO.Swagger.Model.ContactFull
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int?** | Integer ID. Read-only. | [optional] 
**Prefix** | **string** | Salutation, such as Mr, Mrs, or Dr | [optional] 
**FirstName** | **string** | First name or given name | [optional] 
**MiddleName** | **string** | Middle or second name | [optional] 
**LastName** | **string** | Last name or surname | [optional] 
**Suffix** | **string** | Suffix, such as \&quot;Jr.\&quot;, \&quot;Sr.\&quot;, \&quot;II\&quot;, or \&quot;III\&quot; | [optional] 
**Nickname** | **string** | Nickname, or a shortened informal version of the contact&#39;s name | [optional] 
**Company** | **string** | Name of the contact&#39;s company | [optional] 
**PhoneticFirstName** | **string** | Phonetic first name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**PhoneticMiddleName** | **string** | Phonetic middle name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**PhoneticLastName** | **string** | Phonetic last name. Useful for remembering how to pronounce the contact&#39;s name. | [optional] 
**Department** | **string** | Name of the contact&#39;s department | [optional] 
**JobTitle** | **string** | Contact&#39;s job title | [optional] 
**Emails** | [**List&lt;Email&gt;**](Email.md) | Array of Contact Email Objects. See below for details. | [optional] 
**PhoneNumbers** | [**List&lt;PhoneNumberContact&gt;**](PhoneNumberContact.md) | Array of Contact Phone Number Objects. See below for details. | [optional] 
**Addresses** | [**List&lt;AddressListContacts&gt;**](AddressListContacts.md) | Array of Contact Address Objects. See below for details. | [optional] 
**Group** | [**GroupListContacts**](GroupListContacts.md) |  | [optional] 
**CreatedAt** | **int?** | Integer UNIX timestamp when the contact was created. Read-only. | [optional] 
**UpdatedAt** | **int?** | Integer UNIX timestamp when the contact was created. Read-only. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

