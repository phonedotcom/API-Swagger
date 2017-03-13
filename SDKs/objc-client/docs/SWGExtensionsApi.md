# SWGExtensionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountExtension**](SWGExtensionsApi.md#createaccountextension) | **POST** /accounts/{account_id}/extensions | Create an individual extension
[**getAccountExtension**](SWGExtensionsApi.md#getaccountextension) | **GET** /accounts/{account_id}/extensions/{extension_id} | Show details of an individual extension
[**listAccountExtensions**](SWGExtensionsApi.md#listaccountextensions) | **GET** /accounts/{account_id}/extensions | Get a list of extensions visible to the authenticated user or client
[**replaceAccountExtension**](SWGExtensionsApi.md#replaceaccountextension) | **PUT** /accounts/{account_id}/extensions/{extension_id} | Replace an individual extension


# **createAccountExtension**
```objc
-(NSURLSessionTask*) createAccountExtensionWithAccountId: (NSNumber*) accountId
    data: (SWGCreateExtensionParams*) data
        completionHandler: (void (^)(SWGExtensionFull* output, NSError* error)) handler;
```

Create an individual extension

This service shows how to create a virtual extension.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
SWGCreateExtensionParams* data = [[SWGCreateExtensionParams alloc] init]; // Account Extensions Data (optional)

SWGExtensionsApi*apiInstance = [[SWGExtensionsApi alloc] init];

// Create an individual extension
[apiInstance createAccountExtensionWithAccountId:accountId
              data:data
          completionHandler: ^(SWGExtensionFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGExtensionsApi->createAccountExtension: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **data** | [**SWGCreateExtensionParams***](SWGCreateExtensionParams*.md)| Account Extensions Data | [optional] 

### Return type

[**SWGExtensionFull***](SWGExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAccountExtension**
```objc
-(NSURLSessionTask*) getAccountExtensionWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
        completionHandler: (void (^)(SWGExtensionFull* output, NSError* error)) handler;
```

Show details of an individual extension

This service shows the details of an individual Extension.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID

SWGExtensionsApi*apiInstance = [[SWGExtensionsApi alloc] init];

// Show details of an individual extension
[apiInstance getAccountExtensionWithAccountId:accountId
              extensionId:extensionId
          completionHandler: ^(SWGExtensionFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGExtensionsApi->getAccountExtension: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 

### Return type

[**SWGExtensionFull***](SWGExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountExtensions**
```objc
-(NSURLSessionTask*) listAccountExtensionsWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersExtension: (NSArray<NSString*>*) filtersExtension
    filtersName: (NSArray<NSString*>*) filtersName
    sortId: (NSString*) sortId
    sortExtension: (NSString*) sortExtension
    sortName: (NSString*) sortName
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListExtensions* output, NSError* error)) handler;
```

Get a list of extensions visible to the authenticated user or client

This service lists the visible extensions on a given account.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSArray<NSString*>* filtersId = @[@"filtersId_example"]; // ID filter (optional)
NSArray<NSString*>* filtersExtension = @[@"filtersExtension_example"]; // Extension filter (optional)
NSArray<NSString*>* filtersName = @[@"filtersName_example"]; // Name filter (optional)
NSString* sortId = @"sortId_example"; // ID sorting (optional)
NSString* sortExtension = @"sortExtension_example"; // Extension sorting (optional)
NSString* sortName = @"sortName_example"; // Name sorting (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGExtensionsApi*apiInstance = [[SWGExtensionsApi alloc] init];

// Get a list of extensions visible to the authenticated user or client
[apiInstance listAccountExtensionsWithAccountId:accountId
              filtersId:filtersId
              filtersExtension:filtersExtension
              filtersName:filtersName
              sortId:sortId
              sortExtension:sortExtension
              sortName:sortName
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListExtensions* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGExtensionsApi->listAccountExtensions: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **filtersId** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| ID filter | [optional] 
 **filtersExtension** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Extension filter | [optional] 
 **filtersName** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Name filter | [optional] 
 **sortId** | **NSString***| ID sorting | [optional] 
 **sortExtension** | **NSString***| Extension sorting | [optional] 
 **sortName** | **NSString***| Name sorting | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 

### Return type

[**SWGListExtensions***](SWGListExtensions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replaceAccountExtension**
```objc
-(NSURLSessionTask*) replaceAccountExtensionWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    data: (SWGReplaceExtensionParams*) data
        completionHandler: (void (^)(SWGExtensionFull* output, NSError* error)) handler;
```

Replace an individual extension

This service shows how to update an individual extension.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
SWGReplaceExtensionParams* data = [[SWGReplaceExtensionParams alloc] init]; // Account Extensions Data (optional)

SWGExtensionsApi*apiInstance = [[SWGExtensionsApi alloc] init];

// Replace an individual extension
[apiInstance replaceAccountExtensionWithAccountId:accountId
              extensionId:extensionId
              data:data
          completionHandler: ^(SWGExtensionFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGExtensionsApi->replaceAccountExtension: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **data** | [**SWGReplaceExtensionParams***](SWGReplaceExtensionParams*.md)| Account Extensions Data | [optional] 

### Return type

[**SWGExtensionFull***](SWGExtensionFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

