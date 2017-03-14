#import "SWGNumberregionsApi.h"
#import "SWGQueryParamCollection.h"
#import "SWGApiClient.h"
#import "SWGListPhoneNumbersRegions.h"


@interface SWGNumberregionsApi ()

@property (nonatomic, strong, readwrite) NSMutableDictionary *mutableDefaultHeaders;

@end

@implementation SWGNumberregionsApi

NSString* kSWGNumberregionsApiErrorDomain = @"SWGNumberregionsApiErrorDomain";
NSInteger kSWGNumberregionsApiMissingParamErrorCode = 234513;

@synthesize apiClient = _apiClient;

#pragma mark - Initialize methods

- (instancetype) init {
    return [self initWithApiClient:[SWGApiClient sharedClient]];
}


-(instancetype) initWithApiClient:(SWGApiClient *)apiClient {
    self = [super init];
    if (self) {
        _apiClient = apiClient;
        _mutableDefaultHeaders = [NSMutableDictionary dictionary];
    }
    return self;
}

#pragma mark -

-(NSString*) defaultHeaderForKey:(NSString*)key {
    return self.mutableDefaultHeaders[key];
}

-(void) setDefaultHeaderValue:(NSString*) value forKey:(NSString*)key {
    [self.mutableDefaultHeaders setValue:value forKey:key];
}

-(NSDictionary *)defaultHeaders {
    return self.mutableDefaultHeaders;
}

#pragma mark - Api Methods

///
/// 
/// This service lists the quantities of available phone numbers by region.
///  @param filtersCountryCode Country Code filter (optional)
///
///  @param filtersNpa Area Code filter (North America only) (optional)
///
///  @param filtersNxx 2nd set of 3 digits filter (North America only) (optional)
///
///  @param filtersIsTollFree Toll-free status filter (optional)
///
///  @param filtersCity City filter (optional)
///
///  @param filtersProvincePostalCode State or Province filter (optional)
///
///  @param filtersCountryPostalCode Country filter (optional)
///
///  @param sortCountryCode International calling code sorting (optional)
///
///  @param sortNpa Area Code sorting (North America only) (optional)
///
///  @param sortNxx 2nd set of 3 digits sorting (North America) (optional)
///
///  @param sortIsTollFree Toll Free status sorting (optional)
///
///  @param sortCity City sorting (optional)
///
///  @param sortProvincePostalCode State or Province sorting (optional)
///
///  @param sortCountryPostalCode Country sorting (optional)
///
///  @param limit Max results (optional)
///
///  @param offset Results to skip (optional)
///
///  @param fields Field set (optional)
///
///  @param groupBy Fields to group by (supports the same set of fields as the filters and sorting) (optional)
///
///  @returns SWGListPhoneNumbersRegions*
///
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
    completionHandler: (void (^)(SWGListPhoneNumbersRegions* output, NSError* error)) handler {
    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/phone-numbers/available/regions"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];

    NSMutableDictionary* queryParams = [[NSMutableDictionary alloc] init];
    if (filtersCountryCode != nil) {
        queryParams[@"filters[country_code]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersCountryCode format: @"multi"];
        
    }
    if (filtersNpa != nil) {
        queryParams[@"filters[npa]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersNpa format: @"multi"];
        
    }
    if (filtersNxx != nil) {
        queryParams[@"filters[nxx]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersNxx format: @"multi"];
        
    }
    if (filtersIsTollFree != nil) {
        queryParams[@"filters[is_toll_free]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersIsTollFree format: @"multi"];
        
    }
    if (filtersCity != nil) {
        queryParams[@"filters[city]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersCity format: @"multi"];
        
    }
    if (filtersProvincePostalCode != nil) {
        queryParams[@"filters[province_postal_code]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersProvincePostalCode format: @"multi"];
        
    }
    if (filtersCountryPostalCode != nil) {
        queryParams[@"filters[country_postal_code]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersCountryPostalCode format: @"csv"];
        
    }
    if (sortCountryCode != nil) {
        queryParams[@"sort[country_code]"] = sortCountryCode;
    }
    if (sortNpa != nil) {
        queryParams[@"sort[npa]"] = sortNpa;
    }
    if (sortNxx != nil) {
        queryParams[@"sort[nxx]"] = sortNxx;
    }
    if (sortIsTollFree != nil) {
        queryParams[@"sort[is_toll_free]"] = sortIsTollFree;
    }
    if (sortCity != nil) {
        queryParams[@"sort[city]"] = sortCity;
    }
    if (sortProvincePostalCode != nil) {
        queryParams[@"sort[province_postal_code]"] = sortProvincePostalCode;
    }
    if (sortCountryPostalCode != nil) {
        queryParams[@"sort[country_postal_code]"] = sortCountryPostalCode;
    }
    if (limit != nil) {
        queryParams[@"limit"] = limit;
    }
    if (offset != nil) {
        queryParams[@"offset"] = offset;
    }
    if (fields != nil) {
        queryParams[@"fields"] = fields;
    }
    if (groupBy != nil) {
        queryParams[@"group_by"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: groupBy format: @"csv"];
        
    }
    NSMutableDictionary* headerParams = [NSMutableDictionary dictionaryWithDictionary:self.apiClient.configuration.defaultHeaders];
    [headerParams addEntriesFromDictionary:self.defaultHeaders];
    // HTTP header `Accept`
    NSString *acceptHeader = [self.apiClient.sanitizer selectHeaderAccept:@[@"application/json"]];
    if(acceptHeader.length > 0) {
        headerParams[@"Accept"] = acceptHeader;
    }

    // response content type
    NSString *responseContentType = [[acceptHeader componentsSeparatedByString:@", "] firstObject] ?: @"";

    // request content type
    NSString *requestContentType = [self.apiClient.sanitizer selectHeaderContentType:@[@"application/json"]];

    // Authentication setting
    NSArray *authSettings = @[@"apiKey"];

    id bodyParam = nil;
    NSMutableDictionary *formParams = [[NSMutableDictionary alloc] init];
    NSMutableDictionary *localVarFiles = [[NSMutableDictionary alloc] init];

    return [self.apiClient requestWithPath: resourcePath
                                    method: @"GET"
                                pathParams: pathParams
                               queryParams: queryParams
                                formParams: formParams
                                     files: localVarFiles
                                      body: bodyParam
                              headerParams: headerParams
                              authSettings: authSettings
                        requestContentType: requestContentType
                       responseContentType: responseContentType
                              responseType: @"SWGListPhoneNumbersRegions*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGListPhoneNumbersRegions*)data, error);
                                }
                            }];
}



@end
