# mysql的json数据格式

建表语句
```
CREATE TABLE `t_data_json` (
  `id` int(11) DEFAULT NULL,
  `data` json DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8
```