# [viper](https://github.com/spf13/viper)

|实验|简介|说明|
|---|---|---|
|lab001|example|yaml |
|lab002|toml| |
|lab003|环境变量| |
|lab004|flag| |
|lab005|ini| |

## NOTICE
 - 当调用了`v.SetConfigType("ini")`之后，viper还是会读取别的配置文件的，这可能会导致错。
 比如路径下有`config.ini`和`config.json`，会报错`While parsing config: key-value delimiter not found: {`，把`config.json`删除了就可以了