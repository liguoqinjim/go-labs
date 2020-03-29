# 生成数据
这个实验里面的方法可以比较方便的一次设置一行数据

## NOTICE
 - 在excel里面，Z列的后面是AA列
 - `xlsx.SetCellValue("Sheet1", fmt.Sprintf("%c%d", FIRST_COLUMN+n, row), v)`，
为了比较方便的生成A1，B1这样的坐标，使用了rune类型数据来表示列明，然后在fmt里面使用%c，来把rune代表的数字转到对应的字符

## 参考资料
 - https://golang.org/pkg/fmt/

