#import "SWGExpressservicecodesApi.h"
#import "SWGQueryParamCollection.h"
#import "SWGApiClient.h"
#import "SWGExpressServiceCodeFull.h"
#import "SWGListExpressServiceCodes.h"


@interface SWGExpressservicecodesApi ()

@property (nonatomic, strong, readwrite) NSMutableDictionary *mutableDefaultHeaders;

@end

@implementation SWGExpressservicecodesApi

NSString* kSWGExpressservicecodesApiErrorDomain = @"SWGExpressservicecodesApiErrorDomain";
NSInteger kSWGExpressservicecodesApiMissingParamErrorCode = 234513;

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
/// Show details of an account Express Service Code
/// This service shows the details of an Account Express Service Code.
///  @param accountId Account ID 
///
///  @param codeId Device ID 
///
///  @returns SWGExpressServiceCodeFull*
///
-(NSURLSessionTask*) getAccountExpressSrvCodeWithAccountId: (NSNumber*) accountId
    codeId: (NSNumber*) codeId
    completionHandler: (void (^)(SWGExpressServiceCodeFull* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGExpressservicecodesApiErrorDomain code:kSWGExpressservicecodesApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'codeId' is set
    if (codeId == nil) {
        NSParameterAssert(codeId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"codeId"] };
            NSError* error = [NSError errorWithDomain:kSWGExpressservicecodesApiErrorDomain code:kSWGExpressservicecodesApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/express-service-codes/{code_id}"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (codeId != nil) {
        pathParams[@"code_id"] = codeId;
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
                              responseType: @"SWGExpressServiceCodeFull*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGExpressServiceCodeFull*)data, error);
                                }
                            }];
}

///
/// Get the Express Service Code associated with your account in list format
/// See Express Service Codes for more detail.
///  @param accountId Account ID 
///
///  @param filtersId ID filter (optional)
///
///  @returns SWGListExpressServiceCodes*
///
-(NSURLSessionTask*) listAccountExpressSrvCodesWithAccountId: (NSNumber*) accountId
    filtersId: (NSArray<NSString*>*) filtersId
    completionHandler: (void (^)(SWGListExpressServiceCodes* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGExpressservicecodesApiErrorDomain code:kSWGExpressservicecodesApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/express-service-codes"];

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
                              responseType: @"SWGListExpressServiceCodes*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGListExpressServiceCodes*)data, error);
                                }
                            }];
}



@end
