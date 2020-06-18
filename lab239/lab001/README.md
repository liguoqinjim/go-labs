# example

## NOTICE
 - `sentry.Flush(time.Second*20)`，这个方法是把错误上传到sentry的。要是返回false的时候，有可能某些时间还没有上传好。true的时候才是对的

