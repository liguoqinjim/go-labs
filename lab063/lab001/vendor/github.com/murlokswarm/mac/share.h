#ifndef share_h
#define share_h

#import <Cocoa/Cocoa.h>

void Share_Text(const char *cv);
void Share_URL(const char *cv);
void Share_Display(NSArray *items);

#endif /* share_h */