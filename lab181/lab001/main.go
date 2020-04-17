package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
	"github.com/tsuna/gohbase/hrpc"
	"log"
)

var (
	tableName = "FileTable"
	rowKey    = "rowkey1"
	hr        = "---------------------------------"
)

func init() {
	logrus.SetLevel(logrus.FatalLevel)
}

func main() {
	example()
}

func example() {
	client := gohbase.NewClient("hp001:61002")
	defer client.Close()

	//得到完整的一行数据
	log.Println("通过rowkey查询" + hr)
	getRequest, err := hrpc.NewGetStr(context.Background(), tableName, rowKey)
	if err != nil {
		log.Fatalf("hrpc.NewGetStr error:%v", err)
	}
	getRsp, err := client.Get(getRequest)
	if getRsp == nil {
		log.Printf("getRsp is nil")
	} else {
		for _, v := range getRsp.Cells {
			log.Printf("%s:%s", v.Qualifier, v.Value)
		}
	}

	//查询指定的列族
	log.Println("查询指定的列族" + hr)
	family := map[string][]string{"fileInfo": nil}
	getRequest, err = hrpc.NewGetStr(context.Background(), tableName, rowKey, hrpc.Families(family))
	if err != nil {
		log.Fatalf("hrpc.NewGetStr error:%v", err)
	}
	getRsp, err = client.Get(getRequest)
	if getRsp == nil {
		log.Printf("getRsp is nil")
	} else {
		for _, v := range getRsp.Cells {
			log.Printf("%s:%s", v.Qualifier, v.Value)
		}
	}

	//得到指定的qualifier
	log.Println("查询列族中的指定列" + hr)
	family = map[string][]string{"fileInfo": {"name"}}
	getRequest, err = hrpc.NewGetStr(context.Background(), tableName, "rowkey1", hrpc.Families(family))
	if err != nil {
		log.Fatalf("hrpc.NewGetStr error:%v", err)
	}
	getRsp, err = client.Get(getRequest)
	if getRsp == nil {
		log.Printf("getRsp is nil")
	} else {
		for _, v := range getRsp.Cells {
			log.Printf("%s:%s", v.Qualifier, v.Value)
		}
	}

	//scan操作
	log.Println("scan操作" + hr)
	scanRequest, err := hrpc.NewScanStr(context.Background(), tableName)
	if err != nil {
		log.Fatalf("hrpc.NewScanStr error:%v", err)
	}
	scanRsp := client.Scan(scanRequest)

	r, e := scanRsp.Next()
	for ; e == nil; r, e = scanRsp.Next() {
		log.Printf("result:%v", r)
	}

	//scan操作，设置filter
	log.Println("scan操作，设置filter" + hr)
	pFilter := filter.NewPrefixFilter([]byte("row"))
	scanRequest, err = hrpc.NewScanStr(context.Background(), tableName, hrpc.Filters(pFilter))
	if err != nil {
		log.Fatalf("hrpc.NewScanStr error:%v", err)
	}
	scanRsp = client.Scan(scanRequest)

	r, err = scanRsp.Next()
	for ; err == nil; r, err = scanRsp.Next() {
		log.Printf("result:%v", r)
	}
}
