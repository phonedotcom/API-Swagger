#import "SWGCreateMenuParams.h"

@implementation SWGCreateMenuParams

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"name": @"name", @"mainMessage": @"main_message", @"invalidKeypressMessage": @"invalid_keypress_message", @"allowExtensionDial": @"allow_extension_dial", @"keypressWaitTime": @"keypress_wait_time", @"timeoutHandler": @"timeout_handler", @"options": @"options" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"name", @"mainMessage", @"invalidKeypressMessage", @"allowExtensionDial", @"keypressWaitTime", @"timeoutHandler", @"options"];
  return [optionalProperties containsObject:propertyName];
}

@end
