#import "SWGCreateContactParams.h"

@implementation SWGCreateContactParams

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"firstName": @"first_name", @"middleName": @"middle_name", @"lastName": @"last_name", @"prefix": @"prefix", @"phoneticFirstName": @"phonetic_first_name", @"phoneticMiddleName": @"phonetic_middle_name", @"phoneticLastName": @"phonetic_last_name", @"suffix": @"suffix", @"nickname": @"nickname", @"company": @"company", @"department": @"department", @"jobTitle": @"job_title", @"emails": @"emails", @"phoneNumbers": @"phone_numbers", @"addresses": @"addresses", @"group": @"group" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"firstName", @"middleName", @"lastName", @"prefix", @"phoneticFirstName", @"phoneticMiddleName", @"phoneticLastName", @"suffix", @"nickname", @"company", @"department", @"jobTitle", @"emails", @"phoneNumbers", @"addresses", @"group"];
  return [optionalProperties containsObject:propertyName];
}

@end
