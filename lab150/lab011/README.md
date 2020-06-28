#

## NOTICE
 - `ctx.GetCurrentRoute().StaticPath()`，这个方法会返回不带{}的restful api
 - `ctx.GetCurrentRoute().ResolvePath()`，可以得到有%v的path，也可以设置参数替换