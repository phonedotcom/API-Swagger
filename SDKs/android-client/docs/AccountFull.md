
# AccountFull

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Integer** | Account ID. Sometimes referred to as \&quot;Voip ID\&quot; or \&quot;voip_id\&quot;. |  [optional]
**name** | **String** | Name on the account. Read-only. |  [optional]
**username** | **String** | Account user name |  [optional]
**password** | **String** | Account password. For security reason, this is masked as \&quot;**\&quot; |  [optional]
**masterAccount** | [**AccountSummary**](AccountSummary.md) | If this account is a subaccount, this property shows the master account that it belongs to. Otherwise it is NULL. Read-only. Output is an Account Summary Object. |  [optional]
**contact** | [**ContactAccount**](ContactAccount.md) | Contact Object. See below for details. |  [optional]
**billingContact** | [**ContactAccount**](ContactAccount.md) | Contact Object for billing purposes. See below for details. |  [optional]



