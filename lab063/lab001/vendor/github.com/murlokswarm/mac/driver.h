#ifndef driver_h
#define driver_h

#import <Cocoa/Cocoa.h>

#define defer(code)                                                            \
  dispatch_async(dispatch_get_main_queue(), ^{                                 \
                     code})

@interface DriverDelegate : NSObject <NSApplicationDelegate>
@property NSMenu *dock;

- (instancetype)init;
@end

void Driver_Run();
void Driver_Terminate();
void Driver_SetMenuBar(const void *menuPtr);
void Driver_SetDockMenu(const void *dockPtr);
void Driver_SetDockIcon(const char *path);
void Driver_SetDockBadge(const char *str);
void Driver_ShowContextMenu(const void *menuPtr);



#endif /* driver_h */