#include "color.h"

@implementation CIColor (MBCategory)
+ (CIColor *)colorWithHexString:(NSString *)str {
  const char *cstr = [str cStringUsingEncoding:NSASCIIStringEncoding];
  long x = strtol(cstr + 1, NULL, 16);
  return [CIColor colorWithHex:x];
}

+ (CIColor *)colorWithHex:(UInt32)col {
  unsigned char b = col & 0xFF;
  unsigned char g = (col >> 8) & 0xFF;
  unsigned char r = (col >> 16) & 0xFF;

  return [CIColor colorWithRed:(float)r / 255.0f
                         green:(float)g / 255.0f
                          blue:(float)b / 255.0f
                         alpha:1];
}
@end