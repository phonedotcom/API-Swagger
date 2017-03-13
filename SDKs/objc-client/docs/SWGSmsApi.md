# SWGSmsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountSms**](SWGSmsApi.md#createaccountsms) | **POST** /accounts/{account_id}/sms | Send a SMS to one or a group of recipients
[**getAccountSms**](SWGSmsApi.md#getaccountsms) | **GET** /accounts/{account_id}/sms/{sms_id} | Show details of an individual SMS
[**listAccountSms**](SWGSmsApi.md#listaccountsms) | **GET** /accounts/{account_id}/sms | Get a list of SMS messages for an account


# **createAccountSms**
```objc
-(NSURLSessionTask*) createAccountSmsWithAccountId: (NSNumber*) accountId
    data: (SWGCreateSmsParams*) data
        completionHandler: (void (^)(SWGSmsFull* output, NSError* error)) handler;
```

Send a SMS to one or a group of recipients

For more on the input fields, see Intro to SMS.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
SWGCreateSmsParams* data = [[SWGCreateSmsParams alloc] init]; // SMS data

SWGSmsApi*apiInstance = [[SWGSmsApi alloc] init];

// Send a SMS to one or a group of recipients
[apiInstance createAccountSmsWithAccountId:accountId
              data:data
          completionHandler: ^(SWGSmsFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGSmsApi->createAccountSms: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **data** | [**SWGCreateSmsParams***](SWGCreateSmsParams*.md)| SMS data | 

### Return type

[**SWGSmsFull***](SWGSmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAccountSms**
```objc
-(NSURLSessionTask*) getAccountSmsWithAccountId: (NSNumber*) accountId
    smsId: (NSString*) smsId
        completionHandler: (void (^)(SWGSmsFull* output, NSError* error)) handler;
```

Show details of an individual SMS

This service shows the details of an individual sms. See Intro to SMS for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSString* smsId = @"smsId_example"; // SMS ID

SWGSmsApi*apiInstance = [[SWGSmsApi alloc] init];

// Show details of an individual SMS
[apiInstance getAccountSmsWithAccountId:accountId
              smsId:smsId
          completionHandler: ^(SWGSmsFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGSmsApi->getAccountSms: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **smsId** | **NSString***| SMS ID | 

### Return type

[**SWGSmsFull***](SWGSmsFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountSms**
```objc
-(NSURLSessionTask*) listAccountSmsWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersDirection: (NSString*) filtersDirection
    filtersFrom: (NSString*) filtersFrom
    sortId: (NSString*) sortId
    sortCreatedAt: (NSString*) sortCreatedAt
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListSms* output, NSError* error)) handler;
```

Get a list of SMS messages for an account

See Intro to SMS for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSArray<NSString*>* filtersId = @[@"filtersId_example"]; // ID filter (optional)
NSString* filtersDirection = @"filtersDirection_example"; // Direction filter (optional)
NSString* filtersFrom = @"filtersFrom_example"; // Caller ID filter (optional)
NSString* sortId = @"sortId_example"; // ID sorting (optional)
NSString* sortCreatedAt = @"sortCreatedAt_example"; // Sort by created time of message (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGSmsApi*apiInstance = [[SWGSmsApi alloc] init];

// Get a list of SMS messages for an account
[apiInstance listAccountSmsWithAccountId:accountId
              filtersId:filtersId
              filtersDirection:filtersDirection
              filtersFrom:filtersFrom
              sortId:sortId
              sortCreatedAt:sortCreatedAt
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListSms* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGSmsApi->listAccountSms: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **filtersId** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| ID filter | [optional] 
 **filtersDirection** | **NSString***| Direction filter | [optional] 
 **filtersFrom** | **NSString***| Caller ID filter | [optional] 
 **sortId** | **NSString***| ID sorting | [optional] 
 **sortCreatedAt** | **NSString***| Sort by created time of message | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 

### Return type

[**SWGListSms***](SWGListSms.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

