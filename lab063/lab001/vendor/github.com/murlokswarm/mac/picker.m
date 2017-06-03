#include "picker.h"
#include "_cgo_export.h"
#include "driver.h"

void Picker_NewFilePicker(FilePicker__ p) {
  NSOpenPanel *panel = [NSOpenPanel openPanel];
  [panel setAllowsMultipleSelection:p.MultipleSelection];
  [panel setCanChooseDirectories:!p.NoDir];
  [panel setCanChooseFiles:!p.NoFile];

  NSString *ID = [NSString stringWithUTF8String:p.ID];

  NSWindow *currentWindow = NSApp.keyWindow;
  if (currentWindow == nil) {
    defer([panel beginWithCompletionHandler:^(NSInteger result) {
            Picker_FilePickerClosed(ID, result, panel.URLs);
          }];);
    return;
  }

  defer([panel beginSheetModalForWindow:currentWindow
                      completionHandler:^(NSInteger result) {
                        Picker_FilePickerClosed(ID, result, panel.URLs);
                      }];);
}

void Picker_FilePickerClosed(NSString *ID, NSInteger result,
                             NSArray<NSURL *> *URLs) {

  NSMutableArray<NSString *> *filenames = [[NSMutableArray alloc] init];

  if (result == NSFileHandlingPanelOKButton) {
    for (NSURL *url in URLs) {
      [filenames addObject:url.path];
    }
  }

  NSData *jsonData =
      [NSJSONSerialization dataWithJSONObject:filenames options:0 error:nil];
  NSString *jsonString =
      [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];

  onFilePickerClosed((char *)ID.UTF8String, (char *)jsonString.UTF8String);
}