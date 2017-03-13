#import "SWGContactFull.h"

@implementation SWGContactFull

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"_id": @"id", @"prefix": @"prefix", @"firstName": @"first_name", @"middleName": @"middle_name", @"lastName": @"last_name", @"suffix": @"suffix", @"nickname": @"nickname", @"company": @"company", @"phoneticFirstName": @"phonetic_first_name", @"phoneticMiddleName": @"phonetic_middle_name", @"phoneticLastName": @"phonetic_last_name", @"department": @"department", @"jobTitle": @"job_title", @"emails": @"emails", @"phoneNumbers": @"phone_numbers", @"addresses": @"addresses", @"group": @"group", @"createdAt": @"created_at", @"updatedAt": @"updated_at" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"prefix", @"firstName", @"middleName", @"lastName", @"suffix", @"nickname", @"company", @"phoneticFirstName", @"phoneticMiddleName", @"phoneticLastName", @"department", @"jobTitle", @"emails", @"phoneNumbers", @"addresses", @"group", @"createdAt", @"updatedAt"];
  return [optionalProperties containsObject:propertyName];
}

@end
