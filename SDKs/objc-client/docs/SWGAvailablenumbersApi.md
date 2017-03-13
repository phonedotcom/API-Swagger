# SWGAvailablenumbersApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listAvailablePhoneNumbers**](SWGAvailablenumbersApi.md#listavailablephonenumbers) | **GET** /phone-numbers/available | 


# **listAvailablePhoneNumbers**
```objc
-(NSURLSessionTask*) listAvailablePhoneNumbersWithFiltersPhoneNumber: (NSArray<NSString*>*) filtersPhoneNumber
    filtersCountryCode: (NSArray<NSString*>*) filtersCountryCode
    filtersNpa: (NSArray<NSString*>*) filtersNpa
    filtersNxx: (NSArray<NSString*>*) filtersNxx
    filtersXxxx: (NSArray<NSString*>*) filtersXxxx
    filtersCity: (NSArray<NSString*>*) filtersCity
    filtersProvince: (NSArray<NSString*>*) filtersProvince
    filtersCountry: (NSArray<NSString*>*) filtersCountry
    filtersPrice: (NSArray<NSString*>*) filtersPrice
    filtersCategory: (NSArray<NSString*>*) filtersCategory
    sortInternal: (NSString*) sortInternal
    sortPrice: (NSString*) sortPrice
    sortPhoneNumber: (NSString*) sortPhoneNumber
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
        completionHandler: (void (^)(SWGListAvailableNumbers* output, NSError* error)) handler;
```





### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSArray<NSString*>* filtersPhoneNumber = @[@"filtersPhoneNumber_example"]; // Phone number filter (optional)
NSArray<NSString*>* filtersCountryCode = @[@"filtersCountryCode_example"]; // Country Code filter (optional)
NSArray<NSString*>* filtersNpa = @[@"filtersNpa_example"]; // Area Code filter (North America only) (optional)
NSArray<NSString*>* filtersNxx = @[@"filtersNxx_example"]; // 2nd set of 3 digits filter (North America only) (optional)
NSArray<NSString*>* filtersXxxx = @[@"filtersXxxx_example"]; // NANP XXXX filter (optional)
NSArray<NSString*>* filtersCity = @[@"filtersCity_example"]; // City filter (optional)
NSArray<NSString*>* filtersProvince = @[@"filtersProvince_example"]; // State or Province (postal code) filter (optional)
NSArray<NSString*>* filtersCountry = @[@"filtersCountry_example"]; // Country (postal code) filter (optional)
NSArray<NSString*>* filtersPrice = @[@"filtersPrice_example"]; // Price filter (optional)
NSArray<NSString*>* filtersCategory = @[@"filtersCategory_example"]; // Category filter (optional)
NSString* sortInternal = @"sortInternal_example"; // Internal (quasi-random) sorting (optional)
NSString* sortPrice = @"sortPrice_example"; // Price sorting (optional)
NSString* sortPhoneNumber = @"sortPhoneNumber_example"; // Phone number sorting (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)

SWGAvailablenumbersApi*apiInstance = [[SWGAvailablenumbersApi alloc] init];

// 
[apiInstance listAvailablePhoneNumbersWithFiltersPhoneNumber:filtersPhoneNumber
              filtersCountryCode:filtersCountryCode
              filtersNpa:filtersNpa
              filtersNxx:filtersNxx
              filtersXxxx:filtersXxxx
              filtersCity:filtersCity
              filtersProvince:filtersProvince
              filtersCountry:filtersCountry
              filtersPrice:filtersPrice
              filtersCategory:filtersCategory
              sortInternal:sortInternal
              sortPrice:sortPrice
              sortPhoneNumber:sortPhoneNumber
              limit:limit
              offset:offset
              fields:fields
          completionHandler: ^(SWGListAvailableNumbers* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGAvailablenumbersApi->listAvailablePhoneNumbers: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersPhoneNumber** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Phone number filter | [optional] 
 **filtersCountryCode** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Country Code filter | [optional] 
 **filtersNpa** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Area Code filter (North America only) | [optional] 
 **filtersNxx** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| 2nd set of 3 digits filter (North America only) | [optional] 
 **filtersXxxx** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| NANP XXXX filter | [optional] 
 **filtersCity** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| City filter | [optional] 
 **filtersProvince** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| State or Province (postal code) filter | [optional] 
 **filtersCountry** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Country (postal code) filter | [optional] 
 **filtersPrice** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Price filter | [optional] 
 **filtersCategory** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Category filter | [optional] 
 **sortInternal** | **NSString***| Internal (quasi-random) sorting | [optional] 
 **sortPrice** | **NSString***| Price sorting | [optional] 
 **sortPhoneNumber** | **NSString***| Phone number sorting | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 

### Return type

[**SWGListAvailableNumbers***](SWGListAvailableNumbers.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

