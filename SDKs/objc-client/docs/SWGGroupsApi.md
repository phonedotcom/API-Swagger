# SWGGroupsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountExtensionContactGroup**](SWGGroupsApi.md#createaccountextensioncontactgroup) | **POST** /accounts/{account_id}/extensions/{extension_id}/contact-groups | 
[**deleteAccountExtensionContactGroup**](SWGGroupsApi.md#deleteaccountextensioncontactgroup) | **DELETE** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | Delete an addressbook group
[**getAccountExtensionContactGroup**](SWGGroupsApi.md#getaccountextensioncontactgroup) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 
[**listAccountExtensionContactGroups**](SWGGroupsApi.md#listaccountextensioncontactgroups) | **GET** /accounts/{account_id}/extensions/{extension_id}/contact-groups | Show a list of contact groups belonging to an extension
[**replaceAccountExtensionContactGroup**](SWGGroupsApi.md#replaceaccountextensioncontactgroup) | **PUT** /accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id} | 


# **createAccountExtensionContactGroup**
```objc
-(NSURLSessionTask*) createAccountExtensionContactGroupWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    data: (SWGCreateGroupParams*) data
        completionHandler: (void (^)(SWGGroupFull* output, NSError* error)) handler;
```



See Account Contact Groups for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
SWGCreateGroupParams* data = [[SWGCreateGroupParams alloc] init]; // Group name

SWGGroupsApi*apiInstance = [[SWGGroupsApi alloc] init];

// 
[apiInstance createAccountExtensionContactGroupWithAccountId:accountId
              extensionId:extensionId
              data:data
          completionHandler: ^(SWGGroupFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGGroupsApi->createAccountExtensionContactGroup: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **data** | [**SWGCreateGroupParams***](SWGCreateGroupParams*.md)| Group name | 

### Return type

[**SWGGroupFull***](SWGGroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteAccountExtensionContactGroup**
```objc
-(NSURLSessionTask*) deleteAccountExtensionContactGroupWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    groupId: (NSNumber*) groupId
        completionHandler: (void (^)(SWGDeleteGroup* output, NSError* error)) handler;
```

Delete an addressbook group



### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
NSNumber* groupId = @56; // Group ID

SWGGroupsApi*apiInstance = [[SWGGroupsApi alloc] init];

// Delete an addressbook group
[apiInstance deleteAccountExtensionContactGroupWithAccountId:accountId
              extensionId:extensionId
              groupId:groupId
          completionHandler: ^(SWGDeleteGroup* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGGroupsApi->deleteAccountExtensionContactGroup: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **groupId** | **NSNumber***| Group ID | 

### Return type

[**SWGDeleteGroup***](SWGDeleteGroup.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAccountExtensionContactGroup**
```objc
-(NSURLSessionTask*) getAccountExtensionContactGroupWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    groupId: (NSNumber*) groupId
        completionHandler: (void (^)(SWGGroupFull* output, NSError* error)) handler;
```



See Account Contact Groups for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
NSNumber* groupId = @56; // Group ID

SWGGroupsApi*apiInstance = [[SWGGroupsApi alloc] init];

// 
[apiInstance getAccountExtensionContactGroupWithAccountId:accountId
              extensionId:extensionId
              groupId:groupId
          completionHandler: ^(SWGGroupFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGGroupsApi->getAccountExtensionContactGroup: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **groupId** | **NSNumber***| Group ID | 

### Return type

[**SWGGroupFull***](SWGGroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountExtensionContactGroups**
```objc
-(NSURLSessionTask*) listAccountExtensionContactGroupsWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersName: (NSArray<NSString*>*) filtersName
    sortId: (NSString*) sortId
    sortName: (NSString*) sortName
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListGroups* output, NSError* error)) handler;
```

Show a list of contact groups belonging to an extension

See Account Contact Groups for details on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
NSArray<NSString*>* filtersId = @[@"filtersId_example"]; // ID filter (optional)
NSArray<NSString*>* filtersName = @[@"filtersName_example"]; // Name filter (optional)
NSString* sortId = @"sortId_example"; // ID sorting (optional)
NSString* sortName = @"sortName_example"; // Name sorting (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGGroupsApi*apiInstance = [[SWGGroupsApi alloc] init];

// Show a list of contact groups belonging to an extension
[apiInstance listAccountExtensionContactGroupsWithAccountId:accountId
              extensionId:extensionId
              filtersId:filtersId
              filtersName:filtersName
              sortId:sortId
              sortName:sortName
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListGroups* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGGroupsApi->listAccountExtensionContactGroups: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **filtersId** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| ID filter | [optional] 
 **filtersName** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Name filter | [optional] 
 **sortId** | **NSString***| ID sorting | [optional] 
 **sortName** | **NSString***| Name sorting | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 

### Return type

[**SWGListGroups***](SWGListGroups.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replaceAccountExtensionContactGroup**
```objc
-(NSURLSessionTask*) replaceAccountExtensionContactGroupWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    groupId: (NSNumber*) groupId
    data: (SWGCreateGroupParams*) data
        completionHandler: (void (^)(SWGGroupFull* output, NSError* error)) handler;
```



See Account Contact Groups for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* extensionId = @56; // Extension ID
NSNumber* groupId = @56; // Group ID
SWGCreateGroupParams* data = [[SWGCreateGroupParams alloc] init]; // Group name

SWGGroupsApi*apiInstance = [[SWGGroupsApi alloc] init];

// 
[apiInstance replaceAccountExtensionContactGroupWithAccountId:accountId
              extensionId:extensionId
              groupId:groupId
              data:data
          completionHandler: ^(SWGGroupFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGGroupsApi->replaceAccountExtensionContactGroup: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **extensionId** | **NSNumber***| Extension ID | 
 **groupId** | **NSNumber***| Group ID | 
 **data** | [**SWGCreateGroupParams***](SWGCreateGroupParams*.md)| Group name | 

### Return type

[**SWGGroupFull***](SWGGroupFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

