# IO.Swagger.Model.AccountFull
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int?** | Account ID. Sometimes referred to as \&quot;Voip ID\&quot; or \&quot;voip_id\&quot;. | [optional] 
**Name** | **string** | Name on the account. Read-only. | [optional] 
**Username** | **string** | Account user name | [optional] 
**Password** | **string** | Account password. For security reason, this is masked as \&quot;**\&quot; | [optional] 
**MasterAccount** | [**AccountSummary**](AccountSummary.md) | If this account is a subaccount, this property shows the master account that it belongs to. Otherwise it is NULL. Read-only. Output is an Account Summary Object. | [optional] 
**Contact** | [**ContactAccount**](ContactAccount.md) | Contact Object. See below for details. | [optional] 
**BillingContact** | [**ContactAccount**](ContactAccount.md) | Contact Object for billing purposes. See below for details. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

