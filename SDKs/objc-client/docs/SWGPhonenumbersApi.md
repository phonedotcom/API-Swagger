# SWGPhonenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountPhoneNumber**](SWGPhonenumbersApi.md#createaccountphonenumber) | **POST** /accounts/{account_id}/phone-numbers | Add a phone number to an account
[**getAccountPhoneNumber**](SWGPhonenumbersApi.md#getaccountphonenumber) | **GET** /accounts/{account_id}/phone-numbers/{number_id} | Show details of an individual phone number
[**listAccountPhoneNumbers**](SWGPhonenumbersApi.md#listaccountphonenumbers) | **GET** /accounts/{account_id}/phone-numbers | Get a list of phone numbers registered to an account
[**replaceAccountPhoneNumber**](SWGPhonenumbersApi.md#replaceaccountphonenumber) | **PUT** /accounts/{account_id}/phone-numbers/{number_id} | Update the settings for an existing phone number on your account


# **createAccountPhoneNumber**
```objc
-(NSURLSessionTask*) createAccountPhoneNumberWithAccountId: (NSNumber*) accountId
    data: (SWGCreatePhoneNumberParams*) data
        completionHandler: (void (^)(SWGPhoneNumberFull* output, NSError* error)) handler;
```

Add a phone number to an account

See Intro to Account Phone Numbers for more info on the properties to use.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
SWGCreatePhoneNumberParams* data = [[SWGCreatePhoneNumberParams alloc] init]; // Phone Number data (optional)

SWGPhonenumbersApi*apiInstance = [[SWGPhonenumbersApi alloc] init];

// Add a phone number to an account
[apiInstance createAccountPhoneNumberWithAccountId:accountId
              data:data
          completionHandler: ^(SWGPhoneNumberFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGPhonenumbersApi->createAccountPhoneNumber: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **data** | [**SWGCreatePhoneNumberParams***](SWGCreatePhoneNumberParams*.md)| Phone Number data | [optional] 

### Return type

[**SWGPhoneNumberFull***](SWGPhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAccountPhoneNumber**
```objc
-(NSURLSessionTask*) getAccountPhoneNumberWithAccountId: (NSNumber*) accountId
    numberId: (NSNumber*) numberId
        completionHandler: (void (^)(SWGPhoneNumberFull* output, NSError* error)) handler;
```

Show details of an individual phone number

See Intro to Account Phone Numbers for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* numberId = @56; // Number ID

SWGPhonenumbersApi*apiInstance = [[SWGPhonenumbersApi alloc] init];

// Show details of an individual phone number
[apiInstance getAccountPhoneNumberWithAccountId:accountId
              numberId:numberId
          completionHandler: ^(SWGPhoneNumberFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGPhonenumbersApi->getAccountPhoneNumber: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **numberId** | **NSNumber***| Number ID | 

### Return type

[**SWGPhoneNumberFull***](SWGPhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountPhoneNumbers**
```objc
-(NSURLSessionTask*) listAccountPhoneNumbersWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersName: (NSArray<NSString*>*) filtersName
    filtersPhoneNumber: (NSArray<NSString*>*) filtersPhoneNumber
    sortId: (NSString*) sortId
    sortName: (NSString*) sortName
    sortPhoneNumber: (NSString*) sortPhoneNumber
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListPhoneNumbers* output, NSError* error)) handler;
```

Get a list of phone numbers registered to an account

See Intro to Account Phone Numbers for more info on the properties.

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
NSArray<NSString*>* filtersPhoneNumber = @[@"filtersPhoneNumber_example"]; // Phone number filter (optional)
NSString* sortId = @"sortId_example"; // ID sorting (optional)
NSString* sortName = @"sortName_example"; // Name sorting (optional)
NSString* sortPhoneNumber = @"sortPhoneNumber_example"; // Phone number sorting (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGPhonenumbersApi*apiInstance = [[SWGPhonenumbersApi alloc] init];

// Get a list of phone numbers registered to an account
[apiInstance listAccountPhoneNumbersWithAccountId:accountId
              filtersId:filtersId
              filtersName:filtersName
              filtersPhoneNumber:filtersPhoneNumber
              sortId:sortId
              sortName:sortName
              sortPhoneNumber:sortPhoneNumber
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListPhoneNumbers* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGPhonenumbersApi->listAccountPhoneNumbers: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **filtersId** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| ID filter | [optional] 
 **filtersName** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Name filter | [optional] 
 **filtersPhoneNumber** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Phone number filter | [optional] 
 **sortId** | **NSString***| ID sorting | [optional] 
 **sortName** | **NSString***| Name sorting | [optional] 
 **sortPhoneNumber** | **NSString***| Phone number sorting | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 

### Return type

[**SWGListPhoneNumbers***](SWGListPhoneNumbers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replaceAccountPhoneNumber**
```objc
-(NSURLSessionTask*) replaceAccountPhoneNumberWithAccountId: (NSNumber*) accountId
    numberId: (NSNumber*) numberId
    data: (SWGReplacePhoneNumberParams*) data
        completionHandler: (void (^)(SWGPhoneNumberFull* output, NSError* error)) handler;
```

Update the settings for an existing phone number on your account

See Intro to Account Phone Numbers for more info on the properties.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* numberId = @56; // Number ID
SWGReplacePhoneNumberParams* data = [[SWGReplacePhoneNumberParams alloc] init]; // Phone Number data (optional)

SWGPhonenumbersApi*apiInstance = [[SWGPhonenumbersApi alloc] init];

// Update the settings for an existing phone number on your account
[apiInstance replaceAccountPhoneNumberWithAccountId:accountId
              numberId:numberId
              data:data
          completionHandler: ^(SWGPhoneNumberFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGPhonenumbersApi->replaceAccountPhoneNumber: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **numberId** | **NSNumber***| Number ID | 
 **data** | [**SWGReplacePhoneNumberParams***](SWGReplacePhoneNumberParams*.md)| Phone Number data | [optional] 

### Return type

[**SWGPhoneNumberFull***](SWGPhoneNumberFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

