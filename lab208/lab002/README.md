# toml
读取toml配置文件

## NOTICE
 - `v.SetDefault`设置默认值
 - 注意，go-toml库中下划线使用*`toml:"connection_max"`*这种格式，但是viper中使用*`mapstructure:"connection_max"`*，也就是不能使用库本身的。
 - array的例子可以看`lab213/lab002`
 
## 参考资料
 -  https://github.com/spf13/viper#unmarshaling