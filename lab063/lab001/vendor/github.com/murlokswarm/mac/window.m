#include "window.h"
#include "_cgo_export.h"
#include "color.h"

void Window_New(Window__ w) { defer(Window_new(w);); }

void Window_new(Window__ w) {
  NSRect contentRect = NSMakeRect(w.X, w.Y, w.Width, w.Height);
  NSUInteger styleMask =
      NSWindowStyleMaskTitled | NSWindowStyleMaskFullSizeContentView |
      NSWindowStyleMaskClosable | NSWindowStyleMaskMiniaturizable |
      NSWindowStyleMaskResizable;

  if (w.Borderless) {
    styleMask = styleMask & NSWindowStyleMaskBorderless;
  }

  if (w.FixedSize) {
    styleMask = styleMask & ~NSWindowStyleMaskResizable;
  }

  if (w.CloseHidden) {
    styleMask = styleMask & ~NSWindowStyleMaskClosable;
  }

  if (w.MinimizeHidden) {
    styleMask = styleMask & ~NSWindowStyleMaskMiniaturizable;
  }

  NSWindow *win = [[NSWindow alloc] initWithContentRect:contentRect
                                              styleMask:styleMask
                                                backing:NSBackingStoreBuffered
                                                  defer:NO];

  // size.
  NSSize minSize = NSMakeSize(w.MinWidth, w.MinHeight);
  NSSize maxSize = NSMakeSize(w.MaxWidth, w.MaxHeight);
  win.minSize = minSize;
  win.maxSize = maxSize;

  // Background.
  if (w.Vibrancy != NSVisualEffectMaterialAppearanceBased) {
    NSVisualEffectView *visualEffectView =
        [[NSVisualEffectView alloc] initWithFrame:contentRect];

    visualEffectView.material = w.Vibrancy;
    visualEffectView.blendingMode = NSVisualEffectBlendingModeBehindWindow;
    visualEffectView.state = NSVisualEffectStateActive;
    win.contentView = visualEffectView;
  } else {
    NSString *bacgroundColorString =
        [NSString stringWithUTF8String:w.BackgroundColor];

    if (bacgroundColorString.length != 0) {
      CIColor *backgroundColor =
          [CIColor colorWithHexString:bacgroundColorString];
      win.backgroundColor = [NSColor colorWithCIColor:backgroundColor];
    }
  }

  // Window controller.
  NSString *id = [NSString stringWithUTF8String:w.ID];
  WindowController *controller = [[WindowController alloc] initWithID:id];
  controller.window = win;
  win.delegate = controller;
  win.windowController = controller;
  win.windowController.windowFrameAutosaveName =
      [NSString stringWithUTF8String:w.Title];

  // WebView.
  WKWebView *webview =
      Window_NewWebview(controller, [NSString stringWithUTF8String:w.HTML],
                        [NSString stringWithUTF8String:w.ResourcePath]);
  Window_SetWebview(win, webview);
  controller.webview = webview;

  // Titlebar.
  if (w.TitlebarHidden) {
    win.titlebarAppearsTransparent = true;
    TitleBar *titleBar = [[TitleBar alloc] init];
    Window_SetTitleBar(win, titleBar);
  } else {
    win.title = [NSString stringWithUTF8String:w.Title];
  }

  onWindowCreated((void *)CFBridgingRetain(win));
}

WKWebView *Window_NewWebview(WindowController *controller, NSString *HTML,
                             NSString *resourcePath) {
  WKUserContentController *userContentController =
      [[WKUserContentController alloc] init];
  [userContentController addScriptMessageHandler:controller name:@"Call"];

  WKWebViewConfiguration *conf = [[WKWebViewConfiguration alloc] init];
  conf.userContentController = userContentController;

  WKWebView *webview = [[WKWebView alloc] initWithFrame:NSMakeRect(0, 0, 0, 0)
                                          configuration:conf];
  [webview setValue:@(NO) forKey:@"drawsBackground"];
  webview.navigationDelegate = controller;
  webview.UIDelegate = controller;

  // Page loading.
  NSURL *baseURL = [NSURL fileURLWithPath:resourcePath];
  [webview loadHTMLString:HTML baseURL:baseURL];
  return webview;
}

void Window_SetWebview(NSWindow *win, WKWebView *webview) {
  webview.translatesAutoresizingMaskIntoConstraints = NO;
  [win.contentView addSubview:webview];

  [win.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"|[webview]|"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              webview)]];
  [win.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"V:|[webview]|"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              webview)]];
}

void Window_SetTitleBar(NSWindow *win, TitleBar *titleBar) {
  titleBar.translatesAutoresizingMaskIntoConstraints = NO;

  [win.contentView addSubview:titleBar];
  [win.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"|[titleBar]|"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              titleBar)]];
  [win.contentView
      addConstraints:
          [NSLayoutConstraint
              constraintsWithVisualFormat:@"V:|[titleBar(==22)]"
                                  options:0
                                  metrics:nil
                                    views:NSDictionaryOfVariableBindings(
                                              titleBar)]];
}

void Window_Show(const void *ptr) {
  NSWindow *win = (__bridge NSWindow *)ptr;
  defer([win makeKeyAndOrderFront:nil];);
}

void Window_CallJS(const void *ptr, const char *js) {
  NSWindow *win = (__bridge NSWindow *)ptr;
  WindowController *controller = (WindowController *)win.windowController;

  NSString *javaScript = [NSString stringWithUTF8String:js];
  [controller.webview evaluateJavaScript:javaScript completionHandler:nil];
}

NSRect Window_Frame(const void *ptr) {
  NSWindow *win = (__bridge NSWindow *)ptr;
  return win.frame;
}

void Window_Move(const void *ptr, CGFloat x, CGFloat y) {
  NSWindow *win = (__bridge NSWindow *)ptr;
  CGPoint pos = NSMakePoint(x, y);

  defer([win setFrameOrigin:pos];);
}

void Window_Resize(const void *ptr, CGFloat width, CGFloat height) {
  NSWindow *win = (__bridge NSWindow *)ptr;
  CGRect frame = win.frame;
  frame.size.width = width;
  frame.size.height = height;

  defer([win setFrame:frame display:YES];);
}

void Window_Close(const void *ptr) {
  NSWindow *win = (__bridge NSWindow *)ptr;
  defer([win performClose:nil];);
}

@implementation WindowController
- (instancetype)initWithID:(NSString *)ID {
  self.ID = ID;
  return self;
}

- (void)webView:(WKWebView *)webView
    didFinishNavigation:(WKNavigation *)navigation {
  onWindowWebviewLoaded();
}

- (void)userContentController:(WKUserContentController *)userContentController
      didReceiveScriptMessage:(WKScriptMessage *)message {
  if ([message.name isEqual:@"Call"]) {
    NSString *msg = (NSString *)message.body;
    onJSCall((char *)msg.UTF8String);
  }
}

- (void)webView:(WKWebView *)webView
    decidePolicyForNavigationAction:(WKNavigationAction *)navigationAction
                    decisionHandler:
                        (void (^)(WKNavigationActionPolicy))decisionHandler {
  if (navigationAction.navigationType == WKNavigationTypeReload ||
      navigationAction.navigationType == WKNavigationTypeOther) {
    if (navigationAction.targetFrame.request != nil) {
      decisionHandler(WKNavigationActionPolicyCancel);
      return;
    }

    decisionHandler(WKNavigationActionPolicyAllow);
    return;
  }

  NSURL *url = navigationAction.request.URL;
  onWindowWebviewNavigate((char *)self.ID.UTF8String,
                          (char *)url.absoluteString.UTF8String);
  decisionHandler(WKNavigationActionPolicyCancel);
}

- (void)webView:(WKWebView *)webView
    runJavaScriptAlertPanelWithMessage:(NSString *)message
                      initiatedByFrame:(WKFrameInfo *)frame
                     completionHandler:(void (^)(void))completionHandler {
  onJSAlert((char *)message.UTF8String);
  completionHandler();
}

- (void)windowDidMiniaturize:(NSNotification *)notification {
  onWindowMinimize((char *)self.ID.UTF8String);
}

- (void)windowDidDeminiaturize:(NSNotification *)notification {
  onWindowDeminimize((char *)self.ID.UTF8String);
}

- (void)windowDidEnterFullScreen:(NSNotification *)notification {
  onWindowFullScreen((char *)self.ID.UTF8String);
}

- (void)windowDidExitFullScreen:(NSNotification *)notification {
  onWindowExitFullScreen((char *)self.ID.UTF8String);
}

- (void)windowDidMove:(NSNotification *)notification {
  onWindowMove((char *)self.ID.UTF8String, self.window.frame.origin.x,
               self.window.frame.origin.y);
}

- (void)windowDidResize:(NSNotification *)notification {
  onWindowResize((char *)self.ID.UTF8String, self.window.frame.size.width,
                 self.window.frame.size.height);
}

- (void)windowDidBecomeKey:(NSNotification *)notification {
  onWindowFocus((char *)self.ID.UTF8String);
}

- (void)windowDidResignKey:(NSNotification *)notification {
  onWindowBlur((char *)self.ID.UTF8String);
}

- (BOOL)windowShouldClose:(id)sender {
  return onWindowClose((char *)self.ID.UTF8String);
}

- (void)windowWillClose:(NSNotification *)notification {
  onWindowCloseFinal((char *)self.ID.UTF8String);
  CFBridgingRelease((__bridge void *)self.window);
  self.window = nil;
}
@end

@implementation TitleBar
- (void)mouseDragged:(nonnull NSEvent *)theEvent {
  [self.window performWindowDragWithEvent:theEvent];
}

- (void)mouseUp:(NSEvent *)event {
  WindowController *controller =
      (WindowController *)self.window.windowController;
  [controller.webview mouseUp:event];

  if (2 == event.clickCount) {
    [self.window zoom:nil];
  }
}
@end