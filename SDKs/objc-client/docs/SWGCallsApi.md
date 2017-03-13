# SWGCallsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountCalls**](SWGCallsApi.md#createaccountcalls) | **POST** /accounts/{account_id}/calls | Make a phone call


# **createAccountCalls**
```objc
-(NSURLSessionTask*) createAccountCallsWithAccountId: (NSNumber*) accountId
    data: (SWGCreateCallParams*) data
        completionHandler: (void (^)(SWGCallFull* output, NSError* error)) handler;
```

Make a phone call



### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
SWGCreateCallParams* data = [[SWGCreateCallParams alloc] init]; // Call data (optional)

SWGCallsApi*apiInstance = [[SWGCallsApi alloc] init];

// Make a phone call
[apiInstance createAccountCallsWithAccountId:accountId
              data:data
          completionHandler: ^(SWGCallFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGCallsApi->createAccountCalls: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **data** | [**SWGCreateCallParams***](SWGCreateCallParams*.md)| Call data | [optional] 

### Return type

[**SWGCallFull***](SWGCallFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

