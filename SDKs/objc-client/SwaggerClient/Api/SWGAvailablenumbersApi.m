#import "SWGAvailablenumbersApi.h"
#import "SWGQueryParamCollection.h"
#import "SWGApiClient.h"
#import "SWGListAvailableNumbers.h"


@interface SWGAvailablenumbersApi ()

@property (nonatomic, strong, readwrite) NSMutableDictionary *mutableDefaultHeaders;

@end

@implementation SWGAvailablenumbersApi

NSString* kSWGAvailablenumbersApiErrorDomain = @"SWGAvailablenumbersApiErrorDomain";
NSInteger kSWGAvailablenumbersApiMissingParamErrorCode = 234513;

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
/// 
///  @param filtersPhoneNumber Phone number filter (optional)
///
///  @param filtersCountryCode Country Code filter (optional)
///
///  @param filtersNpa Area Code filter (North America only) (optional)
///
///  @param filtersNxx 2nd set of 3 digits filter (North America only) (optional)
///
///  @param filtersXxxx NANP XXXX filter (optional)
///
///  @param filtersCity City filter (optional)
///
///  @param filtersProvince State or Province (postal code) filter (optional)
///
///  @param filtersCountry Country (postal code) filter (optional)
///
///  @param filtersPrice Price filter (optional)
///
///  @param filtersCategory Category filter (optional)
///
///  @param sortInternal Internal (quasi-random) sorting (optional)
///
///  @param sortPrice Price sorting (optional)
///
///  @param sortPhoneNumber Phone number sorting (optional)
///
///  @param limit Max results (optional)
///
///  @param offset Results to skip (optional)
///
///  @param fields Field set (optional)
///
///  @returns SWGListAvailableNumbers*
///
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
    completionHandler: (void (^)(SWGListAvailableNumbers* output, NSError* error)) handler {
    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/phone-numbers/available"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];

    NSMutableDictionary* queryParams = [[NSMutableDictionary alloc] init];
    if (filtersPhoneNumber != nil) {
        queryParams[@"filters[phone_number]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersPhoneNumber format: @"multi"];
        
    }
    if (filtersCountryCode != nil) {
        queryParams[@"filters[country_code]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersCountryCode format: @"multi"];
        
    }
    if (filtersNpa != nil) {
        queryParams[@"filters[npa]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersNpa format: @"multi"];
        
    }
    if (filtersNxx != nil) {
        queryParams[@"filters[nxx]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersNxx format: @"multi"];
        
    }
    if (filtersXxxx != nil) {
        queryParams[@"filters[xxxx]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersXxxx format: @"multi"];
        
    }
    if (filtersCity != nil) {
        queryParams[@"filters[city]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersCity format: @"multi"];
        
    }
    if (filtersProvince != nil) {
        queryParams[@"filters[province]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersProvince format: @"multi"];
        
    }
    if (filtersCountry != nil) {
        queryParams[@"filters[country]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersCountry format: @"multi"];
        
    }
    if (filtersPrice != nil) {
        queryParams[@"filters[price]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersPrice format: @"multi"];
        
    }
    if (filtersCategory != nil) {
        queryParams[@"filters[category]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersCategory format: @"multi"];
        
    }
    if (sortInternal != nil) {
        queryParams[@"sort[internal]"] = sortInternal;
    }
    if (sortPrice != nil) {
        queryParams[@"sort[price]"] = sortPrice;
    }
    if (sortPhoneNumber != nil) {
        queryParams[@"sort[phone_number]"] = sortPhoneNumber;
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
                              responseType: @"SWGListAvailableNumbers*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGListAvailableNumbers*)data, error);
                                }
                            }];
}



@end
