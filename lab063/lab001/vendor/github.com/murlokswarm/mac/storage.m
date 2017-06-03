#include "storage.h"

const char *Storage_Resources() {
  NSBundle *mainBundle = [NSBundle mainBundle];
  return mainBundle.resourcePath.UTF8String;
}

const char *Storage_Home() { return NSHomeDirectory().UTF8String; }

const char *Storage_Support() {
  NSArray *paths = NSSearchPathForDirectoriesInDomains(
      NSApplicationSupportDirectory, NSUserDomainMask, YES);
  NSString *applicationSupportDirectory = [paths firstObject];
  return applicationSupportDirectory.UTF8String;
}

const char *Storage_BundleID() {
  NSBundle *mainBundle = [NSBundle mainBundle];
  return mainBundle.bundleIdentifier.UTF8String;
}