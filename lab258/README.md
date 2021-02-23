# [gen](https://github.com/smallnest/gen) 
安装：`go get -u github.com/smallnest/gen`

## NOTICE
 - `--model=model2`，生成的文件夹名称为model2
 - `model_naming`，生成的struct的name
 - varchar类型要是允许null的话，默认生成类型会是`sql.NullString`，需要修改的话，要自己修改mapping。参考mysql的自定义mapping

## 使用
### Mysql
```
基本配置
gen --sqltype=mysql --connstr="root:$MYSQL_PASSWORD@tcp($MYSQL_HOST:3306)/research_report" --database=research_report --out=. --model=model2 --gorm --json --overwrite --file_naming="{{toUpper .}}" --model_naming="{{ toSnakeCase ( replace . \"t_\" \"\") }}"

自定义mapping
gen --sqltype=mysql --connstr="root:$MYSQL_PASSWORD@tcp(txcent001:3306)/research_report" --database=research_report --out=. --model=model2 --gorm --json --overwrite --file_naming="{{( replace . \"t_\" \"\" 0)}}" --model_naming="{{ toSnakeCase ( replace . \"t_\" \"\") }}" --mapping="extra.json"
```

### template
 - 保存：`gen --save=./templates`

