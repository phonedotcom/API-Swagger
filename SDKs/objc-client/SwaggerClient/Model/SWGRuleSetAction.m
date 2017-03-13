#import "SWGRuleSetAction.h"

@implementation SWGRuleSetAction

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"action": @"action", @"extension": @"extension", @"items": @"items", @"timeout": @"timeout", @"holdMusic": @"hold_music", @"greeting": @"greeting", @"duration": @"duration", @"menu": @"menu", @"queue": @"queue", @"trunk": @"trunk" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"action", @"extension", @"items", @"timeout", @"holdMusic", @"greeting", @"duration", @"menu", @"queue", @"trunk"];
  return [optionalProperties containsObject:propertyName];
}

@end
