#import "SWGCallDetails.h"

@implementation SWGCallDetails

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"startTime": @"start_time", @"type": @"type", @"idValue": @"id_value", @"voipId": @"voip_id", @"voipPhoneId": @"voip_phone_id" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"startTime", @"type", @"idValue", @"voipId", @"voipPhoneId"];
  return [optionalProperties containsObject:propertyName];
}

@end
