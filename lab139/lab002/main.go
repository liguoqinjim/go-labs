package main

import (
	"github.com/beevik/etree"
	"log"
)

func main() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("bookstore.xml"); err != nil {
		log.Fatalf("doc.ReadFromFile error:%v", err)
	}

	root := doc.SelectElement("bookstore")
	log.Println("ROOT element:", root.Tag)

	for _, book := range root.SelectElements("book") {
		log.Println("CHILD element:", book.Tag)

		if title := book.SelectElement("title"); title != nil {
			lang := title.SelectAttrValue("lang", "unknown")
			log.Printf("  TITLE:%s (%s)", title.Text(), lang)
		}

		for _, attr := range book.Attr {
			log.Printf("  ATTR: %s=%s", attr.Key, attr.Value)
		}
	}
}
