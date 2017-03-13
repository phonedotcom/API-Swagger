# SWGTrunksApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountTrunk**](SWGTrunksApi.md#createaccounttrunk) | **POST** /accounts/{account_id}/trunks | Add a trunk record with SIP information
[**deleteAccountTrunk**](SWGTrunksApi.md#deleteaccounttrunk) | **DELETE** /accounts/{account_id}/trunks/{trunk_id} | Delete a trunk from account
[**getAccountTrunk**](SWGTrunksApi.md#getaccounttrunk) | **GET** /accounts/{account_id}/trunks/{trunk_id} | Show details of an individual trunk
[**listAccountTrunks**](SWGTrunksApi.md#listaccounttrunks) | **GET** /accounts/{account_id}/trunks | Get a list of trunks for an account
[**replaceAccountTrunk**](SWGTrunksApi.md#replaceaccounttrunk) | **PUT** /accounts/{account_id}/trunks/{trunk_id} | Replace parameters in a trunk


# **createAccountTrunk**
```objc
-(NSURLSessionTask*) createAccountTrunkWithAccountId: (NSNumber*) accountId
    data: (SWGCreateTrunkParams*) data
        completionHandler: (void (^)(SWGTrunkFull* output, NSError* error)) handler;
```

Add a trunk record with SIP information

For more on the input fields, see Account Trunks.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
SWGCreateTrunkParams* data = [[SWGCreateTrunkParams alloc] init]; // Trunk data

SWGTrunksApi*apiInstance = [[SWGTrunksApi alloc] init];

// Add a trunk record with SIP information
[apiInstance createAccountTrunkWithAccountId:accountId
              data:data
          completionHandler: ^(SWGTrunkFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGTrunksApi->createAccountTrunk: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **data** | [**SWGCreateTrunkParams***](SWGCreateTrunkParams*.md)| Trunk data | 

### Return type

[**SWGTrunkFull***](SWGTrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteAccountTrunk**
```objc
-(NSURLSessionTask*) deleteAccountTrunkWithAccountId: (NSNumber*) accountId
    trunkId: (NSNumber*) trunkId
        completionHandler: (void (^)(SWGDeleteTrunk* output, NSError* error)) handler;
```

Delete a trunk from account

This service deletes a trunk from the account. For more on the properties of trunks, see Account Trunks.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* trunkId = @56; // Trunk ID

SWGTrunksApi*apiInstance = [[SWGTrunksApi alloc] init];

// Delete a trunk from account
[apiInstance deleteAccountTrunkWithAccountId:accountId
              trunkId:trunkId
          completionHandler: ^(SWGDeleteTrunk* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGTrunksApi->deleteAccountTrunk: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **trunkId** | **NSNumber***| Trunk ID | 

### Return type

[**SWGDeleteTrunk***](SWGDeleteTrunk.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAccountTrunk**
```objc
-(NSURLSessionTask*) getAccountTrunkWithAccountId: (NSNumber*) accountId
    trunkId: (NSNumber*) trunkId
        completionHandler: (void (^)(SWGTrunkFull* output, NSError* error)) handler;
```

Show details of an individual trunk

This service shows the details of an individual Trunk.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* trunkId = @56; // Trunk ID

SWGTrunksApi*apiInstance = [[SWGTrunksApi alloc] init];

// Show details of an individual trunk
[apiInstance getAccountTrunkWithAccountId:accountId
              trunkId:trunkId
          completionHandler: ^(SWGTrunkFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGTrunksApi->getAccountTrunk: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **trunkId** | **NSNumber***| Trunk ID | 

### Return type

[**SWGTrunkFull***](SWGTrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountTrunks**
```objc
-(NSURLSessionTask*) listAccountTrunksWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersName: (NSArray<NSString*>*) filtersName
    sortId: (NSString*) sortId
    sortName: (NSString*) sortName
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListTrunks* output, NSError* error)) handler;
```

Get a list of trunks for an account

See Account Trunks for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSArray<NSString*>* filtersId = @[@"filtersId_example"]; // ID filter (optional)
NSArray<NSString*>* filtersName = @[@"filtersName_example"]; // Name filter (optional)
NSString* sortId = @"sortId_example"; // ID sorting (optional)
NSString* sortName = @"sortName_example"; // Name sorting (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGTrunksApi*apiInstance = [[SWGTrunksApi alloc] init];

// Get a list of trunks for an account
[apiInstance listAccountTrunksWithAccountId:accountId
              filtersId:filtersId
              filtersName:filtersName
              sortId:sortId
              sortName:sortName
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListTrunks* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGTrunksApi->listAccountTrunks: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **filtersId** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| ID filter | [optional] 
 **filtersName** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Name filter | [optional] 
 **sortId** | **NSString***| ID sorting | [optional] 
 **sortName** | **NSString***| Name sorting | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 

### Return type

[**SWGListTrunks***](SWGListTrunks.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replaceAccountTrunk**
```objc
-(NSURLSessionTask*) replaceAccountTrunkWithAccountId: (NSNumber*) accountId
    trunkId: (NSNumber*) trunkId
    data: (SWGCreateTrunkParams*) data
        completionHandler: (void (^)(SWGTrunkFull* output, NSError* error)) handler;
```

Replace parameters in a trunk

For more on the input fields, see Account Trunks.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* trunkId = @56; // Trunk ID
SWGCreateTrunkParams* data = [[SWGCreateTrunkParams alloc] init]; // Trunk data

SWGTrunksApi*apiInstance = [[SWGTrunksApi alloc] init];

// Replace parameters in a trunk
[apiInstance replaceAccountTrunkWithAccountId:accountId
              trunkId:trunkId
              data:data
          completionHandler: ^(SWGTrunkFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGTrunksApi->replaceAccountTrunk: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **trunkId** | **NSNumber***| Trunk ID | 
 **data** | [**SWGCreateTrunkParams***](SWGCreateTrunkParams*.md)| Trunk data | 

### Return type

[**SWGTrunkFull***](SWGTrunkFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

