#import "SWGPhoneNumberFull.h"

@implementation SWGPhoneNumberFull

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"_id": @"id", @"name": @"name", @"phoneNumber": @"phone_number", @"blockIncoming": @"block_incoming", @"blockAnonymous": @"block_anonymous", @"route": @"route", @"callerId": @"caller_id", @"smsForwarding": @"sms_forwarding", @"callNotifications": @"call_notifications" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"_id", @"name", @"phoneNumber", @"blockIncoming", @"blockAnonymous", @"route", @"callerId", @"smsForwarding", @"callNotifications"];
  return [optionalProperties containsObject:propertyName];
}

@end
