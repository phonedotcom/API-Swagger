#import "SWGFilterListAvailableNumbers.h"

@implementation SWGFilterListAvailableNumbers

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"phoneNumber": @"phone_number", @"countryCode": @"country_code", @"npa": @"npa", @"nxx": @"nxx", @"xxxx": @"xxxx", @"city": @"city", @"province": @"province", @"country": @"country", @"price": @"price", @"category": @"category" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"phoneNumber", @"countryCode", @"npa", @"nxx", @"xxxx", @"city", @"province", @"country", @"price", @"category"];
  return [optionalProperties containsObject:propertyName];
}

@end
