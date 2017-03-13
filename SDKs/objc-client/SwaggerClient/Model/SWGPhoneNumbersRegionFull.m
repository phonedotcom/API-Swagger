#import "SWGPhoneNumbersRegionFull.h"

@implementation SWGPhoneNumbersRegionFull

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"countryCode": @"country_code", @"npa": @"npa", @"nxx": @"nxx", @"isTollFree": @"is_toll_free", @"city": @"city", @"provincePostalCode": @"province_postal_code", @"countryPostalCode": @"country_postal_code", @"quantity": @"quantity" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"countryCode", @"npa", @"nxx", @"isTollFree", @"city", @"provincePostalCode", @"countryPostalCode", @"quantity"];
  return [optionalProperties containsObject:propertyName];
}

@end
