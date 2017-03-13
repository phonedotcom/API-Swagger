#import "SWGExtensionFull.h"

@implementation SWGExtensionFull

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"_id": @"id", @"name": @"name", @"extension": @"extension", @"fullName": @"full_name", @"usageType": @"usage_type", @"deviceMembership": @"device_membership", @"timezone": @"timezone", @"nameGreeting": @"name_greeting", @"includeInDirectory": @"include_in_directory", @"callerId": @"caller_id", @"localAreaCode": @"local_area_code", @"enableCallWaiting": @"enable_call_waiting", @"enableOutboundCalls": @"enable_outbound_calls", @"voicemail": @"voicemail", @"callNotifications": @"call_notifications", @"route": @"route" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"name", @"extension", @"fullName", @"usageType", @"deviceMembership", @"timezone", @"nameGreeting", @"includeInDirectory", @"callerId", @"localAreaCode", @"enableCallWaiting", @"enableOutboundCalls", @"voicemail", @"callNotifications", @"route"];
  return [optionalProperties containsObject:propertyName];
}

@end
