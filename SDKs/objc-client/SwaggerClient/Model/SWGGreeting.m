#import "SWGGreeting.h"

@implementation SWGGreeting

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"type": @"type", @"alternate": @"alternate", @"standard": @"standard", @"enableLeaveMessagePrompt": @"enable_leave_message_prompt" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"type", @"alternate", @"standard", @"enableLeaveMessagePrompt"];
  return [optionalProperties containsObject:propertyName];
}

@end
