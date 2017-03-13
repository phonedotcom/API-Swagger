#import "SWGContactSubaccount.h"

@implementation SWGContactSubaccount

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"name": @"name", @"address": @"address", @"primaryEmail": @"primary_email", @"alternateEmail": @"alternate_email", @"company": @"company", @"phone": @"phone", @"fax": @"fax" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"address", @"alternateEmail", @"company", @"fax"];
  return [optionalProperties containsObject:propertyName];
}

@end
