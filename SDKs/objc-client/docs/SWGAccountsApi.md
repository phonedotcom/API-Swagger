# SWGAccountsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAccount**](SWGAccountsApi.md#getaccount) | **GET** /accounts/{account_id} | Retrieve details of an individual account
[**listAccounts**](SWGAccountsApi.md#listaccounts) | **GET** /accounts | Get a list of accounts visible to the authenticated user or client


# **getAccount**
```objc
-(NSURLSessionTask*) getAccountWithAccountId: (NSNumber*) accountId
        completionHandler: (void (^)(SWGAccountFull* output, NSError* error)) handler;
```

Retrieve details of an individual account

This service shows the details of an individual account. See Accounts for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID

SWGAccountsApi*apiInstance = [[SWGAccountsApi alloc] init];

// Retrieve details of an individual account
[apiInstance getAccountWithAccountId:accountId
          completionHandler: ^(SWGAccountFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGAccountsApi->getAccount: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 

### Return type

[**SWGAccountFull***](SWGAccountFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccounts**
```objc
-(NSURLSessionTask*) listAccountsWithFiltersId: (NSArray<NSString*>*) filtersId
    sortId: (NSString*) sortId
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListAccounts* output, NSError* error)) handler;
```

Get a list of accounts visible to the authenticated user or client

This service lists the accounts accessible to the authenticated client. In most cases, there will only be one such account. See Accounts for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSArray<NSString*>* filtersId = @[@"filtersId_example"]; // ID filter (optional)
NSString* sortId = @"sortId_example"; // ID sorting (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGAccountsApi*apiInstance = [[SWGAccountsApi alloc] init];

// Get a list of accounts visible to the authenticated user or client
[apiInstance listAccountsWithFiltersId:filtersId
              sortId:sortId
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListAccounts* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGAccountsApi->listAccounts: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

