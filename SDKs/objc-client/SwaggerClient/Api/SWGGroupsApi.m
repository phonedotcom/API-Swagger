#import "SWGGroupsApi.h"
#import "SWGQueryParamCollection.h"
#import "SWGApiClient.h"
#import "SWGCreateGroupParams.h"
#import "SWGDeleteGroup.h"
#import "SWGGroupFull.h"
#import "SWGListGroups.h"


@interface SWGGroupsApi ()

@property (nonatomic, strong, readwrite) NSMutableDictionary *mutableDefaultHeaders;

@end

@implementation SWGGroupsApi

NSString* kSWGGroupsApiErrorDomain = @"SWGGroupsApiErrorDomain";
NSInteger kSWGGroupsApiMissingParamErrorCode = 234513;

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
/// See Account Contact Groups for more info on the properties.
///  @param accountId Account ID 
///
///  @param extensionId Extension ID 
///
///  @param data Group name 
///
///  @returns SWGGroupFull*
///
-(NSURLSessionTask*) createAccountExtensionContactGroupWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    data: (SWGCreateGroupParams*) data
    completionHandler: (void (^)(SWGGroupFull* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'extensionId' is set
    if (extensionId == nil) {
        NSParameterAssert(extensionId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"extensionId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'data' is set
    if (data == nil) {
        NSParameterAssert(data);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"data"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/extensions/{extension_id}/contact-groups"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (extensionId != nil) {
        pathParams[@"extension_id"] = extensionId;
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
                              responseType: @"SWGGroupFull*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGGroupFull*)data, error);
                                }
                            }];
}

///
/// Delete an addressbook group
/// 
///  @param accountId Account ID 
///
///  @param extensionId Extension ID 
///
///  @param groupId Group ID 
///
///  @returns SWGDeleteGroup*
///
-(NSURLSessionTask*) deleteAccountExtensionContactGroupWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    groupId: (NSNumber*) groupId
    completionHandler: (void (^)(SWGDeleteGroup* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'extensionId' is set
    if (extensionId == nil) {
        NSParameterAssert(extensionId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"extensionId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'groupId' is set
    if (groupId == nil) {
        NSParameterAssert(groupId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"groupId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id}"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (extensionId != nil) {
        pathParams[@"extension_id"] = extensionId;
    }
    if (groupId != nil) {
        pathParams[@"group_id"] = groupId;
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
                                    method: @"DELETE"
                                pathParams: pathParams
                               queryParams: queryParams
                                formParams: formParams
                                     files: localVarFiles
                                      body: bodyParam
                              headerParams: headerParams
                              authSettings: authSettings
                        requestContentType: requestContentType
                       responseContentType: responseContentType
                              responseType: @"SWGDeleteGroup*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGDeleteGroup*)data, error);
                                }
                            }];
}

///
/// 
/// See Account Contact Groups for more info on the properties.
///  @param accountId Account ID 
///
///  @param extensionId Extension ID 
///
///  @param groupId Group ID 
///
///  @returns SWGGroupFull*
///
-(NSURLSessionTask*) getAccountExtensionContactGroupWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    groupId: (NSNumber*) groupId
    completionHandler: (void (^)(SWGGroupFull* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'extensionId' is set
    if (extensionId == nil) {
        NSParameterAssert(extensionId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"extensionId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'groupId' is set
    if (groupId == nil) {
        NSParameterAssert(groupId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"groupId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id}"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (extensionId != nil) {
        pathParams[@"extension_id"] = extensionId;
    }
    if (groupId != nil) {
        pathParams[@"group_id"] = groupId;
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
                              responseType: @"SWGGroupFull*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGGroupFull*)data, error);
                                }
                            }];
}

///
/// Show a list of contact groups belonging to an extension
/// See Account Contact Groups for details on the properties.
///  @param accountId Account ID 
///
///  @param extensionId Extension ID 
///
///  @param filtersId ID filter (optional)
///
///  @param filtersName Name filter (optional)
///
///  @param sortId ID sorting (optional)
///
///  @param sortName Name sorting (optional)
///
///  @param limit Max results (optional)
///
///  @param offset Results to skip (optional)
///
///  @param fields Field set (optional)
///
///  @returns SWGListGroups*
///
-(NSURLSessionTask*) listAccountExtensionContactGroupsWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersName: (NSArray<NSString*>*) filtersName
    sortId: (NSString*) sortId
    sortName: (NSString*) sortName
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
    completionHandler: (void (^)(SWGListGroups* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'extensionId' is set
    if (extensionId == nil) {
        NSParameterAssert(extensionId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"extensionId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/extensions/{extension_id}/contact-groups"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (extensionId != nil) {
        pathParams[@"extension_id"] = extensionId;
    }

    NSMutableDictionary* queryParams = [[NSMutableDictionary alloc] init];
    if (filtersId != nil) {
        queryParams[@"filters[id]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersId format: @"multi"];
        
    }
    if (filtersName != nil) {
        queryParams[@"filters[name]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersName format: @"multi"];
        
    }
    if (sortId != nil) {
        queryParams[@"sort[id]"] = sortId;
    }
    if (sortName != nil) {
        queryParams[@"sort[name]"] = sortName;
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
                              responseType: @"SWGListGroups*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGListGroups*)data, error);
                                }
                            }];
}

///
/// 
/// See Account Contact Groups for more info on the properties.
///  @param accountId Account ID 
///
///  @param extensionId Extension ID 
///
///  @param groupId Group ID 
///
///  @param data Group name 
///
///  @returns SWGGroupFull*
///
-(NSURLSessionTask*) replaceAccountExtensionContactGroupWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    groupId: (NSNumber*) groupId
    data: (SWGCreateGroupParams*) data
    completionHandler: (void (^)(SWGGroupFull* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'extensionId' is set
    if (extensionId == nil) {
        NSParameterAssert(extensionId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"extensionId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'groupId' is set
    if (groupId == nil) {
        NSParameterAssert(groupId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"groupId"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'data' is set
    if (data == nil) {
        NSParameterAssert(data);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"data"] };
            NSError* error = [NSError errorWithDomain:kSWGGroupsApiErrorDomain code:kSWGGroupsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/extensions/{extension_id}/contact-groups/{group_id}"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (extensionId != nil) {
        pathParams[@"extension_id"] = extensionId;
    }
    if (groupId != nil) {
        pathParams[@"group_id"] = groupId;
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
                                    method: @"PUT"
                                pathParams: pathParams
                               queryParams: queryParams
                                formParams: formParams
                                     files: localVarFiles
                                      body: bodyParam
                              headerParams: headerParams
                              authSettings: authSettings
                        requestContentType: requestContentType
                       responseContentType: responseContentType
                              responseType: @"SWGGroupFull*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGGroupFull*)data, error);
                                }
                            }];
}



@end
