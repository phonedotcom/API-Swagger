# SWGExpressservicecodesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAccountExpressSrvCode**](SWGExpressservicecodesApi.md#getaccountexpresssrvcode) | **GET** /accounts/{account_id}/express-service-codes/{code_id} | Show details of an account Express Service Code
[**listAccountExpressSrvCodes**](SWGExpressservicecodesApi.md#listaccountexpresssrvcodes) | **GET** /accounts/{account_id}/express-service-codes | Get the Express Service Code associated with your account in list format


# **getAccountExpressSrvCode**
```objc
-(NSURLSessionTask*) getAccountExpressSrvCodeWithAccountId: (NSNumber*) accountId
    codeId: (NSNumber*) codeId
        completionHandler: (void (^)(SWGExpressServiceCodeFull* output, NSError* error)) handler;
```

Show details of an account Express Service Code

This service shows the details of an Account Express Service Code.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* codeId = @56; // Device ID

SWGExpressservicecodesApi*apiInstance = [[SWGExpressservicecodesApi alloc] init];

// Show details of an account Express Service Code
[apiInstance getAccountExpressSrvCodeWithAccountId:accountId
              codeId:codeId
          completionHandler: ^(SWGExpressServiceCodeFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGExpressservicecodesApi->getAccountExpressSrvCode: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **codeId** | **NSNumber***| Device ID | 

### Return type

[**SWGExpressServiceCodeFull***](SWGExpressServiceCodeFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountExpressSrvCodes**
```objc
-(NSURLSessionTask*) listAccountExpressSrvCodesWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
        completionHandler: (void (^)(SWGListExpressServiceCodes* output, NSError* error)) handler;
```

Get the Express Service Code associated with your account in list format

See Express Service Codes for more detail.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSArray<NSString*>* filtersId = @[@"filtersId_example"]; // ID filter (optional)

SWGExpressservicecodesApi*apiInstance = [[SWGExpressservicecodesApi alloc] init];

// Get the Express Service Code associated with your account in list format
[apiInstance listAccountExpressSrvCodesWithAccountId:accountId
              filtersId:filtersId
          completionHandler: ^(SWGListExpressServiceCodes* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGExpressservicecodesApi->listAccountExpressSrvCodes: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **filtersId** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| ID filter | [optional] 

### Return type

[**SWGListExpressServiceCodes***](SWGListExpressServiceCodes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

