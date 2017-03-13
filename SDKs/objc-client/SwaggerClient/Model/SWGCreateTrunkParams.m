#import "SWGCreateTrunkParams.h"

@implementation SWGCreateTrunkParams

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"name": @"name", @"uri": @"uri", @"maxConcurrentCalls": @"max_concurrent_calls", @"maxMinutesPerMonth": @"max_minutes_per_month", @"greeting": @"greeting", @"errorMessage": @"error_message", @"codecs": @"codecs" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"name", @"uri", @"maxConcurrentCalls", @"maxMinutesPerMonth", @"greeting", @"errorMessage", @"codecs"];
  return [optionalProperties containsObject:propertyName];
}

@end
