
# ContactFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Integer** | Integer ID. Read-only. |  [optional]
**prefix** | **String** | Salutation, such as Mr, Mrs, or Dr |  [optional]
**firstName** | **String** | First name or given name |  [optional]
**middleName** | **String** | Middle or second name |  [optional]
**lastName** | **String** | Last name or surname |  [optional]
**suffix** | **String** | Suffix, such as \&quot;Jr.\&quot;, \&quot;Sr.\&quot;, \&quot;II\&quot;, or \&quot;III\&quot; |  [optional]
**nickname** | **String** | Nickname, or a shortened informal version of the contact&#39;s name |  [optional]
**company** | **String** | Name of the contact&#39;s company |  [optional]
**phoneticFirstName** | **String** | Phonetic first name. Useful for remembering how to pronounce the contact&#39;s name. |  [optional]
**phoneticMiddleName** | **String** | Phonetic middle name. Useful for remembering how to pronounce the contact&#39;s name. |  [optional]
**phoneticLastName** | **String** | Phonetic last name. Useful for remembering how to pronounce the contact&#39;s name. |  [optional]
**department** | **String** | Name of the contact&#39;s department |  [optional]
**jobTitle** | **String** | Contact&#39;s job title |  [optional]
**emails** | [**List&lt;Email&gt;**](Email.md) | Array of Contact Email Objects. See below for details. |  [optional]
**phoneNumbers** | [**List&lt;PhoneNumberContact&gt;**](PhoneNumberContact.md) | Array of Contact Phone Number Objects. See below for details. |  [optional]
**addresses** | [**List&lt;AddressListContacts&gt;**](AddressListContacts.md) | Array of Contact Address Objects. See below for details. |  [optional]
**group** | [**GroupListContacts**](GroupListContacts.md) |  |  [optional]
**createdAt** | **Integer** | Integer UNIX timestamp when the contact was created. Read-only. |  [optional]
**updatedAt** | **Integer** | Integer UNIX timestamp when the contact was created. Read-only. |  [optional]



