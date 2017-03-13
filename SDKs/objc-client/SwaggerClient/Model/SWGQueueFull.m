#import "SWGQueueFull.h"

@implementation SWGQueueFull

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"_id": @"id", @"name": @"name", @"greeting": @"greeting", @"holdMusic": @"hold_music", @"maxHoldTime": @"max_hold_time", @"callerIdType": @"caller_id_type", @"ringTime": @"ring_time", @"members": @"members" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"greeting", @"holdMusic", @"maxHoldTime", @"callerIdType", @"ringTime", @"members"];
  return [optionalProperties containsObject:propertyName];
}

@end
