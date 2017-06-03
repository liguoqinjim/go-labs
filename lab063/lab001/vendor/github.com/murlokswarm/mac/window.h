#ifndef window_h
#define window_h

#import <Cocoa/Cocoa.h>
#import <WebKit/WebKit.h>

typedef struct Window__ {
  const char *ID;
  const char *Title;
  CGFloat X;
  CGFloat Y;
  CGFloat Width;
  CGFloat Height;
  CGFloat MinWidth;
  CGFloat MinHeight;
  CGFloat MaxWidth;
  CGFloat MaxHeight;
  const char *BackgroundColor;
  NSVisualEffectMaterial Vibrancy;
  BOOL Borderless;
  BOOL FixedSize;
  BOOL CloseHidden;
  BOOL MinimizeHidden;
  BOOL TitlebarHidden;
  const char *HTML;
  const char *ResourcePath;
} Window__;

@interface WindowController
    : NSWindowController <NSWindowDelegate, WKNavigationDelegate, WKUIDelegate,
                          WKScriptMessageHandler>
@property NSString *ID;
@property(weak) WKWebView *webview;

- (instancetype)initWithID:(NSString *)ID;
@end

@interface TitleBar : NSView
@end

void Window_New(Window__ w);
void Window_new(Window__ w);
WKWebView *Window_NewWebview(WindowController *controller, NSString *HTML,
                             NSString *resourcePath);
void Window_SetWebview(NSWindow *win, WKWebView *webview);
void Window_SetTitleBar(NSWindow *win, TitleBar *titleBar);
void Window_Show(const void *ptr);
void Window_CallJS(const void *ptr, const char *js);
NSRect Window_Frame(const void *ptr);
void Window_Move(const void *ptr, CGFloat x, CGFloat y);
void Window_Resize(const void *ptr, CGFloat width, CGFloat height);
void Window_Close(const void *ptr);

#endif /* window_h */
