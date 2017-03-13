# SWGQueuesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountQueue**](SWGQueuesApi.md#createaccountqueue) | **POST** /accounts/{account_id}/queues | Create a queue
[**deleteAccountQueue**](SWGQueuesApi.md#deleteaccountqueue) | **DELETE** /accounts/{account_id}/queues/{queue_id} | Delete a queue
[**getAccountQueue**](SWGQueuesApi.md#getaccountqueue) | **GET** /accounts/{account_id}/queues/{queue_id} | Show details of an individual queue
[**listAccountQueues**](SWGQueuesApi.md#listaccountqueues) | **GET** /accounts/{account_id}/queues | Get a list of queues for an account
[**replaceAccountQueue**](SWGQueuesApi.md#replaceaccountqueue) | **PUT** /accounts/{account_id}/queues/{queue_id} | Replace a queue


# **createAccountQueue**
```objc
-(NSURLSessionTask*) createAccountQueueWithAccountId: (NSNumber*) accountId
    data: (SWGCreateQueueParams*) data
        completionHandler: (void (^)(SWGQueueFull* output, NSError* error)) handler;
```

Create a queue

For more on the input fields, see Account Queues.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
SWGCreateQueueParams* data = [[SWGCreateQueueParams alloc] init]; // Queue data (optional)

SWGQueuesApi*apiInstance = [[SWGQueuesApi alloc] init];

// Create a queue
[apiInstance createAccountQueueWithAccountId:accountId
              data:data
          completionHandler: ^(SWGQueueFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGQueuesApi->createAccountQueue: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **data** | [**SWGCreateQueueParams***](SWGCreateQueueParams*.md)| Queue data | [optional] 

### Return type

[**SWGQueueFull***](SWGQueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteAccountQueue**
```objc
-(NSURLSessionTask*) deleteAccountQueueWithAccountId: (NSNumber*) accountId
    queueId: (NSNumber*) queueId
        completionHandler: (void (^)(SWGDeleteQueue* output, NSError* error)) handler;
```

Delete a queue

This service a queue from the account. For more information on queue properties, see Account Queues.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* queueId = @56; // Queue ID

SWGQueuesApi*apiInstance = [[SWGQueuesApi alloc] init];

// Delete a queue
[apiInstance deleteAccountQueueWithAccountId:accountId
              queueId:queueId
          completionHandler: ^(SWGDeleteQueue* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGQueuesApi->deleteAccountQueue: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **queueId** | **NSNumber***| Queue ID | 

### Return type

[**SWGDeleteQueue***](SWGDeleteQueue.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAccountQueue**
```objc
-(NSURLSessionTask*) getAccountQueueWithAccountId: (NSNumber*) accountId
    queueId: (NSNumber*) queueId
        completionHandler: (void (^)(SWGQueueFull* output, NSError* error)) handler;
```

Show details of an individual queue

This service shows the details of an individual queue. For more on the input fields, see Account Queues.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* queueId = @56; // Queue ID

SWGQueuesApi*apiInstance = [[SWGQueuesApi alloc] init];

// Show details of an individual queue
[apiInstance getAccountQueueWithAccountId:accountId
              queueId:queueId
          completionHandler: ^(SWGQueueFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGQueuesApi->getAccountQueue: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **queueId** | **NSNumber***| Queue ID | 

### Return type

[**SWGQueueFull***](SWGQueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountQueues**
```objc
-(NSURLSessionTask*) listAccountQueuesWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersName: (NSArray<NSString*>*) filtersName
    sortId: (NSString*) sortId
    sortName: (NSString*) sortName
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListQueues* output, NSError* error)) handler;
```

Get a list of queues for an account

The List Queues service lists all the queues belong to the account. See Account Queues for more info on the properties.

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

SWGQueuesApi*apiInstance = [[SWGQueuesApi alloc] init];

// Get a list of queues for an account
[apiInstance listAccountQueuesWithAccountId:accountId
              filtersId:filtersId
              filtersName:filtersName
              sortId:sortId
              sortName:sortName
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListQueues* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGQueuesApi->listAccountQueues: %@", error);
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

[**SWGListQueues***](SWGListQueues.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replaceAccountQueue**
```objc
-(NSURLSessionTask*) replaceAccountQueueWithAccountId: (NSNumber*) accountId
    queueId: (NSNumber*) queueId
    data: (SWGCreateQueueParams*) data
        completionHandler: (void (^)(SWGQueueFull* output, NSError* error)) handler;
```

Replace a queue

The Replace Queue service replaces the parameters of a queue. For more on the input fields, see Account Queues.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* queueId = @56; // Queue ID
SWGCreateQueueParams* data = [[SWGCreateQueueParams alloc] init]; // Queue data (optional)

SWGQueuesApi*apiInstance = [[SWGQueuesApi alloc] init];

// Replace a queue
[apiInstance replaceAccountQueueWithAccountId:accountId
              queueId:queueId
              data:data
          completionHandler: ^(SWGQueueFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGQueuesApi->replaceAccountQueue: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **queueId** | **NSNumber***| Queue ID | 
 **data** | [**SWGCreateQueueParams***](SWGCreateQueueParams*.md)| Queue data | [optional] 

### Return type

[**SWGQueueFull***](SWGQueueFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

