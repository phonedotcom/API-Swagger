#import "SWGContactsApi.h"
#import "SWGQueryParamCollection.h"
#import "SWGApiClient.h"
#import "SWGContactFull.h"
#import "SWGCreateContactParams.h"
#import "SWGDeleteContact.h"
#import "SWGListContacts.h"


@interface SWGContactsApi ()

@property (nonatomic, strong, readwrite) NSMutableDictionary *mutableDefaultHeaders;

@end

@implementation SWGContactsApi

NSString* kSWGContactsApiErrorDomain = @"SWGContactsApiErrorDomain";
NSInteger kSWGContactsApiMissingParamErrorCode = 234513;

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
/// Add a new address book contact for an extension
/// For more on the input fields, see Account Contacts.
///  @param accountId Account ID 
///
///  @param extensionId Extension ID 
///
///  @param data Contact data (optional)
///
///  @returns SWGContactFull*
///
-(NSURLSessionTask*) createAccountExtensionContactWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    data: (SWGCreateContactParams*) data
    completionHandler: (void (^)(SWGContactFull* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'extensionId' is set
    if (extensionId == nil) {
        NSParameterAssert(extensionId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"extensionId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/extensions/{extension_id}/contacts"];

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
                              responseType: @"SWGContactFull*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGContactFull*)data, error);
                                }
                            }];
}

///
/// 
/// 
///  @param accountId Account ID 
///
///  @param extensionId Extension ID 
///
///  @param contactId Contact ID 
///
///  @returns SWGDeleteContact*
///
-(NSURLSessionTask*) deleteAccountExtensionContactWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    contactId: (NSNumber*) contactId
    completionHandler: (void (^)(SWGDeleteContact* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'extensionId' is set
    if (extensionId == nil) {
        NSParameterAssert(extensionId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"extensionId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'contactId' is set
    if (contactId == nil) {
        NSParameterAssert(contactId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"contactId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id}"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (extensionId != nil) {
        pathParams[@"extension_id"] = extensionId;
    }
    if (contactId != nil) {
        pathParams[@"contact_id"] = contactId;
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
                              responseType: @"SWGDeleteContact*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGDeleteContact*)data, error);
                                }
                            }];
}

///
/// Retrieve the details of an address book contact
/// For more info on the fields shown, see Account Contacts.
///  @param accountId Account ID 
///
///  @param extensionId Extension ID 
///
///  @param contactId Contact ID 
///
///  @returns SWGContactFull*
///
-(NSURLSessionTask*) getAccountExtensionContactWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    contactId: (NSNumber*) contactId
    completionHandler: (void (^)(SWGContactFull* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'extensionId' is set
    if (extensionId == nil) {
        NSParameterAssert(extensionId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"extensionId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'contactId' is set
    if (contactId == nil) {
        NSParameterAssert(contactId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"contactId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id}"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (extensionId != nil) {
        pathParams[@"extension_id"] = extensionId;
    }
    if (contactId != nil) {
        pathParams[@"contact_id"] = contactId;
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
                              responseType: @"SWGContactFull*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGContactFull*)data, error);
                                }
                            }];
}

///
/// Show a list of address book contacts
/// See Account Contacts for more info on the fields in each item.
///  @param accountId Account ID 
///
///  @param extensionId Extension ID 
///
///  @param filtersId ID filter (optional)
///
///  @param filtersGroupId Group filter (optional)
///
///  @param filtersUpdatedAt Updated At filter (optional)
///
///  @param sortId ID sorting (optional)
///
///  @param sortUpdatedAt Updated At sorting (optional)
///
///  @param limit Max results (optional)
///
///  @param offset Results to skip (optional)
///
///  @param fields Field set (optional)
///
///  @returns SWGListContacts*
///
-(NSURLSessionTask*) listAccountExtensionContactsWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    filtersId: (NSArray<NSString*>*) filtersId
    filtersGroupId: (NSArray<NSString*>*) filtersGroupId
    filtersUpdatedAt: (NSArray<NSString*>*) filtersUpdatedAt
    sortId: (NSString*) sortId
    sortUpdatedAt: (NSString*) sortUpdatedAt
    limit: (NSNumber*) limit
    offset: (NSNumber*) offset
    fields: (NSString*) fields
    completionHandler: (void (^)(SWGListContacts* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'extensionId' is set
    if (extensionId == nil) {
        NSParameterAssert(extensionId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"extensionId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/extensions/{extension_id}/contacts"];

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
    if (filtersGroupId != nil) {
        queryParams[@"filters[group_id]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersGroupId format: @"multi"];
        
    }
    if (filtersUpdatedAt != nil) {
        queryParams[@"filters[updated_at]"] = [[SWGQueryParamCollection alloc] initWithValuesAndFormat: filtersUpdatedAt format: @"multi"];
        
    }
    if (sortId != nil) {
        queryParams[@"sort[id]"] = sortId;
    }
    if (sortUpdatedAt != nil) {
        queryParams[@"sort[updated_at]"] = sortUpdatedAt;
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
                              responseType: @"SWGListContacts*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGListContacts*)data, error);
                                }
                            }];
}

///
/// 
/// For more on the input fields, see Account Contacts.
///  @param accountId Account ID 
///
///  @param extensionId Extension ID 
///
///  @param contactId Contact ID 
///
///  @param data Contact data (optional)
///
///  @returns SWGContactFull*
///
-(NSURLSessionTask*) replaceAccountExtensionContactWithAccountId: (NSNumber*) accountId
    extensionId: (NSNumber*) extensionId
    contactId: (NSNumber*) contactId
    data: (SWGCreateContactParams*) data
    completionHandler: (void (^)(SWGContactFull* output, NSError* error)) handler {
    // verify the required parameter 'accountId' is set
    if (accountId == nil) {
        NSParameterAssert(accountId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"accountId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'extensionId' is set
    if (extensionId == nil) {
        NSParameterAssert(extensionId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"extensionId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    // verify the required parameter 'contactId' is set
    if (contactId == nil) {
        NSParameterAssert(contactId);
        if(handler) {
            NSDictionary * userInfo = @{NSLocalizedDescriptionKey : [NSString stringWithFormat:NSLocalizedString(@"Missing required parameter '%@'", nil),@"contactId"] };
            NSError* error = [NSError errorWithDomain:kSWGContactsApiErrorDomain code:kSWGContactsApiMissingParamErrorCode userInfo:userInfo];
            handler(nil, error);
        }
        return nil;
    }

    NSMutableString* resourcePath = [NSMutableString stringWithFormat:@"/accounts/{account_id}/extensions/{extension_id}/contacts/{contact_id}"];

    // remove format in URL if needed
    [resourcePath replaceOccurrencesOfString:@".{format}" withString:@".json" options:0 range:NSMakeRange(0,resourcePath.length)];

    NSMutableDictionary *pathParams = [[NSMutableDictionary alloc] init];
    if (accountId != nil) {
        pathParams[@"account_id"] = accountId;
    }
    if (extensionId != nil) {
        pathParams[@"extension_id"] = extensionId;
    }
    if (contactId != nil) {
        pathParams[@"contact_id"] = contactId;
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
                              responseType: @"SWGContactFull*"
                           completionBlock: ^(id data, NSError *error) {
                                if(handler) {
                                    handler((SWGContactFull*)data, error);
                                }
                            }];
}



@end
