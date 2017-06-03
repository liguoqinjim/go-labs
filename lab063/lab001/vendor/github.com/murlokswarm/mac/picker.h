#ifndef picker_h
#define picker_h

#import <Cocoa/Cocoa.h>

typedef struct FilePicker__ {
  const char *ID;
  BOOL MultipleSelection;
  BOOL NoDir;
  BOOL NoFile;
} FilePicker__;

void Picker_NewFilePicker(FilePicker__ p);
void Picker_FilePickerClosed(NSString *ID, NSInteger result,
                             NSArray<NSURL *> *URLs);

#endif /* picker_h */