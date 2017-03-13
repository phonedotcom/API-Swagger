# SWGDevicesApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createAccountDevice**](SWGDevicesApi.md#createaccountdevice) | **POST** /accounts/{account_id}/devices | Register a generic VoIP device
[**getAccountDevice**](SWGDevicesApi.md#getaccountdevice) | **GET** /accounts/{account_id}/devices/{device_id} | Show details of an individual VoIP device
[**listAccountDevices**](SWGDevicesApi.md#listaccountdevices) | **GET** /accounts/{account_id}/devices | Get a list of VoIP devices associated with your account
[**replaceAccountDevice**](SWGDevicesApi.md#replaceaccountdevice) | **PUT** /accounts/{account_id}/devices/{device_id} | Update the settings for an individual VoIP device


# **createAccountDevice**
```objc
-(NSURLSessionTask*) createAccountDeviceWithAccountId: (NSNumber*) accountId
    data: (SWGCreateDeviceParams*) data
        completionHandler: (void (^)(SWGDeviceFull* output, NSError* error)) handler;
```

Register a generic VoIP device



### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
SWGCreateDeviceParams* data = [[SWGCreateDeviceParams alloc] init]; // Device data (optional)

SWGDevicesApi*apiInstance = [[SWGDevicesApi alloc] init];

// Register a generic VoIP device
[apiInstance createAccountDeviceWithAccountId:accountId
              data:data
          completionHandler: ^(SWGDeviceFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGDevicesApi->createAccountDevice: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **data** | [**SWGCreateDeviceParams***](SWGCreateDeviceParams*.md)| Device data | [optional] 

### Return type

[**SWGDeviceFull***](SWGDeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getAccountDevice**
```objc
-(NSURLSessionTask*) getAccountDeviceWithAccountId: (NSNumber*) accountId
    deviceId: (NSNumber*) deviceId
        completionHandler: (void (^)(SWGDeviceFull* output, NSError* error)) handler;
```

Show details of an individual VoIP device



### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* deviceId = @56; // Device ID

SWGDevicesApi*apiInstance = [[SWGDevicesApi alloc] init];

// Show details of an individual VoIP device
[apiInstance getAccountDeviceWithAccountId:accountId
              deviceId:deviceId
          completionHandler: ^(SWGDeviceFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGDevicesApi->getAccountDevice: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **deviceId** | **NSNumber***| Device ID | 

### Return type

[**SWGDeviceFull***](SWGDeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **listAccountDevices**
```objc
-(NSURLSessionTask*) listAccountDevicesWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersName: (NSArray<NSString*>*) filtersName
    sortId: (NSString*) sortId
    sortName: (NSString*) sortName
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListDevices* output, NSError* error)) handler;
```

Get a list of VoIP devices associated with your account



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

SWGDevicesApi*apiInstance = [[SWGDevicesApi alloc] init];

// Get a list of VoIP devices associated with your account
[apiInstance listAccountDevicesWithAccountId:accountId
              filtersId:filtersId
              filtersName:filtersName
              sortId:sortId
              sortName:sortName
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListDevices* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGDevicesApi->listAccountDevices: %@", error);
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

[**SWGListDevices***](SWGListDevices.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **replaceAccountDevice**
```objc
-(NSURLSessionTask*) replaceAccountDeviceWithAccountId: (NSNumber*) accountId
    deviceId: (NSNumber*) deviceId
    data: (SWGCreateDeviceParams*) data
        completionHandler: (void (^)(SWGDeviceFull* output, NSError* error)) handler;
```

Update the settings for an individual VoIP device



### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSNumber* accountId = @56; // Account ID
NSNumber* deviceId = @56; // Device ID
SWGCreateDeviceParams* data = [[SWGCreateDeviceParams alloc] init]; // Device data (optional)

SWGDevicesApi*apiInstance = [[SWGDevicesApi alloc] init];

// Update the settings for an individual VoIP device
[apiInstance replaceAccountDeviceWithAccountId:accountId
              deviceId:deviceId
              data:data
          completionHandler: ^(SWGDeviceFull* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGDevicesApi->replaceAccountDevice: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **NSNumber***| Account ID | 
 **deviceId** | **NSNumber***| Device ID | 
 **data** | [**SWGCreateDeviceParams***](SWGCreateDeviceParams*.md)| Device data | [optional] 

### Return type

[**SWGDeviceFull***](SWGDeviceFull.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

