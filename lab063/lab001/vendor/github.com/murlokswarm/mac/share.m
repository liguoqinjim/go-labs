#include "share.h"
#include "driver.h"
#include "window.h"

void Share_Text(const char *cv) {
  NSString *v = [NSString stringWithUTF8String:cv];
  NSArray *items = @[ v ];
  Share_Display(items);
}

void Share_URL(const char *cv) {
  NSString *v = [NSString stringWithUTF8String:cv];
  NSURL *url = [NSURL URLWithString:v];
  NSArray *items = @[ url ];
  Share_Display(items);
}

void Share_Display(NSArray *items) {
  defer(WindowController *controller =
            (WindowController *)NSApp.keyWindow.windowController;
        WKWebView *webview = controller.webview;

        NSPoint p = [NSApp.keyWindow mouseLocationOutsideOfEventStream];
        p = [webview convertPoint:p fromView:NSApp.keyWindow.contentView];
        NSRect r = NSMakeRect(p.x, p.y, 1, 1);

        NSSharingServicePicker *picker =
            [[NSSharingServicePicker alloc] initWithItems:items];
        [picker showRelativeToRect:r ofView:webview preferredEdge:NSMinYEdge];);
}
