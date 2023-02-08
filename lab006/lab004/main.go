package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

// 递归获取指定目录下的所有文件名
func GetAllFile(pathname string) ([]string, error) {
	result := []string{}

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := GetAllFile(fullname)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err
			}
			result = append(result, temp...)
		} else {
			result = append(result, fullname)
		}
	}

	return result, nil
}

// 把秒级的时间戳转为time格式
func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func main() {
	// 递归获取目录下的所有文件
	var files []string
	files, _ = GetAllFile("/Users/li/Workspace/go-labs/lab006")

	fmt.Println("目录下的所有文件如下")
	for i := 0; i < len(files); i++ {
		fmt.Println("文件名：", files[i])

		// 获取文件原来的访问时间，修改时间
		finfo, _ := os.Stat(files[i])

		// darwin环境下代码如下
		linuxFileAttr := finfo.Sys().(*syscall.Stat_t)
		fmt.Println("文件创建时间", SecondToTime(linuxFileAttr.Ctimespec.Sec))
		fmt.Println("最后访问时间", SecondToTime(linuxFileAttr.Atimespec.Sec))
		fmt.Println("最后修改时间", SecondToTime(linuxFileAttr.Mtimespec.Sec))

		//linux
		//参考 go-tools/tool093

		// windows下代码如下
		//winFileAttr := finfo.Sys().(*syscall.Win32FileAttributeData)
		//fmt.Println("文件创建时间：", SecondToTime(winFileAttr.CreationTime.Nanoseconds()/1e9))
		//fmt.Println("最后访问时间：", SecondToTime(winFileAttr.LastAccessTime.Nanoseconds()/1e9))
		//fmt.Println("最后修改时间：", SecondToTime(winFileAttr.LastWriteTime.Nanoseconds()/1e9))
	}
}
