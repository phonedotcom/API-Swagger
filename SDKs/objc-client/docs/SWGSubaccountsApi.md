# SWGSubaccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountSubaccount**](SWGSubaccountsApi.md#createaccountsubaccount) | **POST** /accounts/{account_id}/subaccounts | Add a subaccount for the authenticated user or client
[**listAccountSubaccounts**](SWGSubaccountsApi.md#listaccountsubaccounts) | **GET** /accounts/{account_id}/subaccounts | Get a list of subaccounts for the authenticated user or client


# **createAccountSubaccount**
```objc
-(NSURLSessionTask*) createAccountSubaccountWithAccountId: (NSNumber*) accountId
    data: (SWGCreateSubaccountParams*) data
        completionHandler: (void (^)(SWGAccountFull* output, NSError* error)) handler;
```

Add a subaccount for the authenticated user or client

This service shows the details of an individual Subaccount.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
SWGCreateSubaccountParams* data = [[SWGCreateSubaccountParams alloc] init]; // SMS data

SWGSubaccountsApi*apiInstance = [[SWGSubaccountsApi alloc] init];

// Add a subaccount for the authenticated user or client
[apiInstance createAccountSubaccountWithAccountId:accountId
              data:data
          completionHandler: ^(SWGAccountFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGSubaccountsApi->createAccountSubaccount: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **data** | [**SWGCreateSubaccountParams***](SWGCreateSubaccountParams*.md)| SMS data | 

### Return type

[**SWGAccountFull***](SWGAccountFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountSubaccounts**
```objc
-(NSURLSessionTask*) listAccountSubaccountsWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    sortId: (NSString*) sortId
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListAccounts* output, NSError* error)) handler;
```

Get a list of subaccounts for the authenticated user or client

This service lists the Subaccount of the authenticated client. In most cases, there will not be any.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSArray<NSString*>* filtersId = @[@"filtersId_example"]; // ID filter (optional)
NSString* sortId = @"sortId_example"; // ID sorting (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGSubaccountsApi*apiInstance = [[SWGSubaccountsApi alloc] init];

// Get a list of subaccounts for the authenticated user or client
[apiInstance listAccountSubaccountsWithAccountId:accountId
              filtersId:filtersId
              sortId:sortId
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListAccounts* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGSubaccountsApi->listAccountSubaccounts: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **filtersId** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| ID filter | [optional] 
 **sortId** | **NSString***| ID sorting | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 

### Return type

[**SWGListAccounts***](SWGListAccounts.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

