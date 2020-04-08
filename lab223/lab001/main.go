package main

import (
	"github.com/dustin/go-humanize"
	"github.com/dustin/go-humanize/english"
	"log"
	"time"
)

func main() {
	//文件大小
	log.Printf("That file is %s.", humanize.Bytes(82854982)) // That file is 83 MB.

	//时间
	log.Printf("This was touched %s.", humanize.Time(time.Now().Add(time.Hour+time.Minute*15))) // This was touched 7 hours ago.

	//复数
	log.Println(english.PluralWord(1, "object", ""))  // object
	log.Println(english.PluralWord(42, "object", "")) // objects
}
