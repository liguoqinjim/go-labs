#include "driver.h"
#include "_cgo_export.h"
#include "menu.h"

@implementation DriverDelegate
- (instancetype)init {
  self.dock = [[NSMenu alloc] initWithTitle:@""];
  return self;
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
  onLaunch();
}

- (void)applicationDidBecomeActive:(NSNotification *)aNotification {
  onFocus();
}

- (void)applicationDidResignActive:(NSNotification *)aNotification {
  onBlur();
}

- (BOOL)applicationShouldHandleReopen:(NSApplication *)sender
                    hasVisibleWindows:(BOOL)flag {
  onReopen();
  return YES;
}

- (void)application:(NSApplication *)sender
          openFiles:(NSArray<NSString *> *)filenames {
  NSData *jsonData =
      [NSJSONSerialization dataWithJSONObject:filenames options:0 error:nil];
  NSString *jsonString =
      [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];

  onFilesOpen((char *)jsonString.UTF8String);
}

- (void)applicationWillFinishLaunching:(NSNotification *)aNotification {
  NSAppleEventManager *appleEventManager =
      [NSAppleEventManager sharedAppleEventManager];
  [appleEventManager
      setEventHandler:self
          andSelector:@selector(handleGetURLEvent:withReplyEvent:)
        forEventClass:kInternetEventClass
           andEventID:kAEGetURL];
}

- (void)handleGetURLEvent:(NSAppleEventDescriptor *)event
           withReplyEvent:(NSAppleEventDescriptor *)replyEvent {
  NSURL *url =
      [NSURL URLWithString:[[event paramDescriptorForKeyword:keyDirectObject]
                               stringValue]];
  onURLOpen((char *)url.absoluteString.UTF8String);
}

- (NSApplicationTerminateReply)applicationShouldTerminate:
    (NSApplication *)sender {
  return onTerminate();
}

- (void)applicationWillTerminate:(NSNotification *)aNotification {
  onFinalize();
}

- (NSMenu *)applicationDockMenu:(NSApplication *)sender {
  return self.dock;
}
@end

void Driver_Run() {
  [NSApplication sharedApplication];
  [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];

  DriverDelegate *delegate = [[DriverDelegate alloc] init];
  NSApp.delegate = delegate;

  [NSApp run];
}

void Driver_Terminate() { defer([NSApp terminate:NSApp];); }

void Driver_SetMenuBar(const void *menuPtr) {
  Menu *menu = (__bridge Menu *)menuPtr;

  defer(NSApp.mainMenu = menu.Root; [NSApp activateIgnoringOtherApps:YES];);
}

void Driver_SetDockMenu(const void *dockPtr) {
  Menu *menu = (__bridge Menu *)dockPtr;

  defer(DriverDelegate *delegate = NSApp.delegate; delegate.dock = menu.Root;);
}

void Driver_SetDockIcon(const char *path) {
  NSString *p = [NSString stringWithUTF8String:path];

  defer(if (p.length != 0) {
    NSApp.applicationIconImage = [[NSImage alloc] initByReferencingFile:p];
    return;
  } NSApp.applicationIconImage = nil;);
}

void Driver_SetDockBadge(const char *str) {
  NSString *badge = [NSString stringWithUTF8String:str];
  defer([NSApp.dockTile setBadgeLabel:badge];);
}

void Driver_ShowContextMenu(const void *menuPtr) {
  Menu *menu = (__bridge Menu *)menuPtr;

  defer(if (NSApp.keyWindow == nil) { return; }

        NSPoint p = [NSApp.keyWindow mouseLocationOutsideOfEventStream];
        [menu.Root popUpMenuPositioningItem:menu.Root.itemArray[0]
                                 atLocation:p
                                     inView:NSApp.keyWindow.contentView];);
}