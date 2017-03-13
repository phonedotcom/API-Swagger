# SWGNumberregionsApi

All URIs are relative to *https://api.phone.com/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listAvailablePhoneNumberRegions**](SWGNumberregionsApi.md#listavailablephonenumberregions) | **GET** /phone-numbers/available/regions | 


# **listAvailablePhoneNumberRegions**
```objc
-(NSURLSessionTask*) listAvailablePhoneNumberRegionsWithFiltersCountryCode: (NSArray<NSString*>*) filtersCountryCode
    filtersNpa: (NSArray<NSString*>*) filtersNpa
    filtersNxx: (NSArray<NSString*>*) filtersNxx
    filtersIsTollFree: (NSArray<NSString*>*) filtersIsTollFree
    filtersCity: (NSArray<NSString*>*) filtersCity
    filtersProvincePostalCode: (NSArray<NSString*>*) filtersProvincePostalCode
    filtersCountryPostalCode: (NSArray<NSString*>*) filtersCountryPostalCode
    sortCountryCode: (NSString*) sortCountryCode
    sortNpa: (NSString*) sortNpa
    sortNxx: (NSString*) sortNxx
    sortIsTollFree: (NSString*) sortIsTollFree
    sortCity: (NSString*) sortCity
    sortProvincePostalCode: (NSString*) sortProvincePostalCode
    sortCountryPostalCode: (NSString*) sortCountryPostalCode
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
    groupBy: (NSArray<NSString*>*) groupBy
        completionHandler: (void (^)(SWGListPhoneNumbersRegions* output, NSError* error)) handler;
```



This service lists the quantities of available phone numbers by region.

### Example 
```objc
SWGDefaultConfiguration *apiConfig = [SWGDefaultConfiguration sharedConfig];

// Configure API key authorization: (authentication scheme: apiKey)
[apiConfig setApiKey:@"YOUR_API_KEY" forApiKeyIdentifier:@"Authorization"];
// Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//[apiConfig setApiKeyPrefix:@"Bearer" forApiKeyIdentifier:@"Authorization"];


NSArray<NSString*>* filtersCountryCode = @[@"filtersCountryCode_example"]; // Country Code filter (optional)
NSArray<NSString*>* filtersNpa = @[@"filtersNpa_example"]; // Area Code filter (North America only) (optional)
NSArray<NSString*>* filtersNxx = @[@"filtersNxx_example"]; // 2nd set of 3 digits filter (North America only) (optional)
NSArray<NSString*>* filtersIsTollFree = @[@"filtersIsTollFree_example"]; // Toll-free status filter (optional)
NSArray<NSString*>* filtersCity = @[@"filtersCity_example"]; // City filter (optional)
NSArray<NSString*>* filtersProvincePostalCode = @[@"filtersProvincePostalCode_example"]; // State or Province filter (optional)
NSArray<NSString*>* filtersCountryPostalCode = @[@"filtersCountryPostalCode_example"]; // Country filter (optional)
NSString* sortCountryCode = @"sortCountryCode_example"; // International calling code sorting (optional)
NSString* sortNpa = @"sortNpa_example"; // Area Code sorting (North America only) (optional)
NSString* sortNxx = @"sortNxx_example"; // 2nd set of 3 digits sorting (North America) (optional)
NSString* sortIsTollFree = @"sortIsTollFree_example"; // Toll Free status sorting (optional)
NSString* sortCity = @"sortCity_example"; // City sorting (optional)
NSString* sortProvincePostalCode = @"sortProvincePostalCode_example"; // State or Province sorting (optional)
NSString* sortCountryPostalCode = @"sortCountryPostalCode_example"; // Country sorting (optional)
NSNumber* limit = @56; // Max results (optional)
NSNumber* offset = @56; // Results to skip (optional)
NSString* fields = @"fields_example"; // Field set (optional)
NSArray<NSString*>* groupBy = @[@"groupBy_example"]; // Fields to group by (supports the same set of fields as the filters and sorting) (optional)

SWGNumberregionsApi*apiInstance = [[SWGNumberregionsApi alloc] init];

// 
[apiInstance listAvailablePhoneNumberRegionsWithFiltersCountryCode:filtersCountryCode
              filtersNpa:filtersNpa
              filtersNxx:filtersNxx
              filtersIsTollFree:filtersIsTollFree
              filtersCity:filtersCity
              filtersProvincePostalCode:filtersProvincePostalCode
              filtersCountryPostalCode:filtersCountryPostalCode
              sortCountryCode:sortCountryCode
              sortNpa:sortNpa
              sortNxx:sortNxx
              sortIsTollFree:sortIsTollFree
              sortCity:sortCity
              sortProvincePostalCode:sortProvincePostalCode
              sortCountryPostalCode:sortCountryPostalCode
              limit:limit
              offset:offset
              fields:fields
              groupBy:groupBy
          completionHandler: ^(SWGListPhoneNumbersRegions* output, NSError* error) {
                        if (output) {
                            NSLog(@"%@", output);
                        }
                        if (error) {
                            NSLog(@"Error calling SWGNumberregionsApi->listAvailablePhoneNumberRegions: %@", error);
                        }
                    }];
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersCountryCode** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Country Code filter | [optional] 
 **filtersNpa** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Area Code filter (North America only) | [optional] 
 **filtersNxx** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| 2nd set of 3 digits filter (North America only) | [optional] 
 **filtersIsTollFree** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Toll-free status filter | [optional] 
 **filtersCity** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| City filter | [optional] 
 **filtersProvincePostalCode** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| State or Province filter | [optional] 
 **filtersCountryPostalCode** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Country filter | [optional] 
 **sortCountryCode** | **NSString***| International calling code sorting | [optional] 
 **sortNpa** | **NSString***| Area Code sorting (North America only) | [optional] 
 **sortNxx** | **NSString***| 2nd set of 3 digits sorting (North America) | [optional] 
 **sortIsTollFree** | **NSString***| Toll Free status sorting | [optional] 
 **sortCity** | **NSString***| City sorting | [optional] 
 **sortProvincePostalCode** | **NSString***| State or Province sorting | [optional] 
 **sortCountryPostalCode** | **NSString***| Country sorting | [optional] 
 **limit** | **NSNumber***| Max results | [optional] 
 **offset** | **NSNumber***| Results to skip | [optional] 
 **fields** | **NSString***| Field set | [optional] 
 **groupBy** | [**NSArray&lt;NSString*&gt;***](NSString*.md)| Fields to group by (supports the same set of fields as the filters and sorting) | [optional] 

### Return type

[**SWGListPhoneNumbersRegions***](SWGListPhoneNumbersRegions.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

