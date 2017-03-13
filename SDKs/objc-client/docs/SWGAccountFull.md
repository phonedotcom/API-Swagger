# SWGAccountFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**_id** | **NSNumber*** | Account ID. Sometimes referred to as \&quot;Voip ID\&quot; or \&quot;voip_id\&quot;. | [optional] 
**name** | **NSString*** | Name on the account. Read-only. | [optional] 
**username** | **NSString*** | Account user name | [optional] 
**password** | **NSString*** | Account password. For security reason, this is masked as \&quot;**\&quot; | [optional] 
**masterAccount** | [**SWGAccountSummary***](SWGAccountSummary.md) | If this account is a subaccount, this property shows the master account that it belongs to. Otherwise it is NULL. Read-only. Output is an Account Summary Object. | [optional] 
**contact** | [**SWGContactAccount***](SWGContactAccount.md) | Contact Object. See below for details. | [optional] 
**billingContact** | [**SWGContactAccount***](SWGContactAccount.md) | Contact Object for billing purposes. See below for details. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


