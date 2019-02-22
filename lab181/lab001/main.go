package main

import (
	"context"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
	"log"
)

var (
	tableName = "FileTable"
)

func main() {
	example()
}

func example() {
	client := gohbase.NewClient("hp-111:2181")

	//得到完整的一行数据
	log.Println("得到完整的一行数据")
	getRequest, err := hrpc.NewGetStr(context.Background(), tableName, "rowkey1")
	if err != nil {
		log.Fatalf("hrpc.NewGetStr error:%v", err)
	}
	getRsp, err := client.Get(getRequest)

	for _, v := range getRsp.Cells {
		log.Printf("%s:%s", v.Qualifier, v.Value)
	}

	//得到指定的qualifier
	log.Println("得到指定的qualifier")
	family := map[string][]string{"fileInfo": {"name", "type"}}
	getRequest, err = hrpc.NewGetStr(context.Background(), tableName, "rowkey1", hrpc.Families(family))
	if err != nil {
		log.Fatalf("hrpc.NewGetStr error:%v", err)
	}
	getRsp, err = client.Get(getRequest)
	for _, v := range getRsp.Cells {
		log.Printf("%s:%s", v.Qualifier, v.Value)
	}

	//scan操作
	log.Println("scan操作")
	scanRequest, err := hrpc.NewScanStr(context.Background(), tableName)
	if err != nil {
		log.Fatalf("hrpc.NewScanStr error:%v", err)
	}
	scanRsp := client.Scan(scanRequest)

	r, err := scanRsp.Next()
	for ; err == nil; r, err = scanRsp.Next() {
		log.Println(r, err)
	}

	//scan操作，设置filter
	log.Println("scan操作，设置filter")
	pFilter := filter.NewPrefixFilter([]byte("row"))
	scanRequest, err = hrpc.NewScanStr(context.Background(), tableName, hrpc.Filters(pFilter))
	if err != nil {
		log.Fatalf("hrpc.NewScanStr error:%v", err)
	}
	scanRsp = client.Scan(scanRequest)

	r, err = scanRsp.Next()
	for ; err == nil; r, err = scanRsp.Next() {
		log.Println(r, err)
	}
}
