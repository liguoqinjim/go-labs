# 捕获错误

|文件|简介|说明|
|---|---|---|
|lab001| | |
|lab002|panic的坑| |

## NOTICE
 - golang不能捕获到os.Exit，那也就是捕获不到log.Fatal()，因为fatal的底层最后是调用的exit来退出
 - recover只能捕获当前goroutine的panic