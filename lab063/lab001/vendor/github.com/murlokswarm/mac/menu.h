#ifndef menu_h
#define menu_h

#import <Cocoa/Cocoa.h>

typedef struct Menu__ { const char *ID; } Menu__;

typedef struct MenuContainer__ {
  const char *ID;
  const char *Label;
} MenuContainer__;

typedef struct MenuItem__ {
  const char *ID;
  const char *Label;
  const char *Icon;
  const char *Shortcut;
  const char *Selector;
  const char *OnClick;
  BOOL Disabled;
  BOOL Separator;
} MenuItem__;

@interface MenuContainer : NSMenu
@property NSString *ID;
@end

@interface MenuItem : NSMenuItem
@property NSString *ID;
@property NSString *OnClick;
@property BOOL IsSeparator;
@property NSMenuItem *SeparatorItem;

- (void)setSelector:(NSString *)selector;
- (void)setShortcut:(NSString *)shortcut;
- (void)setSeparator;
- (void)clicked:(id)sender;
@end

@interface Menu : NSObject <NSMenuDelegate>
@property NSString *ID;
@property NSMutableDictionary *Elems;
@property MenuContainer *Root;

- (void)dismountElement:(id)elem;
@end

const void *Menu_New(Menu__ m);
void Menu_Mount(const void *ptr, const char *rootID);
void Menu_Show(const void *ptr);
void Menu_Dismount(const void *ptr);
void Menu_MountContainer(const void *ptr, MenuContainer__ container);
void Menu_MountItem(const void *ptr, MenuItem__ item);
void Menu_Associate(const void *ptr, const char *parentID, const char *childID);
void Menu_Clear(const void *ptr);

#endif /* menu_h */