#import "SWGReplaceExtensionParams.h"

@implementation SWGReplaceExtensionParams

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"voicemailGreetingAlternate": @"voicemail[greeting][alternate]", @"nameGreeting": @"name_greeting", @"name": @"name", @"timezone": @"timezone", @"includeInDirectory": @"include_in_directory", @"extension": @"extension", @"enableOutboundCalls": @"enable_outbound_calls", @"usageType": @"usage_type", @"voicemailPassword": @"voicemail[password]", @"fullName": @"full_name", @"enableCallWaiting": @"enable_call_waiting", @"voicemailGreetingStandard": @"voicemail[greeting][standard]", @"voicemailGreetingType": @"voicemail[greeting][type]", @"callerId": @"caller_id", @"localAreaCode": @"local_area_code", @"voicemailEnabled": @"voicemail[enabled]", @"voicemailGreetingEnableLeaveMessagePrompt": @"voicemail[greeting][enable_leave_message_prompt]", @"voicemailTranscription": @"voicemail[transcription]", @"voicemailNotificationsEmails": @"voicemail[notifications][emails]", @"voicemailNotificationsSms": @"voicemail[notifications][sms]", @"callNotificationsEmails": @"call_notifications[emails]", @"callNotificationsSms": @"call_notifications[sms]", @"route": @"route" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"voicemailGreetingAlternate", @"nameGreeting", @"name", @"timezone", @"includeInDirectory", @"extension", @"enableOutboundCalls", @"usageType", @"voicemailPassword", @"fullName", @"enableCallWaiting", @"voicemailGreetingStandard", @"voicemailGreetingType", @"callerId", @"localAreaCode", @"voicemailEnabled", @"voicemailGreetingEnableLeaveMessagePrompt", @"voicemailTranscription", @"voicemailNotificationsEmails", @"voicemailNotificationsSms", @"callNotificationsEmails", @"callNotificationsSms", @"route"];
  return [optionalProperties containsObject:propertyName];
}

@end
