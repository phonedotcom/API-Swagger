#import "SWGMenuFull.h"

@implementation SWGMenuFull

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"_id": @"id", @"name": @"name", @"allowExtensionDial": @"allow_extension_dial", @"keypressWaitTime": @"keypress_wait_time", @"greeting": @"greeting", @"keypressError": @"keypress_error", @"timeoutHandler": @"timeout_handler", @"options": @"options" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"name", @"allowExtensionDial", @"keypressWaitTime", @"greeting", @"keypressError", @"timeoutHandler", @"options"];
  return [optionalProperties containsObject:propertyName];
}

@end
