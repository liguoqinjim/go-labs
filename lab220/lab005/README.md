# 批量设置数据

## NOTICE
 - 得到列名:`excelize.ToAlphaString`，index从0开始
 - xlsx.SetSheetRow，最后一个参数要是pointer，例：`xlsx.SetSheetRow("Sheet1", "B6", &[]interface{}{"1", nil, 2})`。这个需要注意，因为这个函数没有返回error，不是pointer的话只是xslx里面没数据，不会报错。

## 参考资料
 - https://github.com/360EntSecGroup-Skylar/excelize/issues/364