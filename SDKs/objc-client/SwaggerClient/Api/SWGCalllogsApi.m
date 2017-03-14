#import "SWGCalllogsApi.h"
#import "SWGQueryParamCollection.h"
#import "SWGApiClient.h"
#import "SWGCallLogFull.h"
#import "SWGListCallLogs.h"


@interface SWGCalllogsApi ()

@property (nonatomic, strong, readwrite) NSMutableDictionary *mutableDefaultHeaders;

@end

@implementation SWGCalllogsApi

NSString* kSWGCalllogsApiErrorDomain = @"SWGCalllogsApiErrorDomain";
NSInteger kSWGCalllogsApiMissingParamErrorCode = 234513;

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
/// Show details of an individual Call Log entry
/// See Call Logs for more detail.
///  @param accountId Account ID 
///
///  @param callId Call ID 
///
///  @returns SWGCallLogFull*
///
-(NSURLSessionTask*) getAccountCallLogsWithAccountId: (NSNumber*) accountId
    callId: (NSString*) callId
    completionHandler: (void (^)(SWGCallLogFull* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGCalllogsApiErrorDomain code:kSWGCalllogsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'callId' is set
    if (callId == nil) {
        NSParameterAssert(callId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"callId"] };
            NSError* error = [NSError errorWithDomain:kSWGCalllogsApiErrorDomain code:kSWGCalllogsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/call-logs/{call_id}"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (callId != nil) {
        pathParams[@"call_id"] = callId;
    }

    NSMutableDictionary* queryParams = [[NSMutableDictionary alloc] init];
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
                              responseType: @"SWGCallLogFull*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGCallLogFull*)data, error);
                                }
                            }];
}

///
/// Get a list of call details associated with your account
/// See Call Logs for more detail.
///  @param accountId Account ID 
///
///  @param filtersId ID filter (optional)
///
///  @param filtersStartTime Call start time filter (optional)
///
///  @param filtersCreatedAt Call log creation time filter (optional)
///
///  @param filtersDirection Call direction filter: in or out (optional)
///
///  @param filtersCalledNumber Called number (optional)
///
///  @param filtersType Call type, such as 'call', 'fax', 'audiogram' (optional)
///
///  @param filtersExtension Extension filter (optional)
///
///  @param sortId ID sorting (optional)
///
///  @param sortStartTime Sorting by call start time: asc or desc (optional)
///
///  @param sortCreatedAt Sorting by call log creation time: asc or desc (optional)
///
///  @param limit Max results (optional)
///
///  @param offset Results to skip (optional)
///
///  @param fields Field set (optional)
///
///  @returns SWGListCallLogs*
///
-(NSURLSessionTask*) listAccountCallLogsWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersStartTime: (NSArray<NSString*>*) filtersStartTime
    filtersCreatedAt: (NSString*) filtersCreatedAt
    filtersDirection: (NSString*) filtersDirection
    filtersCalledNumber: (NSString*) filtersCalledNumber
    filtersType: (NSString*) filtersType
    filtersExtension: (NSArray<NSString*>*) filtersExtension
    sortId: (NSString*) sortId
    sortStartTime: (NSString*) sortStartTime
    sortCreatedAt: (NSString*) sortCreatedAt
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
    completionHandler: (void (^)(SWGListCallLogs* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGCalllogsApiErrorDomain code:kSWGCalllogsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/call-logs"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }

    NSMutableDictionary* queryParams = [[NSMutableDictionary alloc] init];
    if (filtersId != nil) {
        queryParams[@"filters[id]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersId format: @"multi"];
        
    }
    if (filtersStartTime != nil) {
        queryParams[@"filters[start_time]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersStartTime format: @"multi"];
        
    }
    if (filtersCreatedAt != nil) {
        queryParams[@"filters[created_at]"] = filtersCreatedAt;
    }
    if (filtersDirection != nil) {
        queryParams[@"filters[direction]"] = filtersDirection;
    }
    if (filtersCalledNumber != nil) {
        queryParams[@"filters[called_number]"] = filtersCalledNumber;
    }
    if (filtersType != nil) {
        queryParams[@"filters[type]"] = filtersType;
    }
    if (filtersExtension != nil) {
        queryParams[@"filters[extension]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersExtension format: @"multi"];
        
    }
    if (sortId != nil) {
        queryParams[@"sort[id]"] = sortId;
    }
    if (sortStartTime != nil) {
        queryParams[@"sort[start_time]"] = sortStartTime;
    }
    if (sortCreatedAt != nil) {
        queryParams[@"sort[created_at]"] = sortCreatedAt;
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
                              responseType: @"SWGListCallLogs*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGListCallLogs*)data, error);
                                }
                            }];
}



@end
