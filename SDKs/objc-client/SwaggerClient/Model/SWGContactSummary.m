#import "SWGContactSummary.h"

@implementation SWGContactSummary

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"_id": @"id", @"prefix": @"prefix", @"firstName": @"first_name", @"middleName": @"middle_name", @"lastName": @"last_name", @"suffix": @"suffix", @"nickname": @"nickname", @"company": @"company" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"prefix", @"firstName", @"middleName", @"lastName", @"suffix", @"nickname", @"company"];
  return [optionalProperties containsObject:propertyName];
}

@end
