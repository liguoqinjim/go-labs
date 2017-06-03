#ifndef color_h
#define color_h

#import <QuartzCore/QuartzCore.h>

@interface CIColor (MBCategory)
+ (CIColor *)colorWithHex:(UInt32)col;
+ (CIColor *)colorWithHexString:(NSString *)str;
@end

#endif /* color_h */