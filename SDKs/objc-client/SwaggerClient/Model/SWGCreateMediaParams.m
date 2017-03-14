#import "SWGCreateMediaParams.h"

@implementation SWGCreateMediaParams

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
  return [[JSONKeyMapper alloc] initWithModelToJSONDictionary:@{ @"name": @"name", @"origin": @"origin", @"type": @"type", @"ttsVoice": @"tts_voice", @"ttsText": @"tts_text", @"isTemparary": @"is_temparary", @"expirationDate": @"expiration_date", @"duration": @"duration", @"notes": @"notes", @"randomized": @"randomized" }];
}

/**
 * Indicates whether the property with the given name is optional.
 * If `propertyName` is optional, then return `YES`, otherwise return `NO`.
 * This method is used by `JSONModel`.
 */
+ (BOOL)propertyIsOptional:(NSString *)propertyName {

  NSArray *optionalProperties = @[@"name", @"origin", @"type", @"ttsVoice", @"ttsText", @"isTemparary", @"expirationDate", @"duration", @"notes", @"randomized"];
  return [optionalProperties containsObject:propertyName];
}

@end
