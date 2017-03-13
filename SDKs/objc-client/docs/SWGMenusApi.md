# SWGMenusApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountMenu**](SWGMenusApi.md#createaccountmenu) | **POST** /accounts/{account_id}/menus | Create an individual menu
[**deleteAccountMenu**](SWGMenusApi.md#deleteaccountmenu) | **DELETE** /accounts/{account_id}/menus/{menu_id} | Delete an individual menu
[**getAccountMenu**](SWGMenusApi.md#getaccountmenu) | **GET** /accounts/{account_id}/menus/{menu_id} | Show details of an individual menu
[**listAccountMenus**](SWGMenusApi.md#listaccountmenus) | **GET** /accounts/{account_id}/menus | Get a list of menus for an account
[**replaceAccountMenu**](SWGMenusApi.md#replaceaccountmenu) | **PUT** /accounts/{account_id}/menus/{menu_id} | Replace an individual menu


# **createAccountMenu**
```objc
-(NSURLSessionTask*) createAccountMenuWithAccountId: (NSNumber*) accountId
    data: (SWGCreateMenuParams*) data
        completionHandler: (void (^)(SWGMenuFull* output, NSError* error)) handler;
```

Create an individual menu

This service creates an individual menu. See Account Menus for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
SWGCreateMenuParams* data = [[SWGCreateMenuParams alloc] init]; // Menu data (optional)

SWGMenusApi*apiInstance = [[SWGMenusApi alloc] init];

// Create an individual menu
[apiInstance createAccountMenuWithAccountId:accountId
              data:data
          completionHandler: ^(SWGMenuFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGMenusApi->createAccountMenu: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **data** | [**SWGCreateMenuParams***](SWGCreateMenuParams*.md)| Menu data | [optional] 

### Return type

[**SWGMenuFull***](SWGMenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteAccountMenu**
```objc
-(NSURLSessionTask*) deleteAccountMenuWithAccountId: (NSNumber*) accountId
    menuId: (NSNumber*) menuId
        completionHandler: (void (^)(SWGDeleteMenu* output, NSError* error)) handler;
```

Delete an individual menu

This service shows the details of an individual menu.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* menuId = @56; // Menu ID

SWGMenusApi*apiInstance = [[SWGMenusApi alloc] init];

// Delete an individual menu
[apiInstance deleteAccountMenuWithAccountId:accountId
              menuId:menuId
          completionHandler: ^(SWGDeleteMenu* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGMenusApi->deleteAccountMenu: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **menuId** | **NSNumber***| Menu ID | 

### Return type

[**SWGDeleteMenu***](SWGDeleteMenu.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAccountMenu**
```objc
-(NSURLSessionTask*) getAccountMenuWithAccountId: (NSNumber*) accountId
    menuId: (NSNumber*) menuId
        completionHandler: (void (^)(SWGMenuFull* output, NSError* error)) handler;
```

Show details of an individual menu

This service shows the details of an individual Menu.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* menuId = @56; // Menu ID

SWGMenusApi*apiInstance = [[SWGMenusApi alloc] init];

// Show details of an individual menu
[apiInstance getAccountMenuWithAccountId:accountId
              menuId:menuId
          completionHandler: ^(SWGMenuFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGMenusApi->getAccountMenu: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **menuId** | **NSNumber***| Menu ID | 

### Return type

[**SWGMenuFull***](SWGMenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountMenus**
```objc
-(NSURLSessionTask*) listAccountMenusWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersName: (NSArray<NSString*>*) filtersName
    sortId: (NSString*) sortId
    sortName: (NSString*) sortName
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListMenus* output, NSError* error)) handler;
```

Get a list of menus for an account

See Account Menus for more info on the properties.

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

SWGMenusApi*apiInstance = [[SWGMenusApi alloc] init];

// Get a list of menus for an account
[apiInstance listAccountMenusWithAccountId:accountId
              filtersId:filtersId
              filtersName:filtersName
              sortId:sortId
              sortName:sortName
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListMenus* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGMenusApi->listAccountMenus: %@", error);
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

[**SWGListMenus***](SWGListMenus.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replaceAccountMenu**
```objc
-(NSURLSessionTask*) replaceAccountMenuWithAccountId: (NSNumber*) accountId
    menuId: (NSNumber*) menuId
    data: (SWGReplaceMenuParams*) data
        completionHandler: (void (^)(SWGMenuFull* output, NSError* error)) handler;
```

Replace an individual menu

This service replaces the details of an individual Menu.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* menuId = @56; // Menu ID
SWGReplaceMenuParams* data = [[SWGReplaceMenuParams alloc] init]; // Menu data (optional)

SWGMenusApi*apiInstance = [[SWGMenusApi alloc] init];

// Replace an individual menu
[apiInstance replaceAccountMenuWithAccountId:accountId
              menuId:menuId
              data:data
          completionHandler: ^(SWGMenuFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGMenusApi->replaceAccountMenu: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **menuId** | **NSNumber***| Menu ID | 
 **data** | [**SWGReplaceMenuParams***](SWGReplaceMenuParams*.md)| Menu data | [optional] 

### Return type

[**SWGMenuFull***](SWGMenuFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

