package main

import (
	"github.com/henrylee2cn/pholcus/exec"
	//_ "github.com/pholcus/spider_lib"
	//_ "github.com/liguoqinjim/pholcus_lib"
	//_ "goSpider/demo1/spider"

	_ "github.com/liguoqinjim/pholcus_libs_private"
)

func main() {
	exec.DefaultRun("web")
}
