# 删除文件夹

## NOTICE
 - os.Remove只能删除为空的文件夹，不为空的文件夹需要RemoveAll来删除

## 测试脚本
```shell script
mkdir -p data/1/2
touch data/1/2/3.txt
```

