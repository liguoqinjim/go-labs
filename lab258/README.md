# [gen](https://github.com/smallnest/gen) 
安装：`go get -u github.com/smallnest/gen`

## NOTICE
 - `--model=model2`，生成的文件夹名称为model2
 - `model_naming`，生成的struct的name
 - 

## mysql
```
gen --sqltype=mysql --connstr=root:$MYSQL_PASSWORD@tcp\(host:3306\)/research_report --database=research_report --out=. --model=model2 --gorm --json --overwrite --file_naming="{{toUpper .}}" --model_naming="{{ toSnakeCase ( replace . \"t_\" \"\") }}"
```