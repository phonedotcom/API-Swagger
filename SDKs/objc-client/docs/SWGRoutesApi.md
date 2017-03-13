# SWGRoutesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createRoute**](SWGRoutesApi.md#createroute) | **POST** /accounts/{account_id}/routes | Add a new address book contact for an extension
[**deleteAccountRoute**](SWGRoutesApi.md#deleteaccountroute) | **DELETE** /accounts/{account_id}/routes/{route_id} | 
[**getAccountRoute**](SWGRoutesApi.md#getaccountroute) | **GET** /accounts/{account_id}/routes/{route_id} | Show details of an individual route
[**listAccountRoutes**](SWGRoutesApi.md#listaccountroutes) | **GET** /accounts/{account_id}/routes | Get a list of routes for an account
[**replaceAccountRoute**](SWGRoutesApi.md#replaceaccountroute) | **PUT** /accounts/{account_id}/routes/{route_id} | 


# **createRoute**
```objc
-(NSURLSessionTask*) createRouteWithAccountId: (NSNumber*) accountId
    data: (SWGCreateRouteParams*) data
        completionHandler: (void (^)(SWGRouteFull* output, NSError* error)) handler;
```

Add a new address book contact for an extension

For more on the input fields, see Intro to Routes.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
SWGCreateRouteParams* data = [[SWGCreateRouteParams alloc] init]; // Route data (optional)

SWGRoutesApi*apiInstance = [[SWGRoutesApi alloc] init];

// Add a new address book contact for an extension
[apiInstance createRouteWithAccountId:accountId
              data:data
          completionHandler: ^(SWGRouteFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGRoutesApi->createRoute: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **data** | [**SWGCreateRouteParams***](SWGCreateRouteParams*.md)| Route data | [optional] 

### Return type

[**SWGRouteFull***](SWGRouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteAccountRoute**
```objc
-(NSURLSessionTask*) deleteAccountRouteWithAccountId: (NSNumber*) accountId
    routeId: (NSNumber*) routeId
        completionHandler: (void (^)(SWGDeleteRoute* output, NSError* error)) handler;
```





### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* routeId = @56; // Route ID

SWGRoutesApi*apiInstance = [[SWGRoutesApi alloc] init];

// 
[apiInstance deleteAccountRouteWithAccountId:accountId
              routeId:routeId
          completionHandler: ^(SWGDeleteRoute* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGRoutesApi->deleteAccountRoute: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **routeId** | **NSNumber***| Route ID | 

### Return type

[**SWGDeleteRoute***](SWGDeleteRoute.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAccountRoute**
```objc
-(NSURLSessionTask*) getAccountRouteWithAccountId: (NSNumber*) accountId
    routeId: (NSNumber*) routeId
        completionHandler: (void (^)(SWGRouteFull* output, NSError* error)) handler;
```

Show details of an individual route

This service shows the details of an individual route.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* routeId = @56; // Route ID

SWGRoutesApi*apiInstance = [[SWGRoutesApi alloc] init];

// Show details of an individual route
[apiInstance getAccountRouteWithAccountId:accountId
              routeId:routeId
          completionHandler: ^(SWGRouteFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGRoutesApi->getAccountRoute: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **routeId** | **NSNumber***| Route ID | 

### Return type

[**SWGRouteFull***](SWGRouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountRoutes**
```objc
-(NSURLSessionTask*) listAccountRoutesWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersName: (NSArray<NSString*>*) filtersName
    sortId: (NSString*) sortId
    sortName: (NSString*) sortName
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListRoutes* output, NSError* error)) handler;
```

Get a list of routes for an account

See Intro to Routes for more info on the properties.

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

SWGRoutesApi*apiInstance = [[SWGRoutesApi alloc] init];

// Get a list of routes for an account
[apiInstance listAccountRoutesWithAccountId:accountId
              filtersId:filtersId
              filtersName:filtersName
              sortId:sortId
              sortName:sortName
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListRoutes* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGRoutesApi->listAccountRoutes: %@", error);
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

[**SWGListRoutes***](SWGListRoutes.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replaceAccountRoute**
```objc
-(NSURLSessionTask*) replaceAccountRouteWithAccountId: (NSNumber*) accountId
    routeId: (NSNumber*) routeId
    data: (SWGCreateRouteParams*) data
        completionHandler: (void (^)(SWGRouteFull* output, NSError* error)) handler;
```



For more on the input fields, see Intro to Routes.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* routeId = @56; // Route ID
SWGCreateRouteParams* data = [[SWGCreateRouteParams alloc] init]; // Route data (optional)

SWGRoutesApi*apiInstance = [[SWGRoutesApi alloc] init];

// 
[apiInstance replaceAccountRouteWithAccountId:accountId
              routeId:routeId
              data:data
          completionHandler: ^(SWGRouteFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGRoutesApi->replaceAccountRoute: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **routeId** | **NSNumber***| Route ID | 
 **data** | [**SWGCreateRouteParams***](SWGCreateRouteParams*.md)| Route data | [optional] 

### Return type

[**SWGRouteFull***](SWGRouteFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

