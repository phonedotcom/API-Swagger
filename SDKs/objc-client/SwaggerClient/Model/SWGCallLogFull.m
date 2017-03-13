#import "SWGCallLogFull.h"

@implementation SWGCallLogFull

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"_id": @"id", @"uuid": @"uuid", @"extension": @"extension", @"callerId": @"caller_id", @"calledNumber": @"called_number", @"startTime": @"start_time", @"createdAt": @"created_at", @"direction": @"direction", @"type": @"type", @"callDuration": @"call_duration", @"isMonitored": @"is_monitored", @"callNumber": @"call_number", @"finalAction": @"final_action", @"callRecording": @"call_recording", @"details": @"details", @"callerCnam": @"caller_cnam" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"uuid", @"extension", @"callerId", @"calledNumber", @"startTime", @"createdAt", @"direction", @"type", @"callDuration", @"isMonitored", @"callNumber", @"finalAction", @"callRecording", @"details", @"callerCnam"];
  return [optionalProperties containsObject:propertyName];
}

@end
