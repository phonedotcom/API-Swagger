#import "SWGCreateExtensionParams.h"

@implementation SWGCreateExtensionParams

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"callerId": @"caller_id", @"usageType": @"usage_type", @"allowsCallWaiting": @"allows_call_waiting", @"extension": @"extension", @"includeInDirectory": @"include_in_directory", @"name": @"name", @"fullName": @"full_name", @"timezone": @"timezone", @"nameGreeting": @"name_greeting", @"voicemailGreetingAlternate": @"voicemail[greeting][alternate]", @"localAreaCode": @"local_area_code", @"voicemailGreetingEnableLeaveMessagePrompt": @"voicemail[greeting][enable_leave_message_prompt]", @"voicemailEnabled": @"voicemail[enabled]", @"enableOutboundCalls": @"enable_outbound_calls", @"enableCallWaiting": @"enable_call_waiting", @"voicemailPassword": @"voicemail[password]", @"voicemailGreetingType": @"voicemail[greeting][type]", @"voicemailGreetingStandard": @"voicemail[greeting][standard]", @"voicemailTranscription": @"voicemail[transcription]", @"voicemailNotificationsEmails": @"voicemail[notifications][emails]", @"voicemailNotificationsSms": @"voicemail[notifications][sms]", @"callNotificationsEmails": @"call_notifications[emails]", @"callNotificationsSms": @"call_notifications[sms]" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"callerId", @"usageType", @"allowsCallWaiting", @"extension", @"includeInDirectory", @"name", @"fullName", @"timezone", @"nameGreeting", @"voicemailGreetingAlternate", @"localAreaCode", @"voicemailGreetingEnableLeaveMessagePrompt", @"voicemailEnabled", @"enableOutboundCalls", @"enableCallWaiting", @"voicemailPassword", @"voicemailGreetingType", @"voicemailGreetingStandard", @"voicemailTranscription", @"voicemailNotificationsEmails", @"voicemailNotificationsSms", @"callNotificationsEmails", @"callNotificationsSms"];
  return [optionalProperties containsObject:propertyName];
}

@end
