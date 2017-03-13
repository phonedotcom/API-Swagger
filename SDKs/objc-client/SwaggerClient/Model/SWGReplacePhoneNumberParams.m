#import "SWGReplacePhoneNumberParams.h"

@implementation SWGReplacePhoneNumberParams

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"route": @"route", @"name": @"name", @"blockIncoming": @"block_incoming", @"blockAnonymous": @"block_anonymous", @"callerIdName": @"caller_id[name]", @"callerIdType": @"caller_id[type]", @"smsForwardingType": @"sms_forwarding[type]", @"smsForwardingApplication": @"sms_forwarding[application]", @"smsForwardingExtension": @"sms_forwarding[extension]", @"poolItem": @"pool_item", @"callNotificationsEmails": @"call_notifications[emails]", @"callNotificationsSms": @"call_notifications[sms]" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"route", @"name", @"blockIncoming", @"blockAnonymous", @"callerIdName", @"callerIdType", @"smsForwardingType", @"smsForwardingApplication", @"smsForwardingExtension", @"poolItem", @"callNotificationsEmails", @"callNotificationsSms"];
  return [optionalProperties containsObject:propertyName];
}

@end
