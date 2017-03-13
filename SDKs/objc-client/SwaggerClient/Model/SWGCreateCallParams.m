#import "SWGCreateCallParams.h"

@implementation SWGCreateCallParams

- (instancetype)init {
  self = [super init];
  if (self) {
    // initialize property's default value, if any
    
  }
  return self;
}


/**
 * Maps json key to property name.
 * This method is used by `JSONModel`.
 */
+ (JSONKeyMapper *)keyMapper {
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"callerPhoneNumber": @"caller_phone_number", @"callerExtension": @"caller_extension", @"callerCallerId": @"caller_caller_id", @"callerPrivate": @"caller_private", @"calleePhoneNumber": @"callee_phone_number", @"calleeExtension": @"callee_extension", @"calleeCallerId": @"callee_caller_id", @"calleePrivate": @"callee_private" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"callerPhoneNumber", @"callerExtension", @"callerCallerId", @"callerPrivate", @"calleePhoneNumber", @"calleeExtension", @"calleeCallerId", @"calleePrivate"];
  return [optionalProperties containsObject:propertyName];
}

@end
