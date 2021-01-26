package main

import (
	"flag"
	"github.com/kjk/notionapi"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var token string

func init() {
	pflag.StringVarP(&token, "token", "t", "", "token")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	demo001()
}

//get title
func demo001() {
	client := &notionapi.Client{}
	client.AuthToken = token

	pageID := "3f253e90fea2424aa1d7224b526b04a5"
	//pageID = "c26bee77-d48e-4a3f-ade2-36ff1ae49498"
	page, err := client.DownloadPage(pageID)
	if err != nil {
		log.Fatalf("DownloadPage() failed with %s\n", err)
	}
	log.Printf("%+v", page)

	title := page.Root().Title
	log.Printf("title=[%s]", title)

	//for _, v := range page.CollectionRecords {
	//	log.Printf("%+v", v)
	//	log.Printf("%s", v.Value)
	//}
	//
	//for _, v := range page.CollectionViewRecords {
	//	log.Printf("%+v", v)
	//	log.Printf("%s", v.Value)
	//	log.Printf("%+v", v.CollectionView)
	//}

	//table数据
	for _, v := range page.TableViews {
		log.Printf("%+v", v)
		for _, v2 := range v.Rows {
			log.Printf("column:%+v", v2)
			for _, v3 := range v2.Columns {
				log.Println(v3)
				for _, v4 := range v3 {
					log.Println(v4)
				}
			}
		}

		break
	}
}
