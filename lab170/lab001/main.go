package main

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"github.com/mholt/archiver"
	"io/ioutil"
	"log"
)

func main() {
	example()
}

func example() {
	z := archiver.Zip{
		CompressionLevel:       flate.DefaultCompression,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      false,
		ImplicitTopLevelFolder: false,
	}

	//查看zip
	if err := z.Walk("data.zip", func(f archiver.File) error {
		zfh, ok := f.Header.(zip.FileHeader)
		if ok {
			fmt.Println("Filename:", zfh.Name)
		}

		if !f.IsDir() {
			content, err := ioutil.ReadAll(f)
			if err != nil {
				log.Fatalf("ioutil.ReadAll error:%v", err)
			} else {
				log.Println("content=", string(content))
			}
		}

		return nil
	}); err != nil {
		log.Fatalf("z.Walk error:%v", err)
	}
}
