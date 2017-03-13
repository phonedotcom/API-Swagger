#import "SWGSmsApi.h"
#import "SWGQueryParamCollection.h"
#import "SWGApiClient.h"
#import "SWGCreateSmsParams.h"
#import "SWGListSms.h"
#import "SWGSmsFull.h"


@interface SWGSmsApi ()

@property (nonatomic, strong, readwrite) NSMutableDictionary *mutableDefaultHeaders;

@end

@implementation SWGSmsApi

NSString* kSWGSmsApiErrorDomain = @"SWGSmsApiErrorDomain";
NSInteger kSWGSmsApiMissingParamErrorCode = 234513;

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
/// Send a SMS to one or a group of recipients
/// For more on the input fields, see Intro to SMS.
///  @param accountId Account ID 
///
///  @param data SMS data 
///
///  @returns SWGSmsFull*
///
-(NSURLSessionTask*) createAccountSmsWithAccountId: (NSNumber*) accountId
    data: (SWGCreateSmsParams*) data
    completionHandler: (void (^)(SWGSmsFull* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGSmsApiErrorDomain code:kSWGSmsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'data' is set
    if (data == nil) {
        NSParameterAssert(data);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"data"] };
            NSError* error = [NSError errorWithDomain:kSWGSmsApiErrorDomain code:kSWGSmsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/sms"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
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
    bodyParam = data;

    return [self.apiClient requestWithPath: resourcePath
                                    method: @"POST"
                                pathParams: pathParams
                               queryParams: queryParams
                                formParams: formParams
                                     files: localVarFiles
                                      body: bodyParam
                              headerParams: headerParams
                              authSettings: authSettings
                        requestContentType: requestContentType
                       responseContentType: responseContentType
                              responseType: @"SWGSmsFull*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGSmsFull*)data, error);
                                }
                            }];
}

///
/// Show details of an individual SMS
/// This service shows the details of an individual sms. See Intro to SMS for more info on the properties.
///  @param accountId Account ID 
///
///  @param smsId SMS ID 
///
///  @returns SWGSmsFull*
///
-(NSURLSessionTask*) getAccountSmsWithAccountId: (NSNumber*) accountId
    smsId: (NSString*) smsId
    completionHandler: (void (^)(SWGSmsFull* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGSmsApiErrorDomain code:kSWGSmsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'smsId' is set
    if (smsId == nil) {
        NSParameterAssert(smsId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"smsId"] };
            NSError* error = [NSError errorWithDomain:kSWGSmsApiErrorDomain code:kSWGSmsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/sms/{sms_id}"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (smsId != nil) {
        pathParams[@"sms_id"] = smsId;
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
                              responseType: @"SWGSmsFull*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGSmsFull*)data, error);
                                }
                            }];
}

///
/// Get a list of SMS messages for an account
/// See Intro to SMS for more info on the properties.
///  @param accountId Account ID 
///
///  @param filtersId ID filter (optional)
///
///  @param filtersDirection Direction filter (optional)
///
///  @param filtersFrom Caller ID filter (optional)
///
///  @param sortId ID sorting (optional)
///
///  @param sortCreatedAt Sort by created time of message (optional)
///
///  @param limit Max results (optional)
///
///  @param offset Results to skip (optional)
///
///  @param fields Field set (optional)
///
///  @returns SWGListSms*
///
-(NSURLSessionTask*) listAccountSmsWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersDirection: (NSString*) filtersDirection
    filtersFrom: (NSString*) filtersFrom
    sortId: (NSString*) sortId
    sortCreatedAt: (NSString*) sortCreatedAt
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
    completionHandler: (void (^)(SWGListSms* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGSmsApiErrorDomain code:kSWGSmsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/sms"];

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
    if (filtersDirection != nil) {
        queryParams[@"filters[direction]"] = filtersDirection;
    }
    if (filtersFrom != nil) {
        queryParams[@"filters[from]"] = filtersFrom;
    }
    if (sortId != nil) {
        queryParams[@"sort[id]"] = sortId;
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
                              responseType: @"SWGListSms*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGListSms*)data, error);
                                }
                            }];
}



@end
