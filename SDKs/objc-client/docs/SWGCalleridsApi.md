# SWGCalleridsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getCallerIds**](SWGCalleridsApi.md#getcallerids) | **GET** /accounts/{account_id}/extensions/{extension_id}/caller-ids | Show the Caller ID options a given extension can use


# **getCallerIds**
```objc
-(NSURLSessionTask*) getCallerIdsWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    filtersNumber: (NSArray<NSString*>*) filtersNumber
    filtersName: (NSArray<NSString*>*) filtersName
    sortNumber: (NSString*) sortNumber
    sortName: (NSString*) sortName
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListCallerIds* output, NSError* error)) handler;
```

Show the Caller ID options a given extension can use

Get Caller ID

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
NSArray<NSString*>* filtersNumber = @[@"filtersNumber_example"]; // Number filter (optional)
NSArray<NSString*>* filtersName = @[@"filtersName_example"]; // Name filter (optional)
NSString* sortNumber = @"sortNumber_example"; // Number sorting (optional)
NSString* sortName = @"sortName_example"; // Name sorting (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGCalleridsApi*apiInstance = [[SWGCalleridsApi alloc] init];

// Show the Caller ID options a given extension can use
[apiInstance getCallerIdsWithAccountId:accountId
              extensionId:extensionId
              filtersNumber:filtersNumber
              filtersName:filtersName
              sortNumber:sortNumber
              sortName:sortName
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListCallerIds* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGCalleridsApi->getCallerIds: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **filtersNumber** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Number filter | [optional] 
 **filtersName** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Name filter | [optional] 
 **sortNumber** | **NSString***| Number sorting | [optional] 
 **sortName** | **NSString***| Name sorting | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 

### Return type

[**SWGListCallerIds***](SWGListCallerIds.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

