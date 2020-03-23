# [fsnotify](https://github.com/fsnotify/fsnotify)

|实验|简介|说明|
|---|---|---|
|lab001|example| |
|lab002|自己fork的库，有linux的CloseWrite|监控ftp|

## NOTICE
 - 这个库现在还不能监控linux上的write_close，因为这个库是想做跨平台的，windows上没有write_close。
 所以要在linux上监控，要修改下源代码。具体看`https://github.com/fsnotify/fsnotify/pull/313`
 - linux上的有linux的CloseWrite使用自己修改的库，`https://github.com/liguoqinjim/fsnotify/v1`