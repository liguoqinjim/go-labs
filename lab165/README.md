# [bloom](https://github.com/willf/bloom)

|实验|简介|说明|
|---|---|---|
|lab001|example| |
|lab002|demo|参数说明|
|lab003|demo|数据持久化|

## NOTICE
 - bloom filter中，n是多少个元素，m是多少bit(单位就是bit)，k是多少个hash。New的时候调用的是m而不是n,NewWithEstimates调用的是n
 - n较大的时候，比如20000000，分配内存会花些时间，mac上的话大概10秒
 - 持久化之后的文件大小，就是m的大小转换到mb就可以了。